# CRM FORMULA

Go Project to execute salesforce formulas

- [mux dependency](github.com/gorilla/mux)

```bash
# Clone the repository
git clone https://github.com/ismailSf-rgb/crmformula.git

# Navigate to the project directory
cd crmformula

# Install dependencies
go mod tidy

#execute

go run main.go

#details
kind of formulas to use 

POST http://localhost:8000/formulas

body: 

{"rawFormula": "AND(EQUAL(LOWER(\"abA?\"),\"aba?\"), EQUAL(UPPER(\"abA?\"),\"ABB?\"))", "timeZone":"UTC", "fieldValue": {"city": "Casablanca", "country": "Maroc"}}

or

{"rawFormula":  "IF(3.0 = number,country, city)", "timeZone":"UTC", "fieldValue": {"city": "Casablanca", "country": "Maroc"}}


to get all operators
GET http://localhost:8000/operators

the available operators are {ADDMONTH, ADD, OR, AND, IF, TEXT, LEFT, RIGHT, MULT, TRUE, FALSE, NOT, DIV, MOD, ISBLANK, BLANKVALUE, CONTAINS, EXP, COS, SIN, LEN, MAX, MIN, MORE_EQ, LESS_EQ, MORE, LESS, EQUAL, NOW, CEILING, ROUND, SQRT, LOG, LOG10, ABS, FIND, TRIM, TODAY, BEGINS, CONCAT, BR, YEAR, MONTH, DAY, WEEKDAY, HOUR, MINUTE, FLOOR, UPPER, LOWER, REGEX}


