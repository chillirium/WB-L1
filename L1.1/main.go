package main

import "fmt"

type Human struct {
	Age  int
	Name string
}

func (h *Human) getAge() int {
	return h.Age
}

func (h *Human) getName() string {
	return h.Name
}

type Action struct {
	Human
}

func (a *Action) whoDid() {
	fmt.Println(a.Name, "сделал")
}

func main() {

	w := Action{
		Human: Human{
			Age:  26,
			Name: "Валерий",
		},
	}

	//пользуемся методом встроенного Human через экземпляр Action
	fmt.Println(w.getAge(), w.Human.Age, w.Age)
	fmt.Println(w.getName(), w.Human.Name, w.Name)

	//собственный метод Action
	w.whoDid()
}
