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
	P_id string 
	Estimation float64
}

type Authorization struct{
	Login string
	Password string
}
