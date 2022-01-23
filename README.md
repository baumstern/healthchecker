healthchecker
-------------

healthchecker monitors a liveness of various blockchain (e.g. Ethereum, Klaytn...etc)

## Quickstart

Run server which listen to 8080 port:

```
go run main.go
```

Then go to http://localhost:8080 via your web browser (e.g. Chrome, Safari...etc)


Alternatively, you can build the binary and then run it.

Build command will generate the binary named `healthchecker`:

```
go build
```

Then, execute below command to run:

```
./healthchecker
```

## API endpoint

### GET `watch?network=<blockchain_name>`
 
request:

```
curl -X GET localhost:8080/watch?network=ethereum
```

response:

```
{"block_num":14064245,"timestamp":"2022-01-24T06:06:00.399725+09:00"}
```

#### Supported blockchain

* Ethereum `watch?network=ethereum`
* Klaytn `watch?network=klaytn`