pipeline {
  agent any
  options { timestamps(); ansiColor('xterm') }
  stages {
    stage('Checkout') {
      steps {
        checkout scm
        echo 'Repo checked out. Ready to build!'
      }
    }
  }
}
