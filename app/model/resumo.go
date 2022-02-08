package model

type Resumo struct {
	TotalReceitasMes  float64    `json:"total_receitas_mes"`
	TotalDespesasMes  float64    `json:"total_despesas_mes"`
	SaldoFinalMes     float64    `json:"saldo_final_mes"`
	DespesaByCategory *[]Despesa `json:"despesa_por_categoria"`
}
