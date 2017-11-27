CHECK=' \033[32m✔\033[39m'
HR=\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#
IBLACK='\e[1;30m'
GREEN='\e[0;32m'
NC='\e[0m'              # No Color

VERSION=`git log  -n 1 --format="%h-%cI"`  #`git describe`
GO=go
# GOLIB=$(shell find go-lib -type f -name \*.go)


define golintx
	find $(1) -maxdepth 1 -name '[^.#]*.go' ! -name '*_test.go' ! -name '*_string.go' ! -name 'benchmark_*.go'  ! -name '*_mock.go' | xargs golint;
endef

.PHONY: setup-dev \
	go-generate validate-go-generate-is-uptodate \
	lint lint-go


setup-dev: install-deps
	@echo -en $(IBLACK)fetching server GO lib dependencies... $(NC)
	@echo "*** (Don't forget to add GOBIN to your PATH) ***"
	@go get -v -u github.com/golang/lint/golint \
		github.com/kisielk/errcheck \
		honnef.co/go/tools/cmd/megacheck \
		github.com/golang/mock \
		github.com/mdempsky/unconvert \
		github.com/mibk/dupl
#		rsc.io/gt

# check https://github.com/govend/govend for dependency management
install-deps:
	@go get -v github.com/kardianos/govendor
	@echo -e $(IBLACK)checking if all dependencies are installed... $(NC)
	@govendor sync
	@echo -e "> dependencies check completed" $(CHECK)


###############################
# vendoring

vendor-check:
	@govendor list +external

vendor-add-missing:
	@govendor add +external

###############################
# generating

go-generate:
	@echo -e $(IBLACK)running go generate on go code... $(NC)
	@$(GO) generate ./go-lib/... ./cmd/...

# validate-go-generate-is-uptodate:
# 	@./test/validate_go_generate_is_up_to_date.sh || exit 1

abigen-backstage:
	@node scripts/abigen.js submodules/contract-deployments/backstage/

###############################
# linting

lint: lint-go

lint-go:
	@echo -e $(IBLACK)linting go libs and apps... $(NC)
	@for d in $(shell find go-lib cmd -name '*.go' ! -name '*mock.go' | sort -u); do \
		$(call golintx,$$d) done;
	@echo -e $(IBLACK)$(GO) vet go-lib... cmd...$(NC)
	@$(GO) tool vet -all go-lib
	@echo -e $(IBLACK)errcheck go-lib... cmd... $(NC)
	@errcheck ./go-lib/... ./cmd/...

lint-go-mega:
	@megacheck ./cmd/... ./go-lib/...
	@unconvert -v ./go-lib/... ./cmd/...
	@dupl -v -t 90 ./go-lib ./cmd

###############################
# testing

test:
	@go test ./go-lib/... ./cmd/...


###############################
# building

define _build
	@go generate ./go-lib/...
# GOOS=darwin GOOS=windows GOARCH=amd64
	@GOBIN=`pwd`/bin go install -v \
		-ldflags "-X bitbucket.org/sweetbridge/oracles/go-lib/setup.GitVersion=$(VERSION) -w" \
		./cmd/$(1)
	@echo -e "> build completed" $(CHECK)
endef

build:
	@$(call _build,"...")

build-direct-pledge:
	@$(call _build,"direct-pledge")

build-tranche-queue:
	@$(call _build,"tranche-queue")

build-swc-distributor:
	@$(call _build,"swc-distributor")


###############################
# docker

docker-mk-builder:
	@docker build -t oracle-builder ./docker

docker-run-builder:
	docker container start -a oracle-builder_1 || \
		docker run -v ${PWD}:/go/src/bitbucket.org/sweetbridge/oracles -t --name oracle-builder_1 oracle-builder

docker-bash:
	docker container start -a oracle-bash  || \
		docker run -v ${PWD}/bin:/root/bin -it --name oracle-bash alpine:3.6
