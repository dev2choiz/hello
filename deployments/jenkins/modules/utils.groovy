def ws = pwd()
def moduleVars = load "${ws}/deployments/jenkins/modules/variables.groovy"

def checkout(appTag) {
    echo "in checkout"
    echo moduleVars.repoURL
    echo appTag
    checkout scm: [
            $class: 'GitSCM',
            userRemoteConfigs: [[url: moduleVars.repoURL]],
            branches: [[name: appTag]]
    ], poll: false
    def commitShort = sh(script: 'git rev-parse HEAD', returnStdout: true).trim().substring(0,7)
    echo "git commit: ${commitShort}"
    return commitShort
}

return this
