package main
import (
    a "github.com/amantaiaitolkyn/GO/Assigment2/pkg"
	// b "github.com/amantaiaitolkyn/GO/Assigment2/obj"
	"fmt"
	// "encoding/json"
	"github.com/gorilla/mux"
    // "log"
    "net/http"
)
func main(){
	// Register("Aiii","Aman","mylog","asasasa")
	// RegisterForSeller("Aiii","Aman","mylog","asasasa","122")
	// AddItem("A1","bag",12000,"122")
	// rateTheProduct("A1",2)
	// rateTheProduct("A1",3)
	// RegisterForSeller("Almasova","Aluu","my_log","afght","123")
	// RegisterForSeller("Amanova","Lia","myllog","1234","124")
	// AddItem("A2","pen",10000,"123")
	// AddItem("A3","copybook",20000,"124")
	// FilteringByPrice(13000,20000)
	// a.FilteringByRating(0,1)
	// a.Authorizationn("Aiiids","asasadsdsa")
	router := mux.NewRouter()


    router.HandleFunc("/user", a.Register)
	router.HandleFunc("/seller", a.RegisterForSeller)
	router.HandleFunc("/item", a.AddItem)
	router.HandleFunc("/search", a.SearchItem)
	router.HandleFunc("/authorization", a.Authorizationn)
	router.HandleFunc("/rate", a.RateTheProduct)
	router.HandleFunc("/filteringByPrice", a.FilteringByPrice)
	router.HandleFunc("/filteringByRating", a.FilteringByRating)

    fmt.Println("Server at 8080")
    errr := http.ListenAndServe(":7575", router)
	if errr != nil {
        fmt.Println(errr)
    }
	// rateTheProduct("A1",3.9)
	// SearchItem("bag")
    // defer db.Close()

}