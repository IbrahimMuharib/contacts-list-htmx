package types

var id = -1

type Contact struct {
	Name  string
	Email string
	Id    int
}

func NewContact(name, email string) Contact {
	id++
	return Contact{Name: name, Email: email, Id: id}
}

func (d Data) HasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func (d Data) IndexOf(id int) int {
	for i, contact := range d.Contacts {
		if contact.Id == id {
			return i
		}
	}
	return -1
}

type Data struct {
	Contacts []Contact
}

func newData() Data {
	return Data{Contacts: []Contact{NewContact("a", "a@g.com")}}
}
