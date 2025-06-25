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

export async function trocarSenhaRequisicaoEsqueceuSenha(nova_senha, credencial) {
    nova_senha = { nova_senha }
    const url = new URL(API_ENDERECO+"usuario/esqueceuSenha")
    url.searchParams.append('credencial', credencial)

    const response = await fetch(url, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(nova_senha)
    })

    if (!response.ok) {
            const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status}` }));
            notificar(`Erro ao modificar senha!\n`+errorData.message, "error", 3000);
            throw new Error(errorData.message);
    }

    notificar(`Senha modificada com sucesso!\n`, "success", 3000);
    setTimeout(() => {
        window.location.replace("../auth/LoginPagina.html")
    }, 1500)
}

export async function emailRequisicaoRecuperarSenha(credencial) {
    const url = new URL(API_ENDERECO + "codigo/email");
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

export async function codigoRecebidoRequisicao(codigo, credencial) {
    const url = new URL(API_ENDERECO+"codigo");

    const req = {codigo, credencial}

    const response = await fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json'},
        body: JSON.stringify(req)
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status}` }));
        notificar(`Código incorreto e/ou expirado!\n`+errorData.message, "error", 3000);
        return false;
    }
    else {
        notificar(`Validação concluida!\n`, "success", 3000);
        return true;
    }
}