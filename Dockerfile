FROM alpine:3.15
MAINTAINER dwhitena

# Add gophernotes
ADD . /go/src/github.com/gopherdata/gophernotes/

# Install Jupyter and gophernotes.
RUN set -x \
    # install python and dependencies
    && apk update \
    && apk --no-cache \
        --repository http://dl-4.alpinelinux.org/alpine/v3.15/community \
        --repository http://dl-4.alpinelinux.org/alpine/v3.15/main \
        --arch=x86_64 add \
        ca-certificates \
        g++ \
        gcc \
        git \
        libffi-dev \
        pkgconfig \
        python3 python3-dev \
        py3-pip \
        py3-pyzmq \
        mercurial \
        mesa-dev \
        musl-dev \
        su-exec \
        zeromq-dev \
    && pip3 install --upgrade pip==21.3.1 \
    && ln -s /usr/bin/python3.9 /usr/bin/python \
    ## install Go
    && apk --update-cache \
        --arch=x86_64 add \
        go \
    ## jupyter notebook
    && ln -s /usr/include/locale.h /usr/include/xlocale.h \
    && pip3 install jupyter notebook pyzmq tornado ipykernel \
    ## install gophernotes
    && cd /go/src/github.com/gopherdata/gophernotes \
    && GOPATH=/go GO111MODULE=on go install . \
    && cp /go/bin/gophernotes /usr/local/bin/ \
    && mkdir -p ~/.local/share/jupyter/kernels/gophernotes \
    && cp -r ./kernel/* ~/.local/share/jupyter/kernels/gophernotes \
    && cd - \
    ## clean
    && find /usr/lib/python3.9 -name __pycache__ | xargs rm -r \
    && rm -rf \
        /root/.[acpw]* \
        ipaexg00301* \
    && rm -rf /var/cache/apk/*

# Set GOPATH.
ENV GOPATH /go

EXPOSE 8888
CMD [ "jupyter", "notebook", "--no-browser", "--allow-root", "--ip=0.0.0.0" ]
