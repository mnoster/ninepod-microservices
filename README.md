# ninepod-microservices

## goa

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


Test API from cmd line.

Pass in appropriate required request parameters as flags.

Example:

```
./pods-cli -a 1 -b "entertainment"

```