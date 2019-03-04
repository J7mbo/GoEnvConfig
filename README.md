# GoEnvConfig
*Immutable configuration loaded from environment variables.*

[![Build Status](https://travis-ci.com/J7mbo/GoEnvConfig.svg?branch=master)](https://travis-ci.com/J7mbo/GoEnvConfig)
[![codecov](https://img.shields.io/codecov/c/github/j7mbo/GoEnvConfig.svg?branch=master)](https://codecov.io/gh/J7mbo/GoEnvConfig)
[![GoDoc](https://godoc.org/github.com/J7mbo/GoEnvConfig?status.svg)](https://godoc.org/github.com/J7mbo/GoEnvConfig)
[![Version](https://img.shields.io/github/tag/j7mbo/GoEnvConfig.svg?label=version)](github.com/j7mbo/GoEnvConfig)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE.md)

Automatically load environmental variables into structs with private properties.

## Installation

```bash
go get github.com/j7mbo/goenvconfig
```

## Example

Bash:
```bash
export PORT=1337 
```

Go:
```go
package main

import (
    "github.com/j7mbo/goenvconfig"
    "fmt"
)

type Config struct {
    host     string  `env:"HOME" default:"localhost"`
    port     int     `env:"PORT" default:"8080"`
}

func (c *Config) GetHost() string { return c.host }
func (c *Config) GetPort() int { return c.port }

func main() {
    config := Config{}
    parser := goenvconfig.NewGoEnvParser()
    
    if err := parser.Parse(&config) {
    	panic(err)
    }
    
    fmt.Println(config.GetHost()) // localhost
    fmt.Println(config.GetPort()) // 1337
}
```

## Supported Types

For now the following simple types are supported:

- int
- string

## Why

Just because you want to automatically load environment variables into configuration structs does not mean you should
expose modifiable exported properties on your configuration object. Instead the struct should be immutable with
properties only accessible via getters.

You can either idiomatically create a factory method thereby greatly reducing the simplicity of an automated solution, 
or you do something you're "not supposed to" and use a library that utilises the `reflect` and `unsafe` packages.