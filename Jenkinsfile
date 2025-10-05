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
        bat 'docker build -t studypal-backend:ci .\\Backend'
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

    stage('Code Quality') {
      steps {
        script {
          def failures = []
          // (A) gofmt – report files that need formatting
          def s1 = bat(returnStatus: true, script: '''
            cd Backend
            docker run --rm -v "%cd%":/src -w /src golang:1.24 ^
              sh -lc "bad=$(gofmt -l .); if [ -n \\"$bad\\" ]; then echo '=== gofmt: files need formatting ==='; echo \\"$bad\\"; exit 1; fi"
          ''')
          if (s1 != 0) { failures << 'gofmt' }
    
          // (B) golangci-lint – static analysis
          def s2 = bat(returnStatus: true, script: '''
            cd Backend
            docker run --rm -v "%cd%":/app -w /app golangci/golangci-lint:v1.59.1 ^
              golangci-lint run --timeout=3m
          ''')
          if (s2 != 0) { failures << 'golangci-lint' }
    
          // (C) Prettier – frontend formatting check
          def s3 = bat(returnStatus: true, script: '''
            cd Frontend
            docker run --rm -v "%cd%":/app -w /app node:20-alpine ^
              sh -lc "npm ci && npx --yes prettier --check \\"src/**/*.{js,ts,vue,css,scss,json,md}\\""
          ''')
          if (s3 != 0) { failures << 'prettier' }
    
          // Always succeed; just log a summary
          if (failures) {
            echo "Code quality checks completed with issues in: ${failures.join(', ')} (non-blocking)."
          } else {
            echo 'Code quality checks passed.'
          }
        }
      }
    }

    stage('Deploy') {
      steps {
        bat '''
        set COMPOSE_PROJECT_NAME=studypal
        docker compose -f docker-compose.yml down --remove-orphans || echo no previous stack
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
      bat 'docker images'
    }
  }
}
