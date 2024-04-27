package gooqu

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

func (exp expression) Words() *word {
	// id = 1
	equal := newWord("=", false)
	exp.columnName.n(equal).n(&exp.columnValue)
	return &exp.columnName
}
