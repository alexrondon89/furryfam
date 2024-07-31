FROM  jenkins/jenkins:lts
USER root
RUN apt-get update && apt-get install -y docker.io wget

# to set the update wizard has already completed. It avoid the initial page "Unlock Jenkins"
RUN echo 2.0 > /usr/share/jenkins/ref/jenkins.install.UpgradeWizard.state
# to handle aspects of software upgrade and ensure that certain post-upgrade processes are not run again.
#RUN echo 2.0 > /usr/share/jenkins/ref/jenkins.install.InstallUtil.lastExecVersion

# to create a directory for Groovy scripts
RUN if [ $(! -d "/usr/share/jenkins/ref/init.groovy.d") ]; then mkdir /usr/share/jenkins/ref/init.groovy.d; fi

# to create an admin user automatically
COPY security.groovy /usr/share/jenkins/ref/init.groovy.d/security.groovy

# to create the pipeline for user service
COPY mock.groovy /usr/share/jenkins/ref/init.groovy.d/mock.groovy

#USER jenkins
RUN jenkins-plugin-cli --plugins \
  blueocean:1.24.7 \
  pipeline-utility-steps:2.10.0
EXPOSE 8080 50000