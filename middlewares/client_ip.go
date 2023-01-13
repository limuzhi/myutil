/*
 * @PackageName: middlewares
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/3 14:58
 */

package middlewares

import (
	"context"
	"net"
	"strings"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gogf/gf/util/gconv"
)

var GlobalClientIpKey = "x-md-global-clientip"

func ClientIPMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if tr.Kind() == transport.KindHTTP {
					//HTTP
					if info, ok := tr.(*http.Transport); ok {
						// http打印信息
						request := info.Request()
						clientIp := clientIP(request)
						ctx = context.WithValue(ctx, GlobalClientIpKey, clientIp)
						ctx = metadata.AppendToClientContext(ctx, GlobalClientIpKey, clientIp)
					}
				}
			}
			reply, err = handler(ctx, req)
			return
		}
	}
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func clientIP(r *http.Request) string {
	ip := strings.TrimSpace(r.Header.Get("cf-connecting-ip"))
	if ip != "" {
		return ip
	}
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip = strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	ip, _, _ = net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	return ip
}

//获取全局上下文的token 非metadata
func GetClientIPFormContext(ctx context.Context) string {
	val := ctx.Value(GlobalClientIpKey)
	if val != nil {
		return gconv.String(val)
	}
	return ""
}

//metadata传递上下文到GRPC服务获取的全局
func GetClientIPFormGrpcContext(ctx context.Context) string {
	if md, ok := metadata.FromServerContext(ctx); ok {
		val := md.Get(GlobalClientIpKey)
		return gconv.String(val)
	}
	return ""
}
