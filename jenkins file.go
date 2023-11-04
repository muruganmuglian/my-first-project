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
		deploy adapters: deploy adapters: [tomcat9(credentialsId: '052325e3-b0e0-4ad4-98fa-fbd1adf25d3e', path: '', url: 'http://localhost:9090')], contextPath: null, war: '**/*.war'
	}
}

}
		}
	}
}
