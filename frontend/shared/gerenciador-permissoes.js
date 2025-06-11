export const Permissao = {
  ACESSO_ATENDIMENTO: "ACESSO_ATENDIMENTO",
  ACESSO_EXAMES: "ACESSO_EXAMES",
  ACESSO_LABORATORIO: "ACESSO_LABORATORIO",
  GESTAO: "GESTAO",
  ADMINISTRADOR: "ADMINISTRADOR",
  TODOS: "TODOS"
};

export function pegarTokenUsuario() {
  return localStorage.getItem('token')
}

export function pegarPermissaoUsuario() {
  const token = pegarTokenUsuario()
  
  if (!token) {
    return null;
  }

  const claims = parseJwt(token);

  return claims ? claims.permissao : null;
}

export function pegarNomeUsuario() {
  const token = pegarTokenUsuario()
  
  if (!token) {
    return null;
  }

  const claims = parseJwt(token);

  return claims ? claims.nome : null;
}

export function pegarUnidadeUsuario() {
  const token = pegarTokenUsuario()

  if (!token) {
    return null;
  }

  const claims = parseJwt(token);

  return claims ? claims.cnes : null;
}

function parseJwt(token) {
  try {
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
  } catch (e) {
    console.error("Erro ao decodificar o token:", e);
    return null;
  }
}
