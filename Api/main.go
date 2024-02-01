package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type pc struct {
	ID   int    `json:ID`
	Name string `json:Name`
	Cpu  string `json:Cpu`
}

type allPcs []pc

var pcs = allPcs{
	{
		ID:   1,
		Name: "Ips",
		Cpu:  "Intel i5 5700k",
	},
}

func getPcs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pcs)
	w.Header().Set("Content-Type", "application/json")
}

func createPc(w http.ResponseWriter, r *http.Request) {
	var newPc pc

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}

	json.Unmarshal(reqBody, &newPc)

	newPc.ID = len(pcs) + 1
	pcs = append(pcs, newPc)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPc)

}

func getPc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pcID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, pc := range pcs {
		if pc.ID == pcID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pc)
		}
	}

}

func deletePc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pcID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "invalid ID")
		return
	}

	for i, pc := range pcs {
		if pc.ID == pcID {
			pcs = append(pcs[:i], pcs[i+1:]...)
			fmt.Fprintf(w, "The pcs was delete succesfully")
		}
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to api")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/pcs", getPcs).Methods("GET")
	router.HandleFunc("/pcs", createPc).Methods("POST")
	router.HandleFunc("/pcs/{id}", getPc).Methods("GET")
	router.HandleFunc("/pcs/{id}", deletePc).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))
}
