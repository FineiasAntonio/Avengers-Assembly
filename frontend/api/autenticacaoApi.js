import { API_ENDERECO } from "../environment/environment.js";
import { mostarNotificacao } from "../shared/notificacao.js";

export async function loginRequisicao(credencial, senha) {
    const credenciais = { credencial, senha }

    const response = await fetch(API_ENDERECO + "auth/login", {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(credenciais)
    })

    if (!response.ok) {
            const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status}` }));
            mostarNotificacao(`Erro ao autenticar\n`+errorData.message, "error", 3000);
            throw new Error(errorData.message);
    }

    const token = await response.json()

    localStorage.setItem("token", token)
    window.location.replace("../pages/inicioPagina/inicioPagina.html")
}