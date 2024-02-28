package main




type allAranas []Aranas

var aranas = allAranas{
	{1, "Tarantula", "kidd keo"},
	{2, "Arana normi", "Ariana Grande"},
}


func getAranas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pcs)
	w.Header().Set("Content-Type", "application/json")
}

func main(
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/aranas", getPcs).Methods("GET")
)