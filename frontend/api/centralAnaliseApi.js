import { API_ENDERECO } from "../environment/environment.js";
import "../shared/interceptor.js";
import { mostarNotificacao } from "../shared/notificacao.js";

export async function pegarDadosQtdPacientes(funcao) {
    const url = new URL(API_ENDERECO + 'graficos');
    url.searchParams.append('funcao', funcao);
    
    const response = await fetch(url, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.message} - ${response.body}` }));
        mostarNotificacao(errorData.message, "error", 3000)
        throw new Error(errorData.message);
    }

    return await response.json();
}

export async function pegarDadosQuantidadePacientesPorRegiao() {
    const url = new URL(API_ENDERECO + 'mapa');

    const response = await fetch(url, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' }
    });

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }));
        mostarNotificacao(errorData.message, "error", 3000)
        throw new Error(errorData.message);
    }
    return await response.json();
}