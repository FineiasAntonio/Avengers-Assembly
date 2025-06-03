package dto

type RequisicaoNovaSenha struct {
	NovaSenha string `json:"nova_senha"`
}

type CredenciaisUsuario struct {
	Credencial string `json:"credencial"`
	Senha      string `json:"senha"`
}
