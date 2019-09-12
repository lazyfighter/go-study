package reflect

import "log"

type person struct {
	name    string
	age     int
	hobbies []string
	mapTest map[string]interface{}
	boy     *struct{}
}

func (p *person) SetHobby(hobbies ...string) {
	for _, hobby := range hobbies {
		log.Println(hobby)
		p.hobbies = append(p.hobbies, hobby)
	}
}

func (p *person) GetHobby() []string {
	return p.hobbies
}

func (p *person) SetName(name string) {
	p.name = name
}

func (p *person) GetName() string {
	return p.name
}

func (p *person) SetAge(age int) {
	p.age = age
}

func (p *person) GetAge() int {
	return p.age
}

func (p person) Test() {
	log.Println(p.age)
}
