package domain

type Credentials struct {
	UserId   string
	Username string
	Password string
}

func NewCredentials(userId, username, password string) *Credentials {
	return &Credentials{
		UserId:   userId,
		Username: username,
		Password: password,
	}
}
