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
        _CF_TIMEOUT: "30s",
    ]
}

// Generic function to deploy a cloud function
def deployCloudFunction(appTag, env, sourcePath, funcName, goModule, addSubst) {
    googleCloudBuild   \
        credentialsId: moduleVars.gcloudCredentialsId,
            source: local('.'),
            request: file('deployments/cloudbuild/deploy-functions/cloudbuild.yaml'),
            substitutions: getCbSubstitutions(appTag, env) + [
                    _FUNCTION_SOURCE_PATH: sourcePath,
                    _FUNCTION_NAME       : funcName,
                    _FUNCTION_GO_MODULE  : goModule,
                    _ENTRYPOINT          : 'Execute',
                    _TRIGGER_TOPIC       : funcName,
            ] + addSubst
}

// deploy "function1" cloud function
def deployFunction1(appTag, env) {
    deployCloudFunction(appTag, env, 'cmd/functions/function1', 'hello-function1', 'github.com/dev2choiz/hello/cmd/function/function1', [])
}

// deploy "pg-migration" cloud function
def deployPgMigration(appTag, env) {
    deployCloudFunction(appTag, env, 'cmd/functions/migration', 'pg-migration', 'github.com/dev2choiz/hello/cmd/function/migration', [])
}

return this
