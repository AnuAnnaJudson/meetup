package newmodels

type User struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Meetups []*Meetup `json:"meetups"`
}
