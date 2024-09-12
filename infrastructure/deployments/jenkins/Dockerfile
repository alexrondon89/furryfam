FROM  jenkins/jenkins:lts
USER root
#instala tambien el daemon que no es necesario, to check a liger installing
RUN apt-get update && apt-get install -y docker.io wget

RUN jenkins-plugin-cli --plugins \
    git \
    github \
    git-client \
    github-branch-source \
    scm-api \
    branch-api \
    cloudbees-folder \
    credentials \
    ssh-credentials \
    jackson2-api \
    structs \
    jsch \
    blueocean \
    pipeline-utility-steps

# to set the update wizard has already completed. It avoid the initial page "Unlock Jenkins"
RUN echo 2.0 > /usr/share/jenkins/ref/jenkins.install.UpgradeWizard.state
# to handle aspects of software upgrade and ensure that certain post-upgrade processes are not run again.
#RUN echo 2.0 > /usr/share/jenkins/ref/jenkins.install.InstallUtil.lastExecVersion

# to create a directory for Groovy scripts
RUN if [ $(! -d "/usr/share/jenkins/ref/init.groovy.d") ]; then mkdir /usr/share/jenkins/ref/init.groovy.d; fi

# to create an admin user automatically
COPY ./jenkins/security.groovy /usr/share/jenkins/ref/init.groovy.d/security.groovy
# to create the pipeline for services
COPY ./jenkins/user-pipeline.groovy /usr/share/jenkins/ref/init.groovy.d/user-pipeline.groovy

EXPOSE 8080 50000