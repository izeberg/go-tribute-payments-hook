package settings

import (
	"fmt"
	"os"
	"strconv"
)

func parseInt64(key, dft string) int64 {
	value := os.Getenv(key)
	if value == "" {
		value = dft
	}
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(fmt.Errorf("%s: %s (%s)", key, err, value))
	}
	return v
}
func getStr(key, dft string) string {
	value := os.Getenv(key)
	if value == "" {
		value = dft
	}
	return value
}

var (
	FetchRetryCount        = int(parseInt64("retry_count", "3"))
	AppID                  = int32(parseInt64("telegram_api_id", ""))
	AppHash                = getStr("telegram_api_hash", "")
	SessionPath            = getStr("telegram_session_path", "./session")
	BotUsername            = getStr("telegram_incoming_from", "@tribute")
	ForwardTo              = parseInt64("telegram_forward_to", "0")
	WebhookURL             = getStr("webhook_url", "")
	WebhookLogin           = getStr("webhook_login", "")
	WebhookPassword        = getStr("webhook_password", "")
	WebhookSignatureKey    = getStr("webhook_signature_key", "")
	WebhookSignatureHeader = getStr("webhook_signature_header", "X-Signature")
	WebhookBatch           = getStr("webhook_batch", "0") == "1"
)
