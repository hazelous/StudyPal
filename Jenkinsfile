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
  }

  post {
    always {
      // Show the images that were built
      bat 'docker images'
    }
  }
}
