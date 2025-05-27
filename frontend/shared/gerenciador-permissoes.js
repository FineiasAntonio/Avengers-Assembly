export const Permissao = {
  ACESSO_ATENDIMENTO: "ACESSO_ATENDIMENTO",
  ACESSO_EXAMES: "ACESSO_EXAMES",
  ACESSO_LABORATORIO: "ACESSO_LABORATORIO",
  GESTAO: "GESTAO",
  ADMINISTRADOR: "ADMINISTRADOR",
  TODOS: "TODOS"
};

export function pegarPermissaoUsuario() {
  const permissao = localStorage.getItem('permissao');
  
  return Object.values(Permissao).includes(permissao) 
    ? permissao 
    : null;
}
