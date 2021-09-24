import groovy.transform.Field

//def repoURL = 'git@github.com:dev2choiz/hello.git'
@Field
def helmTag = '3.2.4'
@Field
def projectName = 'hello-api'
@Field
def projectId = "samyn-project4"
@Field
def gcloudCredentialsId = projectId
@Field
def clusterZone = "europe-west1-d"
@Field
def clusterName = "samyn-cluster"
@Field
def namespace   = "lazone"
@Field
def environment = "staging"

return this
