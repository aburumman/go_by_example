package main

import (
	"fmt"
)

type Personel struct{
	Name string
	ID  uint8
	Dept string
}

func (p *Personel) NewPersonnel(name, dept string, id uint8) (Personel){
	return &Personnel{
		Name: name,
		Dept: dept,
		ID: id

	}

func GetPersonnel(id uint) (string, error){
	return fmt.Sprintf("Name: %s \n, Dept: %s \n, ID: %d", )
}
}