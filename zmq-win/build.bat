@echo off
setlocal

if not [%1]==[386]; if not [%1]==[amd64] (
    echo Usage: %~n0 386^|amd64
    goto QUIT
)
set target=%1

REM get the bat's directory path, and replace \ with /
set mydir=%~dp0
set mydir=%mydir:\=/%

set CGO_CFLAGS=-I %mydir%include
set CGO_LDFLAGS=-L %mydir%lib-%target% -l zmq

go build -tags zmq_4_x github.com/gopherdata/gophernotes

:QUIT
endlocal
echo on
