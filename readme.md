kind of formulas to use 

POST http://localhost:8000/formulas

body: 

{"rawFormula": "AND(EQUAL(LOWER(\"abA?\"),\"aba?\"), EQUAL(UPPER(\"abA?\"),\"ABB?\"))", "timeZone":"UTC", "fieldValue": {"city": "Casablanca", "country": "Maroc"}}

to get all operators
GET http://localhost:8000/operators

the available operators are {ADDMONTH, ADD, OR, AND, IF, TEXT, LEFT, RIGHT, MULT, TRUE, FALSE, NOT, DIV, MOD, ISBLANK, BLANKVALUE, CONTAINS, EXP, COS, SIN, LEN, MAX, MIN, MORE_EQ, LESS_EQ, MORE, LESS, EQUAL, NOW, CEILING, ROUND, SQRT, LOG, LOG10, ABS, FIND, TRIM, TODAY, BEGINS, CONCAT, BR, YEAR, MONTH, DAY, WEEKDAY, HOUR, MINUTE, FLOOR, UPPER, LOWER, REGEX}


