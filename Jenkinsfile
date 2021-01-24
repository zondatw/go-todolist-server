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
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                sh 'go test -cover ./...'
            }
        }
    }
}
