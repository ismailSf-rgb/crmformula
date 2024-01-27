package crmformula

import (
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
)

const (
	ADD        string = "ADD"
	OR         string = "OR"
	AND        string = "AND"
	IF         string = "IF"
	TEXT       string = "TEXT"
	LEFT       string = "LEFT"
	RIGHT      string = "RIGHT"
	MULT       string = "MULT"
	TRUE       string = "TRUE"
	FALSE      string = "FALSE"
	NOT        string = "NOT"
	DIV        string = "DIV"
	MOD        string = "MOD"
	ISBLANK    string = "ISBLANK"
	BLANKVALUE string = "BLANKVALUE"
	CONTAINS   string = "CONTAINS"
	EXP        string = "EXP"
	COS        string = "COS"
	SIN        string = "SIN"
	LEN        string = "LEN"
	MAX        string = "MAX"
	MIN        string = "MIN"
	MORE_EQ    string = "MORE_EQ"
	LESS_EQ    string = "LESS_EQ"
	MORE       string = "MORE"
	LESS       string = "LESS"
	EQUAL      string = "EQUAL"
	NOW        string = "NOW"
	CEILING    string = "CEILING"
	ROUND      string = "ROUND"
	SQRT       string = "SQRT"
	LOG        string = "LOG"
	LOG10      string = "LOG10"
	ABS        string = "ABS"
	FIND       string = "FIND"
	TRIM       string = "TRIM"
	TODAY      string = "TODAY"
	ADDMONTH   string = "ADDMONTH"
	BEGINS     string = "BEGINS"
	CONCAT     string = "CONCAT"
	BR         string = "BR"
	YEAR       string = "YEAR"
	MONTH      string = "MONTH"
	DAY        string = "DAY"
	WEEKDAY    string = "WEEKDAY"
	HOUR       string = "HOUR"
	MINUTE     string = "MINUTE"
	FLOOR      string = "FLOOR"
	UPPER      string = "UPPER"
	LOWER      string = "LOWER"
	REGEX      string = "REGEX"
)

var FormOperators []interface{} = []interface{}{ADDMONTH, ADD, OR, AND, IF, TEXT, LEFT, RIGHT, MULT, TRUE, FALSE, NOT, DIV, MOD, ISBLANK, BLANKVALUE, CONTAINS, EXP, COS, SIN, LEN, MAX, MIN, MORE_EQ, LESS_EQ, MORE, LESS, EQUAL, NOW, CEILING, ROUND, SQRT, LOG, LOG10, ABS, FIND, TRIM, TODAY, BEGINS, CONCAT, BR, YEAR, MONTH, DAY, WEEKDAY, HOUR, MINUTE, FLOOR, UPPER, LOWER, REGEX}

func contains(s []interface{}, str interface{}) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

type BinaryNode struct {
	left   *BinaryNode
	right  *BinaryNode
	value  interface{}
	parent *BinaryNode
}

type BinaryTree struct {
	root *BinaryNode
	size int
}

func (t *BinaryTree) insert(val interface{}, op string) *BinaryTree {
	if t.root == nil && val == op {
		t.root = &BinaryNode{value: val, left: nil, right: nil, parent: nil}
	} else {
		t.root.insert(val, op)
	}
	return t
}

func (t *BinaryTree) reduceBr() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceBr()
	}
	return t
}

func (n *BinaryNode) reduceBr() {

	if n.right == nil && n.left == nil {

		str := "\n"

		n.value = str
		if n.parent != nil {
			n = n.parent
			n.reduceBr()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceConcat() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceConcat()
	}
	return t
}

func (n *BinaryNode) reduceConcat() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceConcat()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceConcat()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceConcat()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceConcat()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceConcat()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceConcat()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		vToFind := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)
		var result string
		if vToFind.Kind() == reflect.String && vright.Kind() == reflect.String {
			var val string
			v := n.value.(string)
			val = v

			if v, ok := n.right.value.(string); ok {
				result = fmt.Sprintf("%s%s", val, v)
			} else {
				panic("String Exception : Not a String")
			}
			n.value = result
			n.right = nil
		} else {
			panic("String Exception : Not a String")
		}
	}

}

func (t *BinaryTree) reduceBegins() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceBegins()
	}
	return t
}

func (n *BinaryNode) reduceBegins() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceBegins()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceBegins()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceBegins()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceBegins()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceBegins()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceBegins()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		if v.Kind() == reflect.String {
			val := false
			val = strings.HasPrefix(n.value.(string), n.right.value.(string))
			n.value = val
			n.right = nil
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceAddMonth() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceAddMonth()
	}
	return t
}

func (n *BinaryNode) reduceAddMonth() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceAddMonth()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceAddMonth()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceAddMonth()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceAddMonth()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceAddMonth()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceAddMonth()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		vToConvert := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)
		var result string
		if vToConvert.Kind() == reflect.String && vright.Kind() == reflect.Int {
			var val string
			v := n.value.(string)
			val = v

			if v, ok := n.right.value.(int); ok {
				val = strings.ReplaceAll(val, " ", "")
				res1, _ := time.Parse(Layout, val)
				endDate := res1.AddDate(0, v, 0)
				result = endDate.Format(DisplayedLayout)
			} else {
				panic("Formula Exception : second parameter must be an integer")
			}
			n.value = result
			n.right = nil
		} else {
			panic("Formula Exception : wrong parameters")
		}
	}

}

