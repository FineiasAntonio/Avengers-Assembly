import { API_ENDERECO } from "../environment/environment.js";
import "../shared/interceptor.js"

export async function AlterarInformacaoRequisicao(cpf, campo, novoValor) {
    const url = new URL(API_ENDERECO + "usuario/alterarInf")
    url.searchParams.append('cpf', cpf)

    const dto = { campo, novoValor }

    const response = await fetch(url, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(dto)
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        throw new Error(errorData.message)
    }

    return await response.json()
}

export async function AlterarSenhaRequisicao(novaSenha) {
    const url = new URL(API_ENDERECO + "usuario")

    const response = await fetch(url, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(novaSenha)
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        throw new Error(errorData.message)
    }

    return await response.json()
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
