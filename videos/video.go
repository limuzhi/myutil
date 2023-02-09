package videos

import (
	"myutils/videos/adapter"
	"strings"
)

func GetShortVideoAdapter(url string) adapter.IVideosInterface {
	if strings.Contains(url, "douyin") {
		return adapter.DouyinAdapter{}
	}
	return nil
}
