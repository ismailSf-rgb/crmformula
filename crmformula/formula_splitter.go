package crmformula

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type FormulaType int64

type Formula struct {
	RawFormula      string `json:"rawFormula"`
	formulaRegex    string
	numberOfOperand int
	formulaType     FormulaType
	operators       []string
	emptyFormula    string
	length          int
	opIndex         map[int]string
	parIndex        map[int]string
	midFormula      []string
	endFormula      []interface{}
	TimeZone        string                 `json:"timeZone"`
	FieldValue      map[string]interface{} `json:"fieldValue"`
}

const (
	Number FormulaType = iota
	Checkbox
	Percent
	Text
	Date
	DateTime
	Any
)

func (f *Formula) GetRawFormula() string {
	return f.RawFormula
}
func (f *Formula) GetTimeZone() string {
	return f.TimeZone
}

func (returnType FormulaType) String() string {
	switch returnType {
	case Number:
		return "Number"
	case Checkbox:
		return "Checkbox"
	case Percent:
		return "Percent"
	case Text:
		return "Text"
	case Date:
		return "Date"
	case DateTime:
		return "DateTime"
	case Any:
		return "Text"
	default:
		return fmt.Sprintf("uknown type")
	}

}

func GetReadOnlyMapType() map[string]FormulaType {
	readOnly := map[string]FormulaType{
		ADD:      Number,
		OR:       Checkbox,
		AND:      Checkbox,
		IF:       Any,
		TEXT:     Text,
		LEFT:     Text,
		RIGHT:    Text,
		MULT:     Number,
		TRUE:     Any,
		FALSE:    Any,
		NOT:      Checkbox,
		DIV:      Number,
		ISBLANK:  Checkbox,
		CONTAINS: Checkbox,
		EXP:      Number,
		COS:      Number,
		LEN:      Number,
		MAX:      Number,
		MIN:      Number,
		MORE_EQ:  Checkbox,
		LESS_EQ:  Checkbox,
		MORE:     Checkbox,
		LESS:     Checkbox,
		EQUAL:    Checkbox,
		NOW:      DateTime,
		TODAY:    Date,
		ADDMONTH: Date,
		CEILING:  Number,
		ROUND:    Number,
		SQRT:     Number,
		LOG:      Number,
		LOG10:    Number,
		ABS:      Number,
		FIND:     Number,
		TRIM:     Text,
		BEGINS:   Checkbox,
		CONCAT:   Text,
		BR:       Text,
		MONTH:    Number,
		YEAR:     Number,
		DAY:      Number,
		WEEKDAY:  Number,
		HOUR:     Number,
		MINUTE:   Number,
		FLOOR:    Number,
		UPPER:    Text,
		LOWER:    Text,
		REGEX:    Checkbox,
	}
	return readOnly
}

func (f *Formula) FindOperators() ([]string, error) {

	var result []string
	//opIndex maps the position of the operator with the formula operator
	var opIndex map[int]string = make(map[int]string)
	rawFormula := f.RawFormula
	if rawFormula == "" {
		return nil, errors.New("parsing formula error (formula empty)")
	}
	rawFormula = strings.ReplaceAll(rawFormula, " ", "")
	//iterate over all existing operators
	for _, op := range FormOperators {
		var OpDoesExist bool = true
		for OpDoesExist {
			//get the position of the operator int the formula
			index := strings.Index(rawFormula, op.(string))
			//if operator doesn't exist in formula iterate over next one
			if index == -1 {
				OpDoesExist = false
				continue
			}
			// if the operator has already a position return an error
			if _, ok := opIndex[index]; ok {
				fmt.Printf("Error : %s from %s in %d index already exist", op.(string), rawFormula, index)
				return nil, errors.New("parsing formula error")
			}
			// otherwise put the operator and its position on the map
			opIndex[index] = op.(string)
			//split the raw formula in two in an array: first what is behind the operator and second all next
			stringArray := strings.SplitN(rawFormula, op.(string), 2)
			//if there are more than two parts in the array then throw an error
			if len(stringArray) != 2 {
				return nil, errors.New("parsing formula error")
			}
			// change the formula by removing the found operator but replace it by spaces of the same length
			//this is to have different position over the iterations
			rawFormula = fmt.Sprintf("%s%s%s", stringArray[0], strings.Repeat(" ", len(op.(string))), stringArray[1])
		}
	}

	f.opIndex = opIndex
	f.emptyFormula = rawFormula
	f.length = len(rawFormula)
	// we iterated over all operators. we make two lists
	//one list with the positions and one with the operators
	keys := make([]int, len(opIndex))
	result = make([]string, len(opIndex))
	i := 0
	// we build an array of all the positions
	for k := range opIndex {
		keys[i] = k
		i++
	}
	// we sort the position array so that we have the proper order of the formula
	sort.Ints(keys)

	// we finaly build the array of the formulas
	k := 0
	for _, key := range keys {
		result[k] = opIndex[key]
		k++
	}
	f.numberOfOperand = len(result)
	f.operators = result
	f.formulaType = GetReadOnlyMapType()[result[0]]
	return result, nil
}

