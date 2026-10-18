package domain

type Subscription struct {
	ID      int     `json:"id"`
	User_id string  `json:"user_id"`
	Title   string  `json:"title"`
	Cost    float64 `json:"cost"`
	Date    int     `json:"date"`
}
