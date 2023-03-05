package obj

type Registration struct {
	Name     string `json:"name"`
	Surname  string `json:"surName"`
	Login    string `json:"login"`
	Password string  `json:"password"`
	Id string `json:"id"`
}
type Item struct{
	Item_id string `json:"itemid"`
	Name string `json:"name"`
	Price int `json:"price"`
	Seller_id string `json:"sellerid"`
	Estimation float64 `json:"estimation"`
}

type Rating struct{
	P_id string `json:"p_id"`
	Estimation float64 `json:"estimation"`
}

type Authorization struct{
	Login string `json:"login"`
	Password string `json:"password"`
}

type JsonResponse struct {
    Type    string `json:"type"`
    Data    []Registration `json:"data"`
    Message string `json:"message"`
}