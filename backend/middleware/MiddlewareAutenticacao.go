package middleware

import (
	"backend/auth"
	"context"
	"net/http"
	"strings"
)

type MiddlewareAutenticacao struct {
	ServicoAutenticacao *auth.ServicoAutenticacao
}

func NewMiddlewareAutenticacao(servicoAutenticacao *auth.ServicoAutenticacao) *MiddlewareAutenticacao {
	return &MiddlewareAutenticacao{
		ServicoAutenticacao: servicoAutenticacao,
	}
}

func (m *MiddlewareAutenticacao) middlewareAutenticacao(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return
		}

		claims, err := m.ServicoAutenticacao.ValidarToken(strings.TrimPrefix(token, "Bearer "))
		if err != nil {
			http.Error(w, "Token inválido: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "usuarioAutenticado", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
