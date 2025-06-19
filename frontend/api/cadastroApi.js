import { API_ENDERECO } from "../environment/environment.js";
import "../shared/interceptor.js"
import { mostarNotificacao } from "../shared/notificacao.js";

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
        mostarNotificacao("Requisição cadastrada com sucesso!", "success", 3000);
        setTimeout(() => {
            window.location.replace("/pages/visualizacaoRequisicao/VisualizacaoRequisicao.html?protocolo=" + protocolo);
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