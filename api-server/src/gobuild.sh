#!/bin/bash
BASEDIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
PKGDIR=$BASEDIR"/../pkg/*"

APPCONFPATH=$BASEDIR"/conf/app.conf"
rm -f $APPCONFPATH
cp $BASEDIR"/conf/"$CHINARUN_API_SERVER_MODE".conf" $APPCONFPATH

parentdir="$(dirname "$BASEDIR")"

if [[ $GOPATH == *$parentdir* ]]
then
	echo $GOPATH
else
	export GOPATH=$GOPATH":"$parentdir
	echo $GOPATH
fi


if [[ $CHINARUN_API_SERVER_MODE == "production" ]] 
then
    go build -ldflags "-w" -o chinarun_api
elif [[ $CHINARUN_API_SERVER_MODE  == "testing" ]]
then
    go build  -o chinarun_api_`date +%Y-%m-%d`
else 
    go build -o chinarun_api
fi

#for debug
#go build -gcflags "-N -l" -o chinarun_api_debug
