pipeline {
    agent any
    tools {
        go 'go_1_13_15'
    }
    environment {
        GOPATH = ''
        CGO_ENABLED = 0
        GO111MODULE = 'on'
    }
    stages {
        stage('Compile') {
            steps {
                sh 'go mod download'
                sh 'make build'
            }
        }
        stage('Test') {
            steps {
                sh 'make test'
            }
        }
    }
}
