#!/bin/bash

# This shell is executed before docker build.
export LD_LIBRARY_PATH=.
./riskControl -c config.yaml