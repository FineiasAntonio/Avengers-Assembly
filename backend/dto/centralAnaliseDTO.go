package dto

type GraficoPacientesDTO struct {
	QuantidadePacientes int `json:"quantidade_pacientes"`
}

type GraficoPacientesPorIdadeDTO struct {
	Total    int `json:"total"`
	Qtd25a30 int `json:"qtd_25_a_30"`
	Qtd30a40 int `json:"qtd_30_a_40"`
	Qtd40a50 int `json:"qtd_40_a_50"`
	Qtd50a60 int `json:"qtd_50_a_60"`
	Qtd60a65 int `json:"qtd_60_a_65"`
}

type GraficoPacientesPorRacaDTO struct {
	Branca   string `json:"branca"`
	Preta    string `json:"preta"`
	Parda    string `json:"parda"`
	Amarela  string `json:"amarela"`
	Indigena string `json:"indigena"`
}

type GraficoPacientesPorEscolaridadeDTO struct {
	Analfabeta            string `json:"analfabeta"`
	FundamentalIncompleto string `json:"fundamental_incompleto"`
	FundamentalCompleto   string `json:"fundamental_completo"`
	MedioIncompleto       string `json:"medio_incompleto"`
	MedioCompleto         string `json:"medio_completo"`
	SuperiorIncompleto    string `json:"superior_incompleto"`
	SuperiorCompleto      string `json:"superior_completo"`
}

type MapaPacientesPorRegiao struct {
	Bairro     string `json:"bairro"`
	Quantidade int    `json:"quantidade"`
}
