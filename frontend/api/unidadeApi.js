import { API_ENDERECO } from "../environment/environment.js";
import "../shared/interceptor.js"

export async function listarUnidade(cnes) {

    const url = new URL(API_ENDERECO + "unidade")
    url.searchParams.append('cnes', cnes)


    const response = await fetch(url, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        throw new Error(errorData.message)
    }

    return await response.json()
}