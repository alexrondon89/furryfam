

import jenkins.model.*
import hudson.model.*
import org.jenkinsci.plugins.workflow.job.WorkflowJob
import org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition

// Crear o actualizar un job tipo Pipeline
def jobName = 'user-service'
def jenkins = Jenkins.instance
def myJob = jenkins.getItem(jobName) ?: jenkins.createProject(WorkflowJob, jobName)
def pipelineScript = '''
pipeline {
    agent any  // Define que el pipeline puede ejecutarse en cualquier agente disponible

    stages {
        stage('Git Pull') {
            steps {
                // Clona o actualiza el repositorio público sin credenciales
                git url: 'https://github.com/alexrondon89/furryfam.git', branch: 'main'
            }
        }
        stage('Prueba') {
            steps {
                echo 'Este es el stage de prueba'
                script {
                    println("Más acciones pueden ser ejecutadas aquí dentro del bloque script.")
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
