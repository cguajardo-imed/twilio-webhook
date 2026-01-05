# Twilio Webhook

Go webhook server for receiving and processing Twilio SMS messages.

## Setup

1. Install dependencies:
   ```
   go mod download
   ```

2. Create `.env` file with Twilio credentials:
   ```
   TWILIO_ACCOUNT_SID=your_account_sid
   TWILIO_AUTH_TOKEN=your_auth_token
   ```

3. Run:
   ```
   go run main.go
   ```

Server listens on `localhost:3000`

## Endpoint

**POST /** - Receives Twilio SMS webhooks, logs message details and fetches message price.

## Stack

- Go 1.24
- Fiber v2 (web framework)
- Twilio Go SDK
