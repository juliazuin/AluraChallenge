package model

type Despesas struct {
	id        int     `json:"id" gorm:"primary_key"`
	Valor     float64 `json:"valor"`
	Descricao string  `json:"descricao"`
	DataAtual string  `json:"data_atual"`
}
