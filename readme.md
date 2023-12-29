## About
Simple JSON/HTTP commit log service

### How to Run
cd to `.\cmd\server` and run `go run main.go`

### Add Records to Log
Endpoint
`POST localhost:8080`

Request
```json
{
    "record": {
        "value":{valid_Base64_string}
    }
}
```

Response
```json
{
    "offset": {log_offset}
}
```

### Read Records from Log
Endpoint
`GET localhost:8080`

Request
```json
{
    "offset": {log_offset}
}
```

Response
```json
{
    "record": {
        "value":{valid_Base64_string}
    }
}
```