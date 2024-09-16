# Usa una imagen base de Ubuntu
FROM ubuntu:20.04
ARG FILE_NAME

#instala tambien el daemon que no es necesario, to check a liger installing
RUN apt-get update && apt-get install -y docker.io wget

# Instala las dependencias y Ansible (u otro servicio si deseas)
RUN apt-get update && apt-get install -y ansible curl
RUN echo "####### en ansible DOCKERFILE"
RUN ls
COPY ./infrastructure/deployments/ansible/$FILE_NAME.yaml ./
COPY ./services/$FILE_NAME/Dockerfile ./

# Expone el puerto que deseas usar para el health check (por ejemplo, 8080)
EXPOSE 8081

# Define el comando por defecto para ejecutar Ansible o algún otro servicio
CMD [ "ansible", "--version" ]
