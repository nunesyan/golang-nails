package models

type Appointment struct {
	ID         int    `json:"id"`
	User_Id    int    `json:"user_id"`
	Slot_Id    int    `json:"slot_id"`
	Created_at string `json:"created_at"`
}
