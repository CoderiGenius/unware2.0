package speaker



import (
"fmt"

"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
aai "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aai/v20180522"
)

func Speak(name string) {

	credential := common.NewCredential(
		"AKIDwb5OlJEsKcjfSLqkIhvBMPIWZUPEDSzL",
		"LXhgV9d4L3b8UlYhfdR4wgLXadEf0Uo7",
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

	params := `{"Text":"欢迎   `+name+`同学","SessionId":"123","ModelType":1,"VoiceType":0}`
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
}