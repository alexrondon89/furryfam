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
  echo "image ${IMAGE_NAME} not exists... it will be created"
fi

# creating image and container for jenkins service
docker build --no-cache -f ./tmp/deployments/jenkins/Dockerfile -t jenkins-image:latest ./tmp/deployments
docker run -d -v /var/run/docker.sock:/var/run/docker.sock -p 8080:8080 -p 50000:50000 --name "${CONTAINER_NAME}" "${IMAGE_NAME}"

# checking if jenkins container is running without asking for initial token
echo "waiting for Jenkins to start..."
for ((i=1; i<=10; i++)); do
  response="$(curl -s http://localhost:8080/api/json)"
    if echo "$response" | grep -q 'Authentication required'; then
        echo "Jenkins is fully up and running on http://localhost:8080, and the login screen is ready."
        break
    else
        echo "Attempt $i failed... Jenkins is not ready yet. Retrying in 10 seconds..."
        sleep 10
    fi
done