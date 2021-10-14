import groovy.transform.Field

@Field
def repoURL = 'git@github.com:dev2choiz/hello.git'
@Field
def helmTag = '3.2.4'
@Field
def projectName = 'hello-api'
@Field
def projectId = "samyn-project5"
@Field
def gcloudCredentialsId = projectId
@Field
def gcpRegion    = "europe-west1"
@Field
def clusterZone = "europe-west1-d"
@Field
def clusterName = "samyn-cluster"
@Field
def namespace   = "lazone"
@Field
def kmsKeyringName = "kr-hello"
@Field
def kmsKeyName = "key-default-hello"

return this
