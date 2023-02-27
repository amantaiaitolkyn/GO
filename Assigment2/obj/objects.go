package obj

type Registration struct {
	Name     string
	Surname  string
	Login    string
	Password string
	Id string
}
type Item struct{
	Item_id string 
	Name string 
	Price int
	Seller_id string
	Estimation float64
}

type Rating struct{
	p_id string 
	estimation float64
}

type Authorization struct{
	login string
	password string
}