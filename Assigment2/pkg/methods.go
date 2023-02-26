package pkg

func Authorizationn(newlogin,newpassword string){
	user:= Authorization{login: newlogin,password: newpassword}
	rows, err := db.Query("select login,password from MyUsers where login = $1 and password =  $2", user.login,user.password)
    if err != nil {
        panic(err)
		// fmt.Println("You have entered an incorrect login or password, please try again!")
    }else{
		fmt.Println("Okay!")
	}
	rows.Close()
}


func AddItem(id, name string,price int,seller_id string){
	newItem := Item{item_id: id,name: name,price: price, seller_id: seller_id}
	result,err := db.Exec("insert into items (item_id, name, price, seller_id, estimation) values ($1, $2, $3, $4, $5)",newItem.item_id, newItem.name, newItem.price, newItem.seller_id, newItem.estimation)
    if err != nil{
        panic(err)
    }
	result.LastInsertId()
	fmt.Println("Successfully added")
}
func SearchItem(name string){
	rows, err := db.Query("select * from items where name like $1", name)
    if err != nil {
        panic(err)
    }
    
    products := []Item{}
     
    for rows.Next(){
        p := Item{}
        err := rows.Scan(&p.item_id, &p.name, &p.price, &p.seller_id, &p.estimation)
        if err != nil{
            fmt.Println(err)
            continue
        }
        products = append(products, p)
    }

	rows1, err := db.Query("select * from seller where id = (select seller_id from items where name like $1)", name)
    if err != nil {
        panic(err)
    }
	
    sellers := []Registration{}
	for rows1.Next(){
        p := Registration{}
        err := rows1.Scan(&p.surname, &p.name, &p.login, &p.password, &p.id)
        if err != nil{
            fmt.Println(err)
            continue
        }
        sellers = append(sellers, p)
    }
	fmt.Println("Searching product:")
    for _, p := range products{
        fmt.Println(p.item_id, p.name, p.price, p.seller_id, p.estimation)
    }
	fmt.Println("Seller of this product:")
	for _, p := range sellers{
        fmt.Println(p.surname, p.name)
    }
}

func  Register(newName , newSurname, newLogin, newPassword string) {
	newRegister := Registration{name: newName, surname: newSurname, login: newLogin, password: newPassword}
	result,err := db.Exec("insert into MyUsers (name, surname, login, password) values ($1, $2, $3, $4)",newRegister.name, newRegister.surname, newRegister.login, newRegister.password)
    if err != nil{
        panic(err)
    }
	result.LastInsertId()
}
func  RegisterForSeller(newName , newSurname, newLogin, newPassword, newid string) {
	newRegister := Registration{name: newName, surname: newSurname,login: newLogin, password: newPassword, id: newid}
	result,err := db.Exec("insert into Seller (name, surname, login, password,id) values ($1, $2, $3, $4, $5)",newRegister.name, newRegister.surname, newRegister.login, newRegister.password, newRegister.id)
    if err != nil{
        panic(err)
    }
	result.LastInsertId()
}

func rateTheProduct(id string,newestimation float64){
	newRating := Rating{p_id: id, estimation: newestimation}
	result,err := db.Exec("insert into rating (p_id,estimation) values ($1, $2)",newRating.p_id, newRating.estimation)
    if err != nil{
        panic(err)
    }
	result.LastInsertId()
	rows, err := db.Query("select * from rating where p_id like $1", id)
    if err != nil {
        panic(err)
    }
    
	rating := []Rating{}
     
    for rows.Next(){
        p := Rating{}
        err := rows.Scan(&p.p_id, &p.estimation)
        if err != nil{
            fmt.Println(err)
            continue
        }
        rating = append(rating, p)
    }
	var sum float64
	for _, p := range rating{
        sum += p.estimation
    }
	rate := sum / float64(len(rating))
	row1, err := db.Query("update items set estimation = $1 where item_id = $2", rate, id)
	if err != nil{
        panic(err)
    }
	row1.Close()
}

func FilteringByPrice(startPrice,endPrice int){
	rows, err := db.Query("select * from items where price BETWEEN $1 and $2", startPrice,endPrice)
    if err != nil {
        panic(err)
    }
    
    products := []Item{}
     
    for rows.Next(){
        p := Item{}
        err := rows.Scan(&p.item_id, &p.name, &p.price, &p.seller_id, &p.estimation)
        if err != nil{
            fmt.Println(err)
            continue
        }
        products = append(products, p)
    }
	for _, p := range products{
        fmt.Println(p.item_id, p.name, p.price, p.seller_id, p.estimation)
    }
}
func FilteringByRating(startRating,endRating float64){
	rows, err := db.Query("select * from items where estimation BETWEEN $1 and $2", startRating,endRating)
    if err != nil {
        panic(err)
    }
    
    products := []Item{}
     
    for rows.Next(){
        p := Item{}
        err := rows.Scan(&p.item_id, &p.name, &p.price, &p.seller_id, &p.estimation)
        if err != nil{
            fmt.Println(err)
            continue
        }
        products = append(products, p)
    }
	for _, p := range products{
        fmt.Println(p.item_id, p.name, p.price, p.seller_id, p.estimation)
    }
}