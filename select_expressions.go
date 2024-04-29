package gooqu

import (
	"fmt"
	"strings"
)

type selectExpressions struct {
	exprs []SelectExpression
}

func (exprs selectExpressions) String() string {
	if len(exprs.exprs) == 0 {
		return "*"
	}

	hoge := []string{}
	for _, expr := range exprs.exprs {
		hoge = append(hoge, expr.String())
	}
	return strings.Join(hoge, ", ")
}

type SelectExpression interface {
	String() string
}

type Column struct {
	V string
}

func (c Column) String() string {
	return fmt.Sprintf("`%s`", c.V)
}

func newSelectExpressions(exps ...SelectExpression) selectExpressions {
	var selectExprs selectExpressions
	selectExprs.exprs = append(selectExprs.exprs, exps...)

	return selectExprs
}

type Count struct {
	columnName string
	aliasName  string
}

func (c Count) String() string {
	if c.aliasName == "" {
		return fmt.Sprintf("COUNT(`%s`)", c.columnName)
	} else {
		return fmt.Sprintf("COUNT(`%s`) AS `%s`", c.columnName, c.aliasName)
	}
}

func (c *Count) As(aliasName string) *Count {
	c.aliasName = aliasName
	return c
}

func COUNT(columnName string) *Count {
	return &Count{columnName: columnName}
}
