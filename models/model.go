package models

type Class struct {
	ClassName string `json:"class_name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  int    `json:"capacity"`
}

type ClassBooking struct {
	MemberName string `json:"name"`
	Date       string `json:"date"`
	ClassName  string `json:"class_name"`
}
