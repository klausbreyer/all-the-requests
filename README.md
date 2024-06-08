# All The Requests

This is a simple HTTP server written in Go that logs all incoming requests with details including timestamp, path, headers, and payload. It's useful for debugging callbacks and webhooks.

## Features

- Logs the request path, headers, and body.
- Adds a timestamp to each request log.
- Configurable port via command-line arguments.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

Clone the repository:

```sh
git clone https://github.com/klausbreyer/all-the-requests.git
cd all-the-requests
```

### Usage

Build and run the server:

```sh
go build -o all-the-requests
./all-the-requests -port=8080
```

By default, the server listens on port 8080. You can specify a different port using the `-port` flag.

### Example

Start the server:

```sh
./all-the-requests -port=9090
```

Output:

```
Starting server on :9090
Server is ready to handle requests and display all details.
```

Make a request to the server:

```sh
curl -X POST http://localhost:9090 -d "test payload" -H "Content-Type: text/plain"
```

The server logs:

```
Timestamp: 2023-06-08T15:04:05Z
Path: /
Headers:
Content-Type: text/plain
Payload: test payload
```

### Downloads

You can download the (unsigned) binaries [here](https://github.com/klausbreyer/all-the-requests/tags).

### Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Acknowledgments

- [Go](https://golang.org/) for the programming language.