func (t *BinaryTree) reduceUpper() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceUpper()
	}
	return t
}

func (n *BinaryNode) reduceUpper() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceUpper()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceUpper()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceUpper()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceUpper()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceUpper()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceUpper()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {

		n.value = strings.ToUpper(n.left.value.(string))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceUpper()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceLower() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceLower()
	}
	return t
}

func (n *BinaryNode) reduceLower() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceLower()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceLower()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceLower()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceLower()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceLower()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceLower()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {

		n.value = strings.ToLower(n.left.value.(string))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceLower()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceYear() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceYear()
	}
	return t
}

func (n *BinaryNode) reduceYear() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceYear()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceYear()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceYear()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceYear()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceYear()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceYear()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		val := strings.ReplaceAll(n.left.value.(string), " ", "")
		date, _ := time.Parse(Layout, val)
		n.value = date.Year()
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceYear()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceMonth() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceMonth()
	}
	return t
}

func (n *BinaryNode) reduceMonth() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceMonth()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceMonth()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceMonth()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceMonth()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceMonth()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceMonth()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		val := strings.ReplaceAll(n.left.value.(string), " ", "")
		date, _ := time.Parse(Layout, val)
		n.value = int(date.Month())
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceMonth()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceMonthDay() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceMonthDay()
	}
	return t
}

func (n *BinaryNode) reduceMonthDay() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceMonthDay()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceMonthDay()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceMonthDay()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceMonthDay()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceMonthDay()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceMonthDay()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		val := strings.ReplaceAll(n.left.value.(string), " ", "")
		date, _ := time.Parse(Layout, val)
		n.value = date.Day()
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceMonthDay()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceWeekDay() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceWeekDay()
	}
	return t
}

func (n *BinaryNode) reduceWeekDay() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceWeekDay()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceWeekDay()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceWeekDay()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceWeekDay()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceWeekDay()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceWeekDay()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		val := strings.ReplaceAll(n.left.value.(string), " ", "")
		date, _ := time.Parse(Layout, val)
		n.value = int(date.Weekday())
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceWeekDay()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceHour() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceHour()
	}
	return t
}

func (n *BinaryNode) reduceHour() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceHour()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceHour()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceHour()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceHour()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceHour()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceHour()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		val := strings.ReplaceAll(n.left.value.(string), " ", "")
		date, _ := time.Parse(Layout, val)
		n.value = date.Hour()
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceHour()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceMinute() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceMinute()
	}
	return t
}

func (n *BinaryNode) reduceMinute() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceMinute()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceMinute()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceMinute()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceMinute()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceMinute()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceMinute()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		val := strings.ReplaceAll(n.left.value.(string), " ", "")
		date, _ := time.Parse(Layout, val)
		n.value = date.Minute()
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceMinute()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceToday() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceToday()
	}
	return t
}

func (n *BinaryNode) reduceToday() {

	if n.right == nil && n.left == nil {
		now := time.Now()
		str := now.Format(Layout)

		n.value = str
		if n.parent != nil {
			n = n.parent
			n.reduceToday()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceDay() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceDay()
	}
	return t
}

func (n *BinaryNode) reduceDay() {

	if n.right == nil && n.left == nil {
		now := time.Now()
		today := now.Truncate(24 * time.Hour)
		str := today.Format(Layout)
		dateSlice := strings.Split(str, "at")
		n.value = dateSlice[0] + "at12:00am(UTC)"
		if n.parent != nil {
			n = n.parent
			n.reduceDay()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceEqual() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceEqual()
	}
	return t
}

func (n *BinaryNode) reduceEqual() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceEqual()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceEqual()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceEqual()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceEqual()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceEqual()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceEqual()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)
		var result bool
		if v.Kind() == reflect.Float64 || vright.Kind() == reflect.Float64 {
			var val float64
			if v, ok := n.value.(int); ok {
				val = float64(v)
			} else {
				val = n.value.(float64)
			}
			if v, ok := n.right.value.(int); ok {
				result = (float64(v) == val)
			} else {
				result = (n.right.value.(float64) == val)
			}
			n.value = result
			n.right = nil
		} else if v.Kind() == reflect.Int {
			var val2 int
			val2 = n.value.(int)
			result = (n.right.value.(int) == val2)
			n.value = result
			n.right = nil
		} else if v.Kind() == reflect.String && vright.Kind() == reflect.String {
			var trimedval1 string = strings.TrimSpace(n.value.(string))
			var trimedval2 string = strings.TrimSpace(n.right.value.(string))
			r := strings.NewReplacer(" ", "")
			trimedval1 = r.Replace(trimedval1)
			trimedval2 = r.Replace(trimedval2)
			val2, err0 := time.Parse(Layout, trimedval1)
			val3, err1 := time.Parse(Layout, trimedval2)
			if err1 == nil && err0 == nil {
				n.value = val2.Equal(val3)
			} else {
				n.value = n.value.(string) == n.right.value.(string)
			}
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceLess() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceLess()
	}
	return t
}

