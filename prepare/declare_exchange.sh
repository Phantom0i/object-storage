#!/bin/sh

## 创建rabbitMQ 的exchange

python rabbitmqadmin declare exchange name=apiServers type=fanout
python rabbitmqadmin declare exchange name=dataServers type=fanout