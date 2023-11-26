#!/bin/bash
# TODO: Stop currently running application and all the dependencies using Make

if [ $# -ne 2 ]; then
  echo "Please provide the required arguments:
  1. Run mode
  2. Base directory"
  exit 2
fi

mode=$1
base_dir=$2

export mode=$mode
export base_dir=$base_dir

make down mode=$mode base_dir=$base_dir
