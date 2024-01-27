package crmformula

import (
	"reflect"
	"strings"
	"testing"
)

func TestRunFormulaAdd1(t *testing.T) {
	valop := []interface{}{"ADD", 2, 1, 4, 5, 6}
	got := RunFormula(valop)
	want := 18

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRunFormulaMult1(t *testing.T) {
	valop := []interface{}{MULT, 2, 1, 4, 5, 6}
	got := RunFormula(valop)
	want := 240

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRunFormulaAdd2(t *testing.T) {
	valop := []interface{}{ADD, 2, -1, -4, -5, -6}
	got := RunFormula(valop)
	want := -14

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRunFormulaLeft1(t *testing.T) {
	valop := []interface{}{LEFT, "abcd", 2}
	got := RunFormula(valop)
	want := "ab"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRunFormulaRight1(t *testing.T) {
	valop := []interface{}{RIGHT, "abcd", 2}
	got := RunFormula(valop)
	want := "cd"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRunFormulaIf1(t *testing.T) {
	valop := []interface{}{TRUE, TRUE, 0, 1, 2}
	got := RunFormula(valop)
	want := 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRunFormulaIf2(t *testing.T) {
	valop := []interface{}{FALSE, 1, FALSE, 3, 5}
	got := RunFormula(valop)
	want := 5

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRunFormulaOr1(t *testing.T) {
	valop := []interface{}{OR, "2 < 3", false, false, false, false}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaAnd1(t *testing.T) {
	valop := []interface{}{AND, true, true, true, true, true}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaAdd3(t *testing.T) {
	valop := []interface{}{ADD, 1, ADD, 2, 1}
	got := RunFormula(valop)
	want := 4

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRunFormulaAdd4(t *testing.T) {
	valop := []interface{}{ADD, ADD, 1, 3, ADD, 2, 1}
	got := RunFormula(valop)
	want := 7

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRunFormulaAdd5(t *testing.T) {
	valop := []interface{}{ADD, ADD, 1.0, 3.5, ADD, 2.0, 1.0}
	got := RunFormula(valop)
	want := 7.5

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaAdd6(t *testing.T) {
	valop := []interface{}{ADD, ADD, 1.5, 3, ADD, 2, 1}
	got := RunFormula(valop)
	want := 7.5

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}

func TestRunFormulaMult2(t *testing.T) {
	valop := []interface{}{MULT, 2.5, 1.0, 4.0, 5.0, 6.0}
	got := RunFormula(valop)
	want := 300.0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaNot1(t *testing.T) {
	valop := []interface{}{NOT, true}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaNot2(t *testing.T) {
	valop := []interface{}{NOT, NOT, true}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaNot3(t *testing.T) {
	valop := []interface{}{NOT, NOT, NOT, true}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaDiv1(t *testing.T) {
	valop := []interface{}{DIV, 5, 2}
	got := RunFormula(valop)
	want := 2

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaDiv2(t *testing.T) {
	valop := []interface{}{DIV, 5.0, 2.0}
	got := RunFormula(valop)
	want := 2.5

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMod1(t *testing.T) {
	valop := []interface{}{MOD, 5, 2}
	got := RunFormula(valop)
	want := 1

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMod2(t *testing.T) {
	valop := []interface{}{MOD, 5.5, 2.0}
	got := RunFormula(valop)
	want := 1.5

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaIsBlank1(t *testing.T) {
	valop := []interface{}{ISBLANK, nil}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaIsBlank2(t *testing.T) {
	valop := []interface{}{ISBLANK, "not blank"}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaIsBlank3(t *testing.T) {
	valop := []interface{}{ISBLANK, ""}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaBlankValue1(t *testing.T) {
	valop := []interface{}{BLANKVALUE, "", 2}
	got := RunFormula(valop)
	want := 2

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaBlankValue2(t *testing.T) {
	valop := []interface{}{BLANKVALUE, "", BLANKVALUE, "", 3}
	got := RunFormula(valop)
	want := 3

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaBlankValue3(t *testing.T) {
	valop := []interface{}{BLANKVALUE, "", BLANKVALUE, "not blank", 3}
	got := RunFormula(valop)
	var want interface{} = nil

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaContains1(t *testing.T) {
	valop := []interface{}{CONTAINS, "abcd", "ab"}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaContains2(t *testing.T) {
	valop := []interface{}{CONTAINS, "abcd", "zz"}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaRegex1(t *testing.T) {
	valop := []interface{}{REGEX, "999-99-9999", "[0-9]{3}-[0-9]{2}-[0-9]{4}"}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaRegex2(t *testing.T) {
	valop := []interface{}{REGEX, "abcd", "[0-9]{3}-[0-9]{2}-[0-9]{4}"}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaExponential1(t *testing.T) {
	valop := []interface{}{EXP, 1.23}
	got := RunFormula(valop)
	want := 3.4212295362896734

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaExponential2(t *testing.T) {
	valop := []interface{}{EXP, EXP, 1.23}
	got := RunFormula(valop)
	want := 30.607024342447403

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaExponential3(t *testing.T) {
	valop := []interface{}{EXP, EXP, EXP, 1.23}
	got := RunFormula(valop)
	want := 1.960928594055171e+13

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaCosinus1(t *testing.T) {
	valop := []interface{}{COS, 1.23}
	got := RunFormula(valop)
	want := 0.3342377271245026

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaSinus1(t *testing.T) {
	valop := []interface{}{SIN, 1.23}
	got := RunFormula(valop)
	want := 0.9424888019316975

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLen1(t *testing.T) {
	valop := []interface{}{LEN, "abcd"}
	got := RunFormula(valop)
	want := 4

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLen2(t *testing.T) {
	valop := []interface{}{LEN, "abcdef"}
	got := RunFormula(valop)
	want := 6

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMax1(t *testing.T) {
	valop := []interface{}{MAX, 2, 3}
	got := RunFormula(valop)
	want := 3

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMax2(t *testing.T) {
	valop := []interface{}{MAX, 2, 3, 4, 5, 6}
	got := RunFormula(valop)
	want := 6

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMax3(t *testing.T) {
	valop := []interface{}{MAX, 2, 3.2, 4, 5.3, 6}
	got := RunFormula(valop)
	want := 6.0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMin1(t *testing.T) {
	valop := []interface{}{MIN, 2, 3}
	got := RunFormula(valop)
	want := 2

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMin2(t *testing.T) {
	valop := []interface{}{MIN, 2, 3, 4, 5, 6}
	got := RunFormula(valop)
	want := 2

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMin3(t *testing.T) {
	valop := []interface{}{MIN, 2, 3.2, 4, 5.3, 6}
	got := RunFormula(valop)
	want := 2.0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMoreEq1(t *testing.T) {
	valop := []interface{}{MORE_EQ, 10, 6}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMoreEq2(t *testing.T) {
	valop := []interface{}{MORE_EQ, 10, 6.5}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMoreEq3(t *testing.T) {
	valop := []interface{}{MORE_EQ, 10, 16.5}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLessEq1(t *testing.T) {
	valop := []interface{}{LESS_EQ, 10, 6}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLessEq2(t *testing.T) {
	valop := []interface{}{LESS_EQ, 10, 6.5}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLessEq3(t *testing.T) {
	valop := []interface{}{LESS_EQ, 10, 16.5}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLessEq4(t *testing.T) {
	valop := []interface{}{LESS_EQ, 10, 10.0}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMore1(t *testing.T) {
	valop := []interface{}{MORE, 10, 6}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMore2(t *testing.T) {
	valop := []interface{}{MORE, 10, 6.5}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaMore3(t *testing.T) {
	valop := []interface{}{MORE, 10, 16.5}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLess1(t *testing.T) {
	valop := []interface{}{LESS, 10, 6}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLess2(t *testing.T) {
	valop := []interface{}{LESS, 10, 6.5}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLess3(t *testing.T) {
	valop := []interface{}{LESS, 10, 16.5}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLess4(t *testing.T) {
	valop := []interface{}{LESS, 10, 10.0}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaEqual1(t *testing.T) {
	valop := []interface{}{EQUAL, 10, 6}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaEqual2(t *testing.T) {
	valop := []interface{}{EQUAL, 10, 6.5}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaEqual3(t *testing.T) {
	valop := []interface{}{EQUAL, 10, 16.5}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaEqual4(t *testing.T) {
	valop := []interface{}{EQUAL, 10, 10.0}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaToday1(t *testing.T) {
	valop := []interface{}{NOW}
	got := RunFormula(valop)
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.String {
		if !strings.Contains(got.(string), "(UTC)") && !strings.Contains(got.(string), "at") && strings.Count(got.(string), "-") != 2 {
			t.Errorf("got %v, wanted something else", got)
		}
	}

}

func TestRunFormulaCeiling1(t *testing.T) {
	valop := []interface{}{CEILING, 1.23}
	got := RunFormula(valop)
	want := 2.0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaFloor1(t *testing.T) {
	valop := []interface{}{FLOOR, 1.23}
	got := RunFormula(valop)
	want := 1.0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaFloor2(t *testing.T) {
	valop := []interface{}{FLOOR, -1.23}
	got := RunFormula(valop)
	want := -2.0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaRound1(t *testing.T) {
	valop := []interface{}{ROUND, 1.23}
	got := RunFormula(valop)
	want := 1.0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaRound2(t *testing.T) {
	valop := []interface{}{ROUND, 1.51}
	got := RunFormula(valop)
	want := 2.0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaSqrt1(t *testing.T) {
	valop := []interface{}{SQRT, 25.0}
	got := RunFormula(valop)
	want := 5.0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLog1(t *testing.T) {
	valop := []interface{}{LOG, 25.0}
	got := RunFormula(valop)
	want := 3.2188758248682006

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLog2(t *testing.T) {
	valop := []interface{}{LOG, 5.0}
	got := RunFormula(valop)
	want := 1.6094379124341003

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLog10_1(t *testing.T) {
	valop := []interface{}{LOG10, 5.0}
	got := RunFormula(valop)
	want := 0.6989700043360187

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaAbs1(t *testing.T) {
	valop := []interface{}{ABS, -5.0}
	got := RunFormula(valop)
	want := 5.0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaFind1(t *testing.T) {
	valop := []interface{}{FIND, "p", "Subscription"}
	got := RunFormula(valop)
	want := 7

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaFind2(t *testing.T) {
	valop := []interface{}{FIND, "Subs", "Subscription"}
	got := RunFormula(valop)
	want := 0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaFind3(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as expected")
		}
	}()
	valop := []interface{}{FIND, 2, "Subscription"}
	RunFormula(valop)

}

func TestRunFormulaTrim1(t *testing.T) {
	valop := []interface{}{TRIM, "    Subscription     "}
	got := RunFormula(valop)
	want := "Subscription"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func TestRunFormulaToday2(t *testing.T) {
	valop := []interface{}{TODAY}
	got := RunFormula(valop)
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.String {
		if !strings.Contains(got.(string), "(UTC)") && !strings.Contains(got.(string), "at") && strings.Count(got.(string), "-") != 2 {
			t.Errorf("got %v, wanted something else", got)
		}
	}

}

func TestRunFormulaAddMonth1(t *testing.T) {
	valop := []interface{}{ADDMONTH, "2025-02-01 at 5:06pm (UTC)", 5}
	got := RunFormula(valop)
	want := "2025-07-01 at 5:06pm (UTC)"
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.String {
		if !strings.Contains(got.(string), "(UTC)") && !strings.Contains(got.(string), "at") && strings.Count(got.(string), "-") != 2 {
			t.Errorf("got %v, wanted something else", got)
		}
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func TestRunFormulaAddMonth2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as expected")
		}
	}()
	valop := []interface{}{ADDMONTH, 2, "Subscription"}
	RunFormula(valop)

}

func TestRunFormulaYear1(t *testing.T) {
	valop := []interface{}{YEAR, "2025-02-02at5:06pm(UTC)"}
	got := RunFormula(valop)
	want := 2025
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Int {
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func TestRunFormulaYear2(t *testing.T) {
	valop := []interface{}{YEAR, "2025-02-02 at 5:06pm(UTC)"}
	got := RunFormula(valop)
	want := 2025
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Int {
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func TestRunFormulaMonth1(t *testing.T) {
	valop := []interface{}{MONTH, "2025-02-02at5:06pm(UTC)"}
	got := RunFormula(valop)
	want := 2
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Int {
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func TestRunFormulaMonth2(t *testing.T) {
	valop := []interface{}{MONTH, "2025-02-02 at 5:06pm(UTC)"}
	got := RunFormula(valop)
	want := 2
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Int {
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func TestRunFormulaDay1(t *testing.T) {
	valop := []interface{}{DAY, "2025-02-02at5:06pm(UTC)"}
	got := RunFormula(valop)
	want := 2
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Int {
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func TestRunFormulaDay2(t *testing.T) {
	valop := []interface{}{DAY, "2025-02-03 at 5:06pm(UTC)"}
	got := RunFormula(valop)
	want := 3
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Int {
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func TestRunFormulaWeekDay1(t *testing.T) {
	valop := []interface{}{WEEKDAY, "2023-05-30at5:06pm(UTC)"}
	got := RunFormula(valop)
	want := 2
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Int {
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func TestRunFormulaWeekDay2(t *testing.T) {
	valop := []interface{}{WEEKDAY, "2023-05-31 at 5:06pm(UTC)"}
	got := RunFormula(valop)
	want := 3
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Int {
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func TestRunFormulaHour1(t *testing.T) {
	valop := []interface{}{HOUR, "2023-05-30at6:06pm(UTC)"}
	got := RunFormula(valop)
	want := 18
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Int {
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func TestRunFormulaHour2(t *testing.T) {
	valop := []interface{}{HOUR, "2023-05-31 at 5:06pm(UTC)"}
	got := RunFormula(valop)
	want := 17
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Int {
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func TestRunFormulaMinute1(t *testing.T) {
	valop := []interface{}{MINUTE, "2023-05-30at6:06pm(UTC)"}
	got := RunFormula(valop)
	want := 6
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Int {
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func TestRunFormulaMinute2(t *testing.T) {
	valop := []interface{}{MINUTE, "2023-05-31 at 5:07pm(UTC)"}
	got := RunFormula(valop)
	want := 7
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Int {
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func TestRunFormulaBegins1(t *testing.T) {
	valop := []interface{}{BEGINS, "abcd", "ab"}
	got := RunFormula(valop)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaBegins2(t *testing.T) {
	valop := []interface{}{BEGINS, "abcd", "cd"}
	got := RunFormula(valop)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaConcat1(t *testing.T) {
	valop := []interface{}{CONCAT, "abcd", "ef"}
	got := RunFormula(valop)
	want := "abcdef"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaConcat2(t *testing.T) {
	valop := []interface{}{CONCAT, "abcd", " ef"}
	got := RunFormula(valop)
	want := "abcd ef"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaConcat3(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as expected")
		}
	}()
	valop := []interface{}{CONCAT, 2, "Subscription"}
	RunFormula(valop)

}

func TestRunFormulaConcat4(t *testing.T) {
	valop := []interface{}{CONCAT, "abcd", "ef", "gh", "ij"}
	got := RunFormula(valop)
	want := "abcdefghij"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaBr1(t *testing.T) {
	valop := []interface{}{BR}
	got := RunFormula(valop)
	want := "\n"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaUpper1(t *testing.T) {
	valop := []interface{}{UPPER, "abcd"}
	got := RunFormula(valop)
	want := "ABCD"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaUpper2(t *testing.T) {
	valop := []interface{}{UPPER, "ABCd"}
	got := RunFormula(valop)
	want := "ABCD"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLower1(t *testing.T) {
	valop := []interface{}{LOWER, "ABCD"}
	got := RunFormula(valop)
	want := "abcd"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRunFormulaLower2(t *testing.T) {
	valop := []interface{}{LOWER, "ABCd"}
	got := RunFormula(valop)
	want := "abcd"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
