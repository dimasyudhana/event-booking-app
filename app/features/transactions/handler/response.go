package handler

type TransactionResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    ResponseGetTransaction `json:"data"`
}

type ResponseGetTransaction struct {
	Invoice           string            `json:"invoice"`
	Seller            string            `json:"seller"`
	SEmail            string            `json:"seller_email"`
	Attendee          string            `json:"attendee"`
	AEmail            string            `json:"attendee_email"`
	Title             string            `json:"title"`
	EventDate         string            `json:"event_date"`
	EventTime         string            `json:"event_time"`
	PurchaseStartDate string            `json:"purchase_startdate"`
	PurchaseEndDate   string            `json:"purchase_enddate"`
	Status            string            `json:"status"`
	StatusDate        string            `json:"status_date"`
	ItemDescription   []ResponseTickets `json:"items_description"`
	GrandTotal        uint              `json:"grand_total"`
	PaymentMethod     string            `json:"payment_method"`
}

type ResponseTickets struct {
	TicketCategory string `json:"ticket_category"`
	TicketPrice    uint   `json:"ticket_price"`
	TicketQuantity uint   `json:"ticket_quantity"`
	Subtotal       uint   `json:"subtotal"`
}

type ResponsePayment struct {
	Transaction_ID     string `json:"transaction_id"`
	Order_ID           string `json:"order_id"`
	Gross_Amount       string `json:"gross_amount"`
	Payment_Type       string `json:"payment_type"`
	Bank               string `json:"bank"`
	Transaction_Time   string `json:"transaction_time"`
	Transaction_Status string `json:"transaction_status"`
	VA_Numbers         string `json:"va_numbers"`
}
