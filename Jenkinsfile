pipeline {
    agent any
    tools {
        
    }

    environment {
        // Replace with your Docker registry URL and credentials ID
        DOCKER_REGISTRY = 'https://hub.docker.com'
        DOCKER_REGISTRY_CREDENTIALS_ID = 'registry'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Generate a unique tag based on timestamp
                    def imageTag = "app-${env.BUILD_ID}-${env.BUILD_NUMBER}"

                    // Build the Docker image
                    sh "docker build -t ${env.DOCKER_REGISTRY}/${imageTag} ."

                    // Log in to the Docker registry and push the image
                    withCredentials([usernamePassword(credentialsId: env.DOCKER_REGISTRY_CREDENTIALS_ID, usernameVariable: 'REGISTRY_USER', passwordVariable: 'REGISTRY_PASSWORD')]) {
                        sh "docker login -u ${REGISTRY_USER} -p ${REGISTRY_PASSWORD} ${env.DOCKER_REGISTRY}"
                        sh "docker push ${env.DOCKER_REGISTRY}/${imageTag}"
                    }
                }
            }
        }
    }
}