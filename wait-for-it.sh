#!/usr/bin/env bash
#   Use this script to test if a given TCP host/port are available

# The MIT License (MIT)
# https://github.com/vishnubob/wait-for-it

set -e

TIMEOUT=15
QUIET=0
HOST=""
PORT=""

while [[ $# -gt 0 ]]; do
    case "$1" in
        -h|--host)
            HOST="$2"
            shift 2
            ;;
        -p|--port)
            PORT="$2"
            shift 2
            ;;
        --)
            shift
            break
            ;;
        *)
            break
            ;;
    esac
done

if [[ "$HOST" == "" || "$PORT" == "" ]]; then
    HOSTPORT=($1)
    IFS=: read HOST PORT <<< "$HOSTPORT"
    shift
fi

for i in $(seq $TIMEOUT); do
    nc -z "$HOST" "$PORT" && break
    if [[ $i -eq $TIMEOUT ]]; then
        echo "Timeout after $TIMEOUT seconds waiting for $HOST:$PORT"
        exit 1
    fi
    sleep 1
done

exec "$@"
