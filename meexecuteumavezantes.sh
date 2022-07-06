set -e

export MSYS_NO_PATHCONV=1
starttime=$(date +%s)
CC_SRC_LANGUAGE=${1:-"go"}
CC_SRC_LANGUAGE=`echo "$CC_SRC_LANGUAGE" | tr [:upper:] [:lower:]`

CC_SRC_PATH="/chaincode/"

rm -rf wallet/*
rm -rf gateway/*

pushd ../test-network
./testeComInsercaoDeAtivos.sh
popd


pushd web
go run notificaapi.go
popd
