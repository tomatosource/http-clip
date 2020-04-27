# http clip

Simple http server with 2 routes:
```
GET /paste # returns clipboard
POST /copy # sets clipboard to request body
```

## Installation

```
go get github.com/tomatosource/http-clip
```

## Usage

Start server:

```
go run main.go
```

Paste:

```
curl localhost:8833/paste
```

Copy:

```
curl -d 'example text' localhost:8833/copy
```
