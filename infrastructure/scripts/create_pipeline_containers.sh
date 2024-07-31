#!/bin/bash

CONTAINER_NAME="jenkins-container"
IMAGE_NAME="jenkins-image"

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

# creating image and container for jenkins service
echo "building and running ${CONTAINER_NAME}"
docker build -f ./../../infrastructure/deployments/jenkins/Dockerfile -t jenkins-image:latest ./../../infrastructure/deployments/jenkins/
docker run -d -p 8080:8080 -p 50000:50000 --name "${CONTAINER_NAME}" "${IMAGE_NAME}"

# checking if jenkins container is running without asking for initial token
echo "waiting for Jenkins to start..."
for i in {1..50}; do
    response="$(curl -s http://localhost:8080/api/json)"
  if echo "$response" | grep -q '"useSecurity":false'; then
    echo "Jenkins is fully up and running on http://localhost:8080"
    break
  else
    echo "Attempt ${i} failed... Jenkins is not ready yet. Retrying in 10 seconds..."
    sleep 10
  fi
done