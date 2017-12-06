#!/bin/sh

exec ./payment-forwarder -bcy-key="$BCYKEY" \
                         -bcy-net="$BCYNET" \
                         -contracts="$CONTRACTSPATH" \
                         -eth-network="$ETHNETWORK" \
                         -eth-network-file="$ETHNETWORKCONFIG" \
                         -pk-file="$PKFILEPATH" \
                         -pk-pwd="$PKFILEPWD" \
                         -pk-hex="$PKHEX" \
                         -port="$PORT" \
                         -rollbar="$ROLLBAR" \
                         -tx-timeout="$TXTIMEOUT"
