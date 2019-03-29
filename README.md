# Selection

Side hiring process test

[![CircleCI](https://circleci.com/gh/Elojah/selection/tree/master.svg?style=svg)](https://circleci.com/gh/Elojah/selection/tree/master)

## Installation

OS X & Linux & Windows:

```sh
go get -u github.com/elojah/selection
```
## Usage example
```sh
> docker-compose up # start containers
> make import # import test data into mongo
> curl -k X GET https://127.0.0.1:8080/task/scores?id=0E8dlOR1tqfg31jIR
```

## TODO

- [x] Add tags to get task (API)
- [x] Batch users retrieve in scores calculation (Scorer)
- [ ] Debug HTTPS Docker (API)
- [ ] Add more routes (get all, get by id) (API)
- [ ] Add errors.Wrap for error context
