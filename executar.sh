#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error
set -e

# don't rewrite paths for 
export MSYS_NO_PATHCONV=1
starttime=$(date +%s)
CC_SRC_LANGUAGE=${1:-"go"}
CC_SRC_LANGUAGE=`echo "$CC_SRC_LANGUAGE" | tr [:upper:] [:lower:]`

CC_SRC_PATH="../chaincode/"

# if [ "$CC_SRC_LANGUAGE" = "go" -o "$CC_SRC_LANGUAGE" = "golang" ] ; then
# 	CC_SRC_PATH="../chaincode/fabcar/go/"


# else
# 	echo The chaincode language ${CC_SRC_LANGUAGE} is not supported by this script
# 	echo Supported chaincode languages are: go, java, javascript, and typescript
# 	exit 1
# fi

# clean out any old identites in the wallets
#limpar as antigas wallets

rm -rf wallet/*
rm -rf gateway/*

# launch network; create channel and join peer to channel
pushd ../../test-network
./inicializarRede.sh
./ativandoChaincode.sh
#./meuNetwork.sh deployCC -ccn contratointeligente  -ccv 1 -cci initLedger -ccl ${CC_SRC_LANGUAGE} -ccp ${CC_SRC_PATH}
popd

go run contratointeligente.go
#Executando minha rede 
#./inicializarRede.sh

#cat <<EOF
#
#Total setup execution time : $(($(date +%s) - starttime)) secs ...
#
#Next, use the FabCar applications to interact with the deployed FabCar contract.
#The FabCar applications are available in multiple programming languages.
#Follow the instructions for the programming language of your choice:
#
#
#
#Go:
#
#  Start by changing into the "go" directory:
#    cd go
#
#  Then, install dependencies and run the test using:
#    go run fabcar.go
#
#  The test will invoke the sample client app which perform the following:
#    - Import user credentials into the wallet (if they don't already exist there)
#    - Submit a transaction to create a new car
#    - Evaluate a transaction (query) to return details of this car
#    - Submit a transaction to change the owner of this car
#    - Evaluate a transaction (query) to return the updated details of this car
#
#EOF