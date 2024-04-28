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
