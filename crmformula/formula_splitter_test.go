package crmformula

import (
	"reflect"
	"testing"
)

func TestFindOperators1(t *testing.T) {
	str := "OR(AND(true, false, false))"
	f := Formula{
		RawFormula: str,
	}
	got, err := f.FindOperators()
	want := []string{"OR", "AND"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if !reflect.DeepEqual(f.operators, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if f.numberOfOperand != 2 {
		t.Errorf("number of operand error got %d, wanted %d", f.numberOfOperand, 2)
	}

	lengthGot := f.length
	lengthWant := 25

	if lengthGot != lengthWant {
		t.Errorf("length error got %d, wanted %d", lengthGot, lengthWant)
	}
}

func TestFindOperators2(t *testing.T) {
	str := "OR(OR(AND(true, false, false), true))"
	f := Formula{
		RawFormula: str,
	}
	got, err := f.FindOperators()
	want := []string{"OR", "OR", "AND"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if !reflect.DeepEqual(f.operators, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if f.numberOfOperand != 3 {
		t.Errorf("number of operand error got %d, wanted %d", f.numberOfOperand, 3)
	}

}

func TestFindOperators3(t *testing.T) {
	str := "AND(AND(OR(true, false, false), true))"
	f := Formula{
		RawFormula: str,
	}
	got, err := f.FindOperators()
	want := []string{"AND", "AND", "OR"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if !reflect.DeepEqual(f.operators, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if f.numberOfOperand != 3 {
		t.Errorf("number of operand error got %d, wanted %d", f.numberOfOperand, 3)
	}

}

func TestFindOperatorsAndParenthesis1(t *testing.T) {
	str := "AND(AND(OR(true, false, false), true))"
	f := Formula{
		RawFormula: str,
	}
	got, err := f.FindOperators()
	want := []string{"AND", "AND", "OR"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = f.FindOperatorsAndParenthesis()
	want = []string{"AND", "(", "AND", "(", "OR", "(", ")", ")", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotType := f.formulaType
	wantType := Checkbox

	if !reflect.DeepEqual(gotType, wantType) {
		t.Errorf("type got %v, type wanted %v", got, want)
	}

	t.Logf("opIndex : %v \n", f.opIndex)
	t.Logf("parIndex : %v \n", f.parIndex)
	t.Logf("operators and parenthesis list : %v \n", f.midFormula)
	t.Logf("type of the formula: %v \n", f.formulaType)

}

func TestFindOperatorsAndParenthesis2(t *testing.T) {
	str := "OR(AND(true, false, false))"
	f := Formula{
		RawFormula: str,
	}
	got, err := f.FindOperators()
	want := []string{"OR", "AND"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = f.FindOperatorsAndParenthesis()
	want = []string{"OR", "(", "AND", "(", ")", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	t.Logf("opIndex : %v \n", f.opIndex)
	t.Logf("parIndex : %v \n", f.parIndex)
	t.Logf("operators and parenthesis list : %v \n", f.midFormula)
	t.Logf("empty formula with values only : %v \n", f.emptyFormula)

}

func TestFindOperatorsAndParenthesisError1(t *testing.T) {
	//one more parenthesis
	str := "OR(AND(true, false, false)))"
	f := Formula{
		RawFormula: str,
	}
	got, err := f.FindOperators()
	want := []string{"OR", "AND"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = f.FindOperatorsAndParenthesis()
	wantErr := "compiling formula error"

	if got != nil && !reflect.DeepEqual(err.Error(), wantErr) {
		t.Errorf("got %v, wanted error %v", err.Error(), wantErr)
	}

}

func TestFindOperatorsAndParenthesisAndValues1(t *testing.T) {
	str := "OR(AND(true, false),false)"
	f := Formula{
		RawFormula: str,
	}
	got, err := f.FindOperators()
	want := []string{"OR", "AND"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = f.FindOperatorsAndParenthesis()
	want = []string{"OR", "(", "AND", "(", ")", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotFinal, err := f.Tokenize()
	var wantFinal []interface{} = []interface{}{"OR", "(", "AND", "(", true, false, ")", false, ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}
	t.Logf("final Formula :  %#v\n", gotFinal)
	t.Logf("expected Formula :  %#v\n", wantFinal)
	if !reflect.DeepEqual(gotFinal, wantFinal) {
		t.Errorf("got %v, wanted %v", gotFinal, wantFinal)
	}

}

//for floats doesn' work []interface {}{"ADD", "(", "MULT", "(", 1, 2, ")", 3, 0.5, ")"}
func TestFindOperatorsAndParenthesisAndValues2(t *testing.T) {
	str := "ADD(MULT(1, 2),3.5)"
	f := Formula{
		RawFormula: str,
	}
	got, err := f.FindOperators()
	want := []string{"ADD", "MULT"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = f.FindOperatorsAndParenthesis()
	want = []string{"ADD", "(", "MULT", "(", ")", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotFinal, err := f.Tokenize()
	var wantFinal []interface{} = []interface{}{"ADD", "(", "MULT", "(", 1, 2, ")", 3.5, ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}
	t.Logf("final Formula :  %#v\n", gotFinal)
	t.Logf("expected Formula :  %#v\n", wantFinal)

	if !reflect.DeepEqual(gotFinal, wantFinal) {
		t.Errorf("got %v, wanted %v", gotFinal, wantFinal)
	}

}

func TestFindOperatorsAndParenthesisAndValues3(t *testing.T) {
	str := "LOWER(\"ABCd\")"
	f := Formula{
		RawFormula: str,
	}
	got, err := f.FindOperators()
	want := []string{"LOWER"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = f.FindOperatorsAndParenthesis()
	want = []string{"LOWER", "(", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotFinal, err := f.Tokenize()
	var wantFinal []interface{} = []interface{}{"LOWER", "(", "ABCd", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}
	t.Logf("final Formula :  %#v\n", gotFinal)
	t.Logf("expected Formula :  %#v\n", wantFinal)

	if !reflect.DeepEqual(gotFinal, wantFinal) {
		t.Errorf("got %v, wanted %v", gotFinal, wantFinal)
	}
}

//failing because UTC is put in the next element of the array
func TestFindOperatorsAndParenthesisAndValues4(t *testing.T) {
	str := "ADDMONTH(\"2025-02-01 at 5:06pm\", 5)"
	f := Formula{
		RawFormula: str,
		TimeZone:   "UTC",
	}
	got, err := f.FindOperators()
	want := []string{"ADDMONTH"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = f.FindOperatorsAndParenthesis()
	want = []string{"ADDMONTH", "(", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotFinal, err := f.Tokenize()
	var wantFinal []interface{} = []interface{}{"ADDMONTH", "(", "2025-02-01 at 5:06pm (UTC)", 5, ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}
	t.Logf("final Formula :  %#v\n", gotFinal)
	t.Logf("expected Formula :  %#v\n", wantFinal)

	if !reflect.DeepEqual(gotFinal, wantFinal) {
		t.Errorf("got %v, wanted %v", gotFinal, wantFinal)
	}
}

//Failing because remove spaces
func TestFindOperatorsAndParenthesisAndValues5(t *testing.T) {
	str := "LOWER(\"ABC   d   \")"
	//str = strings.ReplaceAll(str, " ", "/s")
	f := Formula{
		RawFormula: str,
	}
	got, err := f.FindOperators()
	want := []string{"LOWER"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = f.FindOperatorsAndParenthesis()
	want = []string{"LOWER", "(", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotFinal, err := f.Tokenize()
	var wantFinal []interface{} = []interface{}{"LOWER", "(", "ABC   d   ", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}
	t.Logf("final Formula :  %#v\n", gotFinal)
	t.Logf("expected Formula :  %#v\n", wantFinal)

	if !reflect.DeepEqual(gotFinal, wantFinal) {
		t.Errorf("got %v, wanted %v", gotFinal, wantFinal)
	}
}

func TestFindOperatorsAndParenthesisAndValues6(t *testing.T) {
	str := "IF(3.0 = 3, \"result if true\", \"result if false\")"

	f := Formula{
		RawFormula: str,
	}
	got, err := f.FindOperators()
	want := []string{"IF"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = f.FindOperatorsAndParenthesis()
	want = []string{"IF", "(", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotFinal, err := f.Tokenize()
	var wantFinal []interface{} = []interface{}{"IF", "(", "3.0=3", "result if true", "result if false", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}
	t.Logf("final Formula :  %#v\n", gotFinal)
	t.Logf("expected Formula :  %#v\n", wantFinal)

	if !reflect.DeepEqual(gotFinal, wantFinal) {
		t.Errorf("got %v, wanted %v", gotFinal, wantFinal)
	}

	FinalValGot := GetFormulaResult(gotFinal)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}

}

func TestFindOperatorsAndParenthesisAndValues7WithError(t *testing.T) {
	str := "ADD(3,\"result if true\", \"result if false\")"

	f := Formula{
		RawFormula: str,
	}
	got, err := f.FindOperators()
	want := []string{"ADD"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = f.FindOperatorsAndParenthesis()
	want = []string{"ADD", "(", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotFinal, err := f.Tokenize()
	var wantFinal []interface{} = []interface{}{"ADD", "(", 3, "result if true", "result if false", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}
	t.Logf("final Formula :  %#v\n", gotFinal)
	t.Logf("expected Formula :  %#v\n", wantFinal)

	if !reflect.DeepEqual(gotFinal, wantFinal) {
		t.Errorf("got %v, wanted %v", gotFinal, wantFinal)
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as expected")
		}
	}()
	FinalValGot := GetFormulaResult(gotFinal)

	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}

}

func TestFindOperatorsAndParenthesisAndValues8(t *testing.T) {
	mapFieldValues := make(map[string]interface{})
	mapFieldValues["city"] = "\"Example City\""
	mapFieldValues["country"] = "\"Example Country\""
	mapFieldValues["number"] = 3

	str := "IF(3.0 = number,country, city)"

	f := Formula{
		RawFormula: str,
		FieldValue: mapFieldValues,
	}
	FinalValGot, err := f.BuildRegexFromFields()
	FinalValWant1 := `\b(?:city|country|number)\b`
	FinalValWant2 := `\b(?:country|city|number)\b`
	FinalValWant3 := `\b(?:country|number|city)\b`
	FinalValWant4 := `\b(?:city|number|country)\b`
	FinalValWant5 := `\b(?:number|country|city)\b`
	FinalValWant6 := `\b(?:number|city|country)\b`
	if err != nil {
		t.Errorf("got error %v", err)
	}

	if FinalValGot != FinalValWant1 && FinalValGot != FinalValWant2 && FinalValGot != FinalValWant3 && FinalValGot != FinalValWant4 && FinalValGot != FinalValWant5 && FinalValGot != FinalValWant6 {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant1+" in every possible order")
	}

	formulaGot, err := f.ReplaceFieldsWithValue()

	if err != nil {
		t.Errorf("got error %v", err)
	}

	formulaWant := "IF(3.0 = 3,\"Example Country\", \"Example City\")"

	if formulaGot != formulaWant {
		t.Errorf("got %v, wanted %v", formulaGot, formulaWant)
	}

	got, err := f.FindOperators()
	want := []string{"IF"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = f.FindOperatorsAndParenthesis()
	want = []string{"IF", "(", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotFinal, err := f.Tokenize()
	var wantFinal []interface{} = []interface{}{"IF", "(", "3.0=3", "Example Country", "Example City", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}
	t.Logf("final Formula :  %#v\n", gotFinal)
	t.Logf("expected Formula :  %#v\n", wantFinal)

	if !reflect.DeepEqual(gotFinal, wantFinal) {
		t.Errorf("got %v, wanted %v", gotFinal, wantFinal)
	}

	FinalValGotResult := GetFormulaResult(gotFinal)
	FinalValWantResult := "Example Country"

	if FinalValGotResult != FinalValWantResult {
		t.Errorf("got %v, wanted %v", FinalValGotResult, FinalValWantResult)
	}

}

func TestFindOperatorsAndParenthesisAndValues9(t *testing.T) {
	str := "IF(\"2023-02-01 at 5:06pm\" >= \"2022-02-01 at 5:04pm\",  \"result if true\", \"result if false\")"
	f := Formula{
		RawFormula: str,
		TimeZone:   "UTC",
	}

	got, err := f.FindOperators()
	want := []string{"IF"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = f.FindOperatorsAndParenthesis()
	want = []string{"IF", "(", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotFinal, err := f.Tokenize()
	var wantFinal []interface{} = []interface{}{"IF", "(", "2023-02-01 at 5:06pm (UTC) >= 2022-02-01 at 5:04pm (UTC)", "result if true", "result if false", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}
	t.Logf("final Formula :  %#v\n", gotFinal)
	t.Logf("expected Formula :  %#v\n", wantFinal)

	if !reflect.DeepEqual(gotFinal, wantFinal) {
		t.Errorf("got %v, wanted %v", gotFinal, wantFinal)
	}

	FinalValGotResult := GetFormulaResult(gotFinal)
	FinalValWantResult := "result if true"

	if FinalValGotResult != FinalValWantResult {
		t.Errorf("got %v, wanted %v", FinalValGotResult, FinalValWantResult)
	}

}

func TestFindOperatorsAndParenthesisAndValues10(t *testing.T) {
	str := "IF(\"2021-02-01 at 5:06pm\">=\"2022-02-01 at 5:04pm\", \"result if true\", \"result if false\")"
	f := Formula{
		RawFormula: str,
		TimeZone:   "UTC",
	}

	got, err := f.FindOperators()
	want := []string{"IF"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = f.FindOperatorsAndParenthesis()
	want = []string{"IF", "(", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotFinal, err := f.Tokenize()
	var wantFinal []interface{} = []interface{}{"IF", "(", "2021-02-01 at 5:06pm (UTC) >= 2022-02-01 at 5:04pm (UTC)", "result if true", "result if false", ")"}

	if err != nil {
		t.Errorf("got error %v", err)
	}
	t.Logf("final Formula :  %#v\n", gotFinal)
	t.Logf("expected Formula :  %#v\n", wantFinal)

	if !reflect.DeepEqual(gotFinal, wantFinal) {
		t.Errorf("got %v, wanted %v", gotFinal, wantFinal)
	}

	FinalValGotResult := GetFormulaResult(gotFinal)
	FinalValWantResult := "result if false"

	if FinalValGotResult != FinalValWantResult {
		t.Errorf("got %v, wanted %v", FinalValGotResult, FinalValWantResult)
	}
}
