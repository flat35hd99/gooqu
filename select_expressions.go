package gooqu

type selectExpressions struct {
	exprs []SelectExpression
}

func (expr selectExpressions) Words() *word {
	if len(expr.exprs) == 0 {
		return newWord("*", false)
	}

	// return root
	if len(expr.exprs) == 1 {
		return expr.exprs[0].selectWords()
	}

	root := expr.exprs[0].selectWords()
	last := root
	for _, exp := range expr.exprs[1:] {
		separator := newWord(",", false)
		last.next = separator
		last = separator

		word := exp.selectWords()
		last.next = word
		last = word
	}
	return root
}

type SelectExpression interface {
	/*
		return the head of word list
	*/
	selectWords() *word
}

type Column struct {
	V string
}

func (c Column) selectWords() *word {
	return newWord(c.V, true)
}

func newSelectExpressions(exps ...SelectExpression) selectExpressions {
	var selectExprs selectExpressions
	selectExprs.exprs = append(selectExprs.exprs, exps...)

	return selectExprs
}
