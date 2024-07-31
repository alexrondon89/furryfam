import jenkins.model.*
import hudson.model.*
import org.jenkinsci.plugins.workflow.job.WorkflowJob
import org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition

// Crear o actualizar un job tipo Pipeline
def jobName = 'NewPipeline'
def jenkins = Jenkins.instance
def myJob = jenkins.getItem(jobName) ?: jenkins.createProject(WorkflowJob, jobName)
def pipelineScript = '''
pipeline {
    agent any
    stages {
        stage('Hello') {
            steps {
                echo 'Hello, World!'
            }
        }
    }
}
'''
myJob.definition = new CpsFlowDefinition(pipelineScript, true)
myJob.save()
jenkins.reload()
