package models

type App struct {
	Id int64 `json:"id,omitempty"`

	Employees string `json:"employees,omitempty"`

	Trainees string `json:"trainees,omitempty"`
}
