package main

type User struct {
	ID       string `json:"id" form:"id" query:"id"`
	Name     string `json:"name" form:"name" query:"name"`
	Email    string `json:"email" form:"email" query:"eamil"`
	Password string `json:"password" form:"password" query:"password"`
	Checkstr string `json:"checkstr" form:"checkstr" query:"checkstr"`
	Info     string `json:"info" form:"info" query:"info"`
}

type Email struct {
	Name     string `json:"name" form:"name" query:"name"`
	Status   bool   `json:"status" form:"status" query:"status"`
	Info     string `json:"info" form:"info" query:"info"`
	Checkstr string `json:"checkstr" form:"checkstr" query:"checkstr"`
}
