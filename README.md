![alt tag](https://raw.githubusercontent.com/gophergala2016/gophernotes/master/files/gophernotes2.jpg)

# gophernotes - Go in Notebooks

`gophernotes` is a Go kernel for [Jupyter](http://jupyter.org/) and Apache Zeppelin notebooks.  Finally, we can have a reliable and maintained way to use Go interactively and code in the browser.  Go forth and do data science, or anything else interesting, with go notebooks!

This project came out of the [Gopher Gala](http://gophergala.com/) 2016.  It is based on a REPL called [gore](https://github.com/motemen/gore) and on a, no longer maintained and self-described limited, ipython kernel call [iGo](https://github.com/takluyver/igo).

## Screenshots/Examples:

![alt tag](https://raw.githubusercontent.com/gophergala2016/gophernotes/master/files/screencast.gif)

Examples:
- [Simple Printing and Channels Example](https://github.com/gophergala2016/gophernotes/blob/master/examples/Simple-Example.ipynb)

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

## Gopher Gala Status:

- ~~Rewrite of some of the gore functionality to process messages from the notebook.~~
- ~~Implement more intuitive error handling in the notebook, test the suite of functionality in the Jupyter notebook.~~
- ~~Refactor, loggings~~
- Examples, screenshots, installation
- Troubleshoot errors, unit testing
