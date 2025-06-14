import { API_ENDERECO } from "../environment/environment.js";
import "../shared/interceptor.js"

export async function listarPaciente(pacienteChave) {


    const url = new URL(API_ENDERECO + "paciente")
    url.searchParams.append('paciente', pacienteChave)


    const response = await fetch(url, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
    });

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        throw new Error(errorData.message)
    }

    return await response.json()
}
