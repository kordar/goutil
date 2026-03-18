# goutil

通用工具库集合（配置读取、加解密摘要、UUID/雪花ID、JSON 字段兼容等）。

## 安装

```bash
go get github.com/kordar/goutil
```

## 配置（viper）

```go
package main

import (
	"fmt"
	"github.com/kordar/goutil"
)

func main() {
	if err := goutil.ConfigInitE("./conf.ini"); err != nil {
		panic(err)
	}
	fmt.Println(goutil.GetSectionValue("server", "host"))
}
```

## UUID / 雪花 ID

```go
package main

import (
	"fmt"
	"github.com/kordar/goutil/helper"
)

func main() {
	fmt.Println(helper.UUID())
	fmt.Println(helper.UUIDTime())

	if err := helper.InitSnowFlakeNode("2024-01-01", 1); err != nil {
		panic(err)
	}
	id, err := helper.GetSnowFlakeIdE()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
```

## JWT（HS256）

```go
package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	gjwt "github.com/kordar/goutil/jwt"
)

func main() {
	claims := jwt.MapClaims{"uid": 123}
	token, err := gjwt.GenTokenE(claims, "secret")
	if err != nil {
		panic(err)
	}

	out, err := gjwt.ParseToken(token, "secret")
	if err != nil {
		panic(err)
	}
	fmt.Println(out["uid"])
}
```

## JSON 字段兼容（StrInt）

当后端字段可能是 `"1"` / `1` / `1.0` / `null` 时，可以用 `ext.StrInt` 做兼容解析。

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/kordar/goutil/ext"
)

type Resp struct {
	Code ext.StrInt `json:"code"`
}

func main() {
	var r Resp
	_ = json.Unmarshal([]byte(`{"code":"200"}`), &r)
	fmt.Println(int(r.Code))
}
```

## HTTP（resty 简化封装）

```go
package main

import (
	"github.com/kordar/goutil/resty"
)

func main() {
	var out map[string]any
	_, _ = resty.Get("https://example.com", &out)
}
```
