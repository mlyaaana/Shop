package domain

type Credentials struct {
	UserId   string
	Username string
	Password string
}

func NewCredentials(userid, username, password string) *Credentials {
	return &Credentials{
		UserId:   userid,
		Username: username,
		Password: password,
	}
}