func (f *Formula) FindOperatorsAndParenthesis() ([]string, error) {
	var result []string
	operators := f.operators

	rawFormula := f.RawFormula
	//opIndex maps the position of the parenthesis with the parenthesis
	var parenthesisIndex map[int]string = make(map[int]string)

	if rawFormula == "" || len(operators) == 0 {
		return nil, errors.New("parsing formula error (formula empty)")
	}
	result = make([]string, len(operators)+len(operators)*2)

	rawFormula = strings.ReplaceAll(rawFormula, " ", "")

	for _, el := range strings.Split(rawFormula, "") {
		if el == "(" || el == ")" {
			index := strings.Index(rawFormula, el)
			if index == -1 {
				continue
			}
			if _, ok := parenthesisIndex[index]; ok {
				fmt.Printf("Error : %s from %s in %d index already exist", el, rawFormula, index)
				return nil, errors.New("parsing formula error")
			}

			parenthesisIndex[index] = el

			stringArray := strings.SplitN(rawFormula, el, 2)

			if len(stringArray) != 2 {
				return nil, errors.New("parsing formula error")
			}
			// change the formula by removing the found operator but replace it by spaces of the same length
			//this is to have different position over the iterations
			rawFormula = fmt.Sprintf("%s%s%s", stringArray[0], strings.Repeat(" ", 1), stringArray[1])
		}
	}

	if len(parenthesisIndex) != len(operators)*2 {
		return nil, errors.New("compiling formula error :")
	}

	f.emptyFormula = rawFormula
	f.parIndex = make(map[int]string)
	for k, v := range f.opIndex {
		f.parIndex[k] = v
	}
	for k, v := range parenthesisIndex {
		f.parIndex[k] = v
	}
	keys := make([]int, len(f.parIndex))

	i := 0
	for k := range f.parIndex {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	// we finaly build the array of the formulas
	k := 0
	for _, key := range keys {
		result[k] = f.parIndex[key]
		k++
	}
	f.midFormula = result
	return result, nil
}

func (f *Formula) FindOperatorsAndParenthesisAndValues() ([]interface{}, error) {
	opAndParenthesis := f.midFormula
	fmt.Println(opAndParenthesis)

	var valuesIndex map[int]interface{} = make(map[int]interface{})
	fmt.Println(valuesIndex)

	return nil, nil
}

func (f *Formula) Tokenize() ([]interface{}, error) {
	var formula string = f.RawFormula
	var tokens []interface{}
	var stringDeclared bool = false
	token := ""
	inWord := false
	countDoubleQuotes := 0

	for _, char := range formula {
		if char == '"' {
			countDoubleQuotes++
			if !allDoubloQuotesCloses(countDoubleQuotes, stringDeclared) {
				stringDeclared = true
			} else {
				stringDeclared = false
			}
		}
		if char == '(' || char == ')' {
			if inWord {
				tokens = appendToken(tokens, token, f.TimeZone)
				token = ""
				inWord = false
			}
			tokens = append(tokens, string(char))
		} else if char == ',' {
			if inWord {
				tokens = appendToken(tokens, token, f.TimeZone)
				token = ""
				inWord = false
			}
		} else if char == ' ' && stringDeclared {
			token += string(char)
			continue
		} else if char != ' ' {
			token += string(char)
			inWord = true
		}
	}

	if inWord {
		tokens = appendToken(tokens, token, f.TimeZone)
	}

	if !allDoubloQuotesCloses(countDoubleQuotes, stringDeclared) {
		return nil, errors.New("string not closed in formula")
	}

	return tokens, nil
}

func allDoubloQuotesCloses(numDoubleQuote int, stringDeclared bool) bool {
	return numDoubleQuote%2 == 0
}

func appendToken(tokens []interface{}, token string, timeZone string) []interface{} {
	if num, err := strconv.Atoi(token); err == nil {
		tokens = append(tokens, num)
	} else if num, err := strconv.ParseFloat(token, 64); err == nil {
		tokens = append(tokens, num)
	} else if token == "true" {
		tokens = append(tokens, true)
	} else if token == "false" {
		tokens = append(tokens, false)
	} else if isDateValid(token) {
		tokens = append(tokens, strings.Trim(token, "\"")+" ("+timeZone+")")
	} else if isDateComparisonValid(token) {
		correctedToken, err := correctDateComparison(token)
		if err != nil {
			panic(err)
		}
		tokens = append(tokens, correctedToken)
	} else {
		tokens = append(tokens, strings.Trim(token, "\""))
		//tokens = append(tokens, token)
	}
	return tokens
}

func isDateValid(input string) bool {
	// Define the regular expression pattern
	datePattern := `^"\d{4}-\d{2}-\d{2}at\d{1,2}:\d{2}(am|pm)"$`
	datePattern2 := `^"\d{4}-\d{2}-\d{2} at \d{1,2}:\d{2}(am|pm)"$`
	// Compile the regular expression
	regexpDate := regexp.MustCompile(datePattern)
	regexpDate2 := regexp.MustCompile(datePattern2)
	// Test if the input string matches the pattern
	return regexpDate.MatchString(input) || regexpDate2.MatchString(input)
}

func isDateComparisonValid(input string) bool {
	//pattern := `^"(\d{4}-\d{2}-\d{2}at\d{1,2}:\d{2}[ap]m)"\s*(>=|>|=|<=)\s*"(\d{4}-\d{2}-\d{2}at\d{1,2}:\d{2}[ap]m)"$`
	pattern2 := `^"\d{4}-\d{2}-\d{2} at \d{1,2}:\d{2}(am|pm)"(>=|>|=|<=)"\d{4}-\d{2}-\d{2} at \d{1,2}:\d{2}(am|pm)"$`
	//regexp := regexp.MustCompile(pattern)
	regexp := regexp.MustCompile(pattern2)
	// Test if the input string matches the pattern
	return regexp.MatchString(input)
}

func correctDateComparison(input string) (string, error) {
	//pattern := `^"(\d{4}-\d{2}-\d{2}at\d{1,2}:\d{2}[ap]m)"\s*(>=|>|=|<=)\s*"(\d{4}-\d{2}-\d{2}at\d{1,2}:\d{2}[ap]m)"$`

	pattern2 := `^"(\d{4}-\d{2}-\d{2} at \d{1,2}:\d{2}[ap]m)"(>=|>|=|<=)"(\d{4}-\d{2}-\d{2} at \d{1,2}:\d{2}[ap]m)"$`
	//rege := regexp.MustCompile(pattern)
	regexp2 := regexp.MustCompile(pattern2)
	matches := regexp2.FindStringSubmatch(input)
	fmt.Printf("matches Formula :  %#v\n", matches)
	if len(matches) == 0 || len(matches) != 4 {
		return "", errors.New("parsing formula error")
	}

	date1 := matches[1]
	operator := matches[2]
	date2 := matches[3]

	return fmt.Sprintf("%v (UTC) %v %v (UTC)", date1, operator, date2), nil

}

func (f *Formula) ReplaceFieldsWithValue() (string, error) {
	if _, err := f.BuildRegexFromFields(); err != nil {
		return "", err
	}
	placeholderRegex := regexp.MustCompile(f.formulaRegex)
	replacedStr := placeholderRegex.ReplaceAllStringFunc(f.RawFormula, func(match string) string {
		return fmt.Sprintf("%v", f.FieldValue[match])
	})
	f.RawFormula = replacedStr
	return replacedStr, nil
}

func (f *Formula) BuildRegexFromFields() (string, error) {
	if f.FieldValue == nil {
		return "", errors.New("field and values are not initialized")
	}

	if len(f.FieldValue) == 0 {
		f.formulaRegex = f.RawFormula
		return f.RawFormula, nil
	}

	// Iterate over the map using the sorted keys
	var regex string = `\b(?:`
	var i int = 0
	for key := range f.FieldValue {
		if i < len(f.FieldValue)-1 {
			regex += `` + fmt.Sprintf("%v", key) + `|`
		} else {
			regex += `` + fmt.Sprintf("%v", key)
		}
		i++
	}
	regex += `)\b`
	f.formulaRegex = regex
	return regex, nil
}

func ContainsSpecialChars(inputStr string) bool {
	// Define a regular expression to match the specified characters outside double-quoted strings
	re := regexp.MustCompile(`(=|>=|<>|<=|<|>|!=|\+|-|/|%|\*|\^)|(\".*?\")`)

	// Match special characters or double-quoted strings
	matches := re.FindAllString(inputStr, -1)

	// If there are no matches, or if all matches are double-quoted strings, return false
	if len(matches) == 0 {
		return false
	}

	for _, match := range matches {
		if match[0] != '"' {
			return true
		}
	}

	return false
}
