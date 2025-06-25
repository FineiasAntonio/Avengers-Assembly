import { API_ENDERECO } from "../environment/environment.js";
import "../shared/interceptor.js"
import { notificar } from "../shared/notificacao.js";

export async function CadastroRequisicao(objeto) {
    try {
        const response = await fetch(API_ENDERECO + "requisicaoExame", {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(objeto)
        });

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status}` }));
            throw new Error(errorData.message);
        }
        const protocolo = await response.json();
        notificar("Requisição cadastrada com sucesso!", "success", 3000);
        setTimeout(() => {
            window.location.replace("../visualizacaoRequisicao/VisualizacaoRequisicao.html?protocolo=" + protocolo);
        }, 1500);
    }
    catch (error) {
        window.alert("Algo deu errado!")
        console.log(error);
    }
};

export async function ListarRequisicaoExame(protocolo) {
    const url = new URL(API_ENDERECO + "requisicaoExame");
    url.searchParams.append('protocolo', protocolo);

    try {
        const response = await fetch(url, {
            method: "GET",
            headers: {"Content-Type": "application/json"},
        });

        if (!response.ok) {
            const errorData = await response.json.catch(() => ({message: `Erro HTTP: ${response.status}`}));
            throw new Error(errorData.message);
        }
        return await response.json();
    }

    catch (error) {
        window.alert("Algo deu errado!");
        console.log(error);
    }
}

export async function ExisteRequisicaoExame(protocolo) {
    const url = new URL(API_ENDERECO + "requisicaoExame");
    url.searchParams.append('protocolo', protocolo);

    try {
        const response = await fetch(url, {
            method: "HEAD",
            headers: {"Content-Type": "application/json"},
        });

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({message: `Erro HTTP: ${response.status}`}));
            throw new Error(errorData.message);
        }
        window.location.href = `../visualizacaoRequisicao/VisualizacaoRequisicao.html?protocolo=${encodeURIComponent(protocolo)}`;
    }

    catch (error) {
        window.alert("Esse exame não existe!");
        console.log(error);
    }
}

export async function CadastrarResultadoExame(objeto) {
    try {
        const response = await fetch(API_ENDERECO + "requisicaoExame/resultado", {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(objeto)
        });

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status}` }));
            throw new Error(errorData.message);
        }
        
        notificar("Resultado de exame cadastrado com sucesso!", "success", 3000);
        setTimeout(() => {
            window.location.replace("/pages/visualizacaoRequisicao/VisualizacaoRequisicao.html?protocolo=" + objeto.protocolo_exame);
        }, 1500);
        return true;
    }
    catch (error) {
        notificar("Erro ao cadastrar resultado de exame: " + error.message, "error", 5000);
        console.log(error);
        return false;
    }
}

export async function BuscarResultadoExame(protocolo) {
    try {
        const url = new URL(API_ENDERECO + "requisicaoExame/resultado");
        url.searchParams.append('protocolo', protocolo);

        const response = await fetch(url, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' }
        });

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status}` }));
            throw new Error(errorData.message);
        }
        
        return await response.json();
    }
    catch (error) {
        notificar("Erro ao buscar resultado: " + error.message, "error", 5000);
        console.log(error);
        throw error;
    }
}

