/* groovylint-disable LineLength */
/* groovylint-disable-next-line CompileStatic */
pipeline {
    agent any

    stages {
        /* groovylint-disable-next-line SpaceInsideParentheses */
        stage('Checkout Codebase') {
            steps {
                checkout scm: [$class: 'GitSCM', branches: [[name: '*/main']], userRemoteConfigs: [[credentialsID: 'gitlab-ssh-key', url: 'gitlab.com/ljonllc/slack-bot-jenkins-integration.git']]]
            }
        }

        stage('Build') {
            steps {
                echo 'Building Codebase'
            }
        }

        stage('Test') {
            steps {
                echo 'Running Test on changes'
            }
        }

        stage('Deploy') {
            steps {
                echo 'Done!'
            }
        }
    }

    post {
        always {
            echo 'Sending Slack Message'
            sh "go run slack-notification-go ${BUILD_URL} ${currentBuild.currentResult} ${BUILD_NUMBER} ${JOB_NAME} "
        }
    }
}
