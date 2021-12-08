#!/bin/bash

ROOT_DIR=`realpath $(dirname $0)/..`

# Use BuildKit as docker builder
export DOCKER_BUILDKIT=1

function build_boki {
    docker build -t askiad/boki:sosp-ae \
        -f $ROOT_DIR/dockerfiles/Dockerfile.boki \
        $ROOT_DIR/boki
}

function build_queuebench {
    docker build -t askiad/boki-queuebench:sosp-ae \
        -f $ROOT_DIR/dockerfiles/Dockerfile.queuebench \
        $ROOT_DIR/workloads/queue
}

function build_retwisbench {
    docker build -t askiad/boki-retwisbench:sosp-ae \
        -f $ROOT_DIR/dockerfiles/Dockerfile.retwisbench \
        $ROOT_DIR/workloads/retwis
}

function build_voidbench {
    docker build -t askiad/boki-voidbench:sosp-ae \
        -f $ROOT_DIR/dockerfiles/Dockerfile.voidbench \
        $ROOT_DIR/workloads/void
}

function build_beldibench {
    docker build -t askiad/boki-beldibench:sosp-ae \
        -f $ROOT_DIR/dockerfiles/Dockerfile.beldibench \
        $ROOT_DIR/workloads/workflow
}

function build {
    build_boki
    build_queuebench
    build_retwisbench
    build_voidbench
    build_beldibench
}

function push {
    docker push askiad/boki:sosp-ae
    docker push askiad/boki-queuebench:sosp-ae
    docker push askiad/boki-retwisbench:sosp-ae
    docker push askiad/boki-voidbench:sosp-ae
    docker push askiad/boki-beldibench:sosp-ae
}

case "$1" in
build)
    echo "building..."
    build
    ;;
push)
    echo "pushing..."
    push
    ;;
esac
