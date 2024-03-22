package types

type Page struct {
	Data Data
	Form FormData
}

func NewPage() Page {
	return Page{Data: newData(),
		Form: NewFormData()}
}
