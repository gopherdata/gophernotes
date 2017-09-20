![alt tag](files/gophernotes-logo.png)

# gophernotes - Use Go in Jupyter notebooks and nteract

`gophernotes` is a Go kernel for [Jupyter](http://jupyter.org/) notebooks and [nteract](https://nteract.io/).  It lets you use Go interactively in a browser-based notebook or desktop app.  Use `gophernotes` to create and share documents that contain live Go code, equations, visualizations and explanatory text.  These notebooks, with the live Go code, can then be shared with others via email, Dropbox, GitHub and the [Jupyter Notebook Viewer](http://nbviewer.jupyter.org/). Go forth and do data science, or anything else interesting, with Go notebooks!

**Acknowledgements** - This project utilizes a Go interpreter called [gomacro](https://github.com/cosmos72/gomacro) under the hood to evaluate Go code interactively. The gophernotes logo was designed by the brilliant [Marcus Olsson](https://github.com/marcusolsson) and was inspired by Renee French's original Go Gopher design. 

## Examples

### Jupyter Notebook:  

![](files/jupyter.gif)  

### nteract:

![](files/nteract.gif)

### Example Notebooks (dowload and run them locally, follow the links to view in Github, or use the [Jupyter Notebook Viewer](http://nbviewer.jupyter.org/)):
- [Worker Pools](examples/Worker_Pools.ipynb)
- [Matrix Operations](examples/Matrix_Operations.ipynb)
- [Facial Recognition](examples/Facial_Recognition_MachineBox.ipynb)

## Installation

### Prerequisites

- [Go 1.9+](https://golang.org/doc/install) - including GOPATH/bin added to your PATH (i.e., you can run Go binaries that you `go install`).
- [Jupyter Notebook](http://jupyter.readthedocs.io/en/latest/install.html) or [nteract](https://nteract.io/desktop)
- [ZeroMQ 4.X.X](http://zeromq.org/intro:get-the-software) - for convenience, pre-built Windows binaries (v4.2.1) are included in the zmq-win directory.

### Linux

```sh
$ go install github.com/gopherdata/gophernotes
$ mkdir -p ~/.local/share/jupyter/kernels/gophernotes
$ cp $GOPATH/src/github.com/gopherdata/gophernotes/kernel/* ~/.local/share/jupyter/kernels/gophernotes  
```

To confirm that the `gophernotes` binary is installed and in your PATH, you should see the following when running `gophernotes` directly:

```sh
$ gophernotes
2017/09/20 10:33:12 Need a command line argument specifying the connection file.
```

**Note** - if you have the `JUPYTER_PATH` environmental variable set or if you are using an older version of Jupyter, you may need to copy this kernel config to another directory.  You can check which directories will be searched by executing:
  
```sh
$ jupyter --data-dir
```

### OS X/macOS

### Windows

Make sure you have the MinGW toolchain:

- [MinGW-w64](https://sourceforge.net/projects/mingw-w64/), for 32 and 64 bit Windows
- [MinGW Distro](https://nuwen.net/mingw.html), for 64 bit Windows only

Then:

1. build and install gophernotes (using the pre-built binaries and `zmq-win\build.bat`):

    ```
    REM Download w/o building.
    go get -d github.com/gopherdata/gophernotes
    cd %GOPATH%\src\github.com\gopherdata\gophernotes\zmq-win
    
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
    xcopy %GOPATH%\src\github.com\gopherdata\gophernotes\kernel %APPDATA%\jupyter\kernels\gophernotes /s
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

### Docker

You can try out or run Jupyter + gophernotes without installing anything using Docker. To run a Go notebook that only needs things from the standard library, run: 

```
$ docker run -it -p 8888:8888 gopherdata/gophernotes
```

Or to run a Go notebook with access to common Go data science packages (gonum, gota, golearn, etc.), run:

```
$ docker run -it -p 8888:8888 gopherdata/gophernotes-ds
```

In either case, running this command should output a link that you can follow to access Jupyter in a browser.

## Getting Started

### Jupyter

- If you completed one of the local installs above (i.e., not the Docker install), start the jupyter notebook server:

  ```
  jupyter notebook
  ```

- Select `Go` from the `New` drop down menu.

- Have fun!

### nteract

- Launch nteract.

- From the nteract menu select Language -> Go.

- Have fun!

## Limitations

gophernotes uses [gomacro](https://github.com/cosmos72/gomacro) under the hood to evaluate Go code interactively. You can evaluate most any Go code with gomacro, but there are some limitation, which are discussed in further detail [here](https://github.com/cosmos72/gomacro#current-status).  Most noteably, gophernotes does NOT support:

- unexported struct fields
- interfaces - They can be declared, but nothing more: there is no way to implement them or call their methods
- extracting methods from types - For example time.Duration.String should return a func(time.Duration) string but currently gives an error. Instead extracting methods from objects is supported: time.Duration(1s).String correctly returns a func() string
- goto
- named return values
- named imports

## Troubleshooting

### gophernotes not found

Depending on your environment, you may need to manually change the path to the `gophernotes` executable in `kernel/kernel.json` before copying it to `~/.local/share/jupyter/kernels/gophernotes`.  You can put the **full path** to the `gophernotes` executable here, and you shouldn't have any further issues.

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
