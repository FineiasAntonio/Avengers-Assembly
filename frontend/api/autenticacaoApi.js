import { API_ENDERECO } from "../environment/environment.js";
import { notificar } from "../shared/notificacao.js";

export async function loginRequisicao(credencial, senha) {
    const credenciais = { credencial, senha }

    const response = await fetch(API_ENDERECO + "auth/login", {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(credenciais)
    })

    if (!response.ok) {
            const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status}` }));
            notificar(`Erro ao autenticar\n`+errorData.message, "error", 3000);
            throw new Error(errorData.message);
    }

    const token = await response.json()

    localStorage.setItem("token", token)
    window.location.replace("../pages/inicioPagina/inicioPagina.html")
}

export async function trocarSenhaRequisicaoEsqueceuSenha(novaSenha, credencial) {
    const novaSenhaF = { novaSenha}
    const url = new URL(API_ENDERECO+"usuario")
    url.searchParams.append('credencial', credencial)

    const response = await fetch(API_ENDERECO + "auth/login", {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(novaSenhaF)
    })

    if (!response.ok) {
            const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status}` }));
            notificar(`Erro ao modificar senha!\n`+errorData.message, "error", 3000);
            throw new Error(errorData.message);
    }

    notificar(`Senha modificada com sucesso!\n`, "success", 3000);
}

export async function emailRequisicaoRecuperarSenha(credencial) {
    const url = new URL(API_ENDERECO + "usuario/email");
    url.searchParams.append('credencial', credencial);

    const response = await fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json'},
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status}` }));
        notificar(`Erro ao enviar email\n`+errorData.message, "error", 3000);
        throw new Error(errorData.message);
    }

    notificar(`Email enviado com sucesso!\n`, "success", 3000);
}

export async function codigoRecebidoRequisicao(codigo) {
    const url = new URL(API_ENDERECO+"usuario/codigo");
    url.searchParams.append('codigo', codigo);

    const response = await fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json'},
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status}` }));
        notificar(`Código incorreto!\n`+errorData.message, "error", 3000);
        throw new Error(errorData.message);
    }

    notificar(`Validação concluida!\n`, "success", 3000);
    return true;
}