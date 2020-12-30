package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
)

func WithArgHandler(writer http.ResponseWriter, request *http.Request) {
	value := getAuth()
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(writer, string(value))
}

func getAuth() string {
	key := os.Getenv("key")
	pass := os.Getenv("pass")
	client, err := vod.NewClientWithAccessKey("cn-hangzhou", key, pass)

	request := vod.CreateGetVideoPlayAuthRequest()
	request.Scheme = "https"

	request.VideoId = "e06498bc828946fdb99213a8794a2bae"

	response, err := client.GetVideoPlayAuth(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)

	PlayAuth := response.PlayAuth

	return PlayAuth
}
