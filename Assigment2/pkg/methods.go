package pkg

import (
	"fmt"
	"strconv"
	b "github.com/amantaiaitolkyn/GO/Assigment2/obj"
	"database/sql"	
	_ "github.com/lib/pq"
    "encoding/json"
    // "log"
    "net/http"

)


var db *sql.DB = Db()

func Authorizationn(w http.ResponseWriter, r *http.Request){
    Login := r.FormValue("login")
    Password := r.FormValue("password")

    var response = b.JsonResponse{}

    if Login == "" || Password == ""{
        response = b.JsonResponse{Type: "error", Message: "You are missing login or password parameter."}
    } else {
        db:= Db()
        result, err := db.Query("select * from users where login = $1 and password = $2", Login, Password)
        if err != nil {
            panic(err)
        }
        user := []b.Registration{}
     
        for  result.Next(){
            p := b.Registration{}
            err := result.Scan(&p.Name, &p.Surname, &p.Login, &p.Password)
            if err != nil{
                fmt.Println(err)
                continue
            }
            user = append(user, p)
        }

        if len(user) != 0{
            response = b.JsonResponse{Type: "success", Message: "You have successfully logged in!"}
        }else{
            response = b.JsonResponse{Type: "error", Message: "The data is incorrect, please try again."}
        }
       
    }

    json.NewEncoder(w).Encode(response)
}


func AddItem(w http.ResponseWriter, r *http.Request){
    Item_id := r.FormValue("itemid")
    Name := r.FormValue("name")
    Price := r.FormValue("price")
    Seller_id := r.FormValue("sellerid")

    pp, err := strconv.Atoi(Price)
    if err != nil {
    fmt.Println(err)
    }

    var response = b.JsonResponse{}

    if Name == "" || Price == "" || Seller_id == ""|| Item_id == "" {
        response = b.JsonResponse{Type: "error", Message: "You are missing name or price or seller_id parameter."}
    } else {
        fmt.Println("Successfully added")
        db:= Db()
        fmt.Println("Inserting new item with Name and surname: " + Name +" "+ Price + " and seller_id: " + Seller_id)
        newItem := b.Item{Item_id: Item_id,Name: Name,Price: pp, Seller_id: Seller_id}
	    result,err := db.Exec("insert into items (item_id, name, price, seller_id, estimation) values ($1, $2, $3, $4, $5)",newItem.Item_id, newItem.Name, newItem.Price, newItem.Seller_id, newItem.Estimation)
        result.LastInsertId()
        if err != nil{
            panic(err)
        }

        response = b.JsonResponse{Type: "success", Message: "The item has been inserted successfully!"}
    }

    json.NewEncoder(w).Encode(response)
	
}

func SearchItem(w http.ResponseWriter, r *http.Request){
    Name := r.FormValue("name") 

    var response = b.JsonResponse{}

    if Name == ""{
        response = b.JsonResponse{Type: "error", Message: "You are missing name parameter."}
    } else {
        fmt.Println("A search is underway")
        db:= Db()
        rows, err := db.Query("select * from items where name like $1", Name)
        if err != nil {
            panic(err)
        }
        products := []b.Item{}
     
        for rows.Next(){
            p := b.Item{}
            err := rows.Scan(&p.Item_id, &p.Name, &p.Price, &p.Seller_id, &p.Estimation)
            if err != nil{
                fmt.Println(err)
                continue
            }
            products = append(products, p)
        }
        productss := fmt.Sprintf(" ", products)


        rows1, err := db.Query("select * from seller where id = (select seller_id from items where name like $1)", Name)
        if err != nil {
            panic(err)
        }
        
        sellers := []b.Registration{}
        for rows1.Next(){
            p := b.Registration{}
            err := rows1.Scan(&p.Surname, &p.Name, &p.Login, &p.Password, &p.Id)
            if err != nil{
                fmt.Println(err)
                continue
            }
            sellers = append(sellers,p)
        }
        sellerss := fmt.Sprintf(" ", sellers)
        if len(products) != 0{
            response = b.JsonResponse{Type: "success", Message: "Products: "+ productss + ",Seller: "+ sellerss}
        }else{
            response = b.JsonResponse{Type: "error", Message: "Not found!"}
        }
       
    }

    json.NewEncoder(w).Encode(response)
}

func  Register(w http.ResponseWriter, r *http.Request) {
	Name := r.FormValue("name")
    Surname := r.FormValue("surName")
    Login := r.FormValue("login")
    Password := r.FormValue("password")

    var response = b.JsonResponse{}
    if Name == "" || Surname == "" || Login == ""|| Password == "" {
        response = b.JsonResponse{Type: "error", Message: "You are missing name or surname or login or password parameter."}
    } else {
        db:= Db()
        fmt.Println("Inserting new user with Name and surname: " + Name +" "+ Surname + " and login: " + Login)
        newRegister := b.Registration{Name: Name, Surname: Surname, Login: Login, Password: Password}
        result,err := db.Exec("insert into MyUsers (name, surname, login, password) values ($1, $2, $3, $4)",newRegister.Name, newRegister.Surname, newRegister.Login, newRegister.Password)
        result.LastInsertId()
        if err != nil{
            panic(err)
        }

        response = b.JsonResponse{Type: "success", Message: "The user has been inserted successfully!"}
    }
    json.NewEncoder(w).Encode(response)
}

