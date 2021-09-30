ws = pwd()
moduleVars = load "${ws}/deployments/jenkins/modules/variables.groovy"

def getCbSubstitutions(appTag, env) {
    return [
        _APP_TAG: appTag,
        _APP_ENV: env,
        _GCP_REGION: moduleVars.gcpRegion,
        _KMS_KEYRING_NAME: moduleVars.kmsKeyringName,
        _KMS_KEY_NAME: moduleVars.kmsKeyName,
        _GITHUB_PRIVATE_KEY_ENCRYPTED_PATH: "deployments/security/hello_github_deploy_id_rsa.enc",
        _GITHUB_REPOSITORY: "github.com/dev2choiz/hello",
        _SERVICE_ACCOUNT: "sa-serverless@${moduleVars.projectId}.iam.gserviceaccount.com",
    ]
}

def deployFunction1(appTag, env) {
    googleCloudBuild \
        credentialsId: moduleVars.gcloudCredentialsId,
            source: local('.'),
            request: file('deployments/cloudbuild/deploy-functions/function1/cloudbuild.yaml'),
            //request: file('deployments/cloudbuild/deploy-functions/cloudbuild.yaml'),
            substitutions: getCbSubstitutions(appTag, env) + [
                _FUNCTION_SOURCE_PATH: 'cmd/functions/function1',
                _FUNCTION_NAME: 'hello-function1',
                _FUNCTION_GO_MODULE: 'github.com/dev2choiz/hello/cmd/function/function1',
                _ENTRYPOINT: 'Execute',
                _TRIGGER_TOPIC: 'hello-function1',
            ]
}

def deployPgMigration(appTag, env) {
    googleCloudBuild   \
        credentialsId: moduleVars.gcloudCredentialsId,
            source: local('.'),
            //request: file('deployments/cloudbuild/deploy-functions/pg-migration/cloudbuild.yaml'),
            request: file('deployments/cloudbuild/deploy-functions/cloudbuild.yaml'),
            substitutions: getCbSubstitutions(appTag, env) + [
                _FUNCTION_SOURCE_PATH: 'cmd/functions/migration',
                _FUNCTION_NAME       : 'pg-migration',
                _FUNCTION_GO_MODULE  : 'github.com/dev2choiz/hello/cmd/function/migration',
                _ENTRYPOINT          : 'Execute',
                _TRIGGER_TOPIC       : 'pg-migration',
            ]
}

return this
