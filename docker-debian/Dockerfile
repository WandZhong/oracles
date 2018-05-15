# Dockerfile to build oracles
# Assuming you are in the root directory (oracles):
#   docker build -t oracle-builder ./docker
#   docker run -v ${PWD}:/go/src/bitbucket.org/sweetbridge/oracles -t --name oracle-builder_1 oracle-builder
# all subsequent build should reuse the container:
#   docker container start -a oracle-builder_1

FROM golang:1.9.6  # debian based
LABEL name=oracle-builder

RUN mkdir -p /go/src/bitbucket.org/sweetbridge
WORKDIR /go/src/bitbucket.org/sweetbridge/oracles

RUN apk add --update make git gcc libc-dev
