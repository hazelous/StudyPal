pipeline {
  agent any
  options { timestamps() }
  stages {
    stage('Checkout') {
      steps {
        checkout scm
        echo 'Test Build Successful'
      }
    }
  }
}
