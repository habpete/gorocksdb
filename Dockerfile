FROM golang:1.18.4-bullseye

COPY ./ext ./ext-scripts

RUN cd ./ext-scripts && sh ./rocksdb.sh