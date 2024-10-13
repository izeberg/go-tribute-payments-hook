# Telegram Tribute bot hook

Accepts messages from [@tribute](https://t.me/tribute) bot:
* Forwards new messages into other chat
* Sends new transactions into webhook url
* Initially fetches all transactions

Bot requires login into telegram account.

> The microservice guarantees that each message will be delivered or forwarded at least once. However, it cannot guarantee that the message will only be delivered once. Handling duplicate messages is the responsibility of the webhook receiver (for example, use transaction ID as a unique constraint).

## Docker
Everything could be run in docker, just use image [izeberg/go-tribute-payments-hook](https://hub.docker.com/repository/docker/izeberg/go-tribute-payments-hook).

## Settings
Application could be configured with environment variables defined here [./app/settings/settings.go](./app/settings/settings.go).

Here's example:
```
# get credentials from https://my.telegram.org
telegram_api_id=<api_id>
telegram_api_hash=<api_hash>

# peer id where messages will be accepted to processing
telegram_incoming_from=@trubute

# optional, this is where session will be stored (default /session)
telegram_session_path=/session

# optional, peer id to forward message
telegram_forward_to=1111111111

# optional url where HTTP POST with info will be sent
webhook_url=https://example.com/tribute
# optional basic auth parameters
webhook_login=user
webhook_password=pwd
# optional hmac signature key for webhook's body
webhook_signature_key=supersecretKey
```
