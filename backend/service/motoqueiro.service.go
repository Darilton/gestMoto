package service

import (
	"gestmoto/model"
)
var lastID = 4

var motoqueiros = []model.Motoqueiro{
	{ID: 1, Nome: "Jorge", Telefone: "999999999"},
	{ID: 2, Nome: "Dilson", Telefone: "999999999"},
	{ID: 3, Nome: "Dada", Telefone: "999999999"},
}

func GetAll() []model.Motoqueiro {
		return motoqueiros
}

func GetByID(id int) (model.Motoqueiro, bool) {
		for _, m := range motoqueiros {
			if m.ID == id {
				return m, true
			}
		}
		return model.Motoqueiro{}, false
}

func Create(m model.Motoqueiro) model.Motoqueiro {
	lastID++
	m.ID = lastID
	motoqueiros = append(motoqueiros, m)
	return m
}