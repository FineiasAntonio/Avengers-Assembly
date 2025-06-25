import {loginRequisicao, emailRequisicaoRecuperarSenha, codigoRecebidoRequisicao, trocarSenhaRequisicaoEsqueceuSenha} from "../api/autenticacaoApi.js"
import { notificar } from "../shared/notificacao.js";

document.addEventListener('DOMContentLoaded', function() {
    let credencialEmail = null;
    const botaoLogar = document.getElementById("submit")

    botaoLogar.addEventListener("click", realizarLogin)

    const fundoPreto = document.getElementById("fundoPreto");
    const esqueceuSenha = document.getElementById("esqueceuASenha");
    const esqueceuSenhaLink = document.getElementById("esqueceuSenha");
    
    esqueceuSenhaLink.addEventListener("click", () => {
        fundoPreto.style.display = 'flex';
        esqueceuSenha.style.display = 'flex';
    });

    fundoPreto.addEventListener("click", (e) => {
        if (e.target.id == "fundoPreto") {
            fundoPreto.style.display = 'none';
            esqueceuSenha.style.display = 'none';
        }
    });

    const enviarEmailBtn = document.getElementById("enviarEsqueceuSenha");
    enviarEmailBtn.addEventListener("click", async () => {
        credencialEmail = document.getElementById("entradaCredencial").value;
        
        if (credencialEmail == "") {
            notificar("Credencial inválida!", "warning", 3000);
            return;
        }

        await emailRequisicaoRecuperarSenha(credencialEmail);
    });

    const confirmarCodigoBtn = document.getElementById("confirmarCodigo");

    if (confirmarCodigoBtn) {
        confirmarCodigoBtn.addEventListener("click", async () => {
        const codigo = document.getElementById("entradaCodigo").value;

        if (codigo == "") {
            notificar("Codigo em formato inválido!", "warning", 3000);
            return;
        }
        credencialEmail = document.getElementById("entradaCredencial").value;
        
        const ok = await codigoRecebidoRequisicao(codigo, credencialEmail);

        if (ok) {
            esqueceuSenha.style.display = 'none';

            injetarDivTrocaSenha();

            const trocarSenha = document.getElementById("mudarSenha");
            trocarSenha.addEventListener("click", async () => {
                const novaSenhaValor = document.getElementById("novaSenha").value;

                if (novaSenhaValor == "") {
                    notificar("Insira alguma senha!", "warning", 3000);
                    return;
                }

                await trocarSenhaRequisicaoEsqueceuSenha(novaSenhaValor, credencialEmail);
            })
        }
    });
    }
})

function realizarLogin() {
    const credencial = document.getElementById("credencial").value
    const senha = document.getElementById("senha").value

    loginRequisicao(credencial, senha)
}

function injetarDivTrocaSenha() {
    const div = document.createElement("div");
    div.id = "trocaSenhaDiv";
    div.innerHTML = `
        <h3>Trocar senha</h3>
        <input type="password" id="novaSenha" placeholder="Nova senha" class="entrada"/>
        <button id="mudarSenha" class="bt">Mudar senha</button>
    `;

    const fundoPreto = document.getElementById("fundoPreto");
    if (fundoPreto) {
        fundoPreto.appendChild(div)
    }
}