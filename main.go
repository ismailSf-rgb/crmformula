package main

import (
	"crm/go/crmformula"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/formulas", PostFormula).Methods("POST")
	router.HandleFunc("/operators", GetOperators).Methods("GET")
	router.HandleFunc("/swagger.json", SwaggerHandler)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func PostFormula(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				errorMsg := map[string]string{"error": "Formula Compilation Error : " + err.Error()}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errorMsg)
				return
			}
			// Respond with an error message
			errorMsg := map[string]string{"error": "Formula Compilation Error"}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorMsg)
		}
	}()

	var formula crmformula.Formula
	_ = json.NewDecoder(r.Body).Decode(&formula)
	if formula.GetRawFormula() == "" || formula.GetTimeZone() == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(formula)
		return
	}

	if crmformula.ContainsSpecialChars(formula.GetRawFormula()) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Errorf("the formula %v contains illegal characters", formula.RawFormula).Error())
		return
	}

	_, err := formula.ReplaceFieldsWithValue()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	_, err = formula.FindOperators()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	_, err = formula.FindOperatorsAndParenthesis()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	tokenizedForm, err := formula.Tokenize()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	value := crmformula.GetFormulaResult(tokenizedForm)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(value)

}

func GetOperators(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var filteredOps []string
	for _, op := range crmformula.FormOperators {
		if op == "TRUE" || op == "FALSE" {
			continue
		}
		filteredOps = append(filteredOps, op.(string))
	}
	// Set the HTTP status code to 200 (OK)
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(filteredOps); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

}

func SwaggerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonFilePath := "./swaggerformula.json"
	http.ServeFile(w, r, jsonFilePath)
}
