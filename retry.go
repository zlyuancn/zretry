/*
-------------------------------------------------
   Author :       zlyuan
   date：         2019/5/10
   Description :
-------------------------------------------------
*/

package zretry

import (
	"time"
)

type Retry struct {
	attemptCount int
	interval     time.Duration
	errCallFun   ErrCallbackFunc
}

type Option func(*Retry)

type DoFun func() error

/*回调函数
  nowAttemptCount 当前尝试次数
  remainCount 剩余次数
  err 错误
*/
type ErrCallbackFunc func(nowAttemptCount, remainCount int, err error)

// 创建一个重试器
func NewRetry(options ...Option) *Retry {
	r := &Retry{}
	for _, o := range options {
		o(r)
	}
	return r
}

// 设置间隔时间
func WithInterval(interval time.Duration) Option {
	return func(retry *Retry) {
		retry.interval = interval
	}
}

// 设置最大尝试次数, 0为不限次数
func WithAttemptCount(attemptCount int) Option {
	return func(retry *Retry) {
		retry.attemptCount = attemptCount
	}
}

// 设置错误回调函数, 每次执行时有任何错误都会报告给该函数
func WithErrCallback(errfn ErrCallbackFunc) Option {
	return func(retry *Retry) {
		retry.errCallFun = errfn
	}
}

// 执行一个函数
func (m *Retry) Do(f DoFun) (err error) {
	return DoRetry(m.attemptCount, m.interval, f, m.errCallFun)
}

// 执行一个函数
func DoRetry(attemptCount int, interval time.Duration, f DoFun, errCallFun ErrCallbackFunc) (err error) {
	nowAttemptCount := 0
	for {
		nowAttemptCount++

		err = f()
		if err == nil {
			return
		}
		if errCallFun != nil {
			errCallFun(nowAttemptCount, attemptCount-nowAttemptCount, err)
		}

		if nowAttemptCount >= attemptCount {
			break
		}

		if interval > 0 {
			time.Sleep(interval)
		}
	}
	return
}
