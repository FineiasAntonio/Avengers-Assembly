import { API_ENDERECO } from "../environment/environment.js";
import "../shared/interceptor.js"
import { notificar } from "../shared/notificacao.js"

export async function listarUnidade(cnes) {
    const url = new URL(API_ENDERECO + "unidade")
    url.searchParams.append('cnes', cnes)


    const response = await fetch(url, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
    })

    if (!response.ok) {
        console.log(response.body);
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        throw new Error(errorData.message)
    }

    return await response.json()
}

export async function listarLaboratorio(cnes) {
    const url = new URL(API_ENDERECO + "laboratorio")
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

export async function cadastrarUnidade(unidade, endpoint) {
    const url = new URL(API_ENDERECO + endpoint)

    const response = await fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(unidade),
    })   

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        throw new Error(errorData.message)
    }

    notificar("Unidade cadastrada com sucesso", "success")
    setTimeout(() => {
        window.location.replace("../inicioPagina/inicioPagina.html")
    }, 1500)
}


export async function ExisteUnidade(cnes, endpoint) {
    const url = new URL(API_ENDERECO + endpoint);
    url.searchParams.append("cnes", cnes);

    try {
        const response = await fetch(url, {
            method: "HEAD",
            headers: {"Content-Type": "application/json"},
        });

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({message: `Erro HTTP: ${response.status}`}));
            throw new Error(errorData.message);
        }
        window.location.href = `../cadastroUnidade/visualizarUnidade.html?cnes=${encodeURIComponent(cnes)}`;
    }

    catch (error) {
        window.alert("Essa unidade não existe!");
        console.log(error);
    }
}