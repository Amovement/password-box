import "gorm.io/gorm"

type User struct {
	gorm.Model
	Title       string
	Username    string
	Password    string
	Url         string
	Description string
	Tag         string
}
