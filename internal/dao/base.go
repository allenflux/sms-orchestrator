package dao

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gvalid"
	"strings"
)

func Init() {
	g.DB("default").SetLogger(g.Log())
	gvalid.RegisterRule("unique_column", RuleUniqueName)
}

// RuleUniqueName 检查字段是否唯一
func RuleUniqueName(ctx context.Context, in gvalid.RuleFuncInput) error {
	ruleStr := strings.Split(in.Rule, ":")
	if len(ruleStr) != 2 {
		return fmt.Errorf("unique_column规则定义错误")
	}
	ruleValue := strings.Split(ruleStr[1], ".")
	if len(ruleValue) != 2 {
		return fmt.Errorf("unique_column规则内容定义错误")
	}
	table := ruleValue[0]
	column := ruleValue[1]
	count, err := g.DB().Model(table).Where(column, in.Value).Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if count > 0 {
		return fmt.Errorf("%s已被占用，请重新输入", in.Field)
	}
	return nil
}
