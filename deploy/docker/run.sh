#!/bin/bash

mkdir -p logs
mkdir -p postgres/pg-data

echo $PWD
docker stack deploy --compose-file=docker-compose.yml edo_api