# ninepod-microservices

## goa

v2 docs: [https://github.com/goadesign/goa/blob/v2/docs/Guide.md]

### Development

```
go install

go gen ninepod/design
```

Create examples: 

```
goa example ninepod/design
```

#### Building from the `/cmd` directory:

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
cd cmd/http-cli

go build

./http-cli pods pods --a 7674819484560861460 --b "Entertainment"
```