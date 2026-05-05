package security

var DefaultPublicMethods = []string{
	"/auth.Auth/Auth",
	"/auth.Auth/Create",
	"/auth.Auth/Register",
	"/auth.Auth/Login",
	"/auth.Auth/ValidateToken",
	"/grpc.health.v1.Health/Check",
	"/health.Health/Check",
}
