# helloworld

The helloworld application is an oracle which is a Proof of concept demonstration to interact with the Sweetbridge smart contracts.

The application has two commands:

* `check-user <user address>`  -- will verify if the user is registered in the Root contract
* `register` -- will register user specified by the `-pk` parameter in the Root contract

Usage:

	./bin/helloworld -h

Example (address doesn't have to have "0x" prefix):

	make build && SB_ETH_ROOT="f5af14ece3d5ea1022eb2d5f6cad8afc84f98895" ./helloworld -pk "/home/robert/.../account.json" -pwd password -host "localhost:8545" check-user 0x0f21f6fb13310ac0e17205840a91da93119efbec
