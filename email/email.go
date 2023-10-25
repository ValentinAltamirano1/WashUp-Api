package email

type EmailClientRepository interface {
	ResetPassword(email, uniqueID string) error
}