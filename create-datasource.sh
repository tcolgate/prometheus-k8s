#!/bin/bash

curl -u admin --pass admin "$(minikube service grafana --url)/api/datasources" \
  -X POST -H "Content-Type: application/json" \
  --data-binary '{"name":"default","type":"prometheus","url":"http://prometheus:9090","access":"proxy","isDefault":true}'

