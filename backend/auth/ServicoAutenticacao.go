package auth

import (
	"backend/model"
	"backend/repository"
	"errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type ServicoAutenticacao struct {
	repositorioUsuario *repository.UsuarioRepository
	jwtKey             []byte
}

type Claims struct {
	Nome      string `json:"nome"`
	CPF       string `json:"cpf"`
	Permissao string `json:"permissao"`
	jwt.StandardClaims
}

func (s *ServicoAutenticacao) GerarToken(usuario *model.Usuario) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Nome: usuario.Nome,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenAssinado, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenAssinado, nil
}

func (s *ServicoAutenticacao) ValidarToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}

func (s *ServicoAutenticacao) AutenticarUsuario(credenciais model.CredenciaisUsuario) (string, error) {
	var (
		usuario *model.Usuario
		err     error
	)

	credencialFormatada := strings.ReplaceAll(credenciais.Credencial, ".", "")
	credencialFormatada = strings.ReplaceAll(credencialFormatada, "-", "")
	credencialFormatada = strings.ReplaceAll(credencialFormatada, " ", "")

	if len(credenciais.Credencial) != 11 {
		usuario, err = s.repositorioUsuario.GetUsuarioByRegistro(credencialFormatada)
	} else {
		usuario, err = s.repositorioUsuario.GetUsuarioByCPF(credencialFormatada)
	}
	if err != nil {
		return "", errors.New("credenciais inválidas")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(credenciais.Senha))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", errors.New("credenciais inválidas")
		}
		return "", errors.New("erro ao validar credenciais")
	}

	token, err := s.GerarToken(usuario)
	if err != nil {
		return "", errors.New("erro ao gerar token")
	}

	return token, nil
}
