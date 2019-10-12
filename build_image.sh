#!/bin/sh


./go_build_linux_amd64.sh ./build/dav-proxy

docker build -t 10.10.15.51/fc/dav-proxy:ca49c48 .