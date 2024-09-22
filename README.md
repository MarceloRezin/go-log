# go-log
Utility for log management in golang

## Installation
`go get -u github.com/MarceloRezin/go-log`

## Structure

 _Operation_: Represents a log operation, like warning or info. Contains a prefix and a flag (const in log.go), wich configure what information is logged along with the prefix and message, such as date time or file.

 _Config_: Configuration to build the log wrapper, wich contains a writer (like a file in production or console e in development) and the _operation_ for each log type: info, warning and error.

 _Logger_: The wrapper that contains the ready-to-use log operations: info, warning and error.

## How to use

To get a looger, you can either build a custom one, using the `Init` method, passing the configuration, or using the default setups shown below.

## Default Setups
The lib provides two default setups for most common cases:

- DefaultConsole: Ideal for development, writes to the console, showing date time, and the file name and line number for warning and error.

- DefaultFile: Ideal for production, writes on file (`log.log`, in same directory as the execution), showing date time, and the file name and line number for warning and error.

Additionally, there are other methods to compose custom loggers, such as default operations and writers.

## Testing
To execute tests, use the `make test` command.