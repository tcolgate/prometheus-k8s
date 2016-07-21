#!/bin/bash

kubectl replace -f prometheus-configmap.yaml
curl -XPOST $(minikube service prometheus --url)/-/reload
