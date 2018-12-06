package speaker



import (
"fmt"

"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
aai "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aai/v20180522"
	"io/ioutil"
	"encoding/json"
	"../channel"
	"../producer"

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
	data, _ := ioutil.ReadFile("D:\\go\\workplace\\unware\\config\\config.yml")
	//tencentApi := myconfig{}
	fmt.Println(string(data))
	//err0 := yaml.Unmarshal(data, &tencentApi)
	//if err0!=nil{
	//	fmt.Println(err0)
	//}


	tencent:= myconfig{}
	errJson := json.Unmarshal(data,&tencent)
	if errJson!=nil{
		fmt.Println(errJson)
	}
	//fmt.Println(tencent.Tencent.AppId)
	credential := common.NewCredential(
		tencent.Tencent.AppId,
		tencent.Tencent.AppSecret,
		//"AKIDwb5OlJEsKcjfSLqkIhvBMPIWZUPEDSzL",
		//"LXhgV9d4L3b8UlYhfdR4wgLXadEf0Uo7",

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
	name := nameInterface.(string)
	params := `{"Text":"欢迎   ` + name + `同学","SessionId":"` + name + `","ModelType":2,"VoiceType":0}`
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
	responseStruct := channel.ResponseJson{}
	erro := json.Unmarshal([]byte(response.ToJsonString()), &responseStruct)
	fmt.Println("after json", responseStruct)
	if erro != nil {
		panic(erro)
	} else {
		producer.Deal(responseStruct)
	}
}
