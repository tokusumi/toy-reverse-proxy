# !bin/bash

BASEDIR=$(dirname $0)

cd $BASEDIR/api/vanila-api && go run main.go &
cd $BASEDIR/api/echo-api && go run main.go &
cd $BASEDIR/api/gin-api && go run main.go &

wait