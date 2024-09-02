#!/bin/bash

CONTAINER_NAME="ansible-container"
IMAGE_NAME="ansible-image"

if [ "$(docker ps -a -q -f name="${CONTAINER_NAME}")" ]; then
  echo "container ${CONTAINER_NAME} exists... deleting"
  docker rm -f "${CONTAINER_NAME}"
else
  echo "container ${CONTAINER_NAME} not exists... creating"
fi

if [ "$(docker images -q "${IMAGE_NAME}")" ]; then
  echo "image ${IMAGE_NAME} exists... deleting"
  docker rmi "${IMAGE_NAME}"
else
  echo "image ${IMAGE_NAME} not exists... it will be downloaded"
fi

# creating image and container for ansible service
echo "building and running ${CONTAINER_NAME}"
docker build -f ./../../infrastructure/deployments/ansible/Dockerfile -t "${IMAGE_NAME}":latest ./../../infrastructure/deployments/ansible
docker run -d -p <port>:<port> --name --name "${CONTAINER_NAME}" "${IMAGE_NAME}"

echo "waiting for Jenkins to start..."
for i in {1..50}; do
## search how to check that ansible in online
done