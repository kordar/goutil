# goutil
通用工具库（单仓库单 package），目前包含：

- 字符串：驼峰/下划线互转
- ID：UUID(v4/v1)、Snowflake（需初始化）
- 摘要：SHA1/MD5/CRC32
- DB 连接池：简单的句柄注册与管理（并发安全）

## 安装

```bash
go get github.com/kordar/goutil
```

## 字符串

```go
package main

import (
	"fmt"
	"github.com/kordar/goutil"
)

func main() {
	fmt.Println(goutil.SnakeString("UserName"))   // user_name
	fmt.Println(goutil.CamelString("user_name")) // UserName
}
```

## UUID / Snowflake

```go
package main

import (
	"fmt"
	"github.com/kordar/goutil"
)

func main() {
	fmt.Println(goutil.UUID())
	fmt.Println(goutil.UUIDTime())

	if err := goutil.InitSnowFlakeNode("2024-01-01", 1); err != nil {
		panic(err)
	}
	id, err := goutil.GetSnowFlakeIdE()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
```

## 摘要

```go
package main

import (
	"fmt"
	"github.com/kordar/goutil"
)

func main() {
	fmt.Println(goutil.SHA1("hello"))
	fmt.Println(goutil.Md5V("hello"))
	fmt.Println(goutil.CRC32Str("hello"))
}
```

## DB 连接池

你需要实现 `DBItem` 接口，把连接句柄注册到 pool 后通过 name 获取。

```go
package main

import (
	"github.com/kordar/goutil"
)

type Conn struct{}

type Item struct {
	name string
	conn *Conn
}

func (i *Item) GetName() string        { return i.name }
func (i *Item) GetInstance() interface{} { return i.conn }
func (i *Item) Close() error           { return nil }

func main() {
	pool := goutil.GetDbPool()
	_ = pool.Add(&Item{name: "default", conn: &Conn{}})

	_ = pool.Handle("default")
	_ = pool.Instance("default")
}
```
