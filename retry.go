/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/5/10
   Description :
-------------------------------------------------
*/

package zretry

import (
    "time"
)

type Retry struct {
    interval     time.Duration
    attemptCount int32
    errCallFun   ErrCallbackFunc
}

type Option func(*Retry)

type DoFun func() error
type ErrCallbackFunc func(err error)

//创建一个重试器
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
func WithAttemptCount(attemptCount int32) Option {
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
    return DoRetry(f, m.interval, m.attemptCount, m.errCallFun)
}

//执行一个函数
func DoRetry(f DoFun, interval time.Duration, attemptCount int32, errCallFun ErrCallbackFunc) (err error) {
    for {
        err = f()
        if err == nil {
            return
        } else if errCallFun != nil {
            errCallFun(err)
        }

        attemptCount--
        if attemptCount == 0 {
            break
        }

        if interval > 0 {
            time.Sleep(interval)
        }
    }
    return
}
