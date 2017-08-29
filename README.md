# Sweetbridge Oracles

This repository contains blockchain Oracles used by Sweetbridge project.

## Setting up

1. Install `go >= 1.7` and `node.js` and `make`. Make sure that `go` executable is in your `PATH`
2. Make sure `$GOPATH/bin` (`go env GOPATH`) is part of your `PATH`.
3. Make sure you have an access to the oracles repository.
4. Clone this repository into your `GOPATH` (`go env`). Install dependencies:

	git clone git@bitbucket.org:sweetbridge/oracles.git $GOPATH/src/bitbucket.org/sweetbridge/oracles
	cd ~/go/src/bitbucket.org/sweetbridge/oracles
	make install-deps

4. (alternative). Use `go get` (this will work only if the repo is open) and hack the git config a bit: https://gist.github.com/shurcooL/6927554

    go get bitbucket.org/sweetbridge/oracles/cmd/helloworld


### for development:
This is not required for building application, but required for proper development.

	make setup-dev

### Building

	make build

The executable are in the project root directory (oracles)

If you have issues with `zb` then just use plain `go` command:

	cd cmd/<app name>
	go build


## Applictions

### helloworld

The helloworld application is an oracle which has two commands:

* `check-user <user address>`  -- will verify if the user is registered in the Root contract
* `register` -- will register user specified by the `-pk` parameter in the Root contract

Usage:

	make build
	./helloworld -h

Example (address doesn't have to have "0x" prefix):

	make build && SB_ETH_ROOT="f5af14ece3d5ea1022eb2d5f6cad8afc84f98895" ./helloworld -pk "/home/robert/.../account.json" -pwd password -host "localhost:8545" check-user 0x0f21f6fb13310ac0e17205840a91da93119efbec


## Scripts

### abigen

`scripts/abigen.js` is used to generate Go interfaces based on JSON interface file produced by Truffle. Whenever new contracts are provided you should regenerate Go interfaces

    node scripts/abigen.js <path to the truffle contract builds>

Usually you should just use make:

    make abigen-backsage
