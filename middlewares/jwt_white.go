/*
 * @PackageName: middlewares
 * @Description: 演示DEMO
 * @Author: limuzhi
 * @Date: 2022/12/3 14:50
 */

package middlewares

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

//jwt白名单 ---根据protoc的

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	//TODO
	whiteList["/admin.v1.AdminService/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func NewWhiteListMatcher02() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/v1/public/login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if tr, ok := transport.FromServerContext(ctx); ok {
			if tr.Kind() == transport.KindHTTP {
				if info, ok := tr.(*http.Transport); ok {
					// http打印信息
					operation = info.PathTemplate()
				}
			}
		}
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
