kind of formulas to use 

POST http://localhost:8000/formulas

body: 

{"rawFormula": "AND(EQUAL(LOWER(\"abA?\"),\"aba?\"), EQUAL(UPPER(\"abA?\"),\"ABB?\"))", "timeZone":"UTC", "fieldValue": {"city": "Casablanca", "country": "Maroc"}}

to get all operators
GET http://localhost:8000/operators
