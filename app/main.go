package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/amarnathcjd/gogram/telegram"
	"github.com/natefinch/atomic"
	"go-tribute-api/settings"
	"go-tribute-api/tg"
	"go-tribute-api/tribute"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"
	"time"
)

var newTXLock = &sync.Mutex{}

func fetchNewTransactions(client *telegram.Client) {
	if !newTXLock.TryLock() {
		return
	}
	defer newTXLock.Unlock()

	tributeAuth, err := getTributeAuth(client)
	if err != nil {
		log.Println("Failed to read last known transaction ID", err)
		return
	}

	savedTxID, err := readLastKnownTxID()
	if err != nil {
		log.Println("[WARN] Failed to read last known transaction ID", err)
	}

	transactions, maxTxID, err := tribute.FetchTransactions(tributeAuth, savedTxID)
	if err != nil {
		log.Println("Failed to fetch transactions:", err)
		return
	}

	if err := sendTransactions(transactions); err != nil {
		log.Println("Failed to send transactions:", err)
		return
	}

	if maxTxID > savedTxID {
		if err := saveLastKnownTxID(maxTxID); err != nil {
			log.Println("[WARN] Failed to save last known transaction ID:", err)
		}
	}
}

func sendTransactions(transactions []tribute.Transaction) error {
	if len(transactions) == 0 || settings.WebhookURL == "" {
		log.Println("No new transactions")
		return nil
	}
	if settings.WebhookBatch {
		log.Println("Sending batched transactions to webhook:", len(transactions))
		data, _ := json.Marshal(transactions)
		return sendDataRetry(data)
	} else {
		waiter := sync.WaitGroup{}
		var failed error
		for _, transaction := range transactions {
			waiter.Add(1)
			go func(tx tribute.Transaction) {
				defer waiter.Done()
				log.Println("Sending transaction", tx.ID)
				if err := sendDataRetry(tx); err != nil {
					log.Println("[WARN] Failed to send transaction:", err)
					failed = err
				}
			}(transaction)
		}

		waiter.Wait()
		return failed
	}
}

func sendDataRetry(data interface{}) error {
	if data == nil {
		return nil
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	var signature string
	if settings.WebhookSignatureKey != "" {
		signer := hmac.New(sha256.New, []byte(settings.WebhookSignatureKey))
		signer.Write(body)
		signature = hex.EncodeToString(signer.Sum(nil))
	}

	attempt := 0
	for {
		attempt++
		if err := sendData(body, signature); err == nil {
			return nil
		} else if attempt >= settings.FetchRetryCount {
			return err
		}
		time.Sleep(time.Second / 4)
	}
}

func sendData(body []byte, signature string) error {
	req, err := http.NewRequest("POST", settings.WebhookURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if settings.WebhookLogin != "" || settings.WebhookPassword != "" {
		req.SetBasicAuth(settings.WebhookLogin, settings.WebhookPassword)
	}
	if signature != "" {
		req.Header.Set(settings.WebhookSignatureHeader, signature)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("status code %d", resp.StatusCode)
	}
	return nil
}

func saveLastKnownTxID(txID int64) error {
	return atomic.WriteFile(path.Join(settings.SessionPath, "last_known_tx.dat"), bytes.NewBuffer([]byte(strconv.FormatInt(txID, 10))))
}

func readLastKnownTxID() (int64, error) {
	if fd, err := os.Open(path.Join(settings.SessionPath, "last_known_tx.dat")); err != nil {
		return 0, err
	} else {
		defer fd.Close()
		data, err := io.ReadAll(fd)
		if err != nil {
			return 0, err
		}
		return strconv.ParseInt(string(data), 10, 64)
	}
}

func getTributeAuth(client *telegram.Client) (string, error) {
	if fd, err := os.Open(path.Join(settings.SessionPath, "tribute.auth")); err == nil {
		defer fd.Close()
		data, err := io.ReadAll(fd)
		if err != nil {
			return "", err
		}
		return string(data), nil
	} else {
		webViewURL, err := tg.RequestBotWebView(client, settings.BotUsername)
		if err != nil {
			return "", err
		}
		auth, err := tribute.MakeAuthorizationHeader(webViewURL.URL)
		if err != nil {
			return "", err
		}
		if err := atomic.WriteFile(path.Join(settings.SessionPath, "tribute.auth"), bytes.NewBuffer([]byte(auth))); err != nil {
			log.Println("[WARN] Failed to save tribute.auth:", err)
		}
		return auth, nil
	}
}

func main() {
	client, err := tg.RunningClient()
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Stop()

	botUser, err := client.ResolveUsername(settings.BotUsername)
	if err != nil {
		log.Fatalln(err)
	}

	client.On(telegram.OnMessage, func(m *telegram.NewMessage) error {
		if m.Message.Out {
			return nil
		}
		if settings.ForwardTo != 0 {
			log.Println("Forward message", m.Message.ID, "to", settings.ForwardTo)
			if _, err := m.Client.Forward(settings.ForwardTo, m.Peer, []int32{m.Message.ID}); err != nil {
				log.Println("[WARN] Failed to forward message", m.Message.ID, err)
			}
		}
		fetchNewTransactions(m.Client)
		return nil
	}, telegram.FilterUsers(botUser.(*telegram.UserObj).ID))

	log.Println("Fetching new transactions")
	go fetchNewTransactions(client)

	log.Println("Wait for messages from", settings.BotUsername)
	client.Idle()
}
