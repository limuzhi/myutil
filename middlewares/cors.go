/*
 * @PackageName: middlewares
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/3 15:05
 */

package middlewares

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

func CorsMiddleware() http.ServerOption {
	return http.Filter(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"DNT", "X-Mx-ReqToken", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Authorization", "udid", "appkey", "version", "authenticated", "cookie", "token"}),
		handlers.ExposedHeaders([]string{"DNT", "X-Mx-ReqToken", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Authorization", "udid", "appkey", "version", "authenticated", "cookie", "token"}),
		handlers.OptionStatusCode(204),
	))
}
