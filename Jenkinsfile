pipeline {
  agent any
  options { timestamps() }
  stages {
    stage('Checkout') {
      steps {
        checkout scm
        echo 'Repo checked out. Ready to build!'
      }
    }
  }
}
