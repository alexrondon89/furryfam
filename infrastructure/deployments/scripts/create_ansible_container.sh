#!/bin/bash

ANSIBLE_CONTAINER_NAME="${1}-ansible-container"
ANSIBLE_IMAGE_NAME="${1}-ansible-image"

if [ "$(docker ps -a -q -f name="${ANSIBLE_CONTAINER_NAME}")" ]; then
  echo "container ${ANSIBLE_CONTAINER_NAME} exists... deleting"
  docker rm -f "${ANSIBLE_CONTAINER_NAME}"
else
  echo "container ${ANSIBLE_CONTAINER_NAME} not exists... creating"
fi

if [ "$(docker images -q "${ANSIBLE_IMAGE_NAME}")" ]; then
  echo "image ${ANSIBLE_IMAGE_NAME} exists... deleting"
  docker rmi "${ANSIBLE_IMAGE_NAME}"
else
  echo "image ${ANSIBLE_IMAGE_NAME} not exists... it will be created"
fi

# creating image and container for ansible service
echo "building and running ${ANSIBLE_CONTAINER_NAME}"
echo "####### en ansible script"
ls
echo "####### en ansible script"
docker build --build-arg FILE_NAME=$1 --no-cache -t "$ANSIBLE_IMAGE_NAME":latest -f ./infrastructure/deployments/ansible/Dockerfile .

echo "executing $ANSIBLE_CONTAINER_NAME $ANSIBLE_IMAGE_NAME ..."
docker run -v /var/run/docker.sock:/var/run/docker.sock -d --name "${ANSIBLE_CONTAINER_NAME}" "${ANSIBLE_IMAGE_NAME}:latest"

echo "building $1 image"
docker exec "$ANSIBLE_CONTAINER_NAME" docker build -t $1:latest -f ./Dockerfile .

echo "login in dockerhub and pushing image"
docker exec "$ANSIBLE_CONTAINER_NAME" docker login -u alexrondon89 -p Cr1sa!3x8960
docker exec "$ANSIBLE_CONTAINER_NAME" docker tag $1:latest alexrondon89/$1:latest

echo "pushing image in dockerhub"
docker exec "$ANSIBLE_CONTAINER_NAME" docker push alexrondon89/$1:latest
