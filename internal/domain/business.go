package domain

type Business struct {
	ID          string // UUID
	Name        string
	Description string
	Category    string
	Address     string
	ContactInfo string
	Photos      []string // URLs to business photos
	Hours       string   // e.g., JSON: {"Monday": "9am-5pm", "Tuesday": "Closed", ...}
	Latitude    float64
	Longitude   float64
	OwnerID     string // Foreign key to User (business owner)
}

type Category struct {
	ID   string // UUID
	Name string
}
