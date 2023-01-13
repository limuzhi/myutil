/*
 * @PackageName: middlewares
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/3 15:01
 */

package middlewares

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"time"
)

func OperlogMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		// 开始时间
		startTime := time.Now()
		if tr, ok := transport.FromServerContext(ctx); ok {
			if tr.Kind() == transport.KindHTTP {
				//HTTP
				if info, ok := tr.(*http.Transport); ok {
					// http打印信息
					request := info.Request()
					fmt.Println(request)
				}
			}
		}
		//TODO 业务处理
		//上文请求
		reply, err = handler(ctx, req)
		//下文输出体
		//TODO 业务处理
		endTime := time.Now()
		timeCost := endTime.Sub(startTime).String()
		fmt.Println(timeCost)
		return
	}
}
