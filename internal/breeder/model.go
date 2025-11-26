package breeder

// Breeder represents a pet breeder
type Breeder struct {
	ID          int    `json:"id"`
	BreederName string `json:"breeder_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	ProvState   string `json:"prov_state"`
	Country     string `json:"country"`
	Zip         string `json:"zip"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Active      int    `json:"active"`
}
