/*
 * @PackageName: db
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 10:36
 */

package db

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

const (
	callBackBeforeName = "cps:before"
	callBackAfterName  = "cps:after"
	startTime          = "_start_time"
)

// TracePlugin defines gorm plugin used to trace sql.
type TracePlugin struct {
	log *log.Helper
}

// Name returns the name of trace plugin.
func (op *TracePlugin) Name() string {
	return "tracePlugin"
}

// Initialize initialize the trace plugin.
func (op *TracePlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前
	_ = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, op.before)
	_ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, op.before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, op.before)
	_ = db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, op.before)
	_ = db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, op.before)
	_ = db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, op.before)

	// 结束后
	_ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, op.after)
	_ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, op.after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, op.after)
	_ = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, op.after)
	_ = db.Callback().Row().After("gorm:row").Register(callBackAfterName, op.after)
	_ = db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, op.after)

	return
}

func (op *TracePlugin) before(db *gorm.DB) {
	db.InstanceSet(startTime, time.Now())
}

func (op *TracePlugin) after(db *gorm.DB) {
	_ts, isExist := db.InstanceGet(startTime)
	if !isExist {
		return
	}

	ts, ok := _ts.(time.Time)
	if !ok {
		return
	}

	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	log.Info(sql)
	log.Infof("sql cost time: %fs", time.Since(ts).Seconds())
}
