#!/bin/bash

#Check that the environment variable has been set correctly
if [ -z "$SECRETS_BUCKET_NAME" ]; then
  echo >&2 'error: missing SECRETS_BUCKET_NAME environment variable'
  exit 1
fi

# Load the S3 secrets file contents into the environment variables
eval $(aws s3 cp s3://${SECRETS_BUCKET_NAME}/payment-forwarder.txt - | sed 's/^/export /')
aws s3 cp s3://${SECRETS_BUCKET_NAME}/payment-forwarder.txt /

#Copy networks file
aws s3 cp s3://${SECRETS_BUCKET_NAME}/eth-networks.json .
ETHNETWORKCONFIG=eth-networks.json

#Copy contracts directory
aws s3 sync s3://${SECRETS_BUCKET_NAME}/contract-deployments /contracts
CONTRACTSPATH=contracts



ls /
ls /contracts

echo  ./payment-forwarder -bcy-key="$BCYKEY" \
                         -bcy-net="$BCYNET" \
                         -contracts="$CONTRACTSPATH" \
                         -eth-network="$ETHNETWORK" \
                         -eth-network-file="$ETHNETWORKCONFIG" \
                         -pk-hex="$PKHEX" \
                         -port="$PORT" \
			 -btc-pool="$BTCPOOL"\
			 -eth-pool="$ETCPOOL"\
                         -tx-timeout="$TXTIMEOUT"

tail -f /bin/true
exec ./payment-forwarder -bcy-key="$BCYKEY" \
                         -bcy-net="$BCYNET" \
                         -contracts="$CONTRACTSPATH" \
                         -eth-network="$ETHNETWORK" \
                         -eth-network-file="$ETHNETWORKCONFIG" \
                         -pk-hex="$PKHEX" \
			 -btc-pool="$BTCPOOL"\
			 -eth-pool="$ETCPOOL"\
                         -port="$PORT" \
                         -tx-timeout="$TXTIMEOUT"

