package models

type StudentDB struct {
	ID          string `json:"id"`
	Fullname    string `json:"fullname"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Regist_date string `json:"regist_date"`
	Major       string `json:"major"`
	Hobby       string `json:"hobby"`
}

type Hobbies struct {
	ID    string `json:"id"`
	Hobby string `json:"hobby"`
}

type Majors struct {
	ID    string `json:"id"`
	Major string `json:"major"`
}

type Student struct {
	ID          string   `json:"id"`
	Fullname    string   `json:"fullname"`
	Gender      string   `json:"gender"`
	Age         int      `json:"age"`
	Regist_date string   `json:"regist_date"`
	MajorID     string   `json:"major_id"`
	HobbyID     []string `json:"hobby_id"`
	Major       string   `json:"major"`
	Hobbies     []string `json:"hobbies"`
}
