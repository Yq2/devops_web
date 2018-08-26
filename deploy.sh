#! /bin/sh

kill -9 $(pgrep webserver)

cd /opt/go/gopath/src/github.com/Yq2/

git clone https://github.com/Yq2/devops_web.git

cd devops_web/

cd ./webserver/

rm -rf main
rm -rf webserver

go build main.go

chmod 777 main

./main &