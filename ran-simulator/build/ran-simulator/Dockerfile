# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

ARG ONOS_BUILD_VERSION=undefined

FROM onosproject/golang-build:v1.3.0 as build

ENV GO111MODULE=on
ARG ONOS_MAKE_TARGET=build

COPY Makefile go.mod go.sum /go/src/github.com/onosproject/ran-simulator/
COPY cmd/ /go/src/github.com/onosproject/ran-simulator/cmd/
COPY pkg/ /go/src/github.com/onosproject/ran-simulator/pkg/
COPY vendor/ /go/src/github.com/onosproject/ran-simulator/vendor/
COPY build/build-tools /go/src/github.com/onosproject/ran-simulator/build/build-tools/

RUN cd /go/src/github.com/onosproject/ran-simulator && GOFLAGS=-mod=vendor make ${ONOS_MAKE_TARGET}

FROM frolvlad/alpine-glibc:alpine-3.12

USER nobody

COPY --from=build /go/src/github.com/onosproject/ran-simulator/build/_output/ransim /usr/local/bin/ransim
COPY --from=build /go/src/github.com/onosproject/ran-simulator/pkg/servicemodel/kpm2/datasets /usr/local/datasets
# COPY pkg/config/*.yaml /etc/onos/config/

EXPOSE 5150
ENTRYPOINT ["ransim"]
