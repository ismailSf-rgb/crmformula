package crmformula

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMostNestedArray1(t *testing.T) {
	valop := []interface{}{"OR", "(", "ADD", "(", 2, 1, 3, ")", true, ")"}
	got := FindMostNestedArray(valop)
	want := []interface{}{"ADD", 2, 1, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFindMostNestedArray2(t *testing.T) {
	valop := []interface{}{"ADD", "(", 2, 1, 3, ")"}
	got := FindMostNestedArray(valop)
	want := []interface{}{"ADD", 2, 1, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFindMostNestedArrayAndReplace1(t *testing.T) {
	valop := []interface{}{"OR", "(", "ADD", "(", 2, 1, 3, ")", true, ")"}
	got1, got2 := FindMostNestedArrayAndReplace(valop)
	want1 := []interface{}{"ADD", 2, 1, 3}
	want2 := []interface{}{"OR", "(", "?", true, ")"}

	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("got %v, wanted %v", got1, want1)
	}
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %v, wanted %v", got2, want2)
	}
}

func TestFindMostNestedArrayAndEvaluate(t *testing.T) {
	valop := []interface{}{"OR", "(", "ADD", "(", 2, 1, 3, ")", true, ")"}
	got := FindMostNestedArrayAndEvaluate(valop)
	want := []interface{}{"OR", "(", 6, true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetFormulaResult1(t *testing.T) {
	valop := []interface{}{"OR", "(", "AND", "(", true, false, false, ")", true, ")"}
	got := FindMostNestedArrayAndEvaluate(valop)
	want := []interface{}{"OR", "(", false, true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := true

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetFormulaResult2(t *testing.T) {
	valop := []interface{}{"IF", "(", "AND", "(", true, false, false, ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if false"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult3(t *testing.T) {
	valop := []interface{}{"IF", "(", "AND", "(", true, true, true, ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult4(t *testing.T) {
	valop := []interface{}{"IF", "(", "5 >= 3", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult5(t *testing.T) {
	valop := []interface{}{"IF", "(", "5 = 3", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if false"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult6(t *testing.T) {
	valop := []interface{}{"IF", "(", "1 <= 3", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult7(t *testing.T) {
	valop := []interface{}{"IF", "(", "AND", "(", "3 < 5", true, true, ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult8(t *testing.T) {
	valop := []interface{}{"IF", "(", "3.5 >= 3", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult9(t *testing.T) {
	valop := []interface{}{"IF", "(", "2.5 <= 3", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult10(t *testing.T) {
	valop := []interface{}{"IF", "(", "3.0 = 3", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult11(t *testing.T) {
	valop := []interface{}{"IF", "(", "AND", "(", "3.5 < 5", true, true, ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult12(t *testing.T) {
	valop := []interface{}{"IF", "(", "AND", "(", "2023-02-01 at 5:06pm (UTC) >= 2022-02-01 at 5:04pm (UTC)", true, true, ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult13(t *testing.T) {
	valop := []interface{}{"IF", "(", "AND", "(", "2023-02-01 at 5:06pm (UTC) = 2023-02-01 at 5:06pm (UTC)", true, true, ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult14(t *testing.T) {
	valop := []interface{}{"IF", "(", "2023-02-01 at 5:06pm (UTC) >= 2022-02-01 at 5:04pm (UTC)", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult141(t *testing.T) {
	valop := []interface{}{"IF", "(", "2023-02-01 at 5:06pm (UTC) <= 2022-02-01 at 5:04pm (UTC)", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if false"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult15(t *testing.T) {
	valop := []interface{}{"IF", "(", "2023-02-01 at 5:06pm (UTC) = 2023-02-01 at 5:06pm (UTC)", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult151(t *testing.T) {
	valop := []interface{}{"IF", "(", "2023-02-01 at 5:06pm (UTC) = 2023-02-02 at 5:06pm (UTC)", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if false"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult16(t *testing.T) {
	valop := []interface{}{"IF", "(", "2020-02-01 at 5:06pm (UTC) <= 2023-02-01 at 5:06pm (UTC)", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult17(t *testing.T) {
	valop := []interface{}{"IF", "(", "2020-02-01 at 5:06pm (UTC) < 2023-02-01 at 5:06pm (UTC)", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult18(t *testing.T) {
	valop := []interface{}{"IF", "(", "CONTAINS", "(", "abcd", "ab", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult19(t *testing.T) {
	valop := []interface{}{"MORE", "(", "MAX", "(", 2, 1, ")", 3, ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := false

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult20(t *testing.T) {
	valop := []interface{}{"LESS", "(", "MAX", "(", 2, 1, ")", 3, ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := true

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestTransformOp1(t *testing.T) {
	valop := []interface{}{MAX, "(", 2, 1, ")", "<", 3}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{LESS, "(", MAX, "(", 2, 1, ")", 3, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp2(t *testing.T) {
	valop := []interface{}{OR, "(", MAX, "(", 2, 1, ")", "<", 3, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", LESS, "(", MAX, "(", 2, 1, ")", 3, ")", ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp3(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, MAX, "(", 2, 1, ")", "<", 3, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, LESS, "(", MAX, "(", 2, 1, ")", 3, ")", ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp4(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, MAX, "(", 2, 1, ")", "<", 3, true, true, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, LESS, "(", MAX, "(", 2, 1, ")", 3, ")", true, true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp5(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, MAX, "(", 2, 1, ")", "<", 3, true, true, MIN, "(", 2, 1, ")", "<", 3, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, LESS, "(", MAX, "(", 2, 1, ")", 3, ")", true, true, LESS, "(", MIN, "(", 2, 1, ")", 3, ")", ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp6(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, MAX, "(", 2, 1, ")", ">", 3, true, true, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, MORE, "(", MAX, "(", 2, 1, ")", 3, ")", true, true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp7(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, MAX, "(", 2, 1, ")", "=", 3, true, true, MIN, "(", 2, 1, ")", "=", 3, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, EQUAL, "(", MAX, "(", 2, 1, ")", 3, ")", true, true, EQUAL, "(", MIN, "(", 2, 1, ")", 3, ")", ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp8(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, MAX, "(", 2, 1, ")", ">=", 3, true, true, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, MORE_EQ, "(", MAX, "(", 2, 1, ")", 3, ")", true, true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp9(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, MAX, "(", 2, 1, ")", "<=", 3, true, true, MIN, "(", 2, 1, ")", "<=", 3, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, LESS_EQ, "(", MAX, "(", 2, 1, ")", 3, ")", true, true, LESS_EQ, "(", MIN, "(", 2, 1, ")", 3, ")", ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp9_1(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, MAX, "(", 2, 1, ")", "<", 3, true, true, MIN, "(", 2, 1, ")", "<", 3, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, LESS, "(", MAX, "(", 2, 1, ")", 3, ")", true, true, LESS, "(", MIN, "(", 2, 1, ")", 3, ")", ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp10(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, 3, "<", MAX, "(", 4, 1, ")", ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, LESS, "(", 3, MAX, "(", 4, 1, ")", ")", ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp11(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, 3, "<", MAX, "(", 4, 1, 3, ")", ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, LESS, "(", 3, MAX, "(", 4, 1, 3, ")", ")", ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp12(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, 3, "<", MAX, "(", MAX, "(", 4, 1, 3, ")", ")", ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, LESS, "(", 3, MAX, "(", MAX, "(", 4, 1, 3, ")", ")", ")", ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp13(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, 3, "<", MAX, "(", MAX, "(", 4, 1, 3, ")", ")", true, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, LESS, "(", 3, MAX, "(", MAX, "(", 4, 1, 3, ")", ")", ")", true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp14(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, 3, "<", MAX, "(", MAX, "(", 4, 1, 3, ")", ")", true, true, true, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, LESS, "(", 3, MAX, "(", MAX, "(", 4, 1, 3, ")", ")", ")", true, true, true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp15(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, 3, "<", MAX, "(", MAX, "(", 4, 1, 3, ")", ")", true, true, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, LESS, "(", 3, MAX, "(", MAX, "(", 4, 1, 3, ")", ")", ")", true, true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp16(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, 3, ">", MAX, "(", MAX, "(", 4, 1, 3, ")", ")", true, true, true, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, MORE, "(", 3, MAX, "(", MAX, "(", 4, 1, 3, ")", ")", ")", true, true, true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp17(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, 3, "<=", MAX, "(", MAX, "(", 4, 1, 3, ")", ")", true, true, true, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, LESS_EQ, "(", 3, MAX, "(", MAX, "(", 4, 1, 3, ")", ")", ")", true, true, true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp18(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, 3, "=", MAX, "(", MAX, "(", 4, 1, 3, ")", ")", true, true, true, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, EQUAL, "(", 3, MAX, "(", MAX, "(", 4, 1, 3, ")", ")", ")", true, true, true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp19(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, 3, "=", MAX, "(", MAX, "(", 4, 1, 3, ")", ")", true, true, true, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, EQUAL, "(", 3, MAX, "(", MAX, "(", 4, 1, 3, ")", ")", ")", true, true, true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp20(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, 3, "=", MAX, "(", MAX, "(", 4, 1, 3, ")", 2, ")", true, true, true, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, EQUAL, "(", 3, MAX, "(", MAX, "(", 4, 1, 3, ")", 2, ")", ")", true, true, true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestTransformOp21(t *testing.T) {
	valop := []interface{}{OR, "(", true, true, 3, "=", MAX, "(", MAX, "(", MAX, "(", 4, 1, 3, ")", 2, ")", 5, ")", true, true, true, ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	got := valop
	want := []interface{}{OR, "(", true, true, EQUAL, "(", 3, MAX, "(", MAX, "(", MAX, "(", 4, 1, 3, ")", 2, ")", 5, ")", ")", true, true, true, ")"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetFormulaResult21(t *testing.T) {
	valop := []interface{}{"IF", "(", "MAX", "(", 2, 1, ")", "<", 3, "result if true", "result if false", ")"}
	valop = Reformat(valop)
	fmt.Println(valop)
	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult22(t *testing.T) {
	valop := []interface{}{"OR", "(", "MAX", "(", 2, 5, ")", "<", 3, ")"}
	valop = Reformat(valop)
	FinalValGot := GetFormulaResult(valop)
	FinalValWant := false

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult23(t *testing.T) {
	valop := []interface{}{"IF", "(", "LESS", "(", "2020-02-01 at 5:06pm (UTC)", "2023-02-01 at 5:06pm (UTC)", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult24(t *testing.T) {
	valop := []interface{}{"IF", "(", "EQUAL", "(", "2020-02-01 at 5:06pm (UTC)", "2020-02-01 at 5:06pm (UTC)", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult25(t *testing.T) {
	valop := []interface{}{"IF", "(", "EQUAL", "(", "2020-02-01 at 5:06pm (UTC)", "2023-02-01 at 5:06pm (UTC)", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if false"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult26(t *testing.T) {
	valop := []interface{}{"IF", "(", MORE, "(", "2020-02-01 at 5:06pm (UTC)", "2023-02-01 at 5:06pm (UTC)", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if false"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult27(t *testing.T) {
	valop := []interface{}{"IF", "(", LESS_EQ, "(", "2020-02-01 at 5:06pm (UTC)", "2020-02-01 at 5:06pm (UTC)", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult28(t *testing.T) {
	valop := []interface{}{"IF", "(", MORE_EQ, "(", "2020-02-01 at 5:06pm (UTC)", "2023-02-01 at 5:06pm (UTC)", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if false"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult29(t *testing.T) {
	valop := []interface{}{"IF", "(", MORE_EQ, "(", NOW, "(", ")", "2021-02-01 at 5:06pm (UTC)", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult30(t *testing.T) {
	valop := []interface{}{"IF", "(", LESS, "(", NOW, "(", ")", "2025-02-01 at 5:06pm (UTC)", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult31(t *testing.T) {
	valop := []interface{}{"IF", "(", MORE_EQ, "(", "2020-02-01 at 5:06pm (UTC)", NOW, "(", ")", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if false"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult32(t *testing.T) {
	valop := []interface{}{"IF", "(", LESS_EQ, "(", "2020-02-01 at 5:06pm (UTC)", NOW, "(", ")", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult33(t *testing.T) {
	valop := []interface{}{"IF", "(", "2020-02-01 at 5:06pm (UTC)", "<=", NOW, "(", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult34(t *testing.T) {
	valop := []interface{}{"LOG", "(", 25.0, ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := 3.2188758248682006

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult35(t *testing.T) {
	valop := []interface{}{"LOG", "(", "LOG", "(", 25.0, ")", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := 1.1690321758870559

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult36(t *testing.T) {
	valop := []interface{}{"ADD", "(", "FIND", "(", "g", "lego", ")", 25, ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := 27

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult37(t *testing.T) {
	valop := []interface{}{"IF", "(", MORE_EQ, "(", TODAY, "(", ")", "2021-02-01 at 5:06pm (UTC)", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult38(t *testing.T) {
	valop := []interface{}{"IF", "(", LESS, "(", TODAY, "(", ")", "2025-02-01 at 5:06pm (UTC)", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult39(t *testing.T) {
	valop := []interface{}{"IF", "(", BEGINS, "(", "abcd", "ab", ")", "result if true", "result if false", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "result if true"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult40(t *testing.T) {
	valop := []interface{}{CONCAT, "(", "abcd", "ef", BR, "(", ")", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "abcdef\n"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult41(t *testing.T) {
	valop := []interface{}{CONCAT, "(", "abcd", "ef", BR, "(", ")", "gh", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := "abcdef\ngh"

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult42(t *testing.T) {
	valop := []interface{}{YEAR, "(", TODAY, "(", ")", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := 2023

	if FinalValGot.(int) < FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult43(t *testing.T) {
	valop := []interface{}{YEAR, "(", NOW, "(", ")", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := 2023

	if FinalValGot.(int) < FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult44(t *testing.T) {
	valop := []interface{}{ADD, "(", YEAR, "(", "2025-02-02at5:06pm(UTC)", ")", YEAR, "(", NOW, "(", ")", ")", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := 4048

	if FinalValGot.(int) < FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult45(t *testing.T) {
	valop := []interface{}{ADD, "(", MONTH, "(", "2025-02-02 at 5:06pm(UTC)", ")", MONTH, "(", "2025-05-02 at 5:06pm(UTC)", ")", ")"}

	FinalValGot := GetFormulaResult(valop)
	FinalValWant := 7

	if FinalValGot != FinalValWant {
		t.Errorf("got %v, wanted %v", FinalValGot, FinalValWant)
	}
}

func TestGetFormulaResult46(t *testing.T) {
	valop := []interface{}{EQUAL, "(", LOWER, "(", "\"abA\"", ")", "\"aba\"", ")"}

	got := GetFormulaResult(valop)
	want := true

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
