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
        docker run --rm -e GOTOOLCHAIN=local -v "%cd%":/src -w /src golang:1.24 ^
          go test ./internal/health -v
        '''
      }
    }

    stage('Deploy') {
      steps {
        bat '''
        set COMPOSE_PROJECT_NAME=studypal
        docker compose -f docker-compose.yml down -v --remove-orphans || echo no previous stack
    
        rem Clean up any old manually created containers that might clash
        docker rm -f studypal_backend 2>NUL || echo no manual backend
        docker rm -f studypal_frontend 2>NUL || echo no manual frontend
    
        docker compose -f docker-compose.yml up -d --pull never
        '''
    
        // verify ports match compose (8081 & 3000)
        bat 'powershell -Command "$ok=(Test-NetConnection -ComputerName localhost -Port 8000).TcpTestSucceeded; if (-not $ok) { Write-Error \\"Backend port 8000 not listening.\\"; exit 1 }"'
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
