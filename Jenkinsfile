pipeline {
    environment {
        hostName = "code-vps"
        //TODO: Change to your own host ip
        host = "HOST_IP"

        //default branchs
        branch = "stg"

        //TODO: Change to your own service
        service = "SERVICE_NAME"
        version = "$branch-0.0.$BUILD_NUMBER"

        //TODO: Change to your own group and service
        registry = "registry.gitlab.com/{GROUP}/{PROJECT}/$service"
        registryCredentials = "code-gitlab-registry-cred"
        dockerImage = ""
    }

    agent any

    stages {
        stage('Determine version and branch env') {
            steps {
                script {
                    if ("$BRANCH_NAME" == "production") {
                        branch = "prod"
                        version = "$branch-0.0.$BUILD_NUMBER" 
                    }else if ("$BRANCH_NAME" == "staging"){
                        branch = "stg"
                        version = "$branch-0.0.$BUILD_NUMBER" 
                    }else {
                        branch = "dev"
                        version = "$branch-0.0.$BUILD_NUMBER"

                    } 
                }
            }
        } 
        stage('Clone repository from branch staging') {
            when {
                branch "staging"
            } 
            steps {
                echo 'Cloning repository from branch ' + env.BRANCH_NAME
                checkout scm
            } 
        }

        stage('Build') {
            steps {
                echo 'Building image ...'
                script {
                    dockerImage = docker.build registry
                }
            }
        }
        stage('Publish') {
            steps {
                echo 'Publishing image ...'
                script {
                    docker.withRegistry("https://" + registry   , registryCredentials) {
                        dockerImage.push("$version")
                        dockerImage.push("latest")
                    }
                }
            }
        }
        stage('Remove unused docker image') {
            steps {
                echo 'Removing unused image ...'
                sh "docker rmi $registry:$version"
                sh "docker rmi $registry:latest"
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying ...'
                script {
                    def remote = [:]
                    remote.name = "$hostName"
                    remote.host = "$host"
                    remote.allowAnyHosts = true
                    withCredentials([usernamePassword(credentialsId: "code-eclaim-vps-cred",usernameVariable: "userName", passwordVariable: "password")]) {
                        remote.user = userName
                        remote.password = password

                        //TODO: Updated following your manual command deployment in server
                        def composePure = "./eclaim-devops/services/docker-compose.$service"+".yml" 
                        def compose = "./eclaim-devops/services/docker-compose.$branch" + ".$service" + ".yml" 
                        sshCommand remote: remote, command: "cp $composePure $compose && \
                        sed -i -e 's/IMAGE_VERSION/$version/g' $compose"
                        sshCommand remote: remote, command: "sudo docker stack deploy -c $compose eclaim-$service-$branch --with-registry-auth && \
                        rm -rf $compose", sudo: true
                    }
                }
            }
        }
    }
} 

  