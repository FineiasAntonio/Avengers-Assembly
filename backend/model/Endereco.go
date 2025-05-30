package model

type Endereco struct {
	ID              int     `json:"-" db:"id"`
	Logradouro      string  `json:"logradouro" db:"logradouro"`
	Numero          string  `json:"numero" db:"numero"`
	Complemento     *string `json:"complemento" db:"complemento"`
	CodMunicipio    *string `json:"cod_municipio" db:"codmunicipio"`
	Municipio       string  `json:"municipio" db:"municipio"`
	Bairro          string  `json:"bairro" db:"bairro"`
	UF              string  `json:"uf" db:"uf"`
	CEP             string  `json:"cep" db:"cep"`
	PontoReferencia *string `json:"ponto_referencia" db:"pontoreferencia"`
}
