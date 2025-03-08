package furniture

import (
	"fmt"
	"sync"
)

type Table struct {
	Name string
}

var tableInstance *Table
var once sync.Once

func GetInstance() *Table {
	once.Do(func() {
		fmt.Println("Its first time GetInstance is called. Creating a singleton..")
		tableInstance = &Table{Name: "default table"}
	})
	return tableInstance
}

func (t *Table) GetName() string {
	return t.Name
}

func (t *Table) SetName(name string) {
	t.Name = name
}
