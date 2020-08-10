package main

// begin ->
// order service
// stock service -> Rollback service (out_of_stock_event) -> user service
// payment service (in_stock_event)                       -> user service
// delivery service -> end

//order service
type OrderCreatedEvent struct {
	UUID      string // random
	OrderID   string // service id
	OrderType int
	Status    string // pending
	Price     int
}

// stock service
type OutOfStockEvent struct {
	UUID    string
	OrderID string
	Status  string //out_of_stock
	Price   int
}

type InStockEvent struct {
	Status string //in_stock
}

//payment service
type BilledOrderService struct {
	UUID    string
	OrderID string
	Status  string //order_billed
	Address string //park avenue 1317, NYC, NY
}

type PaymentFailedEvent struct {
	UUID    string
	OrderID string
	Status  string //payment_failed
	Reason  string //insufficient_funds	| cc_vencida |
}

// delivery service
type OrderDeliveredEvent struct {
	UUID          string
	OrderID       string
	ReferenceCode string // DHL, UPS ->
	TotalAmount   int    //price + taxes
	Status        string //order_delivered
}
