This documentation is available on [GitHub](https://github.com/pspiagicw/hotshot) or [my website](https://falconite.xyz/hotshot)

# hotshot

`hotshot` is a LISP interpreter, borrowing features from Lua.


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
 - [Internals](#internals)
    - [lexer](#lexer)
    - [parser](#parser)
    - [eval](#eval)
    - [compiler](#compiler)
    - [vm](#vm)

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

# Development

Anybody is free to develop `hotshot`. You will need knowledge of Go and preferably understanding of *Compiler Theory*
.

# Contribution

This project is under heavy development and contributions are highly appreciated.
A lot of decisions are already taken regarding the language, but a lot of them are still remaining.
Hope you can join us in making them.

# concepts

`hotshot` is ultimately a LISP based langauge, so expect crazy amounts of brackets.


## expressions

`hotshot` supports expression, even `5` is a valid hotshot program.

Technically this means, hotshot doesn't have return statements.

There are few simple expression types, these also form the data-types.
- Integer (`5` or `123`), these evaluate to the integers themselves.
- String (`"this is a string"`)
- Booleans (`true` `false`)

You can perform operations on these data-types.

- Arithmetic `(+ 1 2)`, `(- 2 1)`, `(*)`

## variables

You can use `(let <name> <value>)` to store variables, that is the only way of updating those variables too!


For example

```lisp
(let name "pspiagicw")
```

## flow-control

It supports `if`, `if-else`, `while` and `cond` statements.

```lisp
(echo (if (not (= number 1))
        "Number is 1"
        "Number is not 1"))
```

```lisp
(while (> number 0)
       (do
         (echo "The value is:" number)
         (let number (- number 1))))
```

```lisp
(cond 
    ((= number 1) "Number is 1")
    ((= number 2) "Number is 2")
    (true "Number is something else"))
```

```lisp
(let number 3)
(while (> number 0)
       (do
         (echo (if (= number 3) "The number is 3!"))
         (let number (- number 1))
         (echo (cond 
              ((= number 1) "Number is 1")
              ((= number 2) "Number is 2")
              (true "Number is something else")))))
```

## functions

Obviously it's functional-programming, so here you have functions.

```lisp
(fn hello () 
    "Hello, hotshooter!")
```

```lisp
(fn helloName (name) 
    (concat "Hello " name))
```

```lisp
(let greet (lambda (name) (concat "Hello " name)))
(greet "hotshooter!")
```


```lisp

(fn arithmetic (operation x y) (operation x y))

(fn add (x y) (+ x y))

(echo (arithmetic add 3 4))

(echo (arithmetic (lambda (x y) (+ x y)) 3 4))
```

## builtins

We have lots of built in functions, you can add even your own.







