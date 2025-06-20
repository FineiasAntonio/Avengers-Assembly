import { API_ENDERECO } from "../environment/environment.js";
import "../shared/interceptor.js"
import { mostarNotificacao } from "../shared/notificacao.js";

export async function AlterarInformacaoRequisicao(cpf, campo, novo_valor) {
    const url = new URL(API_ENDERECO + "usuario/alterarInf")
    url.searchParams.append('cpf', cpf)

    const dto = { campo, novo_valor }

    const response = await fetch(url, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(dto)
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        throw new Error(errorData.message)
    }
    mostarNotificacao(`${campo} alterado com sucesso!`, "sucess", 3000);
    setTimeout(() => {
        window.location.replace("../perfilUsuario/perfilUsuario.html?cpf=" + cpf);
    }, 1500);
}

export async function AlterarSenhaRequisicao(nova_senha) {
    const url = new URL(API_ENDERECO + "usuario")
    nova_senha = { nova_senha };

    const response = await fetch(url, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(nova_senha)
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        throw new Error(errorData.message)
    }
    mostarNotificacao(`Senha alterada com sucesso!`, "sucess", 3000);
    setTimeout(() => {
        window.location.replace("../perfilUsuario/perfilUsuario.html");
    }, 1500);
}

export async function listarUsuario(registro) {
    const url = new URL(API_ENDERECO + "usuario")
    url.searchParams.append('registro', registro)

    const response = await fetch(url, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' }
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        throw new Error(errorData.message)
    }

    return await response.json()
}
