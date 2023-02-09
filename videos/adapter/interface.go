package adapter

import "fmt"

type IVideosInterface interface {
	GetShortVideoInfo(url string) (*ShortVideoInfoResponse, error)
}

//因为平台众多，所以返回的参数不固定，但 title, cover, url 一定会有

type ShortVideoInfoResponse struct {
	Title    string     `json:"title"`  //视频标题
	Author   string     `json:"author"` //作者
	Uid      string     `json:"uid"`
	Avatar   string     `json:"avatar"`   //作者头像
	Like     int        `json:"like"`     //点赞数
	Time     int        `json:"time"`     //发布时间
	Cover    string     `json:"cover"`    //封面图
	Url      string     `json:"url"`      //视频地址
	MusicUrl string     `json:"musicUrl"` //音乐地址
	Music    *MusicInfo `json:"music"`    //音乐信息
}

type MusicInfo struct {
	Author string `json:"author"` //音乐作者
	Avatar string `json:"avatar"` //音乐作者头像
	Url    string `json:"url"`    //音乐地址
}

func (r *ShortVideoInfoResponse) Info() {
	fmt.Println("title:", r.Title, "--cover:", r.Cover, "--url:", r.Url)
}
