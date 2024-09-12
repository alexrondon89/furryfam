# Usa una imagen base de Ubuntu
FROM ubuntu:20.04
ARG FILE_NAME

# Instala las dependencias y Ansible (u otro servicio si deseas)
RUN apt-get update && apt-get install -y ansible curl

COPY ./$FILE_NAME.yaml .
# Expone el puerto que deseas usar para el health check (por ejemplo, 8080)
EXPOSE 8080

# Define el comando por defecto para ejecutar Ansible o algún otro servicio
CMD [ "ansible", "--version" ]
