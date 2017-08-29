CHECK=' \033[32m✔\033[39m'
HR=\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#
IBLACK='\e[1;30m'
GREEN='\e[0;32m'
NC='\e[0m'              # No Color

VERSION=`git describe`
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
		github.com/golang/mock
#		rsc.io/gt

# check https://github.com/govend/govend for dependency management
install-deps:
	@go get -v github.com/govend/govend
	@echo -e $(IBLACK)checking if all dependencies are installed... $(NC)
	@govend
	@echo -e $(CHECK)


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
	@errcheck -ignoretests ./go-lib/... ./cmd/...
# TODO: integrate github.com/mibk/dupl  and  github.com/opennota/check  and github.com/dominikh/go-simple

lint-go-mega:
	megacheck ./cmd/... ./go-lib/...

###############################
# building

build:
#  -ldflags "-X bitbucket.org/sweetbridge/oracles/go-lib/setup.GitVersion=$(VERSION) -w"
	@GOBIN=`pwd`/bin go install -v ./cmd/...