func (n *BinaryNode) reduceLess() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceLess()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceLess()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceLess()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceLess()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceLess()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceLess()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)
		var result bool
		if v.Kind() == reflect.Float64 || vright.Kind() == reflect.Float64 {
			var val float64
			if v, ok := n.value.(int); ok {
				val = float64(v)
			} else {
				val = n.value.(float64)
			}
			if v, ok := n.right.value.(int); ok {
				result = (float64(v) > val)
			} else {
				result = (n.right.value.(float64) > val)
			}
			n.value = result
			n.right = nil
		} else if v.Kind() == reflect.Int {
			var val2 int
			val2 = n.value.(int)
			result = (n.right.value.(int) > val2)
			n.value = result
			n.right = nil
		} else if v.Kind() == reflect.String && vright.Kind() == reflect.String {
			var trimedval1 string = strings.TrimSpace(n.value.(string))
			var trimedval2 string = strings.TrimSpace(n.right.value.(string))
			r := strings.NewReplacer(" ", "")
			trimedval1 = r.Replace(trimedval1)
			trimedval2 = r.Replace(trimedval2)
			val2, err0 := time.Parse(Layout, trimedval1)
			val3, err1 := time.Parse(Layout, trimedval2)
			if err1 == nil && err0 == nil {
				n.value = val2.Before(val3)
			}
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceFind() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceFind()
	}
	return t
}

func (n *BinaryNode) reduceFind() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceFind()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceFind()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceFind()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceFind()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceFind()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceFind()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		vToFind := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)
		var result int
		if vToFind.Kind() == reflect.String && vright.Kind() == reflect.String {
			var val string
			v := n.value.(string)
			val = v

			if v, ok := n.right.value.(string); ok {
				result = strings.Index(v, val)
			} else {
				panic("String Exception : Not a String")
			}
			n.value = result
			n.right = nil
		} else {
			panic("String Exception : Not a String")
		}
	}

}

func (t *BinaryTree) reduceLessEq() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceLessEq()
	}
	return t
}

func (n *BinaryNode) reduceLessEq() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceLessEq()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceLessEq()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceLessEq()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceLessEq()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceLessEq()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceLessEq()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)
		var result bool
		if v.Kind() == reflect.Float64 || vright.Kind() == reflect.Float64 {
			var val float64
			if v, ok := n.value.(int); ok {
				val = float64(v)
			} else {
				val = n.value.(float64)
			}
			if v, ok := n.right.value.(int); ok {
				result = (float64(v) >= val)
			} else {
				result = (n.right.value.(float64) >= val)
			}
			n.value = result
			n.right = nil
		} else if v.Kind() == reflect.Int {
			var val2 int
			val2 = n.value.(int)
			result = (n.right.value.(int) >= val2)
			n.value = result
			n.right = nil
		} else if v.Kind() == reflect.String && vright.Kind() == reflect.String {
			var trimedval1 string = strings.TrimSpace(n.value.(string))
			var trimedval2 string = strings.TrimSpace(n.right.value.(string))
			r := strings.NewReplacer(" ", "")
			trimedval1 = r.Replace(trimedval1)
			trimedval2 = r.Replace(trimedval2)
			val2, err0 := time.Parse(Layout, trimedval1)
			val3, err1 := time.Parse(Layout, trimedval2)
			if err1 == nil && err0 == nil {
				n.value = val2.Equal(val3) || val2.Before(val3)
			}
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceMore() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceMore()
	}
	return t
}

func (n *BinaryNode) reduceMore() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceMore()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceMore()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceMore()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceMore()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceMore()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceMore()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)
		var result bool
		if v.Kind() == reflect.Float64 || vright.Kind() == reflect.Float64 {
			var val float64
			if v, ok := n.value.(int); ok {
				val = float64(v)
			} else {
				val = n.value.(float64)
			}
			if v, ok := n.right.value.(int); ok {
				result = (float64(v) < val)
			} else {
				result = (n.right.value.(float64) < val)
			}
			n.value = result
			n.right = nil
		} else if v.Kind() == reflect.Int {
			var val2 int
			val2 = n.value.(int)
			result = (n.right.value.(int) <= val2)
			n.value = result
			n.right = nil
		} else if v.Kind() == reflect.String && vright.Kind() == reflect.String {
			var trimedval1 string = strings.TrimSpace(n.value.(string))
			var trimedval2 string = strings.TrimSpace(n.right.value.(string))
			r := strings.NewReplacer(" ", "")
			trimedval1 = r.Replace(trimedval1)
			trimedval2 = r.Replace(trimedval2)
			val2, err0 := time.Parse(Layout, trimedval1)
			val3, err1 := time.Parse(Layout, trimedval2)
			if err1 == nil && err0 == nil {
				n.value = val2.After(val3)
			}
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceMoreEq() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceMoreEq()
	}
	return t
}

