# Golang api tester cli

Console CLI for api rest stress testing

## Build


### Linux: 

go build cmd/main.go  -o tester-linux

### Windows 32 bits:

GOOS=windows GOARCH=386 go build -o tester-32.exe cmd/main.go 

### windows 64 bits:

GOOS=windows GOARCH=amd64 go build -o tester-64.exe cmd/main.go 


## How to use

There are 2 mandatory files:

### tester.json: 
This file must contain configuration related to the endpoint:
```
{
  "method": "POST",
  "url": "http://localhost:8080/myEndpoint",
  "jwtHeader": "jwt-header-key",
  "jwtHeaderValue": "jwtHeaderValue",
  "jsonBodyPath": "./body.json"
}
```



### body.json: 
Json body for your request:
```
{
  "guid": "my guid",
  "someValue": "someValue"
}
```

Then you can run the executable file.

The cli has two options:

1- Stress test: N concurrent requests consuming the endpoint at the same time.

2- Interval stress test: You can define number of iterations and interval between them.


[Download binaries](https://github.com/deidelson/go-api-tester/releases/download/v0.0.1/tester.zip)

