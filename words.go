package gooqu

import (
	"fmt"
	"strconv"
	"strings"
)

type word struct {
	next    *word
	value   string
	escaped bool
}

type worder interface {
	/*
		return the head of word list
	*/
	Words() *word
}

func newWord(value interface{}, escaped bool) *word {
	var v string
	switch value := value.(type) {
	case string:
		v = value
	case int:
		v = strconv.Itoa(value)
	case []int:
		// FIXME: 全然ちゃう
		v = fmt.Sprintf("%v", value)
	case []string:
		// FIXME: 全然ちゃう
		v = strings.Join(value, "")
	default:
		// FIXME: warnかerror入れたほうがいいかも判断
		v = fmt.Sprintf("%v", value)
	}
	return &word{
		value:   v,
		escaped: escaped,
	}
}

func (w *word) n(next worder) *word {
	w.next = (next).Words()
	return w.last()
}

func (w *word) last() *word {
	if w.next == nil {
		return w
	} else {
		return w.next.last()
	}
}

func (w word) String() string {
	var v string
	var separator string
	if w.value != "" {
		separator = " "
	}

	if w.escaped {
		v = "`" + w.value + "`"
	} else {
		v = w.value
	}

	if w.next != nil {
		succession := w.next.String()
		return v + separator + succession
	} else {
		return v
	}
}

func (w *word) Words() *word {
	return w
}
