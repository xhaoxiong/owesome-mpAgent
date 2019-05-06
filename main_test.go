/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE main
 */

package main

import (
	"testing"
)

func TestA(t *testing.T) {
	s := `"{\"cmds\":[{\"cmd\":\"rm -rf ccsu-micro-platform-projects\",\"dir\":\"$work_dir\"},{\"cmd\":\"git clone $repo_url\",\"dir\":\"$work_dir\"},{\"cmd\":\"mvn install\",\"dir\":\"$work_dir/ccsu-micro-platform-projects/common-code\"},{\"cmd\":\"mvn clean package -P$profile\",\"dir\":\"$work_dir/ccsu-micro-platform-projects/$service_name\"},{\"cmd\":\"mv target/*.jar /var/jars/\",\"dir\":\"$work_dir/ccsu-micro-platform-projects/$service_name\"},{\"cmd\":\"docker run -d --name $service_name -p $port:$port -v /var/jars:/var/jars -v /var/logs:/var/logs --restart=always java:openjdk-8-jre-alpine java -jar /var/jars/$service_name-$version.jar $options\",\"dir\":\"/var/jars/\"}],\"action\":\"build\",\"params\":{\"repo_url\":\"https://github.com/notobject/ccsu-micro-platform-projects.git\",\"port\":8761,\"service_name\":\"ccsu-register-server\",\"work_dir\":\"/tmp/build\",\"profile\":\"prod\",\"options\":\"--server.port=$port --eureka.instance.ip-address=$HOST --eureka.client.service-url.defaultZone=$REGISTER_CENTER\",\"version\":\"1.0.0.RELEASE\"},\"taskId\":\"5144f2e29c4549cb9492891be08897b5\"}"`
	t.Log(string(s[1 : len(s)-1]))

}
