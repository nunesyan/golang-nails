package models

type Work struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Duration_Minutes int    `json:"duration_minutes"`
}
