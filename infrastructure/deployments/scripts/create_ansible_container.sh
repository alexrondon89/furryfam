#!/bin/bash

CONTAINER_NAME="$1-container"
IMAGE_NAME="$1-image"

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
  echo "image ${IMAGE_NAME} not exists... it will be created"
fi

# creating image and container for ansible service
echo "building and running ${CONTAINER_NAME} in ${ENVIRONMENT}"
echo "docker build --build-arg FILE_NAME=$1 --no-cache -t $IMAGE_NAME:latest -f Dockerfile ."
docker build --build-arg FILE_NAME=$1 --no-cache -t "$IMAGE_NAME":latest -f Dockerfile .

echo "executing $CONTAINER_NAME $IMAGE_NAME ..."
docker run --rm --name "${CONTAINER_NAME}" "${IMAGE_NAME}" ansible-playbook ./$1.yaml