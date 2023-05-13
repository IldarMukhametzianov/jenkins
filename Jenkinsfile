pipeline {

  environment {
    dockerimagename = "ildarmukhametzyanov/app-local-ip-jenk"
    dockerImage = ""
  }

  agent any

  stages {

    stage('Checkout Source') {
      steps {
        git 'https://github.com/IldarMukhametzianov/jenkins.git'
      }
    }

    

    

    stage('Deploying container to Kubernetes') {
      steps{
        withKubeConfig([credentialsId: 'minikube']) {
          sh 'kubectl apply -f deployment.yaml'
        }
      }
    }

  }

}
