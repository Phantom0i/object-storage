#!/bin/sh

export RABBITMQ_SERVER=ampq://test:test@127.0.0.1:5672

# api service
LISTEN_ADDRESS=:12345
ES_SERVER=127.0.0.1

# data service
STORAGE_ROOT=/tmp
