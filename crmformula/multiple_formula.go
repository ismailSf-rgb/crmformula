package crmformula

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	Layout          = "2006-01-02at3:04pm(UTC)"
	DisplayedLayout = "2006-01-02 at 3:04pm (UTC)"
)

func FindMostNestedArray(valop []interface{}) []interface{} {
	countOp := 0
	firstTermOp := 1
	n := len(valop) - 1
	lastTermOp := n
	for n >= 0 {
		if valop[n] == ")" {
			countOp++
			lastTermOp = n
			n--
			continue
		}
		if valop[n] == "(" {
			countOp--
			firstTermOp = n
			break
		}
		n--
	}
	firstTermOp++
	fmt.Println(valop[firstTermOp:lastTermOp])
	var s []interface{}
	s = append(s, valop[firstTermOp-2])
	slice := valop[firstTermOp:lastTermOp]

	for _, el := range slice {
		s = append(s, el)
	}
	fmt.Println(s)
	return s
}

func ParseIf(derivedOp interface{}) string {
	return parseIf(derivedOp)
}

func CheckFormulaIntAndOr(logic string) bool {
	derivedOp := strings.TrimSpace(logic)
	r := strings.NewReplacer(" ", "")
	derivedOp = r.Replace(derivedOp)
	if strings.Contains(derivedOp, ">=") {
		arr := strings.Split(derivedOp, ">=")
		match1, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[0])
		match2, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[1])
		if match1 && match2 {
			a, _ := strconv.ParseFloat(arr[0], 64)
			b, _ := strconv.ParseFloat(arr[1], 64)
			if a >= b {
				return true
			} else {
				return false
			}
		} else {
			arr := strings.Split(derivedOp, ">=")
			t0, err0 := time.Parse(Layout, arr[0])
			t1, err1 := time.Parse(Layout, arr[1])
			if err1 == nil && err0 == nil {
				return t0.After(t1) || t0.Equal(t1)
			} else {
				return false
			}
		}
	}
	if strings.Contains(derivedOp, "<=") {
		arr := strings.Split(derivedOp, "<=")
		match1, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[0])
		match2, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[1])
		if match1 && match2 {
			a, _ := strconv.ParseFloat(arr[0], 64)
			b, _ := strconv.ParseFloat(arr[1], 64)
			if a <= b {
				return true
			} else {
				return false
			}
		} else {
			arr := strings.Split(derivedOp, "<=")
			t0, err0 := time.Parse(Layout, arr[0])
			t1, err1 := time.Parse(Layout, arr[1])
			if err1 == nil && err0 == nil {
				return t0.Before(t1) || t0.Equal(t1)
			} else {
				return false
			}
		}
	}
	if strings.Contains(derivedOp, "<") {
		arr := strings.Split(derivedOp, "<")
		match1, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[0])
		match2, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[1])
		if match1 && match2 {
			a, _ := strconv.ParseFloat(arr[0], 64)
			b, _ := strconv.ParseFloat(arr[1], 64)
			if a < b {
				return true
			} else {
				return false
			}
		} else {
			arr := strings.Split(derivedOp, "<")
			t0, err0 := time.Parse(Layout, arr[0])
			t1, err1 := time.Parse(Layout, arr[1])
			if err1 == nil && err0 == nil {
				return t0.Before(t1)
			} else {
				return false
			}
		}
	}
	if strings.Contains(derivedOp, ">") {
		arr := strings.Split(derivedOp, ">")
		match1, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[0])
		match2, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[1])
		if match1 && match2 {
			a, _ := strconv.ParseFloat(arr[0], 64)
			b, _ := strconv.ParseFloat(arr[1], 64)
			if a > b {
				return true
			} else {
				return false
			}
		} else {
			arr := strings.Split(derivedOp, ">")
			t0, err0 := time.Parse(Layout, arr[0])
			t1, err1 := time.Parse(Layout, arr[1])
			if err1 == nil && err0 == nil {
				return t0.After(t1)
			} else {
				return false
			}
		}
	}
	if strings.Contains(derivedOp, "=") {
		arr := strings.Split(derivedOp, "=")
		match1, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[0])
		match2, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[1])
		if match1 && match2 {
			a, _ := strconv.ParseFloat(arr[0], 64)
			b, _ := strconv.ParseFloat(arr[1], 64)
			if a == b {
				return true
			} else {
				return false
			}
		} else {
			arr := strings.Split(derivedOp, "=")
			t0, err0 := time.Parse(Layout, arr[0])
			t1, err1 := time.Parse(Layout, arr[1])
			if err1 == nil && err0 == nil {
				return t0.Equal(t1)
			} else {
				return false
			}
		}
	}
	return false
}

