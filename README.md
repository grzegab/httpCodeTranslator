# Http Code Translator
Used for gateway when http request need to be translated into gRPC.

## Usage
Just create object and use it:
```go
    t := Translator{httpCode: 200}
	gCode, gMessage := t.httpToGRPC()
```

## Author
Greg, Lord of PHP Solutions, Master of Golang Services, Warden of RESTful Gates,
Keeper of gRPC Streams, Protector of SQL Queries, Overlord of RabbitMQ Channels, and Guardian of Redis Keys