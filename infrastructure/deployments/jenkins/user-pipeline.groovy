

import jenkins.model.*
import hudson.model.*
import org.jenkinsci.plugins.workflow.job.WorkflowJob
import org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition

// Crear o actualizar un job tipo Pipeline
def jobName = 'user-server'
def jenkins = Jenkins.instance
def myJob = jenkins.getItem(jobName) ?: jenkins.createProject(WorkflowJob, jobName)
def pipelineScript = '''
pipeline {
    agent any  // Define que el pipeline puede ejecutarse en cualquier agente disponible
    environment{
        SERVER_NAME = 'user-server'
    }

    stages {
        stage('Git Pull') {
            steps {
                // Clona o actualiza el repositorio público sin credenciales
                git url: 'https://github.com/alexrondon89/furryfam.git', branch: 'main'
            }
        }
        stage('create docker image') {
            steps {
                sh 'docker build --no-cache -t $SERVER_NAME:latest -f $WORKSPACE/services/$SERVER_NAME/Dockerfile $WORKSPACE/services/$SERVER_NAME/'
                sh 'docker login -u alexrondon89 -p Cr1sa!3x8960'
                sh 'docker tag $SERVER_NAME:latest alexrondon89/$SERVER_NAME:latest'
                sh 'docker push alexrondon89/$SERVER_NAME:latest'
            }
        }
        stage('Run Ansible') {
            steps {
                script {
                    // Listar los archivos copiados para verificar que deploy.yaml está presente
                    sh 'ls -l $WORKSPACE'

                    // Dar permisos de ejecución al archivo copiado (por si no los tiene)
                    sh 'chmod +x $WORKSPACE/infrastructure/deployments/scripts/create_ansible_container.sh'

                    // Ejecutar el script copiado
                    sh '$WORKSPACE/infrastructure/deployments/scripts/create_ansible_container.sh $SERVER_NAME'
                }
            }
        }
    }

    post {
        always {
            echo 'Este mensaje siempre se mostrará independientemente del resultado del Pipeline'
        }
        success {
            echo 'Este mensaje solo se mostrará si el Pipeline se ejecuta exitosamente'
        }
        failure {
            echo 'Este mensaje solo se mostrará si el Pipeline falla'
        }
    }
}
'''
myJob.definition = new CpsFlowDefinition(pipelineScript, true)
myJob.save()
jenkins.reload()
