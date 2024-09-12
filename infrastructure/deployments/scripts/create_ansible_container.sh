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
if [ "$ENVIRONMENT" != 'local' ]; then
  echo "building and running ${CONTAINER_NAME} in ${ENVIRONMENT}"
  docker build --no-cache -t "$IMAGE_NAME":latest -f Dockerfile .
else
  echo "building and running ${CONTAINER_NAME} locally"
  docker build --no-cache -f ./../../../infrastructure/deployments/ansible/Dockerfile -t "$IMAGE_NAME":latest ./../../../infrastructure/deployments/ansible/
fi

echo "executing $CONTAINER_NAME $IMAGE_NAME ..."
docker run --rm --name "${CONTAINER_NAME}" "${IMAGE_NAME}" ansible-playbook ./deploy.yaml