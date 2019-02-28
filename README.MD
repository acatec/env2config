# env2config

Loads env vars from a .env file and current environment to your config struct.

Vars from .env file will override current environmental variables.

thanks to https://github.com/joho/godotenv and https://github.com/caarlos0/env

``env2config`` is just wrapper on over these perfect libraries

## Installation

As a library

```shell
go get github.com/acatec/env2config
```

or using ``dep``

```shell
 dep ensure -add github.com/acatec/env2config
```


## Usage

Add your application configuration to your `.env` file in the root of your project:

```shell
PORT=8080
APP_NAME=My go app
```

Then in your Go app you can do something like

```go
package main

import (
    "github.com/acatec/env2config"
)

type Config struct {
	Port string `env:"PORT" envDefault:"80"`
	AppName string `env:"APP_NAME" envDefault:"noname"`
}

func main() {
	
	var cfg Config
	_ = env2config.PrepareConfig(&cfg)

    // now do something with cfg.Port and cfg.AppName
}
```


And if you have few env files:

```go
_ = env2config.PrepareConfig(&cfg, "first.env", "second.env")

```