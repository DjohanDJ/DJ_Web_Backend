// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewUser struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	UserPassword  string `json:"user_password"`
	ChannelName   string `json:"channel_name"`
	UserImage     string `json:"user_image"`
	ChannelBanner string `json:"channel_banner"`
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
}

type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	UserPassword  string `json:"user_password"`
	ChannelName   string `json:"channel_name"`
	UserImage     string `json:"user_image"`
	ChannelBanner string `json:"channel_banner"`
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
}