func checkFormulaInt(logic string) string {
	derivedOp := strings.TrimSpace(logic)
	r := strings.NewReplacer(" ", "")
	derivedOp = r.Replace(derivedOp)
	if strings.Contains(derivedOp, ">=") {
		arr := strings.Split(derivedOp, ">=")
		match1, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[0])
		match2, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[1])
		if match1 && match2 {
			a, _ := strconv.ParseFloat(arr[0], 64)
			b, _ := strconv.ParseFloat(arr[1], 64)
			if a >= b {
				return "TRUE"
			} else {
				return "FALSE"
			}
		} else {
			arr := strings.Split(derivedOp, ">=")
			t0, err0 := time.Parse(Layout, arr[0])
			t1, err1 := time.Parse(Layout, arr[1])
			if err1 == nil && err0 == nil {
				if t0.After(t1) || t0.Equal(t1) {
					return "TRUE"
				}
				return "FALSE"
			} else {
				return "FALSE"
			}
		}
	}
	if strings.Contains(derivedOp, "<=") {
		arr := strings.Split(derivedOp, "<=")
		match1, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[0])
		match2, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[1])
		if match1 && match2 {
			a, _ := strconv.ParseFloat(arr[0], 64)
			b, _ := strconv.ParseFloat(arr[1], 64)
			if a <= b {
				return "TRUE"
			} else {
				return "FALSE"
			}
		} else {
			arr := strings.Split(derivedOp, "<=")
			t0, err0 := time.Parse(Layout, arr[0])
			t1, err1 := time.Parse(Layout, arr[1])
			if err1 == nil && err0 == nil {
				if t0.Before(t1) || t0.Equal(t1) {
					return "TRUE"
				}
				return "FALSE"
			} else {
				return "FALSE"
			}
		}
	}
	if strings.Contains(derivedOp, "<") {
		arr := strings.Split(derivedOp, "<")
		match1, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[0])
		match2, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[1])
		if match1 && match2 {
			a, _ := strconv.ParseFloat(arr[0], 64)
			b, _ := strconv.ParseFloat(arr[1], 64)
			if a < b {
				return "TRUE"
			} else {
				return "FALSE"
			}
		} else {
			arr := strings.Split(derivedOp, "<")
			t0, err0 := time.Parse(Layout, arr[0])
			t1, err1 := time.Parse(Layout, arr[1])
			if err1 == nil && err0 == nil {
				if t0.Before(t1) {
					return "TRUE"
				}
				return "FALSE"
			} else {
				return "FALSE"
			}
		}
	}
	if strings.Contains(derivedOp, ">") {
		arr := strings.Split(derivedOp, ">")
		match1, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[0])
		match2, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[1])
		if match1 && match2 {
			a, _ := strconv.ParseFloat(arr[0], 64)
			b, _ := strconv.ParseFloat(arr[1], 64)
			if a > b {
				return "TRUE"
			} else {
				return "FALSE"
			}
		} else {
			arr := strings.Split(derivedOp, ">")
			t0, err0 := time.Parse(Layout, arr[0])
			t1, err1 := time.Parse(Layout, arr[1])
			if err1 == nil && err0 == nil {
				if t0.After(t1) {
					return "TRUE"
				}
				return "FALSE"
			} else {
				return "FALSE"
			}
		}
	}
	if strings.Contains(derivedOp, "=") {
		arr := strings.Split(derivedOp, "=")
		match1, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[0])
		match2, _ := regexp.MatchString(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`, arr[1])
		if match1 && match2 {
			a, _ := strconv.ParseFloat(arr[0], 64)
			b, _ := strconv.ParseFloat(arr[1], 64)
			if a == b {
				return "TRUE"
			} else {
				return "FALSE"
			}
		} else {
			arr := strings.Split(derivedOp, "=")
			t0, err0 := time.Parse(Layout, arr[0])
			t1, err1 := time.Parse(Layout, arr[1])
			if err1 == nil && err0 == nil {
				if t0.Equal(t1) {
					return "TRUE"
				}
				return "FALSE"
			} else {
				return "FALSE"
			}
		}
	}
	return "FALSE"
}

func parseIf(derivedOp interface{}) string {
	if derivedOp == true {
		return "TRUE"
	}
	if derivedOp == false {
		return "FALSE"
	}
	if str, ok := derivedOp.(string); ok {
		var logicalOp string = str
		return checkFormulaInt(logicalOp)
	}
	return "FALSE"
}

func manageIf(valop []interface{}) []interface{} {
	if valop[0] == "IF" {
		var res []interface{} = make([]interface{}, 3)
		res[0] = parseIf(valop[1])
		res[1] = valop[2]
		res[2] = valop[3]
		return res
	}
	return valop
}

func FindMostNestedArrayAndReplace(valop []interface{}) ([]interface{}, []interface{}) {
	countOp := 0
	firstTermOp := 1
	n := len(valop) - 1
	lastTermOp := n
	for n >= 0 {
		if valop[n] == ")" {
			countOp++
			lastTermOp = n
			n--
			continue
		}
		if valop[n] == "(" {
			countOp--
			firstTermOp = n
			break
		}
		n--
	}
	firstTermOp++
	var s []interface{}
	s = append(s, valop[firstTermOp-2])
	slice := valop[firstTermOp:lastTermOp]

	for _, el := range slice {
		s = append(s, el)
	}

	subFormula := manageIf(s)

	var newValop []interface{}

	for index, el := range valop {
		if index == (firstTermOp - 2) {
			newValop = append(newValop, "?")
		}
		if index >= (firstTermOp-2) && index <= lastTermOp {
			continue
		} else {
			newValop = append(newValop, el)
		}
	}
	fmt.Println("new array ", newValop)

	return subFormula, newValop
}

func FindMostNestedArrayAndEvaluate(valop []interface{}) []interface{} {
	subFormula, pFormula := FindMostNestedArrayAndReplace(valop)
	eval := RunFormula(subFormula)
	indexOfEval := -1
	for i, el := range pFormula {
		if el == "?" {
			indexOfEval = i
		}
	}
	pFormula[indexOfEval] = eval
	fmt.Println("new array : ", pFormula)
	return pFormula
}

func GetFormulaResult(valop []interface{}) interface{} {
	if len(valop) == 1 {
		return valop[0]
	}
	valop = Reformat(valop)
	subFormula, pFormula := FindMostNestedArrayAndReplace(valop)
	eval := RunFormula(subFormula)
	indexOfEval := -1
	for i, el := range pFormula {
		if el == "?" {
			indexOfEval = i
		}
	}
	pFormula[indexOfEval] = eval

	return GetFormulaResult(pFormula)
}

func Reformat(valop []interface{}) []interface{} {
	var result []interface{}
	for contains(valop, "<=") {
		result = TransformOpLessEq(valop)
		valop = result
	}
	for contains(valop, ">=") {
		result = TransformOpMoreEq(valop)
		valop = result
	}
	for contains(valop, "<") {
		result = TransformOpLess(valop)
		valop = result
	}
	for contains(valop, ">") {
		result = TransformOpMore(valop)
		valop = result
	}
	for contains(valop, "=") {
		result = TransformOpEqual(valop)
		valop = result
	}

	return valop
}

func TransformOpLess(valop []interface{}) []interface{} {
	var result []interface{}
	result = make([]interface{}, len(valop)+2)
	var lessLeft bool = false
	n := len(valop) - 1
	for n > 0 {
		if valop[n] == "<" && valop[n-1] == ")" {
			lessLeft = true
			valop[n] = valop[n+1]
			valop[n+1] = ")"
			n--
			continue
		}
		if valop[n] == "(" && n > 0 && lessLeft {
			result[n] = "("
			result[n-1] = LESS

			for i := range valop {
				if i < (n - 1) {
					result[i] = valop[i]
				}
				if result[2+i] != nil {
					continue
				} else {
					result[2+i] = valop[i]
				}
			}
			return result
		}
		n--
	}

	result = TransformOpLessRight(valop)
	return result
}

func TransformOpLessRight(valop []interface{}) []interface{} {

	size := len(valop)
	var result []interface{}
	result = make([]interface{}, size+2)
	var countOp1 int = 0
	var countOp2 int = 0
	var startCounting bool = false
	//var stopCounting bool = false
	//var index int = 0
	for i := range valop {

		if valop[i] == "<" && valop[i+2] == "(" {
			result[i+1] = valop[i-1]
			result[i] = "("
			result[i-1] = LESS
			result[i+2] = valop[i+1]
			countOp1++
			startCounting = true
			i += 2
			n := i
			for n < size-1 && startCounting {
				if valop[n] == "(" {
					countOp1++
					result[n+1] = valop[n]
					n++
					continue
				}
				if valop[n] == ")" {
					countOp2++
					result[n+1] = valop[n]
					n++
					continue
				}
				if countOp1-countOp2 != 1 {
					result[n+1] = valop[n]
				} else {
					fmt.Printf("countOP1: %v and CountOp2:%v and n: %v \n", countOp1, countOp2, n)
					result[n+1] = ")"
					result[n+2] = valop[n]
					countOp2++
					n += 2
					continue
				}
				n++
			}
			if result[len(result)-2] == nil {
				result[len(result)-2] = valop[len(valop)-2]
			}
			result[len(result)-1] = ")"

		}
		if valop[i] != "<" && !startCounting {
			result[i] = valop[i]
			i++
			continue
		}

		return result

	}
	return result
}

func TransformOpLessEq(valop []interface{}) []interface{} {
	var result []interface{}
	result = make([]interface{}, len(valop)+2)
	var lessLeft bool = false
	n := len(valop) - 1
	for n > 0 {
		if valop[n] == "<=" && valop[n-1] == ")" {
			lessLeft = true
			valop[n] = valop[n+1]
			valop[n+1] = ")"
			n--
			continue
		}
		if valop[n] == "(" && n > 0 && lessLeft {
			result[n] = "("
			result[n-1] = LESS_EQ

			for i := range valop {
				if i < (n - 1) {
					result[i] = valop[i]
				}
				if result[2+i] != nil {
					continue
				} else {
					result[2+i] = valop[i]
				}
			}
			return result
		}
		n--
	}

	result = TransformOpLessEqRight(valop)
	return result
}

func TransformOpLessEqRight(valop []interface{}) []interface{} {

	size := len(valop)
	var result []interface{}
	result = make([]interface{}, size+2)
	var countOp1 int = 0
	var countOp2 int = 0
	var startCounting bool = false
	//var stopCounting bool = false
	//var index int = 0
	for i := range valop {

		if valop[i] == "<=" && valop[i+2] == "(" {
			result[i+1] = valop[i-1]
			result[i] = "("
			result[i-1] = LESS_EQ
			result[i+2] = valop[i+1]
			countOp1++
			startCounting = true
			i += 2
			n := i
			for n < size-1 && startCounting {
				if valop[n] == "(" {
					countOp1++
					result[n+1] = valop[n]
					n++
					continue
				}
				if valop[n] == ")" {
					countOp2++
					result[n+1] = valop[n]
					n++
					continue
				}
				if countOp1-countOp2 != 1 {
					result[n+1] = valop[n]
				} else {
					fmt.Printf("countOP1: %v and CountOp2:%v and n: %v \n", countOp1, countOp2, n)
					result[n+1] = ")"
					result[n+2] = valop[n]
					countOp2++
					n += 2
					continue
				}
				n++
			}
			if result[len(result)-2] == nil {
				result[len(result)-2] = valop[len(valop)-2]
			}
			result[len(result)-1] = ")"

		}
		if valop[i] != "<=" && !startCounting {
			result[i] = valop[i]
			i++
			continue
		}

		return result

	}
	return result
}

func TransformOpMore(valop []interface{}) []interface{} {
	var result []interface{}
	result = make([]interface{}, len(valop)+2)
	var lessLeft bool = false
	n := len(valop) - 1
	for n > 0 {
		if valop[n] == ">" && valop[n-1] == ")" {
			lessLeft = true
			valop[n] = valop[n+1]
			valop[n+1] = ")"
			n--
			continue
		}
		if valop[n] == "(" && n > 0 && lessLeft {
			result[n] = "("
			result[n-1] = MORE

			for i := range valop {
				if i < (n - 1) {
					result[i] = valop[i]
				}
				if result[2+i] != nil {
					continue
				} else {
					result[2+i] = valop[i]
				}
			}
			return result
		}
		n--
	}

	result = TransformOpMoreRight(valop)
	return result
}

func TransformOpMoreRight(valop []interface{}) []interface{} {

	size := len(valop)
	var result []interface{}
	result = make([]interface{}, size+2)
	var countOp1 int = 0
	var countOp2 int = 0
	var startCounting bool = false
	//var stopCounting bool = false
	//var index int = 0
	for i := range valop {

		if valop[i] == ">" && valop[i+2] == "(" {
			result[i+1] = valop[i-1]
			result[i] = "("
			result[i-1] = MORE
			result[i+2] = valop[i+1]
			countOp1++
			startCounting = true
			i += 2
			n := i
			for n < size-1 && startCounting {
				if valop[n] == "(" {
					countOp1++
					result[n+1] = valop[n]
					n++
					continue
				}
				if valop[n] == ")" {
					countOp2++
					result[n+1] = valop[n]
					n++
					continue
				}
				if countOp1-countOp2 != 1 {
					result[n+1] = valop[n]
				} else {
					fmt.Printf("countOP1: %v and CountOp2:%v and n: %v \n", countOp1, countOp2, n)
					result[n+1] = ")"
					result[n+2] = valop[n]
					countOp2++
					n += 2
					continue
				}
				n++
			}
			if result[len(result)-2] == nil {
				result[len(result)-2] = valop[len(valop)-2]
			}
			result[len(result)-1] = ")"

		}
		if valop[i] != ">" && !startCounting {
			result[i] = valop[i]
			i++
			continue
		}

		return result

	}
	return result
}

func TransformOpMoreEq(valop []interface{}) []interface{} {
	var result []interface{}
	result = make([]interface{}, len(valop)+2)
	var lessLeft bool = false
	n := len(valop) - 1
	for n > 0 {
		if valop[n] == ">=" && valop[n-1] == ")" {
			lessLeft = true
			valop[n] = valop[n+1]
			valop[n+1] = ")"
			n--
			continue
		}
		if valop[n] == "(" && n > 0 && lessLeft {
			result[n] = "("
			result[n-1] = MORE_EQ

			for i := range valop {
				if i < (n - 1) {
					result[i] = valop[i]
				}
				if result[2+i] != nil {
					continue
				} else {
					result[2+i] = valop[i]
				}
			}
			return result
		}
		n--
	}

	result = TransformOpMoreEqRight(valop)
	return result
}

func TransformOpMoreEqRight(valop []interface{}) []interface{} {

	size := len(valop)
	var result []interface{}
	result = make([]interface{}, size+2)
	var countOp1 int = 0
	var countOp2 int = 0
	var startCounting bool = false
	//var stopCounting bool = false
	//var index int = 0
	for i := range valop {

		if valop[i] == ">=" && valop[i+2] == "(" {
			result[i+1] = valop[i-1]
			result[i] = "("
			result[i-1] = MORE_EQ
			result[i+2] = valop[i+1]
			countOp1++
			startCounting = true
			i += 2
			n := i
			for n < size-1 && startCounting {
				if valop[n] == "(" {
					countOp1++
					result[n+1] = valop[n]
					n++
					continue
				}
				if valop[n] == ")" {
					countOp2++
					result[n+1] = valop[n]
					n++
					continue
				}
				if countOp1-countOp2 != 1 {
					result[n+1] = valop[n]
				} else {
					fmt.Printf("countOP1: %v and CountOp2:%v and n: %v \n", countOp1, countOp2, n)
					result[n+1] = ")"
					result[n+2] = valop[n]
					countOp2++
					n += 2
					continue
				}
				n++
			}
			if result[len(result)-2] == nil {
				result[len(result)-2] = valop[len(valop)-2]
			}
			result[len(result)-1] = ")"

		}
		if valop[i] != ">=" && !startCounting {
			result[i] = valop[i]
			i++
			continue
		}

		return result

	}
	return result
}

func TransformOpEqual(valop []interface{}) []interface{} {
	var result []interface{}
	result = make([]interface{}, len(valop)+2)
	var lessLeft bool = false
	n := len(valop) - 1
	for n > 0 {
		if valop[n] == "=" && valop[n-1] == ")" {
			lessLeft = true
			valop[n] = valop[n+1]
			valop[n+1] = ")"
			n--
			continue
		}
		if valop[n] == "(" && n > 0 && lessLeft {
			result[n] = "("
			result[n-1] = EQUAL

			for i := range valop {
				if i < (n - 1) {
					result[i] = valop[i]
				}
				if result[2+i] != nil {
					continue
				} else {
					result[2+i] = valop[i]
				}
			}
			return result
		}
		n--
	}

	result = TransformOpEqRight(valop)
	return result
}

func TransformOpEqRight(valop []interface{}) []interface{} {

	size := len(valop)
	var result []interface{}
	result = make([]interface{}, size+2)
	var countOp1 int = 0
	var countOp2 int = 0
	var startCounting bool = false
	//var stopCounting bool = false
	//var index int = 0
	for i := range valop {

		if valop[i] == "=" && valop[i+2] == "(" {
			result[i+1] = valop[i-1]
			result[i] = "("
			result[i-1] = EQUAL
			result[i+2] = valop[i+1]
			countOp1++
			startCounting = true
			i += 2
			n := i
			for n < size-1 && startCounting {
				if valop[n] == "(" {
					countOp1++
					result[n+1] = valop[n]
					n++
					continue
				}
				if valop[n] == ")" {
					countOp2++
					result[n+1] = valop[n]
					n++
					continue
				}
				if countOp1-countOp2 != 1 {
					result[n+1] = valop[n]
				} else {
					fmt.Printf("countOP1: %v and CountOp2:%v and n: %v \n", countOp1, countOp2, n)
					result[n+1] = ")"
					result[n+2] = valop[n]
					countOp2++
					n += 2
					continue
				}
				n++
			}
			if result[len(result)-2] == nil {
				result[len(result)-2] = valop[len(valop)-2]
			}
			result[len(result)-1] = ")"

		}
		if valop[i] != "=" && !startCounting {
			result[i] = valop[i]
			i++
			continue
		}

		return result

	}
	return result
}
