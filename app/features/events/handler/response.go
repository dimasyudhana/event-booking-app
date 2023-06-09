package handler

type ResponseGetEvents struct {
	ID            uint   `json:"event_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Hosted_by     string `json:"hosted_by"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	Status        string `json:"status"`
	Category      string `json:"category"`
	Location      string `json:"location"`
	Event_picture string `json:"event_picture"`
}

type GetEventResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    ResponseGetEvent `json:"data"`
} 

type ResponseGetAttendances struct {
	ID           uint   `json:"event_id"`
	Title        string `json:"title"`
    Description  string `json:"description"` 
	HostedBy     string `json:"hosted_by"` 
	Date         string `json:"date"` 
	Time         string `json:"time"`
	Status       string `json:"status"`
	Category     string `json:"category"`
    Location     string `json:"location"` 
    Image        string `json:"event_picture"` 
	Transactions []ResponseTransactions `json:"attendances"` 
	Reviews      []ResponseReviews       `json:"reviews"`
}

type ResponseGetEvent struct {
	ID            uint                   `json:"event_id"`
	Title         string                 `json:"title"`
	Description   string                 `json:"description"`
	Hosted_by     string                 `json:"hosted_by"`
	Date          string                 `json:"date"`
	Time          string                 `json:"time"`
	Status        string                 `json:"status"`
	Category      string                 `json:"category"`
	Location      string                 `json:"location"`
	Event_picture string                 `json:"event_picture"`
	Transactions  []ResponseTransactions `json:"attendances"`
	Reviews       []ResponseReviews      `json:"reviews"`      
} 

type ResponseTransactions struct {
	Username    string `json:"username"`
	UserPicture string `json:"user_picture"`
}

type ResponseReviews struct {
	Username    string `json:"username"`
	UserPicture string `json:"user_picture"`
	Review      string `json:"review"`
}

type ResponseGetAttendance struct { 
	ID           uint   `json:"event_id"`
	Title        string `json:"title"`
    Description  string `json:"description"` 
	HostedBy     string `json:"hosted_by"` 
	Date         string `json:"date"` 
	Time         string `json:"time"`
	Status       string `json:"status"`
	Category     string `json:"category"`
    Location     string `json:"location"` 
    Image        string `json:"event_picture"`
}

type EventResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    EventData `json:"data"`
}

type EventData struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	HostedBy    string           `json:"hosted_by"`
	Date        string           `json:"date"`
	Time        string           `json:"time"`
	Status      string           `json:"status"`
	Category    string           `json:"category"`
	Location    string           `json:"location"`
	Picture     string           `json:"event_picture"`
	Tickets     []TicketResponse `json:"tickets"`
}

type TicketResponse struct {
	Category string `json:"ticket_category"`
	Price    uint   `json:"ticket_price"`
	Quantity uint   `json:"ticket_quantity"`
}

type EventsResponse struct {
	Code       int                 `json:"code"`
	Message    string              `json:"message"`
	Data       []ResponseGetEvents `json:"data"`
	Pagination Pagination          `json:"pagination"`
}

type Getattendance struct {
	Code       int                      `json:"code"` 
	Message    string                   `json:"message"` 
	Data       []ResponseGetAttendances `json:"data"`
	Pagination Pagination               `json:"pagination"`
}

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

type ResponseUpdateEvent struct { 
	ID             uint  `json:"event_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Hosted_by     string `json:"hosted_by"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	Status        string `json:"status"`
	Category      string `json:"category"`
	Location      string `json:"location"`
	Event_picture string `json:"event_picture"`
}
