#!groovy

parameters([
    string(name: 'VERSION', defaultValue: 'dev', description: 'branch or tag'),
])

node {
    version = params.VERSION
    def repoURL = 'git@github.com:dev2choiz/hello.git'
    //def repoURL = 'https://github.com/dev2choiz/hello'
    def helmTag = '3.2.4'
    def projectName = 'hello'
    def gcloudCredentialsId = "samyn-project2"
    def projectId = "samyn-project2"
    def clusterZone = "europe-west1-d"
    def clusterName = "samyn-cluster"
    def commitShort = ""
    sh "echo \"###########################\""
    sh "echo version=${version}"
    sh "echo scm=${scm}"

    stage("Checkout") {
        checkout scm
        sh "ls -la ${pwd()}"
    }

   echo "start building version ${version}"
   checkout scm: [
      $class: 'GitSCM',
      userRemoteConfigs: [[url: repoURL]],
      branches: [[name: version]]
      ], poll: false
   commitShort = sh(script: 'git rev-parse HEAD', returnStdout: true).trim().substring(0,7)

    echo "build docker image tag:  ${commitShort}"
    stage('Build image') {
      googleCloudBuild \
        credentialsId: gcloudCredentialsId,
        source: local('.'),
        request: file('deployments/cloudbuild/build-image.yaml'),
        substitutions: [
            _APP_NAME: projectName,
            _COMMIT_SHORT: commitShort,
            _PROJECT_ID: projectId,
        ]
    }
    echo "deploy docker image tag:  ${commitShort}"
    /* stage('Deployment') {
      googleCloudBuild \
        credentialsId: gcloudCredentialsId,
        source: local('.'),
        request: file('deployments/cloudbuild/deployment.yaml'),
        substitutions: [
            _CLUSTER_ZONE: clusterZone,
            _CLUSTER_NAME: clusterName,
            _APP_NAME: projectName,
            _COMMIT_SHORT: commitShort,
            _PROJECT_ID: projectId,
            _HELM_TAG: helmTag,
        ]
    } */
}
