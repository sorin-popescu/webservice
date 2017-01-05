package developers

type Developer struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Language string `json:"language"`
	Floor    int    `json:"floor"`
}

func (d *Developer) FindAll() map[int]Developer {
	return Developers
}

func (d *Developer) GetByID(id *int) Developer {
	return Developers[*id]
}

func (d *Developer) AddOne(developer *Developer) map[int]Developer {
	index := len(Developers)
	Developers[index] = *developer
	return Developers
}

var Developers = map[int]Developer{
	1: {
		Name:     "Charlie",
		Age:      23,
		Language: "Go",
		Floor:    5,
	},
	2: {
		Name:     "Bill",
		Age:      23,
		Language: "Java",
		Floor:    5,
	},
	3: {
		Name:     "Kimbley",
		Age:      28,
		Language: "PHP",
		Floor:    2,
	},
	4: {
		Name:     "Roger",
		Age:      34,
		Language: "PHP",
		Floor:    2,
	},
	5: {
		Name:     "Alex",
		Age:      28,
		Language: "Go",
		Floor:    5,
	},
	6: {
		Name:     "Micheal",
		Age:      28,
		Language: "C++",
		Floor:    2,
	},
	7: {
		Name:     "Jill",
		Age:      28,
		Language: "Python",
		Floor:    2,
	},
}
