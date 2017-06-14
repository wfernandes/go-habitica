[![Build
Status](https://travis-ci.org/wfernandes/go-habitica.svg?branch=master)](https://travis-ci.org/wfernandes/go-habitica)

# go-habitica
golang client for habitica, a free habit building productivity app.

It is a client for the [v3 API](https://habitica.com/apidoc/).

## NOTE
There is support for quite a few endpoints regarding the tasks API.
If there is an endpoint that you'd like to see supported by the client feel
free to create a Github issue or a pull request :)

## Tests
### Unit
Run the unit tests as follows,
```
    go test ./... -short
```
### Integration
The integration tests hits the actual API. You will need to set environment
variables `API_TOKEN` and `USER_ID` provided by Habitica.
```
export API_TOKEN="some-api-token"
export USER_ID="some-user-id"
go test ./...
```

## Purpose
- I need this client for another IoT related project that I want to work on.
- I wanted to learn the idiomatic way of writing golang clients.
  This client is heavily influenced for today's notion of "best practices" for building clients.
  I'm using [go-github](https://github.com/google/go-github) as a guide.

## Project
I'm going to create tasks in [this project](https://github.com/wfernandes/go-habitica/projects/1)
to keep track of the work and progress.
