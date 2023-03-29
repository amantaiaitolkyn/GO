package main

import (
	// "database/sql"
	"fmt"
	"strconv"

	// "gorm.io/driver/sqlite"
	// _ "github.com/lib/pq"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Listofbook struct {
	gorm.Model
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
}

type JsonResponse struct {
	Type    string       `json:"type"`
	Data    Listofbook `json:"data"`
	Message string       `json:"message"`
	Datas []Listofbook `json:"datas"`
}

func DB() *gorm.DB {
	dsn := "host=localhost user=postgres password=601246 dbname=golang sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	} else {
		fmt.Println("ok")
	}
	return db
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	db := DB()
	var response = JsonResponse{}
	book := Listofbook{}
	db.Where("id=? and id != '' ", id).Find(&book)
	fmt.Println(book)
	if book.Id != ""{
		response = JsonResponse{Type: "success", Message: "Books: " , Data: book}
	} else {
		response = JsonResponse{Type: "error", Message: "Not found!"}
	}
	json.NewEncoder(w).Encode(response)
}

func updateById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	db := DB()
	var response = JsonResponse{}
	book := Listofbook{}
	db.Where("id=? and id != '' ", id).Find(&book)
	book.Title = r.FormValue("title")
	book.Description = r.FormValue("desc")
	db.Save(&book)
	fmt.Println(book)
	if book.Id != ""{
		response = JsonResponse{Type: "success", Message: "Updated" , Data: book}
	} else {
		response = JsonResponse{Type: "error", Message: "Not found!"}
	}
	json.NewEncoder(w).Encode(response)
}

func searchByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	db := DB()
	var response = JsonResponse{}
	book := []Listofbook{}
	db.Where("title = ?", title).Find(&book)
	fmt.Println(book)
	if len(book) !=0 {
		response = JsonResponse{Type: "success", Message: "Found" , Datas: book}
	} else {
		response = JsonResponse{Type: "error", Message: "Not found!"}
	}
	json.NewEncoder(w).Encode(response)
}

func deleteById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	db := DB()
	var response = JsonResponse{}
	book := Listofbook{}
	db.Where("id=? and id != '' ", id).Find(&book)
	deleted:=book
	db.Delete(&book, id)
	fmt.Println(book)
	if deleted.Id!=""{
		response = JsonResponse{Type: "success", Message: "Deleted" , Data: deleted}
	} else {
		response = JsonResponse{Type: "error", Message: "Not found!"}
	}
	json.NewEncoder(w).Encode(response)
}
func addBook(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	description := r.FormValue("desc")
	cost:=r.FormValue("cost")
	costt, err := strconv.Atoi(cost)
    if err != nil {
        fmt.Println(err)
    }

	db := DB()
	var response = JsonResponse{}
	if id == "" || title == "" || description == "" || costt == 0{
        response = JsonResponse{Type: "error", Message: "You are missing login or password parameter."}
    } else {
		db.Create(&Listofbook{Id: id,Title: title,Description: description,Cost: costt})
		response = JsonResponse{Type: "success", Message: "Added"}
	}
	json.NewEncoder(w).Encode(response)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	db := DB()
	var response = JsonResponse{}
	book := []Listofbook{}
	db.Find(&book)
	fmt.Println(book)
	response = JsonResponse{Type: "success", Message: "Books:" , Datas: book}
	json.NewEncoder(w).Encode(response)
}
func sortInAscOrder(w http.ResponseWriter, r *http.Request){
	db := DB()
	var response = JsonResponse{}
	book := []Listofbook{}
	db.Order("cost asc").Find(&book)
	fmt.Println(book)
	response = JsonResponse{Type: "success", Message: "Books:" , Datas: book}
	json.NewEncoder(w).Encode(response)
}
func sortInDescOrder(w http.ResponseWriter, r *http.Request){
	db := DB()
	var response = JsonResponse{}
	book := []Listofbook{}
	db.Order("cost desc").Find(&book)
	fmt.Println(book)
	response = JsonResponse{Type: "success", Message: "Books:" , Datas: book}
	json.NewEncoder(w).Encode(response)
}

func main() {
	db := DB()
	db.Create(&Listofbook{Id: "18",Title: "fffdf",Description: "sdfsf",Cost: 33303})
	db.Create(&Listofbook{Id: "19",Title: "fff",Description: "sdfsf",Cost: 1000})
	db.Create(&Listofbook{Id: "20",Title: "fff",Description: "sdfsf",Cost: 4000})
	router := mux.NewRouter()
	router.HandleFunc("/getById", getBookById).Methods("GET")
	router.HandleFunc("/getAllBooks", getAllBooks).Methods("GET")
	router.HandleFunc("/updateById", updateById).Methods("PUT")
	router.HandleFunc("/deleteById", deleteById).Methods("DELETE")
	router.HandleFunc("/searchByTitle", searchByTitle).Methods("GET")
	router.HandleFunc("/addBook", addBook).Methods("POST")
	router.HandleFunc("/sortAsc", sortInAscOrder).Methods("GET")
	router.HandleFunc("/sortDesc", sortInDescOrder).Methods("GET")
	fmt.Println("Server at 8080")
    errr := http.ListenAndServe(":7552", router)
	if errr != nil {
        fmt.Println(errr)
    }
}
