![alt tag](https://raw.githubusercontent.com/gopherds/gophernotes/master/files/gophernotes2.jpg)

# gophernotes - Go in Notebooks

`gophernotes` is a Go kernel for [Jupyter](http://jupyter.org/) notebooks.  Finally, we can have a reliable and maintained way to use Go interactively and code in the browser.  Use `gophernotes` to create and share documents that contain live Go code, equations, visualizations and explanatory text.  These notebooks, with the live Go code, can then be shared with others via email, Dropbox, GitHub and the [Jupyter Notebook Viewer](http://nbviewer.jupyter.org/).  Go forth and do data science, or anything else interesting, with go notebooks!

This project came out of the [Gopher Gala](http://gophergala.com/) 2016.  It is inspired by a REPL called [gore](https://github.com/motemen/gore) and by a limited kernel called [iGo](https://github.com/takluyver/igo), which is no longer maintained.

## Screenshots/Examples

### Simple interactive use:

![alt tag](https://raw.githubusercontent.com/gopherds/gophernotes/master/files/screencast.gif)

### Example Notebooks (dowload and run them locally, follow the links to view in Github, or use the [Jupyter Notebook Viewer](http://nbviewer.jupyter.org/)):
- [Simple Printing and Channels](https://github.com/gopherds/gophernotes/blob/master/examples/Simple-Example.ipynb)
- [Pattern Recognition with Golearn](https://github.com/gopherds/gophernotes/blob/master/examples/Pattern-Recognition.ipynb)
- [Feed Forward, Recurrent Neural Nets](https://github.com/gopherds/gophernotes/blob/master/examples/Feed-Forward-Recurrent-NN.ipynb)
- [Time Parsing, Formatting](https://github.com/gopherds/gophernotes/blob/master/examples/Time-Formatting-Parsing.ipynb)
- [Stateful Goroutines](https://github.com/gopherds/gophernotes/blob/master/examples/Stateful-Goroutines.ipynb)
- [Worker Pools](https://github.com/gopherds/gophernotes/blob/master/examples/Worker-Pools.ipynb)

## Installation

### Docker

- Pull down and run the [latest image](https://hub.docker.com/r/dwhitena/gophernotes/):

  ```
  docker pull dwhitena/gophernotes:latest
  docker run --name gophernotes --net host -d dwhitena/gophernotes:latest
  ```

*Possible issues* - For OSX Docker Machine / Dlite users, you may need to set the IP to `0.0.0.0` instead of the default  `localhost` with:

  ```
  docker run -p 8888:8888 -d dwhitena/gophernotes jupyter notebook --no-browser --ip=0.0.0.0
  ```

*Note* - this is a pretty large image, because it contains a full distribution of [Anaconda](http://docs.continuum.io/anaconda/index) plus the add ons of gophernotes.  However, with this image, you can create Go notebooks, Python notebooks, text files, run ipython in the shell, etc.

### Local, Linux

Make sure you have the following dependencies:

  - [Go](https://golang.org/) 1.5+
  - Jupyter (see [here](http://jupyter.readthedocs.org/en/latest/install.html) for more details on installing jupyter)
  - [ZeroMQ](http://zeromq.org/) (2.2.X or 4.x)

Then:

1. Create a workspace and setup your `GOPATH`, see https://golang.org/doc/code.html#GOPATH

2. Install `goimports` if you haven't already:

  ```
  go get golang.org/x/tools/cmd/goimports
  ```

3. Get the kernel:
  - with ZeroMQ 2.2.x:

    ```
    go get github.com/gopherds/gophernotes
    ```
  
  - with ZeroMQ 4.x:

    ```
    go get -tags zmq_4_x github.com/gopherds/gophernotes
    ```

4. Create a directory for the new kernel config:

  ```
  mkdir -p ~/.local/share/jupyter/kernels/gophernotes
  ```
  
  Note, depending on which version of jupyter you are using and if you are using Anaconda, you may need to copy to `~/.ipython` rather than `~/.local/share`:
  
  ```
  mkdir ~/.ipython/kernels/gophernotes
  cp -r $GOPATH/src/github.com/gopherds/gophernotes/kernel/* ~/.ipython/kernels/gophernotes/
  ```
5. Copy the kernel config into the `~/.local/jupyter` directory:

  ```
  cp -r $GOPATH/src/github.com/gopherds/gophernotes/kernel/* ~/.local/share/jupyter/kernels/gophernotes
  ```
  
  Note, if you have the `JUPYTER_PATH` environmental variable set or if you are using an older version of Jupyter, you may need to copy this kernel config to another directory.  You can check which directories will be searched by executing:
  
  ```
  jupyter --data-dir
  ```

### Local, OSX

Make sure you have the following dependencies:

  - [Go](https://golang.org/) 1.5+
  - Jupyter (see [here](http://jupyter.readthedocs.org/en/latest/install.html) for more details on installing jupyter)
  - [ZeroMQ](http://zeromq.org/) (2.2.X or 4.x)

Then: 

1. Install goimports, if not already installed:

  ```
  go get golang.org/x/tools/cmd/goimports
  ```

2. Install gophernotes:
  - with ZeroMQ 2.2.x:

    ```
    go get github.com/gopherds/gophernotes
    ```
  
  - with ZeroMQ 4.x:

    ```
    go get -tags zmq_4_x github.com/gopherds/gophernotes
    ```
  
  - if you get this error:
  
    ```
    # pkg-config --cflags libzmq libzmq libzmq libzmq
    Package libzmq was not found in the pkg-config search path.
    Perhaps you should add the directory containing `libzmq.pc'
    to the PKG_CONFIG_PATH environment variable
    No package 'libzmq' found
    ```
    
    then:
    
    ```
    export PKG_CONFIG_PATH=/usr/local/Cellar/zeromq22/lib/pkgconfig/
    ```

3. Copy the kernel config:

  ```
  mkdir -p ~/Library/Jupyter/kernels/gophernotes
  cp -r $GOPATH/src/github.com/gopherds/gophernotes/kernel/* ~/Library/Jupyter/kernels/gophernotes
  ```
  
  Note, if you have the `JUPYTER_PATH` environmental variable set or if you are using an older version of Jupyter, you may need to copy this kernel config to another directory.  You can check which directories will be searched by executing:
  
  ```
  jupyter --data-dir
  ```

4. Update `~/Library/Jupyter/kernels/gophernotes/kernel.json` with the FULL PATH to your gophernotes binary (in $GOPATH/bin).  For example:

  ```
  {
      "argv": [
        "/Users/<your username>/go/bin/gophernotes",
        "{connection_file}"
        ],
      "display_name": "Go",
      "language": "go",
      "name": "go"
  }
  ```

### Local, Windows

Make sure you have the following dependencies:

  - [Go](https://golang.org/) 1.5+ with cgo enabled
  - MinGW toolchain, such as:
    - [MinGW-w64](https://sourceforge.net/projects/mingw-w64/), for 32 and 64 bit Windows
    - [MinGW Distro](https://nuwen.net/mingw.html), for 64 bit Windows only
  - Jupyter (see [here](http://jupyter.readthedocs.org/en/latest/install.html) for more details on installing jupyter)
  - [ZeroMQ](http://zeromq.org/) (2.2.X or 4.x); for convenience, pre-built binaries (v4.2.1) are included in the zmq-win directory

Then: 

1. Install goimports, if not already installed:

  ```
  go get golang.org/x/tools/cmd/goimports
  ```

2. Build and install gophernotes (using the pre-built binaries and `zmq-win\build.bat`):

    ```
    REM Download w/o building.
    go get -d github.com/gopherds/gophernotes
    cd %GOPATH%\src\github.com\gopherds\gophernotes\zmq-win
    
    REM Build x64 version.
    build.bat amd64
    move gophernotes.exe %GOPATH%\bin
    copy lib-amd64\libzmq.dll %GOPATH%\bin
    
    REM Build x86 version.
    build.bat 386
    move gophernotes.exe %GOPATH%\bin
    copy lib-386\libzmq.dll %GOPATH%\bin
    ```

3. Copy the kernel config:

  ```
  mkdir %APPDATA%\jupyter\kernels\gophernotes
  xcopy %GOPATH%\src\github.com\gopherds\gophernotes\kernel %APPDATA%\jupyter\kernels\gophernotes /s
  ```
  
  Note, if you have the `JUPYTER_PATH` environmental variable set or if you are using an older version of Jupyter, you may need to copy this kernel config to another directory.  You can check which directories will be searched by executing:
  
  ```
  jupyter --data-dir
  ```

4. Update `%APPDATA%\jupyter\kernels\gophernotes\kernel.json` with the FULL PATH to your gophernotes.exe (in %GOPATH%\bin), unless it's already on the PATH.  For example:

  ```
  {
      "argv": [
        "C:\\gopath\\bin\\gophernotes.exe",
        "{connection_file}"
        ],
      "display_name": "Go",
      "language": "go",
      "name": "go"
  }
  ```


## Getting Started

- If you completed one of the local installs above (i.e., not the Docker install), start the jupyter notebook:

  ```
  jupyter notebook
  ```

- If you have a docker install, point a browser at `http://localhost:8888/`.

- Select `Golang` from the `New` drop down menu.

- Have Fun!


## Troubleshooting

### gophernotes not found
- Depending on your environment, you may need to manually change the path to the `gophernotes` executable in `kernel/kernel.json` before copying it to `~/.local/share/jupyter/kernels/gophernotes`.  You can put the **full path** to the `gophernotes` executable here, and you shouldn't have any further issues.

### Print outputs string length followed by `<nil>`

Expressions in the top level scope of the notebook are printed out by default. This means that calling
```go
fmt.Println("Hello world")
```
the string `"Hello world"` is printed and then the two return values `11` and `nil` (bytes written, and error) are also printed.

Instead, if at the top level of the notebook, simply evaluate to a string such as using `fmt.Sprintf` or `"Hello " + "world"`

### "Kernel error" in a running notebook

```
Traceback (most recent call last):
  File "/usr/local/lib/python2.7/site-packages/notebook/base/handlers.py", line 458, in wrapper
    result = yield gen.maybe_future(method(self, *args, **kwargs))
  File "/usr/local/lib/python2.7/site-packages/tornado/gen.py", line 1008, in run
    value = future.result()
  ...
  File "/usr/local/Cellar/python/2.7.11/Frameworks/Python.framework/Versions/2.7/lib/python2.7/subprocess.py", line 1335, in _execute_child
    raise child_exception
OSError: [Errno 2] No such file or directory
```

Stop jupyter, if it's already running.

Add a symlink to `/go/bin/gophernotes` from your path to the gophernotes executable. If you followed the instructions above, this will be:

```
sudo ln -s $HOME/go/bin/gophernotes /go/bin/gophernotes
```

Restart jupyter, and you should now be up and running.


## Custom Commands
Some of the custom commands from the [gore](https://github.com/motemen/gore) REPL have carried over to `gophernotes`.  Note, in particular, the syntax for importing packages:

```
:print                  Show current source (currently prints to the terminal where the notebook server is running)
:write [<filename>]     Write out current source to file
:help                   List commands
:containerize           Build a Docker image that executes the compiled Go code (must have Docker installed)
```

## Licenses

`gophernotes` was created by [Daniel Whitenack](http://www.datadan.io/), and is licensed under an [MIT-style License](LICENSE.md).

The Golang Gopher image adapted for the gophernotes logo was created by [Takuya Ueda](http://u.hinoichi.net) and is licensed under the Creative Commons 3.0 Attributions license.
