/*
 * @PackageName: tick
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 14:10
 */

package tick

import "time"

//即时执行
func ImmediatelyTick(f func(), d time.Duration) {
	t := time.NewTicker(d)
	f()
	go func() {
		for {
			<-t.C
			f()
		}
		t.Stop()
	}()
}

//延迟执行
func DelayTick(f func(), d time.Duration) {
	t := time.NewTicker(d)
	go func() {
		for {
			<-t.C
			f()
		}
		t.Stop()
	}()
}
