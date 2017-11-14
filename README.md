# Sweetbridge Oracles

This repository contains blockchain Oracles used by Sweetbridge project.

## Setting up

Unless you are using docker environment to build the applications (explained below) you have to set-up the following environment:

1. Install `go >= 1.7` and `node.js` and `make`. Make sure that `go` executable is in your `PATH`
2. Make sure `$GOPATH/bin` (`go env GOPATH`) is part of your `PATH`.
3. Make sure you have an access to the oracles repository.
4. Clone this repository into your `GOPATH` (`go env`). Install dependencies:

	git clone git@bitbucket.org:sweetbridge/oracles.git $GOPATH/src/bitbucket.org/sweetbridge/oracles
	cd ~/go/src/bitbucket.org/sweetbridge/oracles
	make install-deps

4. (alternative). Use `go get` (this will work only if the repo is open) and hack the git config a bit: https://gist.github.com/shurcooL/6927554

    go get bitbucket.org/sweetbridge/oracles/cmd/<oracle-name>

Please refer to `setup-quickstart.sh` if you want to install it on debian / ubuntu.

### Development:
This is not required for building application, but required for proper development.

	make setup-dev

### Building

Whenever you want to sync dependencies (so if you don't follow the development process) you have to make sure that everything is synced. However if you are sure that you don't need it, or you are just repeating builds you can skip it.

	# skip the first step if you don't need to sync dependencies
	make install-deps
	make build

The executables are stored in the `bin` directory.

Furthermore there are handy application specific jobs (to build only a single application). To do that just look at one of `make build-` jobs.

#### Release, CI

Again, for CI, and release to make sure that dependencies are always synced. The simple spell to build for release is:

	make install-deps build


#### Building using Docker

We provide consistent environment to build the applications using Docker containers.
Everything is wrapped in the make jobs.
Firstly you need to build the image. This is suppose to be done only once:

	make docker-mk-builder

Next, whenever you want to build the applications:

	make docker-run-builder


## Applications

All applications are listed packaged in the `cmd/` sub-directories. Each package contains the README.md file with the application description.

All applications are compiled into self-contained binaries (no external dependencies needed). Each application support the help argument to display the usage description:

	./bin/<app_name> -h


### Common ethereum parameters

Oracles share the following, ethereum-specific obligatory parameters:

* `-pk-hex` -- the plaintext hex-encoded private key.
* `-pk-file` + `-pk-pwd` -- the JSON encrypted private key. Required if the `-pk-hex` is not specified.
* `-eth-network` -- the network name which oracles is connecting to (use `-h` to list all network names).
* `-contracts` -- the path to the directory containing contract schemas. The contracts have to be deployed on the network before using them, and each schema used by the application have to have contract address (per network) written in. Furthermore the address in schema has to reflect the contract version described by the schema. Shortly speaking: don't manipulate the schemas unless you know what are you doing. Usually the schema in `submodules/contract-deployments/backstage/` reflect to the working contracts deployed on the backstage.


## Production notes

All applications should have defined following parameters:

+ `SB_ENV` - environment variable specifying stage (eg: production, backstage, testing ...)
+ `-rollbar` - flag with rollbar token.


## Scripts

### abigen

`scripts/abigen.js` is used to generate Go interfaces based on JSON interface file produced by Truffle. Whenever new contracts are provided you should regenerate Go interfaces

    node scripts/abigen.js <path to the truffle contract builds>

Usually it's enough to use `make`:

    make abigen-backsage
