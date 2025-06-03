package model

type Endereco struct {
	EnderecoID int `json:"endereco_id" db:"endereco_id"`
	Logradouro      string  `json:"logradouro" db:"logradouro"`
	Numero          string  `json:"numero" db:"numero"`
	Complemento     *string `json:"complemento" db:"complemento"`
	Bairro          string  `json:"bairro" db:"bairro"`
	CodMunicipio    string `json:"cod_municipio" db:"codmunicipio"`
	Municipio       string  `json:"municipio" db:"municipio"`
	UF              string  `json:"uf" db:"uf"`
	CEP             string  `json:"cep" db:"cep"`
	PontoReferencia *string `json:"ponto_referencia" db:"pontoreferencia"`
}
