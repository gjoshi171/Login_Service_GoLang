package models

//Person is a representation of a person
type Person struct {
	//Userid    bson.ObjectId `bson:"userid" json:"id"`
	//	ID        bson. `bson:"_id" json:"_id"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Password  string `bson:"password" json:"password"`
	Username  string `bson:"userName" json:"userName"`
	Email     string `bson:"email" json:"email"`
}

func (p Person) GetFirstName() string {
	return p.FirstName
}

func (p *Person) SetFirstName(FirstName string) {
	p.FirstName = FirstName
}
func (p Person) GetLastName() string {
	return p.LastName
}

func (p *Person) SetLastName(LastName string) {
	p.LastName = LastName
}
func (p *Person) GetPassword() string {
	return p.Password
}

func (p *Person) SetPassword(Password string) {
	p.Password = Password
}
func (p *Person) GetUsername() string {
	return p.Username
}

func (p *Person) SetUsername(Username string) {
	p.Username = Username
}
func (p Person) GetEmail() string {
	return p.Email
}

func (p *Person) SetEmail(Email string) {
	p.Email = Email
}
