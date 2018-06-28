package structs

import "fmt"

//Platzi is an interface of PlatziCourse and PlatziCareer
type Platzi interface {
	Subscribe(name string)
}

func InterfaceTest() {
	platziCourse := PlatziCourse{Name: "Go", Slug: "go", Skills: []string{"1", "2"}}
	platziCareer := new(PlatziCareer)
	platziCareer.Name = "GoCareer"
	platziCareer.Slug = "go"
	callSubscribe(platziCourse)
	callSubscribe(platziCareer)
	fmt.Println(platziCourse)
}

func callSubscribe(p Platzi) {
	p.Subscribe("Jean")
}
