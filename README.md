![alt tag](https://raw.githubusercontent.com/gophergala2016/gophernotes/master/files/gophernotes2.jpg)

# gophernotes - Go in Notebooks

`gophernotes` is a Go kernel for [Jupyter](http://jupyter.org/) notebooks.  Finally, we can have a reliable and maintained way to use Go interactively and code in the browser.  Use `gophernotes` to create and share documents that contain live Go code, equations, visualizations and explanatory text.  These notebooks, with the live Go code, can then be shared with others via email, Dropbox, GitHub and the [Jupyter Notebook Viewer](http://nbviewer.jupyter.org/).  Go forth and do data science, or anything else interesting, with go notebooks!

This project came out of the [Gopher Gala](http://gophergala.com/) 2016.  It is based on a REPL called [gore](https://github.com/motemen/gore) and on a, no longer maintained and self-described limited, ipython kernel call [iGo](https://github.com/takluyver/igo).

## Screenshots/Examples

### Simple interactive use:

![alt tag](https://raw.githubusercontent.com/gophergala2016/gophernotes/master/files/screencast.gif)

### Story telling and pattern recognition with Markdown and Golang:

![alt tag](https://raw.githubusercontent.com/gophergala2016/gophernotes/master/files/pr-screenshot.png)

### Example Notebooks (they render quite nice in Github, or you can use the [Jupyter Notebook Viewer](http://nbviewer.jupyter.org/)):
- [Simple Printing and Channels](https://github.com/gophergala2016/gophernotes/blob/master/examples/Simple-Example.ipynb)
- [Pattern Recognition with Golearn](https://github.com/gophergala2016/gophernotes/blob/master/examples/Pattern-Recognition.ipynb)
- [Feed Forward, Recurrent Neural Nets](https://github.com/gophergala2016/gophernotes/blob/master/examples/Feed-Forward-Recurrent-NN.ipynb)
- [Time Parsing, Formatting](https://github.com/gophergala2016/gophernotes/blob/master/examples/Time-Formatting-Parsing.ipynb)
- [Stateful Goroutines](https://github.com/gophergala2016/gophernotes/blob/master/examples/Stateful-Goroutines.ipynb)

## Jupyter Installation/Usage

- Make sure jupyter notebook is installed.  See [here](http://jupyter.readthedocs.org/en/latest/install.html) for more details.
- Get the kernel:

  ```
  go get github.com/gophergala2016/gophernotes
  ```

- Create a directory for the new kernel config:

  ```
  mkdir -p ~/.ipython/kernels/gophernotes
  ```

- Copy the kernel config into the `.ipython` directory:

  ```
  cp -r $GOPATH/src/github.com/gophergala2016/gophernotes/kernel/* ~/.ipython/kernels/gophernotes
  ```

- Start the jupyter notebook:

  ```
  jupyter notebook
  ```

- Select `Golang` from the `New` drop down menu.
- Have Fun!

Possible Issues:
- Depending on your environment, you may need to manually change the path to the `gophernotes` executable in `kernel/kernel.json` before copying it to `~/.ipython/kernels/gophernotes`.  You can put the **full path** to the `gophernotes` executable here, and you shouldn't have any further issues.

## Custom commands
Some of the custom commands from the [gore](https://github.com/motemen/gore) REPL have carried over to `gophernotes`.  Note in particular the syntax for importing packages:

```
:import <package path>  Import package
:print                  Show current source
:write [<filename>]     Write out current source to file
:help                   List commands
```

## Gopher Gala Status

- ~~Rewrite of some of the gore functionality to process messages from the notebook.~~
- ~~Implement more intuitive error handling in the notebook, test the suite of functionality in the Jupyter notebook.~~
- ~~Refactor, loggings~~
- ~~Examples, screenshots, installation~~
- ~~Troubleshoot errors, unit testing~~
- Godep
- Dockerfile?
