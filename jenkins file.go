pipeline{
	agent any
tools {
	maven 'local_maven
	stages{
		stage('build'){
			steps{
				sh 'maven clan package'
		
	}
	stage ( 'deploy to tomcat server'){
		post{
			success
			{
				echo "archiving the artifacts"
				archivingArtifacts artifacts: '**war/target/*.war'
			}
		}
	}
stage ('deploy to tomcat server'){
	steps{
		deploy adapters: [tomcat9(path: '', url: 'https://downloads.apache.org/tomcat/tomcat-9/v9.0.89/bin/apache-tomcat-9.0.89-deployer.zip.sha512')], contextPath: null, war: '**/\'.war'
	}
}

}
		}
	}
}