func (n *BinaryNode) reduceMoreEq() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceMoreEq()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceMoreEq()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceMoreEq()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceMoreEq()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceMoreEq()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceMoreEq()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)
		var result bool
		if v.Kind() == reflect.Float64 || vright.Kind() == reflect.Float64 {
			var val float64
			if v, ok := n.value.(int); ok {
				val = float64(v)
			} else {
				val = n.value.(float64)
			}
			if v, ok := n.right.value.(int); ok {
				result = (float64(v) <= val)
			} else {
				result = (n.right.value.(float64) <= val)
			}
			n.value = result
			n.right = nil
		} else if v.Kind() == reflect.Int {
			var val2 int
			val2 = n.value.(int)
			result = (n.right.value.(int) <= val2)
			n.value = result
			n.right = nil
		} else if v.Kind() == reflect.String && vright.Kind() == reflect.String {
			var trimedval1 string = strings.TrimSpace(n.value.(string))
			var trimedval2 string = strings.TrimSpace(n.right.value.(string))
			r := strings.NewReplacer(" ", "")
			trimedval1 = r.Replace(trimedval1)
			trimedval2 = r.Replace(trimedval2)
			val2, err0 := time.Parse(Layout, trimedval1)
			val3, err1 := time.Parse(Layout, trimedval2)
			if err1 == nil && err0 == nil {
				n.value = val2.Equal(val3) || val2.After(val3)
			}
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceMin() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceMin()
	}
	return t
}

func (n *BinaryNode) reduceMin() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceMin()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceMin()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceMin()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceMin()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceMin()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceMin()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)

		if v.Kind() == reflect.Float64 || vright.Kind() == reflect.Float64 {
			var val float64
			if v, ok := n.value.(int); ok {
				val = float64(v)
			} else {
				val = n.value.(float64)
			}
			if v, ok := n.right.value.(int); ok {
				if float64(v) < val {
					val = float64(v)
				}
			} else {
				if n.right.value.(float64) < val {
					val = n.right.value.(float64)
				}
			}
			n.value = val
			n.right = nil
		} else {
			var val2 int
			val2 = n.value.(int)
			if n.right.value.(int) < val2 {
				val2 = n.right.value.(int)
			}
			n.value = val2
			n.right = nil
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceMax() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceMax()
	}
	return t
}

func (n *BinaryNode) reduceMax() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceMax()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceMax()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceMax()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceMax()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceMax()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceMax()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)

		if v.Kind() == reflect.Float64 || vright.Kind() == reflect.Float64 {
			var val float64
			if v, ok := n.value.(int); ok {
				val = float64(v)
			} else {
				val = n.value.(float64)
			}
			if v, ok := n.right.value.(int); ok {
				if float64(v) > val {
					val = float64(v)
				}
			} else {
				if n.right.value.(float64) > val {
					val = n.right.value.(float64)
				}
			}
			n.value = val
			n.right = nil
		} else {
			var val2 int
			val2 = n.value.(int)
			if n.right.value.(int) > val2 {
				val2 = n.right.value.(int)
			}
			n.value = val2
			n.right = nil
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceSin() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceSin()
	}
	return t
}

func (n *BinaryNode) reduceSin() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceSin()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceSin()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceSin()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceSin()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceSin()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceSin()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = math.Sin(n.left.value.(float64))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceSin()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceTrim() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceTrim()
	}
	return t
}

func (n *BinaryNode) reduceTrim() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceTrim()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceTrim()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceTrim()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceTrim()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceTrim()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceTrim()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = strings.TrimSpace(n.left.value.(string))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceTrim()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceAbs() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceAbs()
	}
	return t
}

func (n *BinaryNode) reduceAbs() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceAbs()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceAbs()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceAbs()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceAbs()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceAbs()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceAbs()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = math.Abs(n.left.value.(float64))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceAbs()
		}
		return
	}

	if n.parent == nil {
		return
	}

}
func (t *BinaryTree) reduceLog() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceLog()
	}
	return t
}

func (n *BinaryNode) reduceLog() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceLog()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceLog()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceLog()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceLog()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceLog()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceLog()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = math.Log(n.left.value.(float64))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceLog()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceSqrt() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceSqrt()
	}
	return t
}

func (n *BinaryNode) reduceSqrt() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceSqrt()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceSqrt()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceSqrt()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceSqrt()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceSqrt()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceSqrt()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = math.Sqrt(n.left.value.(float64))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceSqrt()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceRound() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceRound()
	}
	return t
}

func (n *BinaryNode) reduceRound() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceRound()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceRound()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceRound()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceRound()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceRound()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceRound()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = math.Round(n.left.value.(float64))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceRound()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceCeiling() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceCeiling()
	}
	return t
}

