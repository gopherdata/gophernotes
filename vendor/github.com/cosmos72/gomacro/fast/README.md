## gomacro - A Go interpreter with Lisp-like macros

The package `fast` contains a faster reimplementation of gomacro interpreter.

To learn about gomacro, download, compile and use it, please refer to the main [README.md](../README.md)

## Current Status

STABLE.

## Features and limitations

See [../doc/features-and-limitations.md](../doc/features-and-limitations.md)


## Misc TODO notes

* contact github.com/neugram/ng author?
* when importing a package, reuse compiled .so if exists already?
* gomacro FILE: execute all the init() functions, then execute main() if (re)defined and package == "main"
* try to run Go compiler tests
