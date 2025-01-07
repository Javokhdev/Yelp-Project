package domain

type Review struct {
	ID         string // UUID
	UserID     string // Foreign key to User
	BusinessID string // Foreign key to Business
	Rating     int    // 1 to 5 stars
	Comment    string
	Photos     []string // URLs to review photos
}

type Rating struct {
	BusinessID   string  // Foreign key to Business
	AvgRating    float64 // Average rating
	TotalRatings int
}
