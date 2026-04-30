package mapper

import (
	"github.com/kaansari/ceerat-platform/packages/ceerat-contracts/domain"
	authpb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/auth"
	patientpb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/patient"
)

func UserFromProto(in *authpb.User) *domain.User {
	if in == nil {
		return nil
	}
	return &domain.User{
		ID:       in.Id,
		Name:     in.Name,
		Company:  in.Company,
		Email:    in.Email,
		Password: in.Password,
		Token:    in.Token,
	}
}

func UserToProto(in *domain.User) *authpb.User {
	if in == nil {
		return nil
	}
	return &authpb.User{
		Id:       in.ID,
		Name:     in.Name,
		Company:  in.Company,
		Email:    in.Email,
		Password: in.Password,
		Token:    in.Token,
	}
}

func UsersToProto(in []*domain.User) []*authpb.User {
	out := make([]*authpb.User, 0, len(in))
	for _, u := range in {
		out = append(out, UserToProto(u))
	}
	return out
}

func PatientFromProto(in *patientpb.Patient) *domain.Patient {
	if in == nil {
		return nil
	}
	return &domain.Patient{
		ID:              in.Id,
		FirstName:       in.Fname,
		LastName:        in.Lname,
		DOB:             in.Dob,
		DOS:             in.Dos,
		Location:        in.Location,
		ICDCodes:        in.Icdcodes,
		COVIDTest:       in.CovidTest,
		COVIDTestResult: in.CovidTestResult,
		RSVTest:         in.RsvTest,
		RSVTestResult:   in.RsvTestResult,
		StrepTest:       in.StrepTest,
		StrepTestResult: in.StrepTestResult,
		FluTest:         in.FluTest,
		FluTestResult:   in.FluTestResult,
	}
}

func PatientToProto(in *domain.Patient) *patientpb.Patient {
	if in == nil {
		return nil
	}
	return &patientpb.Patient{
		Id:              in.ID,
		Fname:           in.FirstName,
		Lname:           in.LastName,
		Dob:             in.DOB,
		Dos:             in.DOS,
		Location:        in.Location,
		Icdcodes:        in.ICDCodes,
		CovidTest:       in.COVIDTest,
		CovidTestResult: in.COVIDTestResult,
		RsvTest:         in.RSVTest,
		RsvTestResult:   in.RSVTestResult,
		StrepTest:       in.StrepTest,
		StrepTestResult: in.StrepTestResult,
		FluTest:         in.FluTest,
		FluTestResult:   in.FluTestResult,
	}
}

func PatientsToProto(in []*domain.Patient) []*patientpb.Patient {
	out := make([]*patientpb.Patient, 0, len(in))
	for _, p := range in {
		out = append(out, PatientToProto(p))
	}
	return out
}
