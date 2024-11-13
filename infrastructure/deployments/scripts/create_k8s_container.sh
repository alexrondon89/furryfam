#!/bin/bash

K8S_CONTAINER_NAME="${1}-k8s-container"
K8S_CONTAINER_IMAGE="${1}-k8s-image"

if [ -n "$(docker ps -a -q -f name=$K8S_CONTAINER_NAME)" ]; then
  echo "$K8S_CONTAINER_NAME is running"
  exit 0
else
  if [ "$(docker images -q $K8S_CONTAINER_IMAGE)" ]; then
      echo "deleting $K8S_CONTAINER_NAME image..."
      docker rmi -f ${K8S_CONTAINER_IMAGE}
  fi
  echo "creating $K8S_CONTAINER_NAME..."
  docker build --no-cache -t $K8S_CONTAINER_IMAGE:latest -f ./infrastructure/k8s/Dockerfile .
  echo "executing $K8S_CONTAINER_NAME $K8S_CONTAINER_IMAGE ..."
  docker run -v /var/run/docker.sock:/var/run/docker.sock -d --name "${K8S_CONTAINER_NAME}" "${K8S_CONTAINER_IMAGE}:latest"
fi
