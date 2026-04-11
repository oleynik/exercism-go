pipeline {
    agent {
        docker {
            image 'golang:1.22-bookworm'
            args  '-u root:root'
        }
    }
    options {
        timestamps()
        timeout(time: 30, unit: 'MINUTES')
        buildDiscarder(logRotator(numToKeepStr: '20'))
    }
    stages {
        stage('Build') {
            steps {
                sh '''
                    for d in */; do
                        (cd "$d" && go build ./...) || exit 1
                    done
                '''
            }
        }
        stage('Test') {
            steps {
                sh '''
                    for d in */; do
                        (cd "$d" && go test ./...) || exit 1
                    done
                '''
            }
        }
    }
}
