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
