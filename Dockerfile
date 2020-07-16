FROM golang:1.14-buster AS builder

COPY fuzzers/ /root/fuzzers/

RUN chmod +x /root/fuzzers/build.sh && \
 /root/fuzzers/build.sh