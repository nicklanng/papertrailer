package events

type EventResponse struct {
	MinimumID string  `json:"min_id"`
	MaximumID string  `json:"max_id"`
	Events    []Event `json:"events"`
}

type Event struct {
	ID                string `json:"id"`
	Message           string `json:"message"`
	DisplayRecievedAt string `json:"display_received_at"`
	Hostname          string `json:"hostname"`
}
