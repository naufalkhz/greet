package greet

type Person struct {
	ID    int
	Name  string
	Hobby string
}

func (p Person) ToString() string {
	return ("Hallo nama saya " + p.Name + ", Hobi saya adalah " + p.Hobby)
}

func (p *Person) SetNama(name string) {
	p.Name = name
}
func (p *Person) SetHobby(hobby string) {
	p.Hobby = hobby
}
func (p *Person) SetId(id int) {
	p.ID = id
}

func (p Person) GetNama() string {
	return p.Name
}
func (p Person) GetId() int {
	return p.ID
}
func (p Person) GetHobby() string {
	return p.Hobby
}

func HaiPerson() Person {
	return Person{}
}
