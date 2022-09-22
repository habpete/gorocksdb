#!/usr/bin/env bash

ROCKSDB_VERSION=7.5.3

curl -OL https://github.com/facebook/rocksdb/archive/refs/tags/v$ROCKSDB_VERSION.tar.gz && \
    tar xvzf v$ROCKSDB_VERSION.tar.gz && cd rocksdb-$ROCKSDB_VERSION && PORTABLE=1 make static_lib