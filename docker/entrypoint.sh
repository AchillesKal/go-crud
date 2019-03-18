#!/bin/bash
cp -r /go/src/cache/vendor/. /go/src/app/vendor/

exec "$@"