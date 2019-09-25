package models

type Result struct {
	Person *Person `json:"address,omitempty"`
}

func (r Result) GetPerson() *Person {
	return r.Person
}

func (r *Result) SetPerson(Person *Person) {
	r.Person = Person
}
