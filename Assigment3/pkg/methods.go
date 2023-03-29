package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	b "github.com/amantaiaitolkyn/GO/Assigment3/obj"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)
var db *gorm.DB = DB()

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var response = b.JsonResponse{}
	book := b.Listofbook{}
	db.Where("id=? and id != '' ", id).Find(&book)
	fmt.Println(book)
	if book.Id != ""{
		response = b.JsonResponse{Type: "success", Message: "Books: ", Data: book}
	} else {
		response = b.JsonResponse{Type: "error", Message: "Not found!"}
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var response = b.JsonResponse{}
	book := b.Listofbook{}
	db.Where("id=? and id != '' ", id).Find(&book)
	book.Title = r.FormValue("title")
	book.Description = r.FormValue("desc")
	db.Save(&book)
	fmt.Println(book)
	if book.Id != ""{
		response = b.JsonResponse{Type: "success", Message: "Updated" , Data: book}
	} else {
		response = b.JsonResponse{Type: "error", Message: "Not found!"}
	}
	json.NewEncoder(w).Encode(response)
}

func SearchByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	var response = b.JsonResponse{}
	book := []b.Listofbook{}
	db.Where("title = ?", title).Find(&book)
	fmt.Println(book)
	if len(book) !=0 {
		response = b.JsonResponse{Type: "success", Message: "Found" , Datas: book}
	} else {
		response = b.JsonResponse{Type: "error", Message: "Not found!"}
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var response = b.JsonResponse{}
	book := b.Listofbook{}
	db.Where("id=? and id != '' ", id).Find(&book)
	deleted:=book
	db.Delete(&book, id)
	fmt.Println(book)
	if deleted.Id!=""{
		response = b.JsonResponse{Type: "success", Message: "Deleted" , Data: deleted}
	} else {
		response = b.JsonResponse{Type: "error", Message: "Not found!"}
	}
	json.NewEncoder(w).Encode(response)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	description := r.FormValue("desc")
	cost:=r.FormValue("cost")
	costt, err := strconv.Atoi(cost)
    if err != nil {
        fmt.Println(err)
    }
	var response = b.JsonResponse{}
	if id == "" || title == "" || description == "" || costt == 0{
        response = b.JsonResponse{Type: "error", Message: "You are missing login or password parameter."}
    } else {
		db.Create(&b.Listofbook{Id: id,Title: title,Description: description,Cost: costt})
		response = b.JsonResponse{Type: "success", Message: "Added"}
		fmt.Println("Successfully added!")
	}
	json.NewEncoder(w).Encode(response)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var response = b.JsonResponse{}
	book := []b.Listofbook{}
	db.Find(&book)
	fmt.Println(book)
	response = b.JsonResponse{Type: "success", Message: "Books:" , Datas: book}
	json.NewEncoder(w).Encode(response)
}
func SortInAscOrder(w http.ResponseWriter, r *http.Request){
	var response = b.JsonResponse{}
	book := []b.Listofbook{}
	db.Order("cost asc").Find(&book)
	fmt.Println(book)
	response = b.JsonResponse{Type: "success", Message: "Books:" , Datas: book}
	json.NewEncoder(w).Encode(response)
}

func SortInDescOrder(w http.ResponseWriter, r *http.Request){
	var response = b.JsonResponse{}
	book := []b.Listofbook{}
	db.Order("cost desc").Find(&book)
	fmt.Println(book)
	response = b.JsonResponse{Type: "success", Message: "Books:" , Datas: book}
	json.NewEncoder(w).Encode(response)
}