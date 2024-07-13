package user

import "time"

type User struct {
	ID           string      `json:"id"`
	Type         string      `json:"type"`
	FirstName    string      `json:"first_name"`
	LastName     string      `json:"last_name"`
	Username     string      `json:"username"`
	Bio          string      `json:"bio"`
	Photo        string      `json:"photo"`
	Description  string      `json:"description"`
	LinkedChatID string      `json:"linked_chat_id"`
	Private      bool        `json:"private"`
	Preferences  Preferences `json:"preferences"`
	LastUpdated  time.Time   `json:"last_updated"`
}

type Preferences struct {
	UserID         string   `json:"-"`
	Locale         string   `json:"locale"`
	ReportSchedule string   `json:"report_schedule"`
	ReportCities   []string `json:"report_cities"`
}

type UserActionLog struct {
	UserID    string    `json:"user_id"`
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
	Details   map[string]interface{}
}
