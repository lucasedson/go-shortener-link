# Go Shortener Link

A simple Link shortening Rest API

## Installation

```bash
go get -u github.com/luizalabs/go-shortener-link
```

## Usage

```bash
go run main.go
```

## Routes

- GET /api/v1/links - Get all links
- POST /api/v1/link - Create a new link
- GET /api/v1/link/:shortLink - Get a link by short link
- DELETE /api/v1/link/:shortLink - Delete a link by short link
- GET /:shortLink - Redirect to external link