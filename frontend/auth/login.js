import {loginRequisicao} from "../api/autenticacaoApi.js"

document.addEventListener('DOMContentLoaded', function() {
    
    const botaoLogar = document.getElementById("submit")

    botaoLogar.addEventListener("click", realizarLogin)

    const fundoPreto = document.getElementById("fundoPreto");
    const esqueceuSenha = document.getElementById("esqueceuASenha");
    const esqueceuSenhaLink = document.getElementById("esqueceuSenha");
    
    esqueceuSenhaLink.addEventListener("click", () => {
        fundoPreto.style.display = 'flex';
        esqueceuSenha.style.display = 'flex';
    });

    fundoPreto.addEventListener("click", () => {
        fundoPreto.style.display = 'none';
        esqueceuSenha.style.display = 'none';
    });
})

function realizarLogin() {
    const credencial = document.getElementById("credencial").value
    const senha = document.getElementById("senha").value

    loginRequisicao(credencial, senha)
}
