FROM alpine:3.5
MAINTAINER dwhitena

# Add gophernotes
ADD . /go/src/github.com/gopherdata/gophernotes/

# Install Jupyter and gophernotes.
RUN set -x \
    # install python and dependencies
    && apk update \
    && apk --no-cache \
        --repository http://dl-4.alpinelinux.org/alpine/v3.7/community \
        --repository http://dl-4.alpinelinux.org/alpine/v3.7/main \
        --arch=x86_64 add \
        ca-certificates \
        python3 \
        su-exec \
        gcc \
        git \
        py3-zmq \
        pkgconfig \
        zeromq-dev \
        musl-dev \
    && pip3 install --upgrade pip==9.0.3 \
    && ln -s /usr/bin/python3.6 /usr/bin/python \
    ## install Go
    && apk --update-cache --allow-untrusted \
        --repository http://dl-4.alpinelinux.org/alpine/edge/community \
        --arch=x86_64 add \
        go \
    ## jupyter notebook
    && ln -s /usr/include/locale.h /usr/include/xlocale.h \
    ### fix pyzmq to v16.0.2 as that is what is distributed with py3-zmq
    ### pin down the tornado and ipykernel to compatible versions
    && pip3 install jupyter notebook pyzmq==16.0.2 tornado==4.5.3 ipykernel==4.8.1 \
    ## install gophernotes
    && cd /go/src/github.com/gopherdata/gophernotes \
    && GOPATH=/go GO111MODULE=on go install . \
    && cp /go/bin/gophernotes /usr/local/bin/ \
    && mkdir -p ~/.local/share/jupyter/kernels/gophernotes \
    && cp -r ./kernel/* ~/.local/share/jupyter/kernels/gophernotes \
    && cd - \
    ## clean
    && find /usr/lib/python3.6 -name __pycache__ | xargs rm -r \
    && rm -rf \
        /root/.[acpw]* \
        ipaexg00301* \
    && rm -rf /var/cache/apk/*

# Set GOPATH.
ENV GOPATH /go

EXPOSE 8888
CMD [ "jupyter", "notebook", "--no-browser", "--allow-root", "--ip=0.0.0.0" ]
