# Go Backend API

An enterprise-level, yet simple, Go API that provides endpoints for basic operations.

## Features

- Clean architecture with separation of concerns
- JSON API responses
- Graceful server shutdown
- Structured logging
- Configuration management

## Getting Started

### Prerequisites

- Go 1.16+

### Installation

1. Clone the repository
   ```bash
   git clone https://github.com/yourusername/gobackend.git
   cd gobackend
   ```

2. Build the application
   ```bash
   go build -o api ./cmd/api
   ```

3. Run the application
   ```bash
   ./api
   ```

### Quick Start

To quickly run the application without building:
```bash
go run ./cmd/api
```

## API Usage

### Get Current Date and Time

```
GET http://127.0.0.1:3999/api/datetime
```

Example Response:
```json
{
  "current_time": "2023-08-01T14:30:45+01:00",
  "timestamp": 1690899045,
  "utc": "2023-08-01T13:30:45Z"
}
```

