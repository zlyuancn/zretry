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
    "github.com/zlyuancn/zretry"
    "time"
)

func main() {
    retry := zretry.NewRetry(zretry.WithInterval(1e9), zretry.WithAttemptCount(3))
    err := retry.Do(func() error {
        fmt.Println(time.Now())
        return errors.New("故意报错")
    })
    if err != nil {
        panic(err)
    }
}
```
