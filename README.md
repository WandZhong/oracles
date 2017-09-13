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


## Applications

All applications are listed packaged in the `cmd/` sub-directories. Each package contains the README.md file with the application description.

To build all applications simply run the following command:

	make build

All applications are compiled into self-contained binaries (no external dependencies needed). Each application support the help argument to display the usage description:

	./bin/<app_name> -h


## Scripts

### abigen

`scripts/abigen.js` is used to generate Go interfaces based on JSON interface file produced by Truffle. Whenever new contracts are provided you should regenerate Go interfaces

    node scripts/abigen.js <path to the truffle contract builds>

Usually you should just use make:

    make abigen-backsage
