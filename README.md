This documentation is available on [GitHub](https://github.com/pspiagicw/hotshot) or [my website](https://falconite.xyz/hotshot)

# Contents

 - [Hotshot](#hotshot)
    - [status](#status)
    - [getting started](#getting-started)
    - [testing](#testing)
    - [development](#development)
    - [contribution](#contribution)
 - [Concepts](#concepts)
    - [expressions](#expressions)
    - [variables](#variables)
    - [flow-control](#flow-control)
    - [functions](#functions)
    - [builtin](#builtin)
 - [Data Structures](#data-stucture)
    - [tables](#table)
    - [list](#list)
    - [hash](#hash)
 - [Import](#import)
    - [import](#import)
    - [stdlib](#stdlib)
 - [Internals]()

# hotshot

This is `hotshot`. A LISP based language written in Golang.

It's a dynamically typed language with hint of languages like Lua.
It's built in Go.

It looks like this

```lisp
(fn fizzbuzz (n)
  (cond ((and (= (mod n 3) 0) (= (mod n 5) 0)) "FizzBuzz")
        ((= (mod n 3) 0) "Fizz")
        ((= (mod n 5) 0) "Buzz")
        (true n)))
```

# Status

This language is practically newborn and has a lot of it's features unplanned.
But it's big enough to be considered releasing to the public.

It does have around 40 builtin functions implemented. 
But no you can't use this language in production.

Its designed for reading and experimenting, and maybe small scripts.

# Getting Started

The only way to use the langauge is to manually compile and run it on the CLI.

It has a builtin REPL, along with ability to run scripts.

## Compilation

1. Clone the project

```sh
git clone https://github.com/pspiagicw/hotshot
```

2. Compile using `Go`

```sh
cd hotshot
go build .
```

## Running

Run the binary without any arguments to open the REPL.

```sh
./hotshot
```

To run a script pass it as a argument.

```sh
./hotshot <script-to-run>
```
## Testing

You can run

1. To run all the tests.


```sh
go test ./...
```

2. To run specific tests (`lexer` in this case).

```sh
go test ./lexer
```

3. use the `-v` flag to provide information about all the subtests being run.

```sh
go test -v ./...
```

# Contribution

This project is under heavy development and contributions are highly appreciated.
A lot of decisions are already taken regarding the language, but a lot of them are still remaining.
Hope you can join us in making them.


