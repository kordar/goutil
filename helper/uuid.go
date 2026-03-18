package helper

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"time"
)

// UUID V4 基于随机数
func UUID() string {
	return uuid.NewString()
}

// UUIDTime V1 基于时间
func UUIDTime() string {
	u, err := uuid.NewUUID()
	if err != nil {
		return UUID()
	}
	return u.String()
}

// TODO 雪花ID生成，需初始化 InitSnowFlakeNode
var node *snowflake.Node

func InitSnowFlakeNode(startdate string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startdate)
	if err != nil {
		return fmt.Errorf("parse startdate: %w", err)
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		return fmt.Errorf("new snowflake node: %w", err)
	}

	return
}

// GetSnowFlakeId 生成64位的雪花ID
func GetSnowFlakeId() int64 {
	id, _ := GetSnowFlakeIdE()
	return id
}

func GetSnowFlakeIdE() (int64, error) {
	if node == nil {
		return 0, fmt.Errorf("snowflake node not initialized")
	}
	return node.Generate().Int64(), nil
}
