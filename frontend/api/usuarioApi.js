import { API_ENDERECO } from "../environment/environment.js";
import "../shared/interceptor.js"
import { notificar } from "../shared/notificacao.js";

export async function AlterarInformacaoRequisicao(cpf, campo, novo_valor) {
    const url = new URL(API_ENDERECO + "usuario/alterarInf")
    url.searchParams.append('cpf', cpf)

    const dto = { campo, novo_valor }

    const response = await fetch(url, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(dto)
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        notificar(`Erro ao alterar ${campo}!\n` + errorData.message, "error", 3000);
        throw new Error(errorData.message)
    }
    notificar(`${campo} alterado com sucesso!`, "success", 3000);
    setTimeout(() => {
        window.location.replace("../perfilUsuario/perfilUsuario.html?cpf=" + cpf);
    }, 1500);
}

export async function AlterarSenhaRequisicao(nova_senha, credencial) {
    const url = new URL(API_ENDERECO + "usuario");
    url.searchParams.append("credencial", credencial);
    nova_senha = { nova_senha };

    const response = await fetch(url, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(nova_senha)
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        notificar(`Erro ao alterar senha: \n`+errorData.message, "error", 3000)
        throw new Error(errorData.message)
    }
    notificar(`Senha alterada com sucesso!`, "success", 3000);
    setTimeout(() => {
        window.location.replace("../perfilUsuario/perfilUsuario.html");
    }, 1500);
}

export async function listarUsuario(registro) {
    const url = new URL(API_ENDERECO + "usuario")
    url.searchParams.append('registro', registro)

    const response = await fetch(url, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' }
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        throw new Error(errorData.message)
    }

    return await response.json()
}

export async function cadastrarFuncionario(funcionario) {
    const url = new URL(API_ENDERECO + "usuario")

    const response = await fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(funcionario)
    })

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status} - ${response.body}` }))
        throw new Error(errorData.message)
    }

    notificar("Funcionário cadastrado com sucesso!", "success")
    setTimeout(() => {
        window.location.replace("../inicioPagina/inicioPagina.html")
    }, 1500)
}

export async function ExisteProfissional(registro) {
    const url = new URL(API_ENDERECO + "usuario");
    url.searchParams.append("registro", registro);
    console.log(registro)

    try {
        const response = await fetch(url, {
            method: "HEAD",
            headers: {"Content-Type": "application/json"},
        });

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({message: `Erro HTTP: ${response.status}`}));
            throw new Error(errorData.message);
        }
        window.location.href = `../cadastroFuncionario/visualizarFuncionario.html?registro=${encodeURIComponent(registro)}`;
    }

    catch (error) {
        window.alert("Esse funcionário não existe!");
        console.log(error);
    }
}