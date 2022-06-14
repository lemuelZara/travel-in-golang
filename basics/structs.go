package basics

type Car struct {
	Name  string
	Brand string
}

type Person struct {
	Name, LastName string
	Age            int
	Car            Car
}

var persons []Person = []Person{}

func Add(p Person) Person {
	persons = append(persons, p)

	return persons[len(persons)-1]
}

func (p *Person) Update(Name string) {
	p.Name = Name
}

func GetAll() []Person {
	return persons
}
