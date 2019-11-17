# BenchHog

[![GoDoc](https://godoc.org/github.com/apiotrowski312/BenchHog/lib?status.svg)](https://godoc.org/github.com/apiotrowski312/BenchHog/lib)

[![GoDoc](https://godoc.org/github.com/apiotrowski312/BenchHog/results?status.svg)](https://godoc.org/github.com/apiotrowski312/BenchHog/results)

## app is under development, there is no stable version

BenchHog is simple to use application which allows you to do some performance testing of your app. App aim to provide the simplicity of ApacheBench, but in more performant way with more features like benching multiple endpoints at once.

I tried tried to create app simmilar to ApacheBench in case of

Benchhog is an app created simply because I wanted to learn programing in Go.

## Instalation

### Pre-compiled executables

Not available at the moment

<!-- You can download them [HERE](https://github.com/apiotrowski312/BenchHog/releases). -->

### apt-get

Not available at the moment.

### Source

At the moment, the only way to build and use BenchHog is by building it from source.

Fortunately, it's pretty simple and straightforward. Clone repo with git clone XXX and call make build inside cloned repo. It will create binary inside bin catalog.

## Usage

| flag  | short | default value | description                                                      |
| ----- | ----- | ------------- | ---------------------------------------------------------------- |
| ratio |       | 10            | Number of requests at once (more == better is not always a case) |
|       | r     | 100           | Number of request app should send                                |
| json  |       |               | Provide json to send with POST request                           |

### GET

```bash
./bin/benchHog get -r 3000 --ratio 40 github.com gitlab.com
```

### POST

```bash
./bin/benchHog post -r 3000 --json '{"name": "USER", "password": "123"}' github.com gitlab.com
```
