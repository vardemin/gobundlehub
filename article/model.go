package article

type Article struct {
	Title   string `json:"title" bson:"title" minLen:"2" maxLen:"30" required:"true"`
	Content string `json:"content" bson:"content" maxLen:"1024" required:"true"`
	Image   string `json:"image" bson:"image"`
}
