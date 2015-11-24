# Daily stars

Review stared GitHub repositories.

![microservice-ready](https://img.shields.io/badge/Microservice-Ready%F0%9F%91%8D-green.svg)

## Install

```sh
go get github.com/rollbrettler/daily-stars/cmd/daily-stars
```

## Usage

### Get request
```
https://daily-stars.herokuapp.com/<username>
```

### Get json request
```
https://daily-stars.herokuapp.com/<username>.json
```

## Development

### Dependecies
- godep

```sh
mkdir -p $GOPATH/src/github.com/rollbrettler/daily-stars
git clone git@github.com:rollbrettler/daily-stars.git $GOPATH/src/github.com/rollbrettler/daily-stars
cd $GOPATH/src/github.com/rollbrettler/daily-stars
godep restore
```
