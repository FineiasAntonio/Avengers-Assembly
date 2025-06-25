import { pegarPermissaoUsuario, Permissao } from '../../shared/gerenciador-permissoes.js';
import '../../shared/interceptor.js'
import { invocarModal } from '../../shared/modal.js';
import { ExisteRequisicaoExame } from '../../api/cadastroApi.js';
import { ExistePaciente } from '../../api/pacienteApi.js';
import { ExisteProfissional } from '../../api/usuarioApi.js';
import { ExisteUnidade } from '../../api/unidadeApi.js';

export const items = [
  {
    href: "",
    title: "Buscar Paciente",
    icon: "buscarPac.png",
    event: buscarPaciente,
    permissoesRequeridas: [Permissao.TODOS]
  },
  {
    href: "../cadastroPaciente/CadastroPaciente.html",
    title: "Cadastrar Paciente",
    icon: "cadastroPac.png",
    permissoesRequeridas: [Permissao.ACESSO_ATENDIMENTO, Permissao.GESTAO]
  },
  {
    href: "../agendamentoExame/agendamentoExame.html",
    title: "Agendar Exame",
    icon: "agendarExm.png",
    permissoesRequeridas: [Permissao.ACESSO_ATENDIMENTO, Permissao.ACESSO_EXAMES, Permissao.GESTAO]
  },
  {
    href: "",
    title: "Buscar Exame",
    icon: "exameBus.png",
    event: buscarExame,
    permissoesRequeridas: [Permissao.TODOS]
  },
  {
    href: "../visualizacaoRequisicao/CadastroRequisicao.html",
    title: "Criar Exame",
    icon: "exameCad.png",
    permissoesRequeridas: [Permissao.ACESSO_EXAMES]
  },
  {
    href: "../cadastroResultadoExame/CadastroResultado.html",
    title: "Cadastrar Resultado Exame",
    icon: "resultado.png",
    permissoesRequeridas: [Permissao.ACESSO_EXAMES]
  },
  {
    href: "../agendamentoExame/agenda.html",
    title: "Agenda",
    icon: "agenda.png",
    permissoesRequeridas: [Permissao.ACESSO_EXAMES]
  },
  {
    href: "../cadastroFuncionario/CadastroFuncionario.html",
    title: "Cadastrar Profissional",
    icon: "profissionalCad.png",
    permissoesRequeridas: [Permissao.GESTAO]
  },
  {
    href: "",
    title: "Buscar Profissional",
    icon: "profissional.png",
    event: buscarProfissional,
    permissoesRequeridas: [Permissao.GESTAO]
  },
  {
    href: "../cadastroUnidade/CadastroUnidade.html",
    title: "Cadastrar Unidade",
    icon: "unidadedesaudeCad.png",
    permissoesRequeridas: [Permissao.GESTAO]
  },
  {
    href: "",
    title: "Buscar Unidade",
    icon: "unidadedesaude.png",
    event: buscarUnidade,
    permissoesRequeridas: [Permissao.GESTAO]
  },
  {
    href: "../centralAnalisePagina/centralAnalisePagina.html",
    title: "Dashboard",
    icon: "dashboard.png",
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
                <img src="../../assets/icons/${item.icon}" alt="${item.title}">
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

function buscarPaciente() {
  const modal = document.getElementById("modal-container-paciente");
  invocarModal(modal);
  document.getElementById("cartaoPaciente").value = ""

  document.getElementById("submitPaciente").addEventListener("click", function(event) {
    const cartaoSus = document.getElementById("cartaoPaciente").value.trim();
    
    if (cartaoSus) {
      ExistePaciente(cartaoSus)
    } else {
      alert("Por favor, insira um cartão SUS válido.");
    }
  })
}

function buscarProfissional() {
  const modal = document.getElementById("modal-container-profissional");
  invocarModal(modal);
  document.getElementById("registroProfissional").value = ""

  document.getElementById("submitProfissional").addEventListener("click", function(event) {
    const registro = document.getElementById("registroProfissional").value.trim();
    if (registro) {
      ExisteProfissional(registro)
    } else {
      alert("Por favor, insira um registro válido.");
    }
  })
}

function buscarUnidade() {
  const modal = document.getElementById("modal-container-unidade");
  invocarModal(modal);
  document.getElementById("cnesUnidade").value = ""

  document.getElementById("submitUnidade").addEventListener("click", async function(event) {
    const cnes = document.getElementById("cnesUnidade").value.trim();
    const tipo = document.getElementById("tipoUnidade").value;
    console.log(tipo)
    if (cnes) {
      if (tipo == "laboratorio") {
        ExisteUnidade(cnes, "laboratorio")

      } else {
        ExisteUnidade(cnes, "unidade")
      }
    } else {
      alert("Por favor, insira um CNES válido.");
    }
  })
}