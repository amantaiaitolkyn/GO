package obj

type Registration struct {
	name     string
	surname  string
	login    string
	password string
	id string
}
type Item struct{
	item_id string 
	name string 
	price int
	seller_id string
	estimation float64
}

type Rating struct{
	p_id string 
	estimation float64
}

type Authorization struct{
	login string
	password string
}