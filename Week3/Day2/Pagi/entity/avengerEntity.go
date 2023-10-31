package entity

type User struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Deposit_Amount int    `json:"deposit_amount"`
}

// type Product struct {
// 	Name_Product string `json:"product"`
// 	Stock        string `json:"Stock"`
// 	Price        int    `json:"price"`
// }

// type Transaction struct {
// 	User_id      int    `json:"user_id"`
// 	ProductID    int    `json:"product_id"`
// 	Quantity     string `json:"quantity"`
// 	Total_Amount int    `json:"total"`
// }