func (n *BinaryNode) reduceCeiling() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceCeiling()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceCeiling()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceCeiling()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceCeiling()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceCeiling()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceCeiling()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = math.Ceil(n.left.value.(float64))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceCeiling()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceFloor() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceFloor()
	}
	return t
}

func (n *BinaryNode) reduceFloor() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceFloor()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceFloor()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceFloor()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceFloor()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceFloor()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceFloor()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = math.Floor(n.left.value.(float64))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceFloor()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceLen() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceLen()
	}
	return t
}

func (n *BinaryNode) reduceLen() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceLen()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceLen()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceLen()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceLen()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceLen()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceLen()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		str := fmt.Sprintf("%v", n.left.value)
		n.value = len([]rune(str))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceLen()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceCos() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceCos()
	}
	return t
}

func (n *BinaryNode) reduceCos() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceCos()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceCos()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceCos()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceCos()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceCos()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceCos()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = math.Cos(n.left.value.(float64))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceCos()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceLog10() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceLog10()
	}
	return t
}

func (n *BinaryNode) reduceLog10() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceLog10()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceLog10()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceLog10()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceLog10()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceLog10()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceLog10()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = math.Log10(n.left.value.(float64))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceLog10()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceContains() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceContains()
	}
	return t
}

func (n *BinaryNode) reduceContains() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceContains()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceContains()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceContains()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceContains()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceContains()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceContains()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		if v.Kind() == reflect.String {
			val := false
			val = strings.Contains(n.value.(string), n.right.value.(string))
			n.value = val
			n.right = nil
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceRegex() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceRegex()
	}
	return t
}

func (n *BinaryNode) reduceRegex() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceRegex()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceRegex()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceRegex()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceRegex()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceRegex()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceRegex()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		if v.Kind() == reflect.String {
			val := false
			val, _ = regexp.MatchString(n.right.value.(string), n.value.(string))
			n.value = val
			n.right = nil
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceExp() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceExp()
	}
	return t
}

func (n *BinaryNode) reduceExp() {

	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceExp()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceExp()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceExp()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceExp()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceExp()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceExp()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = math.Exp(n.left.value.(float64))
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceExp()
		}
		return
	}

	if n.parent == nil {
		return
	}

}

func (t *BinaryTree) reduceAdd() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceAdd()
	}
	return t
}

func (n *BinaryNode) reduceAdd() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceAdd()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceAdd()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceAdd()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceAdd()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceAdd()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceAdd()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)

		if v.Kind() == reflect.Float64 || vright.Kind() == reflect.Float64 {
			var val float64
			if v, ok := n.value.(int); ok {
				val = float64(v)
			} else {
				val = n.value.(float64)
			}
			if v, ok := n.right.value.(int); ok {
				val += float64(v)
			} else {
				val += n.right.value.(float64)
			}
			n.value = val
			n.right = nil
		} else {
			var val2 int
			val2 = n.value.(int)
			val2 += n.right.value.(int)
			n.value = val2
			n.right = nil
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceDiv() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceDiv()
	}
	return t
}

func (n *BinaryNode) reduceDiv() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceDiv()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceDiv()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceDiv()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceDiv()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceDiv()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceDiv()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)
		if v.Kind() == reflect.Float64 || vright.Kind() == reflect.Float64 {
			var val float64
			if v, ok := n.value.(int); ok {
				val = float64(v)
			} else {
				val = n.value.(float64)
			}
			if v, ok := n.right.value.(int); ok {
				val /= float64(v)
			} else {
				val /= n.right.value.(float64)
			}
			n.value = val
			n.right = nil
		} else {
			var val2 int
			val2 = n.value.(int)
			val2 /= n.right.value.(int)
			n.value = val2
			n.right = nil
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceMod() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceMod()
	}
	return t
}

func (n *BinaryNode) reduceMod() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceMod()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceMod()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceMod()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceMod()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceMod()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceMod()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)

		if v.Kind() == reflect.Float64 || vright.Kind() == reflect.Float64 {
			var val float64
			if v, ok := n.value.(int); ok {
				val = float64(v)
			} else {
				val = n.value.(float64)
			}
			if v, ok := n.right.value.(int); ok {
				val = math.Mod(val, float64(v))
			} else {
				val = math.Mod(val, n.right.value.(float64))
			}
			n.value = val
			n.right = nil
		} else {
			var val2 int
			val2 = n.value.(int)
			val2 = val2 % n.right.value.(int)
			n.value = val2
			n.right = nil
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceOr() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceOr()
	}
	return t
}

func (n *BinaryNode) reduceOr() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceOr()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceOr()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceOr()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceOr()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceOr()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceOr()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		val := parseAndOr(n.value)
		valr := parseAndOr(n.right.value)
		val = val || valr
		n.value = val
		n.right = nil
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceAnd() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceAnd()
	}
	return t
}

