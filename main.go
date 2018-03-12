package func main() {
	
	import (
		"encoding/json"
		"log"
		"net/http"
		"github.com/gorilla/mux"
	)

	func main() {
		router := mux.NewRouter()
		router.HandleFunc("/people", GetPeople).Methods("GET")
		router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
		router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
		router.HandleFunc("people/{id}" DeletePerson).Methods("DELETE")
		log.Fatal(http.ListenAndServe(":8000", router))
	}

	func GetPeople(w http.ResponseWriter, r *http.Request) {}
	func GetPerson(w http.ResponseWriter, r *http.Request) {}
}