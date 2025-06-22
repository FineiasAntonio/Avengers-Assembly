import { pegarCpfUsuario } from "../../shared/gerenciador-permissoes.js";
import { mostarNotificacao } from "../../shared/notificacao.js";
import { AlterarInformacaoRequisicao, AlterarSenhaRequisicao, listarUsuario } from "../../api/usuarioApi.js";

document.addEventListener("DOMContentLoaded", async () => {
    const alterarInfBtn = document.getElementById("submitAltInf");

    alterarInfBtn.onclick = async () => {
        const campo = document.getElementById("select").value.toLowerCase();
        const novoValor = document.getElementById("inputT").value;

        if (campo == 'telefone' && novoValor.length != 11) {
            mostarNotificacao(`${campo} invalido!\n Siga o padrão (xx) xxxxx-xxxx`, "error", 3000);
            return
        }
        if (campo == 'email' && !validarEmail(novoValor)) {
            mostarNotificacao(`${campo} invalido!\n`, "error", 3000);
            return
        }

        const cpf = pegarCpfUsuario();
        await AlterarInformacaoRequisicao(cpf, campo, novoValor);
    };

    const alterarInfSen = document.getElementById("submitAltSen");

    alterarInfSen.onclick = async () => {
        const novaSenha = document.getElementById("novaSenha").value;
        await AlterarSenhaRequisicao(novaSenha);
    }

    const campoNomeH = document.getElementById("nomeh");
    const campoNomeI = document.getElementById("nomeI");
    const campoRegistro = document.getElementById("registro");
    const campoCpf = document.getElementById("cpf");
    const campoEmail = document.getElementById("email");
    const campoTelefone = document.getElementById("telefone");

    const campoUSCNES = document.getElementById("USCNES");
    const campoLCNES = document.getElementById("LCNES");

    const cpf = pegarCpfUsuario();
    const response = await listarUsuario(cpf, "paciente");

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
    if (!cpf) return "";
    cpf = cpf.replace(/\D/g, '');

    return cpf.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, "$1.$2.$3-$4");
}

function formatarTelefone(telefone) {
    if (!telefone) return "";

    telefone = telefone.replace(/\D/g, '');

    if (telefone.length === 10) {
        return telefone.replace(/(\d{2})(\d{4})(\d{4})/, "($1) $2-$3");
    } else if (telefone.length === 11) {
        return telefone.replace(/(\d{2})(\d{5})(\d{4})/, "($1) $2-$3");
    } else {
        return telefone;
    }
}

function validarEmail(email) {
    if (!email) return false;

    email = email.trim();

    const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    return regex.test(email);
}