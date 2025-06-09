import { API_ENDERECO } from "../environment/environment.js";
import "../shared/interceptor.js"

export async function CadastroRequisição(objeto, endPoint) {
    try {
        const response = await fetch(API_ENDERECO + endPoint, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(objeto)
        });

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status}` }));
            throw new Error(errorData.message);
        }
    }
    catch (error) {
        window.alert("Algo deu errado!")
        console.log(error);
    }
};

export async function GetDadosCadastrados(parametro, endPoint) {
    const url = new URL(API_ENDERECO+endPoint);
    url.searchParams.append('parametro', parametro);

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