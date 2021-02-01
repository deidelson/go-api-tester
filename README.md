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


