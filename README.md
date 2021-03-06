# Sweetbridge Oracles

This repository contains blockchain Oracles used by Sweetbridge project.

## Setting up

Unless you are using docker environment to build the applications (explained below) you have to set-up the following environment:

1. Install `go >= 1.7` and `node.js` and `make`. Make sure that `go` executable is in your `PATH`
1. For static compilation: `libc-static` (eg: `glibc-static`).
1. Make sure `$GOPATH/bin` (`go env GOPATH`) is part of your `PATH`.
1. Make sure you have an access to the oracles repository.
1. Clone this repository into your `GOPATH` (`go env`). Install dependencies:

	git clone git@bitbucket.org:sweetbridge/oracles.git $GOPATH/src/bitbucket.org/sweetbridge/oracles
	cd ~/go/src/bitbucket.org/sweetbridge/oracles
	make install-deps

1. (alternative). Use `go get` (this will work only if the repo is open) and hack the git config a bit: https://gist.github.com/shurcooL/6927554

    go get bitbucket.org/sweetbridge/oracles/cmd/<oracle-name>

Please refer to `setup-quickstart.sh` if you want to install it on debian / ubuntu.

### Development:
This is not required for building application, but required for proper development.

	make setup-dev install-deps

#### Development workflow

+ We use [Gitlab Flow](https://docs.gitlab.com/ce/workflow/gitlab_flow.html).
+ Squash your commits before merging! Here is [why](https://softwareengineering.stackexchange.com/questions/263164/why-squash-git-commits-for-pull-requests).
+ Write tests.

### Dependencies

+ DB model is defined in the [model](https://bitbucket.org/sweetbridge/model) repository.


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

**NOTE**: For docker build we are using Alpine Linux. This allows is compatible for putting the binaries in a tiny Alpine container. However to run the same binary in other distributions (Fedora, Ubuntu, OpenSUSE) you will need to use `libc` from `musl` (which is using by Alpine during linking).

####  Building using Docker compatible with AWS Linux

For the reasons explained above, the alpine based builds are not working with Debian or Red Hat based distributions. To overcome this we provide alternative docker builders which produce compatible binaries.
The procedure is similar to the one above - just use builder commands with `-debian` suffix, eg: `make docker-run-builder-debian`


## Applications

All applications are build from the main packages in the `cmd/` sub-directories. Each package contains the README.md file with the application description.

All applications are compiled into self-contained binaries (no external dependencies needed except the compatible **libc** version). Each application support the help argument to display the usage description - please use it to find more information about required and optional parameters:

	./bin/<app_name> -h


### Logger

All applications are using structured [log15](https://github.com/robert-zaremba/log15) logger.
Logger is configured by default to construct colored output, which is helpful when printing the logs to the screen or storying in the system journal. However when redirecting the output to the file it's better to disable the colored output. We can configure logger with the following options:

+ `-log-colored=false` - disable the colored output

### Common ethereum parameters

Oracles share the following, ethereum-specific obligatory parameters:

* `-pk-hex` -- a plaintext hex-encoded private key.
* `-pk-file` + `-pk-pwd` -- a JSON encrypted private key. Required if the `-pk-hex` is not specified.
* `-eth-network` -- a network name which oracles is connecting to (use `-h` to list all network names).
* `-eth-network-file` -- a JSON file describing Ethereum networks. It's compatible with `networks` [Truffle](http://truffleframework.com/docs/advanced/configuration) file section.
* `-contracts` -- the path to the directory containing contract schemas. The contracts have to be deployed on the network before using them, and each schema used by the application have to have contract address (per network) written in. Furthermore the address in schema has to reflect the contract version described by the schema. Shortly speaking: don't manipulate the schemas unless you know what are you doing. Usually the schema in `submodules/contract-deployments/backstage/` reflect to the working contracts deployed on the backstage.


## Production notes

All applications should have defined following parameters:

+ `SB_ENV` - environment variable specifying stage (eg: production, backstage, testing ...)
+ `-rollbar` - flag with rollbar token. Has to be specified.

For production release, the command line parameters are specified using environment variables. This are translated to the command line flags using the wrapper shell scripts in the `/docker` directory.


## Scripts

### abigen

`scripts/abigen.js` is used to generate Go interfaces based on JSON interface file produced by Truffle. Whenever new contracts are provided you should regenerate Go interfaces

    node scripts/abigen.js <path to the truffle contract builds>

Usually it's enough to use `make`:

    make abigen-backsage