func (n *BinaryNode) reduceAnd() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceAnd()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceAnd()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceAnd()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceAnd()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceAnd()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceAnd()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		val := parseAndOr(n.value)
		valr := parseAndOr(n.right.value)
		val = val && valr
		n.value = val
		n.right = nil
		return
	}
	if n.parent == nil {
		return
	}
}

func parseAndOr(derivedOp interface{}) bool {
	if derivedOp == true {
		return true
	}
	if derivedOp == false {
		return false
	}
	if str, ok := derivedOp.(string); ok {
		var logicalOp string = str
		return CheckFormulaIntAndOr(logicalOp)
	}
	return false
}

func (t *BinaryTree) reduceMult() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceMult()
	}
	return t
}

func (n *BinaryNode) reduceMult() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceMult()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceMult()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceMult()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceMult()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceMult()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceMult()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		v := reflect.ValueOf(n.value)
		vright := reflect.ValueOf(n.right.value)
		if v.Kind() == reflect.Float64 || vright.Kind() == reflect.Float64 {
			var val float64
			if v, ok := n.value.(int); ok {
				val = float64(v)
			} else {
				val = n.value.(float64)
			}
			if v, ok := n.right.value.(int); ok {
				val *= float64(v)
			} else {
				val *= n.right.value.(float64)
			}
			n.value = val
			n.right = nil
		} else {
			var val2 int
			val2 = n.value.(int)
			val2 *= n.right.value.(int)
			n.value = val2
			n.right = nil
		}
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceLeft() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceLeft()
	}
	return t
}

func (n *BinaryNode) reduceLeft() {
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value.(string)
		n.left = nil
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		val := n.value.(string)
		chars := []rune(val)
		s := chars[:n.right.value.(int)]
		val = string(s)
		n.value = val
		n.right = nil
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceRight() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceRight()
	}
	return t
}

func (n *BinaryNode) reduceRight() {
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value.(string)
		n.left = nil
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		val := n.value.(string)
		chars := []rune(val)
		s := chars[n.right.value.(int):]
		val = string(s)
		n.value = val
		n.right = nil
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceTrue() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceTrue()
	}
	return t
}

func (n *BinaryNode) reduceTrue() {

	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceTrue()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceTrue()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceTrue()
		return
	}

	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = n.left.value
		n.left = nil
		n.right = nil
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceFalse() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceFalse()
	}
	return t
}

func (n *BinaryNode) reduceFalse() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceFalse()
		return
	}

	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceFalse()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceFalse()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) {
		n.value = n.right.value
		n.left = nil
		n.right = nil
		return
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceNot() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceNot()
	}
	return t
}

func (n *BinaryNode) reduceNot() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceNot()
		return
	}
	if n.left != nil && contains(FormOperators, n.left.value) {
		n = n.left
		n.reduceNot()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.left != nil {
		n = n.left
		n.reduceNot()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) && n.left.right != nil {
		n = n.left
		n.reduceNot()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.left != nil {
		n = n.right
		n.reduceNot()
		return
	}
	if n.right != nil && !contains(FormOperators, n.right.value) && n.right.right != nil {
		n = n.right
		n.reduceNot()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = !n.left.value.(bool)
		n.left = nil
		if n.parent != nil {
			n = n.parent
			n.reduceNot()
		}
		return
	}

	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceBlank() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceBlank()
	}
	return t
}

func (n *BinaryNode) reduceBlank() {
	if n.left != nil && !contains(FormOperators, n.left.value) {
		n.value = (n.left.value == nil) || (n.left.value.(string) == "")
		n.left = nil
	}
	if n.parent == nil {
		return
	}
}

func (t *BinaryTree) reduceBlankValue() *BinaryTree {
	if t.root == nil {
		return t
	} else {
		t.root.reduceBlankValue()
	}
	return t
}

func (n *BinaryNode) reduceBlankValue() {
	if n.right != nil && contains(FormOperators, n.right.value) {
		n = n.right
		n.reduceBlankValue()
		return
	}
	if n.left != nil && !contains(FormOperators, n.left.value) {
		check1 := (n.left.value == nil) || (n.left.value.(string) == "")
		check2 := n.right != nil && !contains(FormOperators, n.right.value)
		if check1 && check2 {
			n.value = n.right.value
			n.left = nil
			n.right = nil
			return
		} else {
			n.left = nil
			n.right = nil
			n.value = nil
		}
	}
	if n.parent == nil {
		return
	}
}

func (n *BinaryNode) insert(val interface{}, op string) {

	if n.right != nil && n.right.value == op {
		n = n.right
		n.insert(val, op)
		return
	}

	if n.left != nil && n.left.value == op {
		n = n.left
		n.insert(val, op)
		return
	}

	if n.left == nil {
		n.left = &BinaryNode{value: val, left: nil, right: nil, parent: n}
		n = n.left
		fmt.Println("Node left", n)
		return
	}

	if n.right == nil {
		n.right = &BinaryNode{value: val, left: nil, right: nil, parent: n}
		n = n.right
		fmt.Println("Node right", n)
		return
	}
	if n.parent == nil {
		return
	}
	lookup(n, val, 0, op)

}

