package mapper

import (
	"testing"

	"github.com/kaansari/ceerat-platform/packages/ceerat-contracts/domain"
	authpb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/auth"
	customerpb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/customer"
	patientpb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/patient"
	servicepb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/service"
)

func TestUserFromProtoHandlesNil(t *testing.T) {
	if got := UserFromProto(nil); got != nil {
		t.Fatalf("expected nil, got %#v", got)
	}
}

func TestUserProtoRoundTrip(t *testing.T) {
	in := &authpb.User{
		Id:       "user-1",
		Name:     "Khalid",
		Company:  "Ceerat",
		Email:    "khalid@example.com",
		Password: "secret",
		Token:    "token",
	}

	model := UserFromProto(in)
	if model == nil {
		t.Fatal("expected user model")
	}

	if model.ID != in.Id ||
		model.Name != in.Name ||
		model.Company != in.Company ||
		model.Email != in.Email ||
		model.Password != in.Password ||
		model.Token != in.Token {
		t.Fatalf("user model did not match proto: %#v", model)
	}

	out := UserToProto(model)
	if out.GetId() != in.GetId() ||
		out.GetName() != in.GetName() ||
		out.GetCompany() != in.GetCompany() ||
		out.GetEmail() != in.GetEmail() ||
		out.GetPassword() != in.GetPassword() ||
		out.GetToken() != in.GetToken() {
		t.Fatalf("user proto did not round-trip: %#v", out)
	}
}

func TestUsersToProtoPreservesLengthAndOrder(t *testing.T) {
	users := []*domain.User{
		{ID: "1", Email: "one@example.com"},
		{ID: "2", Email: "two@example.com"},
	}

	got := UsersToProto(users)

	if len(got) != len(users) {
		t.Fatalf("expected %d users, got %d", len(users), len(got))
	}
	if got[0].GetId() != "1" || got[1].GetId() != "2" {
		t.Fatalf("expected converted users to preserve order, got %#v", got)
	}
}

func TestPatientFromProtoHandlesNil(t *testing.T) {
	if got := PatientFromProto(nil); got != nil {
		t.Fatalf("expected nil, got %#v", got)
	}
}

func TestPatientProtoRoundTrip(t *testing.T) {
	in := &patientpb.Patient{
		Id:              "patient-1",
		Fname:           "Jane",
		Lname:           "Doe",
		Dob:             "1990-01-02",
		Dos:             "2024-03-04",
		Location:        "NY",
		Icdcodes:        "A00",
		CovidTest:       true,
		CovidTestResult: false,
		RsvTest:         true,
		RsvTestResult:   true,
		StrepTest:       false,
		StrepTestResult: false,
		FluTest:         true,
		FluTestResult:   false,
	}

	model := PatientFromProto(in)
	if model == nil {
		t.Fatal("expected patient model")
	}

	if model.ID != in.Id ||
		model.FirstName != in.Fname ||
		model.LastName != in.Lname ||
		model.DOB != in.Dob ||
		model.DOS != in.Dos ||
		model.Location != in.Location ||
		model.ICDCodes != in.Icdcodes ||
		model.COVIDTest != in.CovidTest ||
		model.COVIDTestResult != in.CovidTestResult ||
		model.RSVTest != in.RsvTest ||
		model.RSVTestResult != in.RsvTestResult ||
		model.StrepTest != in.StrepTest ||
		model.StrepTestResult != in.StrepTestResult ||
		model.FluTest != in.FluTest ||
		model.FluTestResult != in.FluTestResult {
		t.Fatalf("patient model did not match proto: %#v", model)
	}

	out := PatientToProto(model)
	if out.GetId() != in.GetId() ||
		out.GetFname() != in.GetFname() ||
		out.GetLname() != in.GetLname() ||
		out.GetDob() != in.GetDob() ||
		out.GetDos() != in.GetDos() ||
		out.GetLocation() != in.GetLocation() ||
		out.GetIcdcodes() != in.GetIcdcodes() ||
		out.GetCovidTest() != in.GetCovidTest() ||
		out.GetCovidTestResult() != in.GetCovidTestResult() ||
		out.GetRsvTest() != in.GetRsvTest() ||
		out.GetRsvTestResult() != in.GetRsvTestResult() ||
		out.GetStrepTest() != in.GetStrepTest() ||
		out.GetStrepTestResult() != in.GetStrepTestResult() ||
		out.GetFluTest() != in.GetFluTest() ||
		out.GetFluTestResult() != in.GetFluTestResult() {
		t.Fatalf("patient proto did not round-trip: %#v", out)
	}
}

func TestPatientsToProtoPreservesLengthAndOrder(t *testing.T) {
	patients := []*domain.Patient{
		{ID: "1", FirstName: "One"},
		{ID: "2", FirstName: "Two"},
	}

	got := PatientsToProto(patients)

	if len(got) != len(patients) {
		t.Fatalf("expected %d patients, got %d", len(patients), len(got))
	}
	if got[0].GetId() != "1" || got[1].GetId() != "2" {
		t.Fatalf("expected converted patients to preserve order, got %#v", got)
	}
}

