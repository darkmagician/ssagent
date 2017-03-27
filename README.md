# ssagent

## Introduction

The tool does the followings:

1. Get the free shadowsocks configuration from [ishadow](http://fast.ishadow.online/)
1. Launch the local agent of shadowsocks-go with the configuration.

It is developed based on [shadowsocks-go](https://github.com/shadowsocks/shadowsocks-go). 

## Usage

### Install

```shell
 go get github.com/darkmagician/ssagent
```

### Run

```shell
  ssagent 
```

By default, the sock5 port is 1080.