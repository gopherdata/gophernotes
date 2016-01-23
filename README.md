![alt tag](https://raw.githubusercontent.com/gophergala2016/gophernotes/master/files/gophernotes2.jpg)

# gophernotes - Go in Notebooks

`gophernotes` is a Go kernel for [Jupyter](http://jupyter.org/) and Apache Zeppelin notebooks.  Finally, we can have a reliable and maintained way to use Go interactively and code in the browser.  Go forth and do data science, or anything else interesting, with go notebooks!

This project came out of the [Gopher Gala](http://gophergala.com/) 2016.  It is based on a REPL called [gore](https://github.com/motemen/gore) and on a, no longer maintained and self-described limited, ipython kernel call [iGo](https://github.com/takluyver/igo).

## Screenshots/Examples:

### Simple interactive use in Jupyter

![alt tag](https://raw.githubusercontent.com/gophergala2016/gophernotes/master/files/screencast.gif)

### Telling a story and performing pattern recognition with Markdown and Golang

![alt tag](https://raw.githubusercontent.com/gophergala2016/gophernotes/master/files/pr-screenshot.png)

### Example Notebooks:
- [Simple Printing and Channels Example](https://github.com/gophergala2016/gophernotes/blob/master/examples/Simple-Example.ipynb)
- [Pattern Recognition with Golearn](https://github.com/gophergala2016/gophernotes/blob/master/examples/Pattern-Recognition.ipynb)

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

## Gopher Gala Status:

- ~~Rewrite of some of the gore functionality to process messages from the notebook.~~
- ~~Implement more intuitive error handling in the notebook, test the suite of functionality in the Jupyter notebook.~~
- ~~Refactor, loggings~~
- Examples, screenshots, installation
- Troubleshoot errors, unit testing
