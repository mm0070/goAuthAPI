// In login service file i am going to authentication static login information. I want to make it simple and clear.Typically in that part we can query to database or call API end point to check user credentials dynamically. In this part we just return boolean if user credentials are true or false accordingly.

package service

type LoginService interface {
	LoginUser(email string, password string) bool
}
type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "test@google.com",
		password: "test123",
	}
}
func (info *loginInformation) LoginUser(email string, password string) bool {
	return info.email == email && info.password == password
}