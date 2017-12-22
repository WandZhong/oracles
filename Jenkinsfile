// Copyright (c) 2017 Sweetbridge Stiftung (Sweetbridge Foundation)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

node {

 stage('Clone repository') {
        /* Let's make sure we have the repository cloned to our workspace */
        echo 'Checkout started'
        checkout scm
        echo "Branch name: ${env.BRANCH_NAME}"

    }

 stage('Clone config repo') {
        /* Let's make sure we have the repository cloned to our workspace */
        dir('config') {
          git branch: 'master',
          credentialsId: '9fdaf31c-83e8-46c7-a495-40e7dbafb3a9', url: 'git@bitbucket.org:sweetbridge/config.git'
        }
}

 stage('Docker build container') {
          echo 'creating docker build container'
                    sh 'make docker-mk-builder'
 }

 stage('Docker compile') {
          echo 'compiling oracles with docker'
                    sh 'make docker-run-builder'
 }


 stage('Build docker container') {
        /* This builds the actual image; synonymous to
         * docker build on the command line */
        echo 'Image is being build now...'
        sh 'cp docker/Dockerfile-payment-forwarder Dockerfile'
        app = docker.build('sweetbridge/payment-forwarder')
    }



}
