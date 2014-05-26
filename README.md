# Spark

This is a go CLI wrapper to the [spark.io](https://www.spark.io)


## Why another incomplete CLI?

The spark team has released [spark-cli](https://github.com/spark/spark-cli/) which is based on nodeJS and distributed via npm.
This CLI aims to eventually be a full replacement for spark-cli.


## Example usage

```
$ spark
usage: spark [--help] <command> [<args>]

Available commands are:
    claim      Claims a spark core
    devices    Lists devices for authenticated user
    exec       Calls a function exposed by the core
    info       Displays basic information about the given Core
    login      Log in spark cloud
    read       Reads the value of variables exposed by the spark core
    rename     Renames a core
    tokens     List all access tokens
```

For example, to rename your core:

```
$ spark rename 53ff73065075582132181487 new_name
Successfully renamed core 53ff73065075582132181487 to new_name
```


## TODO

- add tests
- document the spark API endpoints