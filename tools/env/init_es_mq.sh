#!/bin/bash

echo "install es, please"

echo "install rabbitmq, please"

sudo yum install rabbitmq-server
systemctl start rabbitmq-server
rabbitmq-plugins enable rabbitmq_management


wget http://master:15672/cli/rabbitmqadmin

rabbitmqadmin declare exchange name=apiServers type=fanout
rabbitmqadmin declare exchange name=dataServers type=fanout

rabbitmqadmin list exchanges