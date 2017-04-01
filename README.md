[![Build
Status](https://travis-ci.org/wfernandes/go-habitica.svg?branch=master)](https://travis-ci.org/wfernandes/go-habitica)

# go-habitica
golang client for habitica, a free habit building productivity app

## NOTE
This client is not ready yet. I'll do my best to flesh out most of the endpoints.

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
