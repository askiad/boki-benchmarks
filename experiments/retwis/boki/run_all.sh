#!/bin/bash
BASE_DIR=`realpath $(dirname $0)`
ROOT_DIR=`realpath $BASE_DIR/../../..`

HELPER_SCRIPT=$ROOT_DIR/scripts/exp_helper

$HELPER_SCRIPT start-machines --base-dir=$BASE_DIR --spot-instances-waiting-time=300

# $BASE_DIR/run_once.sh cont128_cl1 1
# $BASE_DIR/run_once.sh cont128_cl8 8
# $BASE_DIR/run_once.sh cont128_cl16 16
# $BASE_DIR/run_once.sh cont128_cl32 32
# $BASE_DIR/run_once.sh cont128_cl64 64
# $BASE_DIR/run_once.sh cont128_cl128 128
# $BASE_DIR/run_once.sh cont128_cl256 256
# $BASE_DIR/run_once.sh cont128_cl512 512
# $BASE_DIR/run_once.sh cont128_cl1024 1024
$BASE_DIR/run_once.sh cont128_cl1024_re 1024

$HELPER_SCRIPT stop-machines --base-dir=$BASE_DIR
