# payment-forwarder

The payment-forwarder service creates a `BTC` or `ETH` alias for a specific address. 
This allows the creation of individual contribution addresses for each user in an ICO 
while centralising the contributions in a common treasury. 

## Usage

Run command with `-h` to check the usage:

	./bin/payment-forwarder -h

The program requires path to the contract schema directory (generated by truffle).

usage: `./bin/payment-forwarder <parameters> :`

PARAMETERS:

*  **-bcy-key string** BlockCypher API Token [required]
*  **-bcy-net** string BlockCypher network (main or test) [required]
*  **-contracts** string path to contract schemas directory [required]
*  **-eth-network** string blockchain network name, must be specified in eth-network-file file [required]
*  **-eth-network-file** string path to the file with network configuration [required]
*  **-pk-file pk-hex** path to the private key json file [required if pk-hex not specified]
*  **-pk-hex** private key - takes precedence over `pk-file` + `pk-pwd` [required]
*  **-pk-file** an encrypted PK file [if provided then `-pk-pwd` is required too]
*  **-pk-pwd** key file password [required if `pk-hex` not specified]
*  **-port** string The HTTP listening port (default "8000")
*  **-rollbar** string rollbar token [required in production env]
*  **-tx-timeout** int how many seconds should the daemon wait for the transaction receipt? (default 600)

## Docker

