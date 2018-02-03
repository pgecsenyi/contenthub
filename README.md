# ContentHub

A minimalist, cross-platform static file server that can be used to share smaller files over HTTP.

## Building

Make sure you have an appropriate version of _Go_ installed. Use one of the _run_ scripts or _Visual Studio Code_ to build the program. This will provide you one executable in the `bin` folder. You can also use the `go run` command of course. The _run_ scripts accept one of the following arguments: `build`, `fmt`, `lint`.

## Usage

The program accepts the following command line arguments.

  * `-content`: the path of directory that has to the content to share. Optional, the default value is `data/content`.
  * `-log`: the path of the log file. Optional, the default value is `data/log.txt`.
  * `-port`: the port to listen on. Optional, the default value is `8080`.
  * `-builtin`: use Go's built-in static server instead of the custom mechanism. Optional, the default value is `false`.

## Development environment

  * Go 1.7.4
  * Visual Studio Code 1.19.3
    * Extension: Go 0.6.73
