## gomacro - A Go interpreter with Lisp-like macros

The package `classic` contains the original old, small (and slow) implementation of gomacro interpreter.

For the current `fast` interpreter, see [../fast/README.md](../fast/README.md).

To learn about gomacro, download, compile and use it, please refer to the main [README.md](../README.md)

## Current Status

STABLE.

## Features and limitations

The classic interpreter has some additional limitations with respect to the fast one. Most notably:

* untyped constants and arithmetic on them, as `1<<100`, are evaluated as typed constants.
* types are not accurate when mixing untyped constants with typed values,
  i.e. `uint8(10) + 1` gives `uint64(11)` instead of `uint8(11)`.
* interpreted interfaces are not functional (they can only be declared).
* interpreted types cannot implement compiled interfaces.
* struct tags are ignored.
* support for embedded fields in structs is very limited - they mostly
  work as non-embedded fields.

