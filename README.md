# ninepod-microservices

## goa

v2 docs: [https://github.com/goadesign/goa/blob/v2/docs/Guide.md]

### Development

```
go install

go gen ninepod/design

```

Build from the `/cmd` directory

```
cd cmd/pods-cli
go build

```


Create examples: 

```
goa example ninepod/design

```

Run API Server:

(defaults port :80) 
Override port with `-http-port 8080`
```
cd cmd/http
go build
./http
```


Test the Service running on Server:

Pass in appropriate required request parameters as flags.

Example:

```
./http-cli pods pods --a 7674819484560861460 --b "Entertainment"
```