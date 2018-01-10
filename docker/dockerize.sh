#!/usr/bin/env bash

# this script creates a limited context for the docker builder
# to only include all required dependencies.
# To run the docker build at the end you have to pass 'build' as
# the second argument
#
# usage:  dockerize.sh <app-name> [build]
#

# this script dir:
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ORACLE=$1

echo "building $ORACLE,    DIR = $DIR"


mkdir -p $DIR/tmp/$ORACLE
cp $DIR/../config/eth-networks.json $DIR/tmp/$ORACLE/
cp $DIR/${ORACLE}* $DIR/tmp/$ORACLE/
cp $DIR/../bin/$ORACLE $DIR/tmp/$ORACLE/
cp -r $DIR/../submodules/contract-deployments $DIR/tmp/$ORACLE/contract-deployments
cp $DIR/Dockerfile-$ORACLE $DIR/tmp/$ORACLE/Dockerfile

if [ "$2" = build ]; then
	docker build -t $ORACLE $DIR/tmp/$ORACLE
	rm -Rf $DIR/tmp
fi

echo "Dockerize complete"
