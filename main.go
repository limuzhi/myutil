package main

import (
	"fmt"
	"myutils/videos"
)

func main() {
	url := "https://v.douyin.com/BPfcJX1/"
	adapter := videos.GetShortVideoAdapter(url)
	info, err := adapter.GetShortVideoInfo(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	info.Info()
}
