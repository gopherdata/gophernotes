The command `import` will write here import files for third-party
packages if they are to be imported and statically linked into gophernotes.

This is currently needed on Windows, *BSD, Android etc. (to be exact,
as of Go 1.13 it's needed on all operating system except Linux and Mac OS X).

It is safe to remove files in this directory,
as long as you **keep at least** README.md and a_package.go
