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
        // Backend unit tests (run in a Go container; no Go on host needed)
        // GOTOOLCHAIN=auto to automatically fetch new versions of GO
        bat '''
        cd Backend
        docker run --rm -e GOTOOLCHAIN=auto -v "%cd%":/src -w /src golang:1.24 ^
          go test ./... -v -coverprofile=coverage.out
        docker run --rm -v "%cd%":/src -w /src golang:1.24 ^
          bash -lc "go tool cover -func=coverage.out | tail -n 1 || true"
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
