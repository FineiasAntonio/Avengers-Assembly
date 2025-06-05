import { API_ENDERECO } from "../environment/environment.js";
import "../shared/interceptor.js"

export async function CadastroRequisição(objeto, endPoint) {
    try {
        const response = await fetch(API_ENDERECO + endPoint, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(objeto)
        });

        /* Daria pra retornar a message do response pra controlar melhor o erro
        de cada pagina que utiliza da função?*/
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