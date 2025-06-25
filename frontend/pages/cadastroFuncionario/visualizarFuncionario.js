import { listarUsuario } from "../../api/usuarioApi.js";
import { listarUnidade } from "../../api/unidadeApi.js";

document.addEventListener("DOMContentLoaded", () => {
    const nome = document.getElementById("nome");
    const cpf = document.getElementById("cpf");
    const email = document.getElementById("email");
    const telefone = document.getElementById("telefone");
    
    const registroA = document.getElementById("registro");
    const permissao = document.getElementById("permissao");
    const cnes = document.getElementById("cnesUnidade");
    const tipoUnidade = document.getElementById("tipoUnidade");

    const nomeUnidade = document.getElementById("nomeUnidade");
    const cnpjUnidade = document.getElementById("cnpjUnidade");
    const telefoneUnidade = document.getElementById("telefoneUnidade");
    const enderecoUnidade = document.getElementById("enderecoUnidade");

    const registro = new URLSearchParams(window.location.search).get("registro");

    listarUsuario(registro).then((resultado) => {
        if (nome) nome.value = resultado.nome || "";
        if (cpf) cpf.value = resultado.cpf || "";
        if (telefone) telefone.value = resultado.telefone || "";
        if (email) email.value = resultado.email|| "";

        if (registroA) registroA.value = resultado.registro || "";
        if (permissao) permissao.value = resultado.permissao || "";

        if (cnes) {
            if (resultado.unidade_saude_cnes) {
                cnes.value = resultado.unidade_saude_cnes || "";
            } else {
                cnes.value = resultado.laboratorio_cnes || "";
            }
        }
        if (cnes) {
            if (resultado.unidade_saude_cnes) {
                tipoUnidade.value = "Unidade de Saúde Básica";
            } else {
                tipoUnidade.value = "Laboratório";
            }
        }

        listarUnidade(resultado.unidade_saude_cnes).then((result) => {
            if (nomeUnidade) nomeUnidade.textContent = result.nome || "";
            if (cnpjUnidade) cnpjUnidade.textContent = result.cnpj || "";
            if (telefoneUnidade) telefoneUnidade.textContent = result.telefone || "";
            if (enderecoUnidade) enderecoUnidade.textContent = result.endereco.bairro || "";
        }) 
    }).catch((erro) => {
        console.error("Erro ao buscar funcionário:", erro);
    });
});
