package tg

import (
	"fmt"
	"github.com/amarnathcjd/gogram/telegram"
	"go-tribute-api/settings"
	"os"
	"path"
)

func GetBotByUsername(client *telegram.Client, username string) (*telegram.UserObj, *telegram.UserFull, error) {
	botUser, err := client.ResolveUsername(username)
	if err != nil {
		return nil, nil, err
	}
	botUserObj, ok := botUser.(*telegram.UserObj)
	if !ok {
		return nil, nil, fmt.Errorf("could not resolve user")
	}
	fullBot, err := client.UsersGetFullUser(&telegram.InputUserObj{UserID: botUserObj.ID, AccessHash: botUserObj.AccessHash})
	if err != nil {
		return nil, nil, err
	}
	return botUserObj, fullBot.FullUser, nil
}

func RequestBotWebView(client *telegram.Client, username string) (*telegram.WebViewResultURL, error) {
	botUser, botFull, err := GetBotByUsername(client, username)
	if err != nil {
		return nil, err
	}

	return client.MessagesRequestWebView(&telegram.MessagesRequestWebViewParams{
		Peer:        &telegram.InputPeerUser{UserID: botUser.ID, AccessHash: botUser.AccessHash},
		Bot:         &telegram.InputUserObj{UserID: botUser.ID, AccessHash: botUser.AccessHash},
		URL:         botFull.BotInfo.MenuButton.(*telegram.BotMenuButtonObj).URL,
		FromBotMenu: true,
		Platform:    settings.WebAppPlatform, // Dear tribute, pls don't ban, we need API......
	})
}

func RunningClient() (*telegram.Client, error) {
	os.MkdirAll(settings.SessionPath, os.ModePerm)

	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID:        settings.AppID,
		AppHash:      settings.AppHash,
		Session:      path.Join(settings.SessionPath, "gogram.dat"),
		DisableCache: true,
	})
	if err != nil {
		return nil, err
	}

	if err := client.Start(); err != nil {
		return nil, err
	}

	if err := client.AuthPrompt(); err != nil {
		defer client.Stop()
		return nil, err
	}
	return client, nil
}
