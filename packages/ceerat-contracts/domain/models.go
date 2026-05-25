package domain

// User is the shared business representation of an authenticated user.
// It intentionally has no transport, database, or framework annotations.
type User struct {
	ID       string
	Name     string
	Company  string
	Email    string
	Password string
	Token    string
	Role     string
	Status   string
}

// Patient is the shared business representation of a patient.
// It intentionally has no transport, database, or framework annotations.
type Patient struct {
	ID              string
	FirstName       string
	LastName        string
	DOB             string
	DOS             string
	Location        string
	ICDCodes        string
	COVIDTest       bool
	COVIDTestResult bool
	RSVTest         bool
	RSVTestResult   bool
	StrepTest       bool
	StrepTestResult bool
	FluTest         bool
	FluTestResult   bool
}

// Address is the shared business representation of a customer address.
type Address struct {
	Line1      string
	Line2      string
	City       string
	State      string
	Country    string
	PostalCode string
}

// Customer is the shared business representation of a customer assigned to a user.
type Customer struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Address   *Address
	UserID    string
	User      *User
	CreatedAt string
	UpdatedAt string
}

// Service is the shared business representation of a customer-orderable service.
type Service struct {
	ID           string
	Name         string
	Category     string
	Price        float64
	Type         string
	ScheduleDate string
	StartDate    string
	AgentName    string
	Description  string
	CreatedAt    string
	UpdatedAt    string
	SKU          string
	Active       bool
}

// Product is a service-catalog product exposed through the catalog service boundary.
type Product struct {
	ID          string
	Name        string
	Description string
	SKU         string
	Price       float64
	Active      bool
	CreatedAt   string
	UpdatedAt   string
}

// CustomerService links a customer to an ordered service.
type CustomerService struct {
	ID         string
	CustomerID string
	ServiceID  string
	Customer   *Customer
	Service    *Service
	Status     string
	OrderedAt  string
}

// Cart groups service and product items selected by a customer before checkout.
type Cart struct {
	ID         string
	CustomerID string
	Customer   *Customer
	Items      []*CartItem
	Subtotal   float64
	Total      float64
	CreatedAt  string
	UpdatedAt  string
}

// CartItem stores one selected service or product in a cart.
type CartItem struct {
	ID         string
	CartID     string
	ItemType   string
	ServiceID  string
	ProductID  string
	Service    *Service
	Product    *Product
	Quantity   int32
	UnitPrice  float64
	TotalPrice float64
	Notes      string
	CreatedAt  string
	UpdatedAt  string
}

// Order groups one or more services for a customer.
type Order struct {
	ID           string
	CustomerID   string
	UserID       string
	OrderNumber  string
	Status       string
	ScheduleDate string
	StartDate    string
	DueDate      string
	Subtotal     float64
	Tax          float64
	Total        float64
	Notes        string
	Customer     *Customer
	Services     []*OrderService
	CreatedAt    string
	UpdatedAt    string
}

// OrderService stores the service snapshot captured when an order is created.
type OrderService struct {
	ID           string
	OrderID      string
	ServiceID    string
	ServiceName  string
	Category     string
	Type         string
	UnitPrice    float64
	Quantity     int32
	TotalPrice   float64
	AgentName    string
	ScheduleDate string
	StartDate    string
	DueDate      string
	Service      *Service
	CreatedAt    string
	UpdatedAt    string
}

// CreateOrderServiceInput is the domain input for adding a service to an order.
type CreateOrderServiceInput struct {
	ServiceID    string
	Quantity     int32
	AgentName    string
	ScheduleDate string
	StartDate    string
	DueDate      string
}
