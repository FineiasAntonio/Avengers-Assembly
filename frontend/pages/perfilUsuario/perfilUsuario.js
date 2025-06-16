import { GetDadosCadastrados } from "../../api/cadastroApi.js";
import { pegarCpfUsuario } from "../../shared/gerenciador-permissoes.js";
import { AlterarInformacaoRequisicao, AlterarSenhaRequisicao } from "../../api/usuarioApi.js";

const alterarInfBtn = document.getElementById("submitAltInf");

alterarInfBtn.onclick = () => {
    const campo = document.getElementById("select").value;
    const novoValor = document.getElementById("inputT").value;
    const cpf = pegarCpfUsuario();
    AlterarInformacaoRequisicao(cpf, campo, novoValor);
};

const alterarInfSen = document.getElementById("submitAltSen");

alterarInfSen.onclick = () => {
    const novaSenha = document.getElementById("novaSenha").value;
    AlterarSenhaRequisicao(novaSenha);
}

document.addEventListener("DOMContentLoaded", () => {
    const campoNomeH = document.getElementById("nomeh");
    const campoNomeI = document.getElementById("nomeI");
    const campoRegistro = document.getElementById("registro");
    const campoCpf = document.getElementById("cpf");
    const campoEmail = document.getElementById("email");
    const campoTelefone = document.getElementById("telefone");

    const campoUSCNES = document.getElementById("USCNES");
    const campoLCNES = document.getElementById("LCNES");

    const cpf = pegarCpfUsuario();
    const response = GetDadosCadastrados(cpf, "paciente");

    campoNomeI.textContent = response.nome;
    campoNomeH.textContent = response.nome;
    campoRegistro.textContent = response.registro;
    campoCpf.textContent = formatarCpf(cpf);
    campoEmail.textContent = response.email;
    campoTelefone.textContent = formatarTelefone(response.telefone);
    
    campoUSCNES.textContent = response.unidadesaudecnes;
    campoLCNES.textContent = response.laboratoriocnes;

    const botao1 = document.getElementById("botao1");
    const botao2 = document.getElementById("botao2");
    const overlay1 = document.getElementById("overlay");
    const overlay2 = document.getElementById("overlay2");
    const fechar1 = document.getElementById("fecharTela");
    const fechar2 = document.getElementById("fecharTela2");

    botao1.addEventListener("click", (e) => {
        e.preventDefault();
        overlay1.style.display = "flex";
    });
    botao2.addEventListener('click', (e) => {
        e.preventDefault();
        overlay2.style.display = 'flex';
    });
    fechar1.addEventListener('click', () => {
        overlay1.style.display = 'none';
    });
    fechar2.addEventListener('click', () => {
        overlay2.style.display = 'none';
    });
});

function formatarCpf(cpf) {
    cpf = cpf.replace(/\D/g, '');

    return cpf.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, "$1.$2.$3-$4");
}

function formatarTelefone(telefone) {
    let size = telefone.length;
    telefone = telefone.replace(/\D/g, '');

    if (size === 10) {
        return telefone.replace(/(\d{2})(\d{4})(\d{4})/, "($1) $2-$3")

    } else if (size === 11) {
        return telefone.replace(/(\d{2})(\d{5})(\d{4})/, "($1) $2-$3")

    } else {
        return telefone;
    }
}