#!/bin/sh
# wait-for-it.sh

set -e

host="$1"
port="$2"
shift
shift

until nc -z -v -w30 $host $port; do
  echo "Waiting for $host:$port..."
  sleep 1
done

exec "$@"
