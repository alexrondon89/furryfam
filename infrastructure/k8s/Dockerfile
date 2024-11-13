# Usa una imagen base de Alpine que es ligera
FROM alpine:3.18

# Definir la versión de K3s como una variable de build ARG
ARG K3S_VERSION=v1.24.4+k3s1

# Establecer la variable de entorno para K3s
ENV INSTALL_K3S_VERSION=${K3S_VERSION}

# Instalar dependencias necesarias
RUN apk add --no-cache curl bash openrc

# Iniciar openrc para cumplir con los requisitos de K3s
RUN openrc && touch /run/openrc/softlevel

# Instalar dependencias necesarias
RUN apk add --no-cache curl bash

# Descargar e instalar K3s
RUN curl -sfL https://get.k3s.io | INSTALL_K3S_VERSION=${K3S_VERSION} sh -

# Exponer el puerto para que K3s sea accesible desde la red interna de Docker
EXPOSE 6443

# Iniciar K3s cuando se inicie el contenedor
CMD ["k3s", "server"]
