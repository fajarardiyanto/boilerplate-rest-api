# Boilerplate Rest Api Flt Go Router

**boilerplate-rest-api** is base for generates standard structure build rest api application using [flt-go-router](https://github.com/fajarardiyanto/flt-go-router).

## Getting started
Required Golang 1.17+
### Run
```bash
make run
```

### Build Docker
```bash
make docker-build
```

### Run Docker
```bash
make build-run
```

#### Tips
Maybe it would be better to do some basic code scanning before pushing to the repository.
```sh
# for *.nix users just run gosec.sh
# curl is required
# more information https://github.com/securego/gosec
make scan
```