func  RegisterForSeller(w http.ResponseWriter, r *http.Request) {
    Name := r.FormValue("name")
    Surname := r.FormValue("surName")
    Login := r.FormValue("login")
    Password := r.FormValue("password")
    Id := r.FormValue("id")

    var response = b.JsonResponse{}
    if Name == "" || Surname == "" || Login == ""|| Password == ""|| Id == "" {
        response = b.JsonResponse{Type: "error", Message: "You are missing name or surname or login or password parameter."}
    } else {
        db:= Db()
        fmt.Println("Okay,user has been inserted")
        fmt.Println("Name and surname: " + Name +" "+ Surname + " and login: " + Login)
        newRegister := b.Registration{Name: Name, Surname: Surname, Login: Login, Password: Password,Id:Id}
        result,err := db.Exec("insert into seller (name, surname, login, password,id) values ($1, $2, $3, $4, $5)",newRegister.Name, newRegister.Surname, newRegister.Login, newRegister.Password, newRegister.Id)
        if err != nil{
            panic(err)
        }
        result.LastInsertId()
        resultt,err := db.Exec("insert into MyUsers (name, surname, login, password) values ($1, $2, $3, $4)",newRegister.Name, newRegister.Surname, newRegister.Login, newRegister.Password)
        resultt.LastInsertId()
        if err != nil{
            panic(err)
        }
        response = b.JsonResponse{Type: "success", Message: "The user has been inserted successfully!"}
    }
    json.NewEncoder(w).Encode(response)
}

func RateTheProduct(w http.ResponseWriter, r *http.Request){
    Item_id := r.FormValue("itemid")
    Estimation := r.FormValue("estimation")

    e, err := strconv.ParseFloat(Estimation, 64)
    if err != nil{
        panic(err)
    }
 
    var response = b.JsonResponse{}

    if Item_id == "" || Estimation == "" {
        response = b.JsonResponse{Type: "error", Message: "You are missing item_id or estimation parameter."}
    } else {
        db:= Db()
        newRating := b.Rating{P_id: Item_id, Estimation: e}
        result,err := db.Exec("insert into rating (p_id,estimation) values ($1, $2)",newRating.P_id, newRating.Estimation)
        result.LastInsertId()
        
        if err != nil{
            panic(err)
        }

        rows, err := db.Query("select * from rating where p_id like $1", Item_id)
        if err != nil {
            panic(err)
        }

        rating := []b.Rating{}
        for rows.Next(){
            p := b.Rating{}
            err := rows.Scan(&p.P_id, &p.Estimation)
            if err != nil{
                fmt.Println(err)
                continue
            }
            rating = append(rating, p)
        }

        var sum float64
        for _, p := range rating{
            sum += p.Estimation
        }
        rate := sum / float64(len(rating))
        row1, err := db.Query("update items set estimation = $1 where item_id = $2", rate, Item_id)
        if err != nil{
            panic(err)
        }else{
            response = b.JsonResponse{Type: "success", Message: "The product has been successfully evaluated"}
        }
        row1.Close()
    }
    json.NewEncoder(w).Encode(response)
}

func FilteringByPrice(w http.ResponseWriter, r *http.Request){
    startPrice := r.FormValue("startp")
    endPrice := r.FormValue("endp")

    var response = b.JsonResponse{}

    if startPrice == "" || endPrice == "" {
        response = b.JsonResponse{Type: "error", Message: "You are missing start price  or end price parameter."}
    } else {
        fmt.Println("A search is underway")
        db:= Db()
        rows, err := db.Query("select * from items where price BETWEEN $1 and $2", startPrice,endPrice)
        if err != nil {
            panic(err)
        }
        
        products := []b.Item{}
        
        for rows.Next(){
            p := b.Item{}
            err := rows.Scan(&p.Item_id, &p.Name, &p.Price, &p.Seller_id, &p.Estimation)
            if err != nil{
                fmt.Println(err)
                continue
            }
            products = append(products, p)
        }
        productss := fmt.Sprintf(" ", products)

        if len(products) != 0{
            response = b.JsonResponse{Type: "success", Message: "Products: "+ productss}
        }else{
            response = b.JsonResponse{Type: "error", Message: "Not found!"}
        }
       
    }

    json.NewEncoder(w).Encode(response)
	
}
func FilteringByRating(w http.ResponseWriter, r *http.Request){
    startRating := r.FormValue("startr")
    endRating := r.FormValue("endr")

    var response = b.JsonResponse{}

    if startRating == "" || endRating == "" {
        response = b.JsonResponse{Type: "error", Message: "You are missing start rating  or end rating parameter."}
    } else {
        fmt.Println("A search is underway")
        db:= Db()
        rows, err := db.Query("select * from items where estimation BETWEEN $1 and $2", startRating,endRating)
        if err != nil {
            panic(err)
        }
        
        products := []b.Item{}
        
        for rows.Next(){
            p := b.Item{}
            err := rows.Scan(&p.Item_id, &p.Name, &p.Price, &p.Seller_id, &p.Estimation)
            if err != nil{
                fmt.Println(err)
                continue
            }
            products = append(products, p)
        }
        productss := fmt.Sprintf(" ", products)

        if len(products) != 0{
            response = b.JsonResponse{Type: "success", Message: "Products: "+ productss}
        }else{
            response = b.JsonResponse{Type: "error", Message: "Not found!"}
        }
       
    }

    json.NewEncoder(w).Encode(response)
	
}