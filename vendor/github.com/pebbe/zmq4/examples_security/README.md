These are Go versions of the C examples described in
[Using ZeroMQ Security (part 2)](http://hintjens.com/blog:49) by Pieter Hintjens.
Those C examples use the [zauth module](http://czmq.zeromq.org/manual:zauth)
in the [czmq library](http://czmq.zeromq.org).

There are some differences:

 * The zauth module doesn't handle domains. The Go version does.
 * The zauth module handles files with usernames/passwords and directories with certificates.
   The Go version just uses maps of usernames/passwords and lists of public user keys.
