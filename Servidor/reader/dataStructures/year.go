package dataStructures

type Year struct {
	Year  int
	Meses []Mes
}

func (year Year) ShowMeses() []Mes {
	var allMeses []Mes
	for i := 0; i < len(year.Meses); i++ {
		allMeses = append(allMeses, year.Meses[i])
	}
	return allMeses
}

func (year Year) FindByName(name string) *Mes {
	for i := 0; i < len(year.Meses); i++ {
		if name == year.Meses[i].Nombre {
			return &year.Meses[i]
		}
	}
	return nil
}
