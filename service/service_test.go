package service

import (
	"fmt"
	"github.com/unknwon/goconfig"
	"github.com/ztino/jd_seckill/common"
	"testing"
)

func TestSendMessage(t *testing.T) {
	confFile := "../conf.ini"
	var err error
	if common.Config, err = goconfig.LoadConfigFile(confFile); err != nil {
		t.Fatal("配置文件不存在")
	}
	if err = SendMessage(common.Config, "测试消息", "测试描述"); err != nil {
		t.Fatal(fmt.Sprintf("测试消息失败%s", err))
	}
}
