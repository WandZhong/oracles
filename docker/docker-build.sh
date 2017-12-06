#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ORACLE=$1

mkdir -p $DIR/tmp/$ORACLE
cp $DIR/../bin/$ORACLE $DIR/tmp/$ORACLE/
cp $DIR/$ORACLE-dockerfile $DIR/tmp/$ORACLE/Dockerfile
docker build -t $ORACLE $DIR/tmp/$ORACLE
rm -Rf $DIR/tmp
