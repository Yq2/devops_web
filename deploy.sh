#! /bin/sh

kill -9 $(pgrep webserver)

cd /opt/go/gopath/src/github.com/Yq2/devops_web/

git pull https://github.com/Yq2/devops_web.git

cd ./webserver/

go build main.go

chmod 777 webserver

./webserver &