import { pegarPermissaoUsuario, Permissao } from '../../shared/gerenciador-permissoes.js';
import '../../shared/interceptor.js'

export const items = [
  {
    href: "",
    title: "Buscar Paciente",
    icon: "",
    permissoesRequeridas: [Permissao.TODOS]
  },
  {
    href: "/pages/cadastroPaciente/CadastroPaciente.html",
    title: "Cadastrar Paciente",
    icon: "",
    permissoesRequeridas: [Permissao.ACESSO_ATENDIMENTO, Permissao.GESTAO]
  },
  {
    href: "",
    title: "Ver Agenda Paciente",
    icon: "",
    permissoesRequeridas: [Permissao.ACESSO_ATENDIMENTO, Permissao.ACESSO_EXAMES, Permissao.GESTAO]
  },
  {
    href: "",
    title: "Buscar Exame",
    icon: "",
    permissoesRequeridas: [Permissao.TODOS]
  },
  {
    href: "",
    title: "Criar Exame",
    icon: "",
    permissoesRequeridas: [Permissao.ACESSO_EXAMES]
  },
  {
    href: "",
    title: "Agenda",
    icon: "",
    permissoesRequeridas: [Permissao.ACESSO_EXAMES]
  },
  {
    href: "",
    title: "Cadastrar Profissional",
    icon: "",
    permissoesRequeridas: [Permissao.GESTAO]
  },
  {
    href: "",
    title: "Buscar Profissional",
    icon: "",
    permissoesRequeridas: [Permissao.GESTAO]
  },
  {
    href: "",
    title: "Cadastrar Laboratório",
    icon: "",
    permissoesRequeridas: [Permissao.GESTAO]
  },
  {
    href: "",
    title: "Buscar Laboratório",
    icon: "",
    permissoesRequeridas: [Permissao.GESTAO]
  },
  {
    href: "",
    title: "Cadastrar Unidade de Saúde",
    icon: "",
    permissoesRequeridas: [Permissao.GESTAO]
  },
  {
    href: "",
    title: "Buscar Unidade de Saúde",
    icon: "",
    permissoesRequeridas: [Permissao.GESTAO]
  },
  {
    href: "",
    title: "Dashboard",
    icon: "",
    permissoesRequeridas: [Permissao.GESTAO]
  }
];

function renderMenu() {
  
  const content = document.querySelector('.content');
  content.innerHTML = '';

  const permissao = pegarPermissaoUsuario();

  items.forEach(item => {
    if (item.permissoesRequeridas.includes(permissao) || item.permissoesRequeridas.includes(Permissao.TODOS) || permissao === Permissao.ADMINISTRADOR) {
      const itemElement = document.createElement('div');
      itemElement.addEventListener('click', () => {
        window.location.replace(item.href);
      });
      itemElement.className = 'item';
      itemElement.innerHTML = `
                <img src="./assets/icons/${item.icon}" alt="${item.title}">
                <h2>${item.title}</h2>
            `;
      content.appendChild(itemElement);
    }
  });
}

document.addEventListener('DOMContentLoaded', renderMenu);