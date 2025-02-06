# AWS-SES-Emulator

This is a mock implementation of AWS Simple Email Service (SES) API using Go and Gin. It mimics AWS SES behavior for testing without actually sending emails.

## Features

- **AWS SES API Compatibility**:
  - Matches the request and response structure of AWS SES API, ensuring seamless integration.

- **Email Sending Limits**:
  - **Day 1**: 5 emails/day.
  - **Day 2**: Resets count and allows sending of 5 additional emails.
  - **After 10 emails sent**: Limit increases to 10 emails/day.
  - **After 100 total emails sent**: Limit increases to 100 emails/day.
  - **After 250 total emails sent**: Limit increases to 250 emails/day.

- **Tracks Statistics**:
  - Tracks emails sent in the last 24 hours.
  - Tracks the total number of emails sent.
  - Tracks the max send limit per day, adapting based on progress.

- **Error Handling & Scenarios**:
  - Implements common AWS SES error codes (e.g., quota exceeded, invalid email address).
  - Supports various test scenarios such as:
    - Exceeding sending quotas.
    - Warm-up rules (e.g., gradual sending increases).
    - Temporary failures (e.g., throttling).

## Installation & Setup

### Prerequisites

Before you begin, make sure you have Go installed on your machine. If not, follow the instructions below for your respective platform:

#### Install Go (if not installed)

```bash
# macOS
brew install go

# Linux
sudo apt install golang-go

# Windows
winget install GoLang.Go
```

#### Clone the Repository

```bash
git clone https://github.com/your-repo/aws-ses-emulator.git
cd aws-ses-emulator
```

#### Run the API

```bash
go run main.go
```

By default, the API runs on http://localhost:8080.

## API Endpoints

### 1️⃣ Health Check

- Endpoint: ```GET /health```

- Response: 

```{ "message": "API is running" }```

### 2️⃣ Send Email

- Endpoint: ```POST /send-email```

- Response: 

```{ "message": "mock-message-id-123" }```

- If limit exceeded: 

```{ "error": "Sending Limit Exceeded. Wait for 24 hours!" }```

### 3️⃣ Send Raw Email

- Endpoint: ```POST /send-raw-email```

- Response:

```{ "message": "mock-raw-message-id-456" }```

### 4️⃣ Get Send Quota

- Endpoint: ```GET /get-send-quota```

- Response:

```
{
  "Max24HourSend": 5,
  "SentLast24Hours": 3,
  "TotalEmailCount": 8,
  "MaxSendRate": 1.0
}
```

### 5️⃣ Get Send Statistics

- Endpoint: ```GET /get-send-statistics```

- Response:

```
{
  "DeliveryAttempts": 3,
  "TotalEmailsSent": 8,
  "Bounces": 0,
  "Complaints": 0,
  "Rejects": 0
}
```

## AWS Documentation & References

- [AWS SES API Reference](https://docs.aws.amazon.com/ses/latest/APIReference/)
- [Understanding SES Sending Limits](https://docs.aws.amazon.com/ses/latest/dg/manage-sending-limits.html)


## Common Error Codes & Meanings

## Error Codes

| **Error Code**         | **Meaning**                |
|------------------------|----------------------------|
| 400 BadRequest         | Invalid request structure  |
| 403 AccessDenied       | User is not authorized     |
| 429 TooManyRequests    | Sending limit exceeded     |
| 500 InternalServerError| Unexpected server error    |

## Future Enhancements

- Add logging to track API requests.

- Store email history using database instead of in-memory variables.

- Implement email warm-up rules like AWS SES.

- Create mock bounces and complaints for better testing.

## License

This project is licensed under the Apache 2.0 License. Check out the full details in the [LICENSE](https://github.com/ShamyaHaria/AWS-SES-Emulator?tab=Apache-2.0-1-ov-file) file.