func lookup(n *BinaryNode, val interface{}, i int, op string) {
	fmt.Println("lookup for :", val, "for the ", i, "th time")
	if i == 2 {
		fmt.Println("end of lookup for :", val, "for the last time")
		fmt.Println(n, n.right, n.left, n.parent)
		return
	}

	if n.parent == nil && n.right.value == op {
		n = n.right
		n.insert(val, op)
		return
	}

	if n.parent == nil && n.left.value == op {
		n = n.left
		n.insert(val, op)
		return
	}

	if n.parent != nil && n.parent.value == op && n.parent.right == nil {
		n.parent.right = &BinaryNode{value: val, left: nil, right: nil, parent: n.parent}
		n = n.parent.right
		return
	}
	if n.parent != nil && n.parent.value == op && n.parent.left == nil {
		n.parent.left = &BinaryNode{value: val, left: nil, right: nil, parent: n.parent}
		n = n.parent.left
		return
	}

	if n != nil && n.value == op && n.left == nil {
		n.left = &BinaryNode{value: val, left: nil, right: nil, parent: n.parent}
		n = n.left
		return
	}

	if n != nil && n.value == op && n.right == nil {
		n.right = &BinaryNode{value: val, left: nil, right: nil, parent: n.parent}
		n = n.right
		return
	}

	n = n.parent
	lookup(n, val, i+1, op)

	if n.left != nil && n.left.value != op && n.right != nil && n.right.value != op {
		return
	}

}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.value)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')

}

func checkIfBinary(slice []interface{}) (int, int, error) {
	var countOp int = 0
	var size = len(slice)
	for _, el := range slice {
		if contains(FormOperators, el) {
			countOp++
		}
	}
	if limit := 2*countOp + 1; limit != size {
		numberBinOpRequired := size - countOp - 1
		return numberBinOpRequired, countOp, fmt.Errorf("not a binary tree, you need %d binary operators\n", numberBinOpRequired)
	}

	return 0, 0, nil
}

func fixTree(slice []interface{}, n int, countOp int, op string) []interface{} {
	var stack []interface{}
	var res []interface{}
	var arity = 0
	numberOfOp := countOp
	for i := (len(slice) - 1); i >= 0; i-- {
		val := slice[i]
		if val != op {
			arity++
			stack = append(stack, val)
			if arity == 2 && numberOfOp < n {
				stack = append(stack, op)
				numberOfOp++
				arity = 1
			}
		} else {
			stack = append(stack, val)
			arity = 1
			numberOfOp++
		}
	}

	for len(stack) > 0 {
		n := len(stack) - 1 // Top element
		res = append(res, stack[n])
		stack = stack[:n] // Pop
	}
	return res
}

func buildBinaryTree(slice []interface{}, op string) *BinaryTree {
	tree := &BinaryTree{}

	for _, el := range slice {
		if contains(FormOperators, el) && op != el {
			op = fmt.Sprintf("%v", el)
		}
		tree.insert(el, op)
	}

	return tree
}

func buildFormulaTree(valop []interface{}, op string) *BinaryTree {

	tree := buildBinaryTree(valop, op)
	print(os.Stdout, tree.root, 0, 'M')

	numberBinOpRequired, countOp, err := checkIfBinary(valop)
	fmt.Printf("error : '%v' \n", err)

	if err != nil {
		stack := fixTree(valop, numberBinOpRequired, countOp, op)
		fmt.Printf("fixed tree %v \n", stack)
		if _, _, fixErr := checkIfBinary(valop); fixErr != nil {
			tree = buildBinaryTree(stack, op)
			print(os.Stdout, tree.root, 0, 'M')
		}
	}
	return tree
}

