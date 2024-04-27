package gooqu

import "fmt"

// もっと込み入ったものにできる
type expression struct {
	columnName  word
	columnValue word
}

func newExpression(columnName, columnValue interface{}) expression {
	return expression{
		columnName:  *newWord(columnName, true),
		columnValue: *newWord(columnValue, true),
	}
}

func (exp expression) String() string {
	return fmt.Sprintf("`%s` = %s", exp.columnName, exp.columnValue)
}

func (exp expression) Words() *word {
	// id = 1
	equal := newWord("=", false)
	exp.columnName.n(equal).n(&exp.columnValue)
	return &exp.columnName
}
