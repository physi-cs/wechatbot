package operate

import (
	"fmt"
	"time"

	"github.com/869413421/wechatbot/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ces "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v1/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v1/region"
)

func create_events(error_msg string) {
    ak := config.LoadConfig().AK
    sk := config.LoadConfig().SK

    auth := basic.NewCredentialsBuilder().
        WithAk(ak).
        WithSk(sk).
        Build()

    client := ces.NewCesClient(
        ces.CesClientBuilder().
            WithRegion(region.ValueOf("cn-east-3")).
            WithCredential(auth).
            Build())

    request := &model.CreateEventsRequest{}
	contentDetail:= error_msg
	detailBody := &model.EventItemDetail{
		Content: &contentDetail,
	}
	var listBodybody = []model.EventItem{
        {
            EventName: "wxbot",
            EventSource: "hw.wxbot",
            Time: int64(time.Now().UnixMilli()),
            Detail: detailBody,
        },
    }
	request.Body = &listBodybody
	response, err := client.CreateEvents(request)
	if err == nil {
        fmt.Printf("%+v\n", response)
    } else {
        fmt.Println(err)
    }
}