/*
 * @PackageName: response
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/3 15:13
 */

package response

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/transport/http"
	nethttp "net/http"
	"time"
)

func ResponseEncoder(w nethttp.ResponseWriter, r *nethttp.Request, v interface{}) error {
	if v == nil {
		return nil
	}
	if rd, ok := v.(http.Redirector); ok {
		url, code := rd.Redirect()
		nethttp.Redirect(w, r, url, code)
		return nil
	}
	codec, _ := http.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(v)
	if err != nil {
		return err
	}
	var outRes = make(map[string]interface{})
	err = json.Unmarshal(data, &outRes)
	if err != nil {
		return err
	}
	type response struct {
		Code     int         `json:"code"`
		Data     interface{} `json:"data"`
		Messsage string      `json:"messsage"`
		Ts       int64       `json:"ts"`
	}
	reply := &response{
		Code:     200,
		Data:     outRes,
		Ts:       time.Now().UTC().Unix(),
		Messsage: "success",
	}
	outCodec := encoding.GetCodec("json")
	outData, err := outCodec.Marshal(reply)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(outData)
	if err != nil {
		return err
	}
	return nil
}
