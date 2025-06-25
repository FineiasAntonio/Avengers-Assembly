import { API_ENDERECO } from "../environment/environment.js";
import { notificar } from "../shared/notificacao.js";
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

export async function cadastrarPaciente(paciente) {
    const response = await fetch(API_ENDERECO+"paciente", {
        method: "POST",
        headers: { "Content-Type": "application;json"},
        body: JSON.stringify(paciente)
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        notificar(`Erro ao cadastrar paciente\n`+errorData.message, "error", 3000);
        throw new Error(errorData.message)
    }

    notificar(`Paciente cadastrado com sucesso!\n`, "success", 3000);
}


export async function ExistePaciente(cartao_sus) {
    const url = new URL(API_ENDERECO + "paciente");
    url.searchParams.append("cartao_sus", cartao_sus);

    try {
        const response = await fetch(url, {
            method: "HEAD",
            headers: {"Content-Type": "application/json"},
        });

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({message: `Erro HTTP: ${response.status}`}));
            throw new Error(errorData.message);
        }
        window.location.href = `../cadastroPaciente/visualizacaoPaciente.html?cartao_sus=${encodeURIComponent(cartao_sus)}`;
    }

    catch (error) {
        window.alert("Esse paciente não existe!");
        console.log(error);
    }
}