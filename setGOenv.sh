#!/bin/bash

#export GOPATH=$HOME/go
export GOPATH=$HOME/.go

#export GOROOT=$HOME/go1.X
export GOROOT=/usr/local/opt/go/libexec
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
