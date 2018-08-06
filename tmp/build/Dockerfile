FROM alpine:3.6

RUN adduser -D dynamic-kubelet-operator
USER dynamic-kubelet-operator

ADD tmp/_output/bin/dynamic-kubelet-operator /usr/local/bin/dynamic-kubelet-operator