func TestCustomerProtoRoundTrip(t *testing.T) {
	in := &customerpb.Customer{
		Id:        "customer-1",
		FirstName: "Amina",
		LastName:  "Ansari",
		Email:     "amina@example.com",
		Phone:     "555-0100",
		Address: &customerpb.Address{
			Line1:      "100 Main",
			Line2:      "Suite 2",
			City:       "Chicago",
			State:      "IL",
			Country:    "US",
			PostalCode: "60601",
		},
		UserId:    "user-1",
		User:      &authpb.User{Id: "user-1", Email: "owner@example.com"},
		CreatedAt: "2026-01-01T00:00:00Z",
		UpdatedAt: "2026-01-02T00:00:00Z",
	}

	model := CustomerFromProto(in)
	if model == nil {
		t.Fatal("expected customer model")
	}
	if model.ID != in.Id ||
		model.FirstName != in.FirstName ||
		model.LastName != in.LastName ||
		model.Email != in.Email ||
		model.Phone != in.Phone ||
		model.UserID != in.UserId ||
		model.User.ID != in.User.Id ||
		model.Address.PostalCode != in.Address.PostalCode ||
		model.CreatedAt != in.CreatedAt ||
		model.UpdatedAt != in.UpdatedAt {
		t.Fatalf("customer model did not match proto: %#v", model)
	}

	out := CustomerToProto(model)
	if out.GetId() != in.GetId() ||
		out.GetFirstName() != in.GetFirstName() ||
		out.GetLastName() != in.GetLastName() ||
		out.GetEmail() != in.GetEmail() ||
		out.GetPhone() != in.GetPhone() ||
		out.GetUserId() != in.GetUserId() ||
		out.GetUser().GetId() != in.GetUser().GetId() ||
		out.GetAddress().GetPostalCode() != in.GetAddress().GetPostalCode() ||
		out.GetCreatedAt() != in.GetCreatedAt() ||
		out.GetUpdatedAt() != in.GetUpdatedAt() {
		t.Fatalf("customer proto did not round-trip: %#v", out)
	}
}

func TestServiceProtoRoundTrip(t *testing.T) {
	in := &servicepb.Service{
		Id:           "service-1",
		Name:         "Consultation",
		Category:     "Clinical",
		Price:        125.50,
		Type:         "virtual",
		ScheduleDate: "2026-01-10",
		StartDate:    "2026-01-11",
		AgentName:    "Agent Smith",
		Description:  "Initial visit",
		CreatedAt:    "2026-01-01T00:00:00Z",
		UpdatedAt:    "2026-01-02T00:00:00Z",
	}

	model := ServiceFromProto(in)
	if model == nil {
		t.Fatal("expected service model")
	}
	if model.ID != in.Id ||
		model.Name != in.Name ||
		model.Category != in.Category ||
		model.Price != in.Price ||
		model.Type != in.Type ||
		model.ScheduleDate != in.ScheduleDate ||
		model.StartDate != in.StartDate ||
		model.AgentName != in.AgentName ||
		model.Description != in.Description ||
		model.CreatedAt != in.CreatedAt ||
		model.UpdatedAt != in.UpdatedAt {
		t.Fatalf("service model did not match proto: %#v", model)
	}

	out := ServiceToProto(model)
	if out.GetId() != in.GetId() ||
		out.GetName() != in.GetName() ||
		out.GetCategory() != in.GetCategory() ||
		out.GetPrice() != in.GetPrice() ||
		out.GetType() != in.GetType() ||
		out.GetScheduleDate() != in.GetScheduleDate() ||
		out.GetStartDate() != in.GetStartDate() ||
		out.GetAgentName() != in.GetAgentName() ||
		out.GetDescription() != in.GetDescription() ||
		out.GetCreatedAt() != in.GetCreatedAt() ||
		out.GetUpdatedAt() != in.GetUpdatedAt() {
		t.Fatalf("service proto did not round-trip: %#v", out)
	}
}

func TestCustomerServiceProtoRoundTrip(t *testing.T) {
	in := &servicepb.CustomerService{
		Id:         "customer-service-1",
		CustomerId: "customer-1",
		ServiceId:  "service-1",
		Customer:   &customerpb.Customer{Id: "customer-1", FirstName: "Amina"},
		Service:    &servicepb.Service{Id: "service-1", Name: "Consultation"},
		Status:     "ordered",
		OrderedAt:  "2026-01-03T00:00:00Z",
	}

	model := CustomerServiceFromProto(in)
	if model == nil {
		t.Fatal("expected customer service model")
	}
	if model.ID != in.Id ||
		model.CustomerID != in.CustomerId ||
		model.ServiceID != in.ServiceId ||
		model.Customer.ID != in.Customer.Id ||
		model.Service.ID != in.Service.Id ||
		model.Status != in.Status ||
		model.OrderedAt != in.OrderedAt {
		t.Fatalf("customer service model did not match proto: %#v", model)
	}

	out := CustomerServiceToProto(model)
	if out.GetId() != in.GetId() ||
		out.GetCustomerId() != in.GetCustomerId() ||
		out.GetServiceId() != in.GetServiceId() ||
		out.GetCustomer().GetId() != in.GetCustomer().GetId() ||
		out.GetService().GetId() != in.GetService().GetId() ||
		out.GetStatus() != in.GetStatus() ||
		out.GetOrderedAt() != in.GetOrderedAt() {
		t.Fatalf("customer service proto did not round-trip: %#v", out)
	}
}
