import { API_ENDERECO } from "../environment/environment.js";
import "../shared/interceptor.js"

export async function consultarHorarioOcupado(cnes, data) {

    const url = new URL(`${API_ENDERECO}agendamento`);
    url.searchParams.append('cnes', cnes);
    url.searchParams.append('data', data);


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
