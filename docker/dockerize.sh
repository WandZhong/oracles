#!/usr/bin/env bash

# this script dir:
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ORACLE=$1

mkdir -p $DIR/tmp/$ORACLE
cp $DIR/eth-networks.json $DIR/tmp/$ORACLE/
cp $DIR/${ORACLE}* $DIR/tmp/$ORACLE/
cp $DIR/../bin/$ORACLE $DIR/tmp/$ORACLE/
cp -r $DIR/../submodules/contract-deployments/backstage $DIR/tmp/$ORACLE/contract-schemas
cp $DIR/Dockerfile-$ORACLE $DIR/tmp/$ORACLE/Dockerfile
docker build -t $ORACLE $DIR/tmp/$ORACLE
rm -Rf $DIR/tmp
