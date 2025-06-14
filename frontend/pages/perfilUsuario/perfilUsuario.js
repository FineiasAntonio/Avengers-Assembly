import { pegarCpfUsuario } from "../../shared/gerenciador-permissoes"
import { GetDadosCadastrados } from "../../api/cadastroApi"

window.addEventListener("DOMContentLoaded", () => {
    const campoNome = document.getElementById("nome");
    const campoRegistro = document.getElementById("registro");
    const campoCpf = document.getElementById("cpf");
    const campoEmail = document.getElementById("email");
    const campoTelefone = document.getElementById("telefone");

    const campoUSCNES = document.getElementById("USCNES");
    const campoLCNES = document.getElementById("LCNES");

    const cpf = pegarCpfUsuario();
    const response = GetDadosCadastrados(cpf, "paciente");

    campoNome.textContent = response.nome;
    campoRegistro.textContent = response.registro;
    campoCpf.textContent = formatarCpf(cpf);
    campoEmail.textContent = response.email;
    campoTelefone.textContent = formatarTelefone(response.telefone);
    
    campoUSCNES.textContent = response.unidadesaudecnes;
    campoLCNES.textContent = response.laboratoriocnes;

    const botao1 = document.getElementById("alterarInformacoes");
    const botao2 = document.getElementById("mudarSenha");
    const overlay1 = document.getElementById("overlay");
    const overlay2 = document.getElementById("overlay2");
    const fechar1 = document.getElementById("fecharTela");
    const fechar2 = document.getElementById("fecharTela2");

    botao1.addEventListener('click', (e) => {
        e.preventDefault();
        overlay1.style.display = 'flex';
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
        return telefone.replace(/\(d{2})(\d{5})(\d{4})/, "($1) $2-$3")

    } else {
        return telefone;
    }
}