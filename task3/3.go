package main

// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
//要求 ：
//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//编写Go代码，使用Gorm创建这些模型对应的数据库表

type User struct {
	gorm.Model
	Username  string `gorm:"size:50;not null"`
	Email     string `gorm:"size:100;not null"`
	Password  string `gorm:"size:100;not null"`
	PostCount int    `gorm:"default:0"`
	Posts     []Post
}

type Post struct {
	gorm.Model
	Title         string `gorm:"size:100;not null"`
	Content       string `gorm:"type:text;not null"`
	UserID        int
	CommentStatus string `gorm:"size:20;default:'无评论'"`
	User          User   `gorm:"foreignKey:UserID"`
	Comments      []Comment
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	UserID  int
	PostID  int
	User    User `gorm:"foreignKey:UserID"`
	Post    Post `gorm:"foreignKey:PostID"`
}

func main() {
	_ = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	GetUserPostsAndComments(db, 1)
	GetMostCommentedPost(db)
}

// 基于上述博客系统的模型定义。
//要求 ：
//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//编写Go代码，使用Gorm查询评论数量最多的文章信息

func GetUserPostsAndComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Where("user_id = ?", userID).Preload("Comments").Find(&posts).Error
	return posts, err
}

func GetMostCommentedPost(db *gorm.DB) (*Post, error) {
	var post Post
	query := db.Model(&Comment{}).Select("post_id, COUNT(*) as comment_count").Group("post_id")

	err := db.Joins("JOIN (?) as c on posts.id = c.post_id", query).Order("c.comment_count DESC").First(&post).Error
	return &post, nil
}

// 继续使用博客系统的模型。
//要求 ：
//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
// AfterCreate 钩子 - 创建文章后更新用户的文章计数z

func (p *Post) UpdatePostCount(tx *gorm.DB) (err error) {
	return tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count + 1")).Error
}

func (c *Comment) UpdatePostStatus(tx *gorm.DB) (err error) {
	var count int64
	if err = tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		if err = tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error; err != nil {
			return err
		}
	}
	return nil
}
