FROM alpine:3.5
MAINTAINER dwhitena

# Add gophernotes
ADD . /go/src/github.com/gopherdata/gophernotes/

# Install Jupyter and gophernotes.
RUN set -x \
    # install python and dependencies
    && apk update \
    && apk --no-cache add \
        ca-certificates \
        python3 \
        su-exec \
        gcc \
        git \
        py3-zmq \
        pkgconfig \ 
        zeromq-dev \
        musl-dev \
    && pip3 install --upgrade pip \
    && ln -s /usr/bin/python3.5 /usr/bin/python \
    ## install Go
    && apk --update-cache --allow-untrusted \
        --repository http://dl-4.alpinelinux.org/alpine/edge/community \
        --arch=x86_64 add \
        go \
    ## jupyter notebook 
    && ln -s /usr/include/locale.h /usr/include/xlocale.h \
    && pip3 install jupyter notebook \
    ## install gophernotes
    && GOPATH=/go go install github.com/gopherdata/gophernotes \
    && cp /go/bin/gophernotes /usr/local/bin/ \
    && mkdir -p ~/.local/share/jupyter/kernels/gophernotes \
    && cp -r /go/src/github.com/gopherdata/gophernotes/kernel/* ~/.local/share/jupyter/kernels/gophernotes \
    ## clean
    && find /usr/lib/python3.5 -name __pycache__ | xargs rm -r \
    && rm -rf \
        /root/.[acpw]* \
        ipaexg00301* \
    && rm -rf /var/cache/apk/*

EXPOSE 8888
CMD [ "jupyter", "notebook", "--no-browser", "--allow-root", "--ip=0.0.0.0" ]
