package tribute

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-tribute-api/settings"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"
)

func MakeAuthorizationHeader(webViewURL string) (string, error) {
	webURL, err := url.Parse(webViewURL)
	if err != nil {
		return "", err
	}
	fragmentValues, err := url.ParseQuery(webURL.RawFragment)
	if err != nil {
		return "", err
	}
	tgWebAppData := fragmentValues.Get("tgWebAppData")
	initDataUnsafe, err := url.ParseQuery(tgWebAppData)
	if err != nil {
		return "", err
	}
	var values []string
	for k := range initDataUnsafe {
		if k != "hash" {
			values = append(values, k)
		}
	}
	slices.Sort(values)

	for idx := range values {
		values[idx] = fmt.Sprintf("%s=%s", values[idx], initDataUnsafe.Get(values[idx]))
	}
	initData := base64.StdEncoding.EncodeToString([]byte(strings.Join(values, "\n")))
	authPayload := fmt.Sprintf("1;%s;%s", initData, initDataUnsafe.Get("hash"))
	authKey := base64.StdEncoding.EncodeToString([]byte(authPayload))
	return fmt.Sprintf("TgAuth %s", authKey), nil
}

func FetchTransactions(tributeAuth string, lastKnownTxID int64) ([]Transaction, int64, error) {
	nextFrom := "0"
	var maxTxID int64 = -1
	var transactions []Transaction

	for {
		resp, err := requestTransactionsRetry(tributeAuth, nextFrom)
		if err != nil {
			return transactions, maxTxID, err
		}
		for _, tx := range resp.Transactions {
			if tx.ID > maxTxID {
				maxTxID = tx.ID
			}
			if tx.ID > lastKnownTxID {
				transactions = append(transactions, tx)
			} else {
				break
			}
		}

		if resp.NextFrom == `` {
			break
		} else {
			nextFrom = resp.NextFrom
		}
	}
	return transactions, maxTxID, nil
}

func requestTransactionsRetry(tributeAuth string, nextFrom string) (*TransactionsResponse, error) {
	attempt := 0
	for {
		attempt++
		if resp, err := requestTransactions(tributeAuth, nextFrom); err == nil {
			return resp, nil
		} else if attempt >= settings.FetchRetryCount {
			return nil, err
		}
		time.Sleep(time.Second / 4)
		continue
	}
}

func requestTransactions(tributeAuth string, nextFrom string) (*TransactionsResponse, error) {
	body, _ := json.Marshal(map[string]string{"list": "dashboard_creator", "mode": "creator", "startFrom": nextFrom})
	req, err := http.NewRequest("POST", "https://subscribebot.org/api/v4/dashboard/transactions", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tributeAuth)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	content := &TransactionsResponse{}
	return content, json.Unmarshal(data, content)
}
