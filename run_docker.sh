#!/bin/bash

set -ex

nginx -g daemon off;

./otelcontribcol_linux_amd64 --config /conf/otel-collector-config.yml;