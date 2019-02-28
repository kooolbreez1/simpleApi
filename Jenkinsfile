pipeline {
    agent any 
    stages {
        stage('Pull Source Code') { 
            steps {
                build job: 'Pull Code'
            }
        }
        stage('Deploy Code') { 
            steps {
                 build job: 'Deploy Code'
            }
        }
        stage('Deploy Kong') { 
            steps {
                 build job: 'Deploy Kong'
            }
        }
    }
}