# Spark

This is a go wrapper (CLI + API client)  to the [spark.io](https://www.spark.io) API.


## Why another incomplete CLI?

The spark team has released [spark-cli](https://github.com/spark/spark-cli/) which is based on nodeJS and distributed via npm.
This CLI aims to eventually be a full replacement for spark-cli.


## CLI example usage

```
$ spark
usage: spark [--help] <command> [<args>]

Available commands are:
    claim         Claims a spark core
    devices       Lists devices for authenticated user
    exec          Calls a function exposed by the core
    info          Displays basic information about the given Core
    invalidate    Invalidates an access token. Requires username/password
    login         Log in to spark cloud
    logout        Logout from spark cloud
    read          Reads the value of variables exposed by the spark core
    rename        Renames a core
    tokens        List all access tokens. Requires username/password
```

For example, to rename your core:

```sh
$ spark devices
Error connecting to Spark cloud: You should login first.

$ spark login
Username:  me@example.com
Password: 
Successfully logged in. Access token persisted to: ~/.sparkio

$ spark devices
- Device: spork [53dd73045076535132181487], connected?: false

$ spark rename 53dd73045076535132181487 new_name
Successfully renamed core 53dd73045076535132181487 to new_name
```


## API Client example usage

```golang

client := spark.NewClient(nil)
token, resp, err := client.Tokens.Login(username, password)

devices, resp, err := client.Devices.List()
resp, err := client.Devices.Claim(coreId)

resp, err := client.Devices.Rename(coreId, new_name)
```

## TODO
- add version
- add tests
- document the spark API endpoints


## Thanks

- To [@Mitchellh](https://github.com/mitchellh) for his [cli](https://github.com/mitchellh/cli) package
- To [@armon](https://github.com/armon) for a well documented [example](https://github.com/hashicorp/consul) package of the cli package
- To [@google](https://github.com/google) package for [go-github](https://github.com/google/go-github) from which the spark client is heavily inspired.