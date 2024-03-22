package types

type Contact struct {
	Name  string
	Email string
}

func NewContact(name, email string) Contact {
	return Contact{Name: name, Email: email}
}

func (d Data) HasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

type Data struct {
	Contacts []Contact
}

func newData() Data {
	return Data{Contacts: []Contact{NewContact("a", "a@g.com")}}
}
