#!/bin/bash

if [ $# -ne 6 ]; then
  echo "Please provide the required arguments:
  1. Run mode
  2. Base directory
  3. Container image name
  4. Container image version
  5. Configuration source file name
  6. Configuration source file path"
  exit 6
fi

mode=$1
base_dir=$2
image=$3
version=$4
config_source=$5
config_path=$6

export mode=$mode
export base_dir=$base_dir
export image=$image
export version=$version
export config_source=$config_source
export config_path=$config_path

# call Makefile with unit test & lint checks, build, run
#make lint unit-test
#make build mode=$mode base_dir=$base_dir image=$image version=$version CONFIG_SOURCE=$config_source CONFIG_PATH=$config_path
make run mode=$mode base_dir=$base_dir image=$image version=$version
