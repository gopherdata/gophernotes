The gomacro command `import` will write here import files for third-party
libraries if they are to be imported and statically linked into gomacro.

This is currently needed on non-Linux systems
to allow importing third-party libraries at gomacro prompt.

It is safe to remove files in this directory,
as long as you **keep at least** README.md and a_package.go
