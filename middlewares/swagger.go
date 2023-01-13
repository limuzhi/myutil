/*
 * @PackageName: middlewares
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/3 14:55
 */

package middlewares

import (
	"github.com/go-kratos/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	"github.com/go-kratos/swagger-api/openapiv2"
	"net/http"
)

func OpenApiSwagger() http.Handler {
	openApiHandler := openapiv2.NewHandler(openapiv2.WithGeneratorOptions(
		generator.UseJSONNamesForFields(true),
	))
	return openApiHandler
}
