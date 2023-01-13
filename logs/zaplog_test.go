/*
 * @PackageName: logs
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/3 14:35
 */

package logs

import (
	"github.com/go-kratos/kratos/v2/log"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := log.With(NewLogger(WithInfoFilePath("info")),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)

	aa := log.NewHelper(log.With(logger, "module", "data.af_affiliate"))
	aa.Info("aa")

}
