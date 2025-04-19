package goutil

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"time"
)

// UUID V4 基于随机数
func UUID() string {
	return uuid.New().String()
}

// UUIDTime V1 基于时间
func UUIDTime() string {
	uuid1, uuid1err := uuid.NewUUID()
	if uuid1err != nil {
		return UUID()
	}
	return uuid1.String()
}

// TODO 雪花ID生成，需初始化 InitSnowFlakeNode
var node *snowflake.Node

func InitSnowFlakeNode(startdate string, machineID int64) (err error) {
	var st time.Time
	// 格式化1月2号下午3时4分5秒2006年
	st, err = time.Parse("2006-01-02", startdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

// GetSnowFlakeId 生成64位的雪花ID
func GetSnowFlakeId() int64 {
	return node.Generate().Int64()
}
