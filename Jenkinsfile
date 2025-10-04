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
          go test ./internal/health -v -vet=off
        '''
      }
    }

    stage('Deploy') {
      steps {
        bat '''
        set COMPOSE_PROJECT_NAME=studypal
    
        rem Kill anything else using 8000/3000 so our compose can bind them
        powershell -Command "$ids = docker ps -q --filter \"publish=8000\"; if ($ids) { docker rm -f $ids }"
        powershell -Command "$ids = docker ps -q --filter \"publish=3000\"; if ($ids) { docker rm -f $ids }"
    
        docker compose -f docker-compose.yml down -v --remove-orphans || echo no previous stack
        docker compose -f docker-compose.yml up -d --pull never --force-recreate
        '''
        bat 'docker compose -f docker-compose.yml ps'
        bat 'docker compose -f docker-compose.yml logs backend --tail=120'
        bat 'powershell -Command "$ok=(Test-NetConnection -ComputerName localhost -Port 8000).TcpTestSucceeded; if (-not $ok) { Write-Error \\"Backend 8000 not listening.\\"; exit 1 }"'
        bat 'powershell -Command "$ok=(Test-NetConnection -ComputerName localhost -Port 3000).TcpTestSucceeded; if (-not $ok) { Write-Error \\"Frontend 3000 not listening.\\"; exit 1 }"'
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
