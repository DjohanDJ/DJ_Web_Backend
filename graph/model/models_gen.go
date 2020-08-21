// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID              string `json:"id"`
	VideoID         int    `json:"video_id"`
	CommentParentID int    `json:"comment_parent_id"`
	CommentValue    string `json:"comment_value"`
	UserID          int    `json:"user_id"`
	CommentDate     string `json:"comment_date"`
	Likes           int    `json:"likes"`
	Dislikes        int    `json:"dislikes"`
}

type Membership struct {
	ID              string `json:"id"`
	UserID          int    `json:"user_id"`
	JoinPremiumDate string `json:"join_premium_date"`
	MemberType      string `json:"member_type"`
}

type NewComment struct {
	VideoID         int    `json:"video_id"`
	CommentParentID int    `json:"comment_parent_id"`
	CommentValue    string `json:"comment_value"`
	UserID          int    `json:"user_id"`
	CommentDate     string `json:"comment_date"`
	Likes           int    `json:"likes"`
	Dislikes        int    `json:"dislikes"`
}

type NewMembership struct {
	UserID          int    `json:"user_id"`
	JoinPremiumDate string `json:"join_premium_date"`
	MemberType      string `json:"member_type"`
}

type NewSubscriber struct {
	ChannelID int `json:"channel_id"`
	UserID    int `json:"user_id"`
}

type NewUser struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	UserPassword  string `json:"user_password"`
	ChannelName   string `json:"channel_name"`
	UserImage     string `json:"user_image"`
	ChannelBanner string `json:"channel_banner"`
	ChannelDesc   string `json:"channel_desc"`
	Restriction   string `json:"restriction"`
	Location      string `json:"location"`
	Membership    string `json:"membership"`
	ExpiredMember string `json:"expired_member"`
	JoinDate      string `json:"join_date"`
}

type NewVideo struct {
	ImagePath   string `json:"image_path"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ViewCount   int    `json:"view_count"`
	UploadDate  string `json:"upload_date"`
	VideoPath   string `json:"video_path"`
	UserID      int    `json:"user_id"`
	Restriction string `json:"restriction"`
	CategoryID  int    `json:"category_id"`
	Location    string `json:"location"`
	Publish     string `json:"publish"`
	Premium     string `json:"premium"`
}

type Subscriber struct {
	ID        string `json:"id"`
	ChannelID int    `json:"channel_id"`
	UserID    int    `json:"user_id"`
}

type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	UserPassword  string `json:"user_password"`
	ChannelName   string `json:"channel_name"`
	UserImage     string `json:"user_image"`
	ChannelBanner string `json:"channel_banner"`
	ChannelDesc   string `json:"channel_desc"`
	Restriction   string `json:"restriction"`
	Location      string `json:"location"`
	Membership    string `json:"membership"`
	ExpiredMember string `json:"expired_member"`
	JoinDate      string `json:"join_date"`
}

type Video struct {
	ID          string `json:"id"`
	ImagePath   string `json:"image_path"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ViewCount   int    `json:"view_count"`
	UploadDate  string `json:"upload_date"`
	VideoPath   string `json:"video_path"`
	UserID      string `json:"user_id"`
	Restriction string `json:"restriction"`
	CategoryID  int    `json:"category_id"`
	Location    string `json:"location"`
	Publish     string `json:"publish"`
	Premium     string `json:"premium"`
}
