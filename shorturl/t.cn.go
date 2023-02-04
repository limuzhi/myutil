package shorturl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type TcnUrl struct {
	Cookie    string
	Source    string
	UserId    string
	UserIdInt int64
	Clientid  string
	HeadrMap  map[string]string
}

func NewTcnUrl(cookie string) *TcnUrl {
	return &TcnUrl{
		Cookie:    cookie,
		Source:    "209678993",
		UserId:    "2735327001", //微博安全中心ID
		UserIdInt: 2735327001,
		HeadrMap: map[string]string{
			"Host":            "api.weibo.com",
			"Referer":         "https://api.weibo.com/chat/",
			"Connection":      "keep-alive",
			"Pragma":          "no-cache",
			"Cache-Control":   "no-cache",
			"Accept":          "application/json, text/plain, */*",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.122 Safari/537.36",
			"Sec-Fetch-Site":  "same-origin",
			"Sec-Fetch-Mode":  "cors",
			"Sec-Fetch-Dest":  "empty",
			"Accept-Encoding": "gzip, deflate, br",
			"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		},
	}
}

type ReturnMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type TcnReturn struct {
	Contacts []ContactInfo `json:"contacts"`
}

type ContactInfo struct {
	Message struct {
		ExtText struct {
			SourceMsgId   int    `json:"source_msg_id"`
			MsgSource     int    `json:"msg_source"`
			AutoReply     bool   `json:"autoReply,omitempty"`
			AutoreplyType string `json:"autoreply_type,omitempty"`
			SendFrom      string `json:"send_from,omitempty"`
			Clientid      string `json:"clientid,omitempty"`
		} `json:"ext_text,omitempty"`
	} `json:"message"`
	User struct {
		Idstr string `json:"idstr"`
	} `json:"user"`
}

func (u *TcnUrl) getContactsJson() string {
	url := fmt.Sprintf("https://api.weibo.com/webim/2/direct_messages/contacts.json?special_source=3&add_virtual_user=3,4&is_include_group=0&need_back=0,0&count=50&source=%s&t=%d",
		u.Source, time.Now().UnixMilli())
	res, err := u.curlGet(url, "GET")
	if err != nil {
		return u.returnMsg(-1, err.Error(), "获取失败")
	}
	data := &TcnReturn{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return u.returnMsg(-1, err.Error(), "解析错误")
	}
	info := ContactInfo{}
	if len(data.Contacts) > 0 {
		for _, v := range data.Contacts {
			if v.User.Idstr == u.UserId {
				info = v
			}
		}
	}
	if info.User.Idstr == "" {
		return u.returnMsg(-2, "你未关注此人或未给他发消息", "empty")
	}
	//if info.Message.ExtText.Clientid == "" {
	//	return u.returnMsg(-3, "与此人没有消息往来，无法获取会话id", "empty")
	//}
	u.Clientid = info.Message.ExtText.Clientid
	_ = FilePutContents("1.txt", u.Clientid, 0644)
	return u.Clientid
}

type sendMsgRespone struct {
	Id          int64  `json:"id"`
	Text        string `json:"text"`
	RecipientId int64  `json:"recipient_id"`
}

func (u *TcnUrl) sendMsg(text string) string {
	//clientId, err := FileGetContents("contacts.txt")
	//u.Clientid = clientId
	//if err != nil || u.Clientid == "" {
	//	u.Clientid = u.getContactsJson()
	//}
	apiUrl := "https://api.weibo.com/webim/2/direct_messages/new.json"
	q := url.Values{}
	q.Set("text", text)
	q.Set("uid", u.UserId)
	q.Set("is_encoded", "0")
	q.Set("decodetime", "1")
	q.Set("source", u.Source)
	//q.Set("extensions", "{\"clientid\":\""+u.Clientid+"\"}")
	res, err := u.curlPost(apiUrl, "POST", q.Encode())
	if err != nil {
		return u.returnMsg(-1, err.Error(), "请求失败")
	}

	sendRespone := &sendMsgRespone{}
	err = json.Unmarshal(res, sendRespone)
	if err != nil {
		return u.returnMsg(-1, err.Error(), "解析错误")
	}
	if sendRespone.Id == 0 || sendRespone.RecipientId != u.UserIdInt {
		return u.returnMsg(-1, "发送失败", "empty")
	}
	returnData := make(map[string]string)
	returnData["short_url"] = sendRespone.Text
	returnData["long_url"] = text

	return u.returnMsg(0, "生成成功，消息撤回成功", returnData)
}

func (u *TcnUrl) recallMsg(id int64) bool {
	apiUrl := "https://api.weibo.com/webim/2/direct_messages/recall.json"
	q := url.Values{}
	q.Set("id", strconv.FormatInt(id, 10))
	q.Set("source", u.Source)
	res, err := u.curlPost(apiUrl, "POST", q.Encode())
	if err != nil {
		return false
	}
	sendRespone := &sendMsgRespone{}
	err = json.Unmarshal(res, sendRespone)
	if err != nil {
		return false
	}
	if sendRespone.Id == 0 || sendRespone.Text == "" {
		return false
	}
	return true
}

func (u *TcnUrl) returnMsg(code int, msg string, data interface{}) string {
	out := &ReturnMsg{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	b, _ := json.Marshal(out)
	return string(b)
}

func (u *TcnUrl) curlGet(url, method string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", u.Cookie)
	for k, v := range u.HeadrMap {
		req.Header.Set(k, v)
	}
	req.Header.Set("Accept-Encoding", "application/json; charset=UTF-8")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (u *TcnUrl) curlPost(url, method, params string) ([]byte, error) {
	payload := strings.NewReader(params)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", u.Cookie)
	for k, v := range u.HeadrMap {
		req.Header.Set(k, v)
	}
	req.Header.Set("Accept-Encoding", "application/json; charset=UTF-8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(body))
	return body, nil
}

// FileGetContents 把整个文件读入一个字符串中
func FileGetContents(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	return string(data), err
}

// FilePutContents 把一个字符串写入文件中
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), mode)
}
