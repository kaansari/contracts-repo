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
