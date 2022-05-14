# zretry
###### 重试包

## 获得zretry
`go get -u github.com/zlyuancn/zretry`

## 使用zretry

```go
package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/zlyuancn/zretry"
)

func main() {
	retry := zretry.NewRetry(
		zretry.WithAttemptCount(3),       // 尝试次数
		zretry.WithInterval(time.Second), // 重试间隔时间
		zretry.WithErrCallback(func(nowAttemptCount, remainCount int, err error) {
			fmt.Printf("当前尝试次数: %v, 剩余次数: %v, err: %v\n", nowAttemptCount, remainCount, err)
		}),
	)
	err := retry.Do(func() error {
		fmt.Println(time.Now())
		return errors.New("故意报错")
	})
	if err != nil {
		panic(fmt.Errorf("最终错误: %v", err))
	}
}
```

```go
package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/zlyuancn/zretry"
)

func main() {
	err := zretry.DoRetry(
		3,           // 尝试次数
		time.Second, // 重试间隔时间
		func() error { // 执行函数
			fmt.Println(time.Now())
			return errors.New("故意报错")
		},
		func(nowAttemptCount, remainCount int, err error) { // 错误回调
			fmt.Printf("当前尝试次数: %v, 剩余次数: %v, err: %v\n", nowAttemptCount, remainCount, err)
		},
	)
	if err != nil {
		panic(fmt.Errorf("最终错误: %v", err))
	}
}
```
