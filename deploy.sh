#! /bin/sh

kill -9 $(pgrep webserver)

git pull https://github.com/Yq2/devops_web.git

cd ./webserver/
./webserver &