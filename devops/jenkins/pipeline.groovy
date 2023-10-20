pipeline {
    agent any
    parameters {
        string defaultValue: '10.64.3.183', name: 'IP_ADDRESS'
        string defaultValue: '1.2.2', name: 'PROVIDER_VERSION'
        string defaultValue: 'admin', name: 'USER_NAME'
        string defaultValue: 'a10', name: 'PASSWORD'
        string defaultValue: '0', name: 'TERRAFORM_DESTROY'
        string defaultValue: '/var/lib/jenkins/jobs/state_files', name: 'STATE_DIR_PATH'
        string defaultValue: 'thunder_gslb_group', name: 'EXAMPLE_DIR_PATH'
        string defaultValue: 'https://github.com/a10networks/terraform-provider-thunder.git', name: 'GIT_REPOSITORY'
        string defaultValue: 'master', name: 'GIT_BRANCH'
    }
    stages {
        stage('Clone Repository') {
            steps {
                echo "Cloning repository!!"
                checkout([$class: 'GitSCM',
                          branches: (["$GIT_BRANCH" != 'master'] ? [[name: "$GIT_BRANCH"]] : []),
                          userRemoteConfigs: [[url: "$GIT_REPOSITORY"]]])
            }
        }

        stage('Preparing Environment') {
            steps {
                echo "Changing directory!!"
                dir("examples/resources/$EXAMPLE_DIR_PATH") {
                    echo "Removing state files!!!"
                    sh "rm -rf .terraform/ .terraform.lock.hcl .terraform.lock.hcl Thunder_shared.txt terraform.tfstate thunder.log terraform.tfstate.backup Thunder_L3v.txt"

                    echo "Creating variables.tf file!!"
                    script {
                        def varContent = """
                            variable "dut9049" {
                                description = "Thunder hostname/IP address"
                                type        = string
                                }

                            variable "username" {
                                description = "Thunder username"
                                type        = string
                            }

                            variable "password" {
                                description = "Thunder password"
                                type        = string
                            }
                        """
                        writeFile file: 'variables.tf', text: varContent
                    }

                    echo "Creating terraform.tfvars file!!"

                    script {
                        def tfvarContent = """
                        dut9049 = "${IP_ADDRESS}"
                        username = "${USER_NAME}"
                        password = "${PASSWORD}"
                        """
                        writeFile file: 'terraform.tfvars', text: tfvarContent
                    }

                    echo "Creating provider.tf file!!"

                    script {
                        def providerContent = """
                        terraform {
                            required_providers {
                                thunder = {
                                    source  = "a10networks/thunder"
                                    version = "${PROVIDER_VERSION}"
                                }
                            }
                        }
                        """
                        writeFile file: 'provider.tf', text: providerContent
                    }
                }
            }
        }

        stage('Terraform') {
            steps {
                script {
                    if (params.TERRAFORM_DESTROY == "0") {

                        echo "Running terraform init"
                        dir("examples/resources/$EXAMPLE_DIR_PATH") {
                            sh 'terraform init'
                        }

                        echo "Running terraform plan"
                        dir("examples/resources/$EXAMPLE_DIR_PATH") {
                            sh "terraform plan -var=\"dut9049=${params.IP_ADDRESS}\" -var=\"username=${params.USER_NAME}\" -var=\"password=${params.PASSWORD}\""
                        }

                        echo "Running terraform apply"
                        dir("examples/resources/$EXAMPLE_DIR_PATH") {
                            sh "terraform apply --auto-approve -var=\"dut9049=${params.IP_ADDRESS}\" -var=\"username=${params.USER_NAME}\" -var=\"password=${params.PASSWORD}\""
                        }

                        def filename = "${EXAMPLE_DIR_PATH}"
                        def path = "${STATE_DIR_PATH}/${filename}"
                        echo "${path}"

                        echo "Copying state files"
                        dir("examples/resources/$EXAMPLE_DIR_PATH") {
                            sh "mkdir -p ${path}"
                            sh "cp -r .terraform .terraform.lock.hcl terraform.tfstate provider.tf terraform.tfvars variables.tf ${path}"
                        }


                    } else {
                        def filename = "${EXAMPLE_DIR_PATH}"
                        def path = "$STATE_DIR_PATH/$filename"
                        def example_path = "examples/resources/$EXAMPLE_DIR_PATH"

                        echo "Copying state files"
                        dir("examples/resources/$EXAMPLE_DIR_PATH") {
                            sh "cp -r ${path}/.terraform ."
                            sh "cp -r ${path}/.terraform.lock.hcl ."
                            sh "cp -r ${path}/terraform.tfstate ."
                            sh "cp -r ${path}/provider.tf ."
                            sh "cp -r ${path}/terraform.tfvars ."
                            sh "cp -r ${path}/variables.tf ."
                        }

                        echo "Running terraform destroy"
                        dir("examples/resources/$EXAMPLE_DIR_PATH") {
                            sh "terraform destroy --auto-approve -var=\"dut9049=${params.IP_ADDRESS}\" -var=\"username=${params.USER_NAME}\" -var=\"password=${params.PASSWORD}\""
                        }
                    }
                }
            }
        }
    }
}
