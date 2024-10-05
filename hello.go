package hello

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Greet() string {
	return fmt.Sprintf("Hello %s!!!", p.Name)
}
