# `hotshot!`

This is `hotshot`. A LISP based language aimed to be used as bash scripts. 

## Installation

You need to install the `Go` compiler to build the compiler. It uses a Pure Go implementation and doesn't require any other dependency.

You can even build a complete `static` binary without any dynamic libraries.

After you have the `go` compiler working. Simply run

```sh
go build .
```

This should build the entire project and produce a binary named `hotshot`. 


## Usage

You can run `./hotshot --help` (On Unix) to get more help about it.

You can run `./hotshot` without any arguments to start the REPL. This acts like any other REPL completely supporting readline and even history.
You can type multi-line statements with ease.

If you have written a file with working hotshot code, (optionally) save it with a extension of `.ht` (Not at all needed, but we prefer it).

Use the following command to run the file.

```sh
./hotshot <file>
```

You can use other flags to have debug info or perform other operations with the source code/interpreter.

## Testing

The test are written again in Go's native test runner. 

You can run

`go test ./...` to run all the tests.

Or `go test ./lexer` to run specific tests (`lexer` in this case).

You can use the `-v` flag to provide information about all the subtests being run.

```sh
go test -v ./...
```

## Contribution

This project is under heavy development and contributions are highly appreciated.
A lot of decisions are already taken regarding the language, but a lot of them are still remaining.
Hope you can join us in making them.


