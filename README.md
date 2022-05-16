# The Monkey Language

This is an interpreter written in Go, as described in an interesting book [Writing An Interpreter In Go](https://interpreterbook.com/)

The language looks like this:

```
let five = 5;
let ten = 10;

let add = fn(x, y) {
    x + y;
};

let result = add(five, ten);
```

## Start a REPL

```
go run main.go
```
