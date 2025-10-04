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
        bat '''
        cd Backend
        docker run --rm -e GOTOOLCHAIN=auto -v "%cd%":/src -w /src golang:1.24 ^
          go test ./... -v
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
