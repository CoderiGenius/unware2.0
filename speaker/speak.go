package speaker



import (
"fmt"

"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
aai "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aai/v20180522"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"encoding/json"
)

type myconfig struct {
	Tencent api
}
type api struct {
	AppId string
	AppSecret string
}
type Response struct {
	Audio string
	SessionId string
	RequestId string
}

func Speak(nameInterface interface{}) {
	data, _ := ioutil.ReadFile("config.yml")
	tencentApi := myconfig{}
	yaml.Unmarshal(data,&tencentApi)
	credential := common.NewCredential(
		tencentApi.Tencent.AppId,
		tencentApi.Tencent.AppSecret,
	)
	//type params struct {
	//	Text string,
	//	SessionId string,
	//
	//}
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "aai.ap-beijing.tencentcloudapi.com"
	client, _ := aai.NewClient(credential, "ap-beijing", cpf)

	request := aai.NewTextToVoiceRequest()
	name :=nameInterface .(string)
	params := `{"Text":"欢迎   `+name+`同学","SessionId":"`+name+`","ModelType":1,"VoiceType":0}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.TextToVoice(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
	responseStruct := Response{}
	json.Unmarshal([]byte(response.ToJsonString()),&responseStruct)
}