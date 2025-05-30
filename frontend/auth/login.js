import {loginRequisicao} from "../api/autenticacaoApi.js"

document.addEventListener('DOMContentLoaded', function() {
    
    const botaoLogar = document.getElementById("submit")

    botaoLogar.addEventListener("click", realizarLogin)

})

function realizarLogin() {
    const credencial = document.getElementById("credencial").value
    const senha = document.getElementById("senha").value

    loginRequisicao(credencial, senha)
}
