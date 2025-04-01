# Much-Sender

A much simple SMTP email sender written in Go that accepts JSON POST requests.

## Such Features
- Configurable via TOML file
- Bearer token authentication
- HTML email support
- RESTful API endpoint

## So Installation
```bash
git clone https://github.com/dogeorg/much-sender.git
cd much-sender
go mod init github.com/dogeorg/much-sender
go get

## Much Use Example
```bash

curl -X POST http://localhost:42001/send-email \
     -H "Authorization: Bearer such-secure-token-here" \
     -H "Content-Type: application/json" \
     -d '{"reply_to_email":"from-shibe@domain.com","reply_to_name":"From Shibe","to_email":"to-shibe@domain.com","to_name":"For Shibe","subject":"Such Email Subject","html":"<h1>Hello</h1><p>Shibes</p>"}'
