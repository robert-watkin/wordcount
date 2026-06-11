# wordcount
This is my first go cli project. This is intended to get me up to speed with Go fundamentals.
The script will take in a file or read from stdin and count the number of words in it.

## install

go install github.com/robert-watkin/wordcount@latest

## usage

```sh
wordcount --min <number> --top <number> <file>
```

## example

```sh
wordcount README.md
wordcount --min 5 --top 10 README.md
wordcount // read from stdin
```

## testing

```sh
go test // run tests
go test -v // verbose
```