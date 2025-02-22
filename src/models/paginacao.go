package models

import "math"

type Paginacao struct {
	PaginaAtual       int   `json:"pagina_atual"`
	QuantidadePaginas int   `json:"quantidade_paginas"`
	Limite            int   `json:"limite"`
	Total             int64 `json:"total"`
}

// GetOffset: retorna o offset da páginação
func (p *Paginacao) GetOffset() int {
	return (p.PaginaAtual - 1) * p.Limite
}

// CalcularQuantidadePaginas: calcula a quantidade de páginas. total / limite
func (p *Paginacao) CalcularQuantidadePaginas() {
	p.QuantidadePaginas = int(math.Ceil((float64(p.Total) / float64(p.Limite))))
}
