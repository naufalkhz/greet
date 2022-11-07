package greet

type Person struct {
	ID    int
	Name  string
	Hobby string
}

func (p Person) toString() string {
	return ("Hallo nama saya " + p.Name + ", Hobi saya adalah " + p.Hobby)
}

func (p *Person) setNama(name string) {
	p.Name = name
}
func (p *Person) setHobby(hobby string) {
	p.Hobby = hobby
}
func (p *Person) setId(id int) {
	p.ID = id
}

func (p Person) getNama() string {
	return p.Name
}
func (p Person) getId() int {
	return p.ID
}
func (p Person) getHobby() string {
	return p.Hobby
}

func HaiPerson() Person {
	return Person{}
}