func evaluate(t *BinaryTree, op string) interface{} {
	for t.root.right != nil && t.root.left != nil {
		if op == ADD {
			t.reduceAdd()
			continue
		}
		if op == OR {
			t.reduceOr()
			continue
		}
		if op == AND {
			t.reduceAnd()
			continue
		}
		if op == MULT {
			t.reduceMult()
			continue
		}
		if op == LEFT {
			t.reduceLeft()
			continue
		}
		if op == RIGHT {
			t.reduceRight()
			continue
		}
		if op == TRUE {
			t.reduceTrue()
			continue
		}
		if op == FALSE {
			t.reduceFalse()
			continue
		}
		if op == NOT {
			t.reduceNot()
			continue
		}
		if op == DIV {
			t.reduceDiv()
			continue
		}
		if op == MOD {
			t.reduceMod()
			continue
		}
		if op == ISBLANK {
			t.reduceBlank()
			continue
		}
		if op == BLANKVALUE {
			t.reduceBlankValue()
			continue
		}
		if op == CONTAINS {
			t.reduceContains()
			continue
		}
		if op == EXP {
			t.reduceExp()
			continue
		}
		if op == COS {
			t.reduceCos()
			continue
		}
		if op == SIN {
			t.reduceSin()
			continue
		}
		if op == LEN {
			t.reduceLen()
			continue
		}
		if op == MAX {
			t.reduceMax()
			continue
		}
		if op == MIN {
			t.reduceMin()
			continue
		}
		if op == MORE_EQ {
			t.reduceMoreEq()
			continue
		}
		if op == LESS_EQ {
			t.reduceLessEq()
			continue
		}
		if op == MORE {
			t.reduceMore()
			continue
		}
		if op == LESS {
			t.reduceLess()
			continue
		}
		if op == EQUAL {
			t.reduceEqual()
			continue
		}
		if op == NOW {
			t.reduceToday()
			continue
		}
		if op == CEILING {
			t.reduceCeiling()
			continue
		}
		if op == ROUND {
			t.reduceRound()
			continue
		}
		if op == SQRT {
			t.reduceSqrt()
			continue
		}
		if op == LOG {
			t.reduceLog()
			continue
		}
		if op == LOG10 {
			t.reduceLog10()
			continue
		}
		if op == ABS {
			t.reduceAbs()
			continue
		}
		if op == FIND {
			t.reduceFind()
			continue
		}
		if op == TRIM {
			t.reduceTrim()
			continue
		}
		if op == TODAY {
			t.reduceDay()
			continue
		}
		if op == ADDMONTH {
			t.reduceAddMonth()
			continue
		}
		if op == BEGINS {
			t.reduceBegins()
			continue
		}
		if op == CONCAT {
			t.reduceConcat()
			continue
		}
		if op == BR {
			t.reduceBr()
			continue
		}
		if op == YEAR {
			t.reduceYear()
			continue
		}
		if op == MONTH {
			t.reduceMonth()
			continue
		}
		if op == DAY {
			t.reduceMonthDay()
			continue
		}
		if op == WEEKDAY {
			t.reduceWeekDay()
			continue
		}
		if op == HOUR {
			t.reduceHour()
			continue
		}
		if op == MINUTE {
			t.reduceMinute()
			continue
		}
		if op == FLOOR {
			t.reduceFloor()
			continue
		}
		if op == UPPER {
			t.reduceUpper()
			continue
		}
		if op == LOWER {
			t.reduceLower()
			continue
		}
		if op == REGEX {
			t.reduceRegex()
			continue
		}
	}
	if op == ADD {
		t.reduceAdd()
	} else if op == OR {
		t.reduceOr()
	} else if op == AND {
		t.reduceAnd()
	} else if op == MULT {
		t.reduceMult()
	} else if op == LEFT {
		t.reduceLeft()
	} else if op == RIGHT {
		t.reduceRight()
	} else if op == TRUE {
		t.reduceTrue()
	} else if op == FALSE {
		t.reduceFalse()
	} else if op == NOT {
		t.reduceNot()
	} else if op == DIV {
		t.reduceDiv()
	} else if op == MOD {
		t.reduceMod()
	} else if op == ISBLANK {
		t.reduceBlank()
	} else if op == BLANKVALUE {
		t.reduceBlankValue()
	} else if op == CONTAINS {
		t.reduceContains()
	} else if op == EXP {
		t.reduceExp()
	} else if op == COS {
		t.reduceCos()
	} else if op == SIN {
		t.reduceSin()
	} else if op == LEN {
		t.reduceLen()
	} else if op == MAX {
		t.reduceMax()
	} else if op == MIN {
		t.reduceMin()
	} else if op == MORE_EQ {
		t.reduceMoreEq()
	} else if op == LESS_EQ {
		t.reduceLessEq()
	} else if op == MORE {
		t.reduceMore()
	} else if op == LESS {
		t.reduceLess()
	} else if op == EQUAL {
		t.reduceEqual()
	} else if op == NOW {
		t.reduceToday()
	} else if op == CEILING {
		t.reduceCeiling()
	} else if op == ROUND {
		t.reduceRound()
	} else if op == SQRT {
		t.reduceSqrt()
	} else if op == LOG {
		t.reduceLog()
	} else if op == LOG10 {
		t.reduceLog10()
	} else if op == ABS {
		t.reduceAbs()
	} else if op == FIND {
		t.reduceFind()
	} else if op == TRIM {
		t.reduceTrim()
	} else if op == TODAY {
		t.reduceDay()
	} else if op == ADDMONTH {
		t.reduceAddMonth()
	} else if op == BEGINS {
		t.reduceBegins()
	} else if op == CONCAT {
		t.reduceConcat()
	} else if op == BR {
		t.reduceBr()
	} else if op == YEAR {
		t.reduceYear()
	} else if op == MONTH {
		t.reduceMonth()
	} else if op == DAY {
		t.reduceMonthDay()
	} else if op == WEEKDAY {
		t.reduceWeekDay()
	} else if op == HOUR {
		t.reduceHour()
	} else if op == MINUTE {
		t.reduceMinute()
	} else if op == FLOOR {
		t.reduceFloor()
	} else if op == UPPER {
		t.reduceUpper()
	} else if op == LOWER {
		t.reduceLower()
	} else if op == REGEX {
		t.reduceRegex()
	}

	return t.root.value
}

func RunFormula(valop []interface{}) interface{} {
	t := buildFormulaTree(valop, fmt.Sprintf("%v", valop[0]))
	val := evaluate(t, fmt.Sprintf("%v", valop[0]))
	print(os.Stdout, t.root, 0, 'M')
	return val
}
