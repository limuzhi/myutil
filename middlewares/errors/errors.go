/*
 * @PackageName: middlewares
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/3 15:10
 */

package errors

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	nethttp "net/http"
)

//错误统一处理
func ErrorEncoder(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
	se := FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal((se))
	if err != nil {
		w.WriteHeader(nethttp.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	if se.Code > 99 && se.Code < 600 {
		w.WriteHeader(se.Code)
	} else {
		w.WriteHeader(500)
	}
	_, _ = w.Write(body)
}
