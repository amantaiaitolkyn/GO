package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"github.com/gorilla/mux"
	d "github.com/amantaiaitolkyn/GO/Assigment3/pkg"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getById", d.GetBookById).Methods("GET")
	router.HandleFunc("/getAllBooks", d.GetAllBooks).Methods("GET")
	router.HandleFunc("/updateById", d.UpdateById).Methods("PUT")
	router.HandleFunc("/deleteById", d.DeleteById).Methods("DELETE")
	router.HandleFunc("/searchByTitle", d.SearchByTitle).Methods("GET")
	router.HandleFunc("/addBook", d.AddBook).Methods("POST")
	router.HandleFunc("/sortAsc", d.SortInAscOrder).Methods("GET")
	router.HandleFunc("/sortDesc", d.SortInDescOrder).Methods("GET")
	fmt.Println("Server at 8080")
    errr := http.ListenAndServe(":7551", router)
	if errr != nil {
        fmt.Println(errr)
    }
}
