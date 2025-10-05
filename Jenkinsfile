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
    
          // gofmt report files that need formatting
          def s1 = bat(returnStatus: true, script: '''
            cd Backend
            docker run --rm -v "%cd%":/src -w /src golang:1.24 ^
              sh -lc "bad=$(gofmt -l .); if [ -n \\"$bad\\" ]; then echo '=== gofmt: files need formatting ==='; echo \\"$bad\\"; exit 1; fi"
          ''')
          if (s1 != 0) { failures << 'gofmt' }
    
          // golangci-lint download modules then lint (all in one container)
          def s2 = bat(returnStatus: true, script: '''
            cd Backend
            docker run --rm -e GOTOOLCHAIN=auto -v "%cd%":/app -w /app golangci/golangci-lint:v1.59.1 ^
              sh -lc "go env -w GOPROXY=https://proxy.golang.org,direct; go mod download; golangci-lint run --timeout=5m"
          ''')
          if (s2 != 0) { failures << 'golangci-lint' }
    
          // Prettier frontend formatting check
          def s3 = bat(returnStatus: true, script: '''
            cd Frontend
            docker run --rm -v "%cd%":/app -w /app node:20-alpine ^
              sh -lc "npm ci && npx --yes prettier --check \\"src/**/*.{js,ts,vue,css,scss,json,md}\\""
          ''')
          if (s3 != 0) { failures << 'prettier' }
    
          // Always succeed, just logs a summary
          if (failures) {
            echo "Code quality checks completed with issues in: ${failures.join(', ')} (non-blocking)."
          } else {
            echo 'Code quality checks passed.'
          }
        }
      }
    }

    stage('Security') {
      steps {
        script {
          // Go backend: govulncheck (non-blocking)
          bat '''
          cd Backend
          docker run --rm -e GOTOOLCHAIN=auto -v "%cd%":/src -w /src golang:1.24-alpine ^
            sh -lc "export GOPROXY=https://proxy.golang.org,direct; /usr/local/go/bin/go run golang.org/x/vuln/cmd/govulncheck@latest ./... || true"
          '''
    
          // Frontend: npm audit (non-blocking)
          bat '''
          cd Frontend
          docker run --rm -v "%cd%":/app -w /app node:20-alpine ^
            sh -lc "npm ci >/dev/null 2>&1 || true; npm audit --audit-level=high || true"
          '''
    
          // Container images: Trivy scan via TARs (Windows-friendly, non-blocking)
          // Save locally built images to TAR so Trivy doesn't need Docker socket access
          bat 'docker save studypal-backend:ci  -o backend.tar'
          bat 'docker save studypal-frontend:ci -o frontend.tar'
    
          // Scan the TARs, --exit-code 0 ensures success even if vulnerabilities are found so that build can continue
          bat '''
          docker run --rm -v "%cd%":/work aquasec/trivy:latest ^
            image --input /work/backend.tar --no-progress --severity HIGH,CRITICAL --exit-code 0 --format table
          '''
          bat '''
          docker run --rm -v "%cd%":/work aquasec/trivy:latest ^
            image --input /work/frontend.tar --no-progress --severity HIGH,CRITICAL --exit-code 0 --format table
          '''
    
          echo 'Security checks (non-blocking) finished.'
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
