pipeline {
  agent any
  options { timestamps() }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
        echo 'Checked out.'
      }
    }

    stage('Build') {
      steps {
        // Build Backend image
        bat 'docker build -t studypal-backend:ci .\\Backend'
        // Build Frontend image
        bat 'docker build -t studypal-frontend:ci .\\Frontend'
      }
    }

    stage('Test') {
      steps {
        // Run Go tests inside a docker container
        bat '''
        cd Backend
        docker run --rm -e CGO_ENABLED=1 -v "%cd%":/src -w /src golang:1.24-alpine ^
          sh -lc "apk add --no-cache build-base && go test ./... -v -coverprofile=coverage.out && go tool cover -func=coverage.out | tail -n 1"
        '''
      }
    }
  }

  post {
    always {
      // Show the images that were built
      bat 'docker images'
    }
  }
}
