import { pegarPermissaoUsuario, Permissao } from '../../shared/gerenciador-permissoes.js';
import '../../shared/interceptor.js'
import { invocarModal } from '../../shared/modal.js';
import { ExisteRequisicaoExame } from '../../api/cadastroApi.js';

export const items = [
  {
    href: "",
    title: "Buscar Paciente",
    icon: "",
    permissoesRequeridas: [Permissao.TODOS]
  },
  {
    href: "../cadastroPaciente/CadastroPaciente.html",
    title: "Cadastrar Paciente",
    icon: "",
    permissoesRequeridas: [Permissao.ACESSO_ATENDIMENTO, Permissao.GESTAO]
  },
  {
    href: "../agendamentoExame/agendamentoExame.html",
    title: "Agendar Exame",
    icon: "",
    permissoesRequeridas: [Permissao.ACESSO_ATENDIMENTO, Permissao.ACESSO_EXAMES, Permissao.GESTAO]
  },
  {
    href: "",
    title: "Buscar Exame",
    icon: "",
    event: buscarExame,
    permissoesRequeridas: [Permissao.TODOS]
  },
  {
    href: "../visualizacaoRequisicao/CadastroRequisicao.html",
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
    href: "../centralAnalisePagina/centralAnalisePagina.html",
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
      itemElement.className = 'item';
      itemElement.innerHTML = `
                <img src="./assets/icons/${item.icon}" alt="${item.title}">
                <h2>${item.title}</h2>
            `;

      if (item.event) {
        itemElement.addEventListener('click', item.event);
      } else if (item.href) {
        itemElement.addEventListener('click', () => {
          window.location.replace(item.href);
        });
      }
      
      content.appendChild(itemElement);
    }
  });
}

document.addEventListener('DOMContentLoaded', renderMenu);

function buscarExame() {
  const modal = document.getElementById("modal-container-exame");
  invocarModal(modal);
  document.getElementById("protocoloExame").value = ""

  document.getElementById("submit").addEventListener("click", function(event) {
    const protocolo = document.getElementById("protocoloExame").value.trim();
    
    if (protocolo) {
      ExisteRequisicaoExame(protocolo)
    } else {
      alert("Por favor, insira um protocolo de exame válido.");
    }
  })
}