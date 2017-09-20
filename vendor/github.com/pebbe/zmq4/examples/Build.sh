#!/bin/sh

go get github.com/pborman/uuid

for i in bstar mdapi flcliapi kvsimple kvmsg clone intface
do
    go install github.com/pebbe/zmq4/examples/$i
done

cd `dirname $0`

goos=`go env GOOS`
gobin=`go env GOBIN`
if [ "$gobin" = "" ]
then
    gobin=`go env GOPATH`
    if [ "$gobin" = "" ]
    then
	gobin=`go env GOROOT`
    fi
    gobin=`echo $gobin | sed -e 's/:.*//'`
    gobin=$gobin/bin
fi

dir=$gobin/zmq4-examples

echo Installing examples in $dir

mkdir -p $dir

for i in *.sh
do
    if [ $i != Build.sh ]
    then
	cp -u $i $dir
    fi
done

src=''
for i in *.go
do
    if [ $i = interrupt.go ]
    then
	if [ $goos = windows -o $goos = plan9 ]
	then
	    continue
	fi
    fi
    bin=$dir/`basename $i .go`
    if [ ! -f $bin -o $i -nt $bin ]
    then
	src="$src $i"
    fi
done

libs=`pkg-config --libs-only-L libzmq`
if [ "$libs" = "" ]
then
    for i in $src
    do
	go build -o $dir/`basename $i .go` $i
    done
else
    libs="-r `echo $libs | sed -e 's/-L//; s/  *-L/:/g'`"
    for i in $src
    do
	go build -ldflags="$libs" -o $dir/`basename $i .go` $i
    done
fi
