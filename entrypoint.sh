#!/bin/bash  -x

proto_go_out() {
  protoc --go_out=. "$1"
}


find /app/proto -name '*.proto' -type f -print0 | xargs -0 -I {} bash -c 'protoc --proto_path=/app/proto --go_out=/app/proto "$@"' _ {}
#cp -rp /app/proto/github.com /app/build/vendor

#tail -f /dev/null