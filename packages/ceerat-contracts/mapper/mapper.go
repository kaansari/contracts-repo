package mapper

import (
	"github.com/kaansari/ceerat-platform/packages/ceerat-contracts/domain"
	authpb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/auth"
	customerpb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/customer"
	orderpb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/order"
	patientpb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/patient"
	servicepb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/service"
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

func AddressFromProto(in *customerpb.Address) *domain.Address {
	if in == nil {
		return nil
	}
	return &domain.Address{
		Line1:      in.Line1,
		Line2:      in.Line2,
		City:       in.City,
		State:      in.State,
		Country:    in.Country,
		PostalCode: in.PostalCode,
	}
}

func AddressToProto(in *domain.Address) *customerpb.Address {
	if in == nil {
		return nil
	}
	return &customerpb.Address{
		Line1:      in.Line1,
		Line2:      in.Line2,
		City:       in.City,
		State:      in.State,
		Country:    in.Country,
		PostalCode: in.PostalCode,
	}
}

func CustomerFromProto(in *customerpb.Customer) *domain.Customer {
	if in == nil {
		return nil
	}
	return &domain.Customer{
		ID:        in.Id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Phone:     in.Phone,
		Address:   AddressFromProto(in.Address),
		UserID:    in.UserId,
		User:      UserFromProto(in.User),
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}

func CustomerToProto(in *domain.Customer) *customerpb.Customer {
	if in == nil {
		return nil
	}
	return &customerpb.Customer{
		Id:        in.ID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Phone:     in.Phone,
		Address:   AddressToProto(in.Address),
		UserId:    in.UserID,
		User:      UserToProto(in.User),
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}

func CustomersToProto(in []*domain.Customer) []*customerpb.Customer {
	out := make([]*customerpb.Customer, 0, len(in))
	for _, c := range in {
		out = append(out, CustomerToProto(c))
	}
	return out
}

func ServiceFromProto(in *servicepb.Service) *domain.Service {
	if in == nil {
		return nil
	}
	return &domain.Service{
		ID:           in.Id,
		Name:         in.Name,
		Category:     in.Category,
		Price:        in.Price,
		Type:         in.Type,
		ScheduleDate: in.ScheduleDate,
		StartDate:    in.StartDate,
		AgentName:    in.AgentName,
		Description:  in.Description,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}

func ServiceToProto(in *domain.Service) *servicepb.Service {
	if in == nil {
		return nil
	}
	return &servicepb.Service{
		Id:           in.ID,
		Name:         in.Name,
		Category:     in.Category,
		Price:        in.Price,
		Type:         in.Type,
		ScheduleDate: in.ScheduleDate,
		StartDate:    in.StartDate,
		AgentName:    in.AgentName,
		Description:  in.Description,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}

func ServicesToProto(in []*domain.Service) []*servicepb.Service {
	out := make([]*servicepb.Service, 0, len(in))
	for _, s := range in {
		out = append(out, ServiceToProto(s))
	}
	return out
}

func CustomerServiceFromProto(in *servicepb.CustomerService) *domain.CustomerService {
	if in == nil {
		return nil
	}
	return &domain.CustomerService{
		ID:         in.Id,
		CustomerID: in.CustomerId,
		ServiceID:  in.ServiceId,
		Customer:   CustomerFromProto(in.Customer),
		Service:    ServiceFromProto(in.Service),
		Status:     in.Status,
		OrderedAt:  in.OrderedAt,
	}
}

func CustomerServiceToProto(in *domain.CustomerService) *servicepb.CustomerService {
	if in == nil {
		return nil
	}
	return &servicepb.CustomerService{
		Id:         in.ID,
		CustomerId: in.CustomerID,
		ServiceId:  in.ServiceID,
		Customer:   CustomerToProto(in.Customer),
		Service:    ServiceToProto(in.Service),
		Status:     in.Status,
		OrderedAt:  in.OrderedAt,
	}
}

func CustomerServicesToProto(in []*domain.CustomerService) []*servicepb.CustomerService {
	out := make([]*servicepb.CustomerService, 0, len(in))
	for _, cs := range in {
		out = append(out, CustomerServiceToProto(cs))
	}
	return out
}

func OrderFromProto(in *orderpb.Order) *domain.Order {
	if in == nil {
		return nil
	}
	return &domain.Order{
		ID:           in.Id,
		CustomerID:   in.CustomerId,
		UserID:       in.UserId,
		OrderNumber:  in.OrderNumber,
		Status:       in.Status,
		ScheduleDate: in.ScheduleDate,
		StartDate:    in.StartDate,
		DueDate:      in.DueDate,
		Subtotal:     in.Subtotal,
		Tax:          in.Tax,
		Total:        in.Total,
		Notes:        in.Notes,
		Customer:     CustomerFromProto(in.Customer),
		Services:     OrderServicesFromProto(in.Services),
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}

func OrderToProto(in *domain.Order) *orderpb.Order {
	if in == nil {
		return nil
	}
	return &orderpb.Order{
		Id:           in.ID,
		CustomerId:   in.CustomerID,
		UserId:       in.UserID,
		OrderNumber:  in.OrderNumber,
		Status:       in.Status,
		ScheduleDate: in.ScheduleDate,
		StartDate:    in.StartDate,
		DueDate:      in.DueDate,
		Subtotal:     in.Subtotal,
		Tax:          in.Tax,
		Total:        in.Total,
		Notes:        in.Notes,
		Customer:     CustomerToProto(in.Customer),
		Services:     OrderServicesToProto(in.Services),
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}

func OrdersToProto(in []*domain.Order) []*orderpb.Order {
	out := make([]*orderpb.Order, 0, len(in))
	for _, order := range in {
		out = append(out, OrderToProto(order))
	}
	return out
}

func OrderServiceFromProto(in *orderpb.OrderService) *domain.OrderService {
	if in == nil {
		return nil
	}
	return &domain.OrderService{
		ID:           in.Id,
		OrderID:      in.OrderId,
		ServiceID:    in.ServiceId,
		ServiceName:  in.ServiceName,
		Category:     in.Category,
		Type:         in.Type,
		UnitPrice:    in.UnitPrice,
		Quantity:     in.Quantity,
		TotalPrice:   in.TotalPrice,
		AgentName:    in.AgentName,
		ScheduleDate: in.ScheduleDate,
		StartDate:    in.StartDate,
		DueDate:      in.DueDate,
		Service:      ServiceFromProto(in.Service),
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}

func OrderServiceToProto(in *domain.OrderService) *orderpb.OrderService {
	if in == nil {
		return nil
	}
	return &orderpb.OrderService{
		Id:           in.ID,
		OrderId:      in.OrderID,
		ServiceId:    in.ServiceID,
		ServiceName:  in.ServiceName,
		Category:     in.Category,
		Type:         in.Type,
		UnitPrice:    in.UnitPrice,
		Quantity:     in.Quantity,
		TotalPrice:   in.TotalPrice,
		AgentName:    in.AgentName,
		ScheduleDate: in.ScheduleDate,
		StartDate:    in.StartDate,
		DueDate:      in.DueDate,
		Service:      ServiceToProto(in.Service),
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}

func OrderServicesFromProto(in []*orderpb.OrderService) []*domain.OrderService {
	out := make([]*domain.OrderService, 0, len(in))
	for _, service := range in {
		out = append(out, OrderServiceFromProto(service))
	}
	return out
}

func OrderServicesToProto(in []*domain.OrderService) []*orderpb.OrderService {
	out := make([]*orderpb.OrderService, 0, len(in))
	for _, service := range in {
		out = append(out, OrderServiceToProto(service))
	}
	return out
}

func CreateOrderServiceInputFromProto(in *orderpb.CreateOrderServiceInput) domain.CreateOrderServiceInput {
	if in == nil {
		return domain.CreateOrderServiceInput{}
	}
	return domain.CreateOrderServiceInput{
		ServiceID:    in.ServiceId,
		Quantity:     in.Quantity,
		AgentName:    in.AgentName,
		ScheduleDate: in.ScheduleDate,
		StartDate:    in.StartDate,
		DueDate:      in.DueDate,
	}
}

func CreateOrderServiceInputsFromProto(in []*orderpb.CreateOrderServiceInput) []domain.CreateOrderServiceInput {
	out := make([]domain.CreateOrderServiceInput, 0, len(in))
	for _, service := range in {
		out = append(out, CreateOrderServiceInputFromProto(service))
	}
	return out
}
