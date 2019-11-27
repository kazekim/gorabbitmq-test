#!/usr/bin/env bash
docker run -d --rm -p 15672:5672 --name kazekim-mq --hostname kazekim-rabbitmq rabbitmq:3
