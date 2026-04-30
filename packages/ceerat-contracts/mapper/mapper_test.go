package mapper

import (
	"testing"

	"github.com/kaansari/ceerat-platform/packages/ceerat-contracts/domain"
	authpb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/auth"
	patientpb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/patient"
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
