package orm

type User struct {
	UserID        string `json:"user_id" gorm:"column:user_id;primary_key"` // primary key, wallet address
	UserName      string `json:"user_name" gorm:"column:user_name"`
	TwitterHandle string `json:"twitter_handle" gorm:"column:twitter_handle"`
	TwitterImgUrl string `json:"twitter_img_url" gorm:"column:twitter_img_url"`
}
