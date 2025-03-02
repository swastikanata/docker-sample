type User struct {
	FirstName  string   `json:"firstName"`
	MiddleName string   `json:"middleName"`
	LastName   string   `json:"lastName"`
	Hobbies    []string `json:"hobbies"`
	Address    struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"address"`
	Wallet []struct {
		Type   string `json:"type"`
		Amount int    `json:"amount"`
	} `json:"wallet"`
}
