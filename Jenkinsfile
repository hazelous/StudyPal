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
        docker pull golang:1.24.3
        docker run --rm -e GOTOOLCHAIN=local -v "%cd%":/src -w /src golang:1.24.3 ^
          go test ./... -v
        '''
      }
    }

    stage('Deploy') {
      steps {
        // Clean up any previous stack (ignore errors if nothing is running)
        bat 'docker compose -f docker-compose.yml down -v --remove-orphans || echo no previous stack'

        // Bring up both services using the local images built in Stage 1 (set in the docker-compose.yml file)
        bat 'docker compose -f docker-compose.yml up -d'

        // fail build if ports aren't listening
        bat 'powershell -Command "$ok=(Test-NetConnection -ComputerName localhost -Port 8080).TcpTestSucceeded; if (-not $ok) { Write-Error \\"Backend port 8080 not listening.\\"; exit 1 }"'
        bat 'powershell -Command "$ok=(Test-NetConnection -ComputerName localhost -Port 3000).TcpTestSucceeded; if (-not $ok) { Write-Error \\"Frontend port 3000 not listening.\\"; exit 1 }"'
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
