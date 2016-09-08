FROM continuumio/anaconda

# dependencies
RUN apt-get update
RUN apt-get install -y pkg-config libzmq-dev build-essential

# set up golang
RUN wget https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.5.linux-amd64.tar.gz
ENV PATH /usr/local/go/bin:$PATH
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# install gophernotes
RUN go get golang.org/x/tools/cmd/goimports
RUN go get github.com/gopherds/gophernotes
RUN mkdir -p ~/.ipython/kernels/gophernotes
RUN cp -r $GOPATH/src/github.com/gopherds/gophernotes/kernel/* ~/.ipython/kernels/gophernotes

EXPOSE 8888
CMD ["jupyter", "notebook"]
