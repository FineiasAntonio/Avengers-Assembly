import { listarLaboratorio, listarUnidade } from "../../api/unidadeApi.js";

document.addEventListener("DOMContentLoaded", () => {
    const nome = document.getElementById("nome");
    const tipo = document.getElementById("tipoUnidade");
    const cnpj = document.getElementById("cnpj");
    const telefone = document.getElementById("telefone");

    const cep = document.getElementById("cepEndereco");
    const logradouro = document.getElementById("logradouro");
    const numero = document.getElementById("numero");
    const bairro = document.getElementById("bairro");
    const complemento = document.getElementById("complemento");
    const municipio = document.getElementById("municipio");
    const uf = document.getElementById("uf");

    const cnes = new URLSearchParams(window.location.search).get("cnes");

    listarUnidade(cnes).then((resultado) => {
        console.log(resultado)
        if (nome) nome.value = resultado.nome || "";
        if (cnpj) cnpj.value = resultado.cnpj || "";
        if (tipo) tipo.value = "Unidade Básica de Saúde";
        if (telefone) telefone.value = resultado.telefone || "";

        if (cep) cep.value = resultado.endereco.cep || "";
        if (logradouro) logradouro.value = resultado.endereco.logradouro || "";
        if (numero) numero.value = resultado.endereco.numero || "";
        if (bairro) bairro.value = resultado.endereco.bairro || "";
        if (complemento) complemento.value = resultado.endereco.complemento || "";
        if (municipio) municipio.value = resultado.endereco.municipio || "";
        if (uf) uf.value = resultado.endereco.uf || "";
    }).catch((erro) => {
        listarLaboratorio(cnes).then((resultado) => {
                console.log(resultado)
                if (nome) nome.value = resultado.nome || "";
                if (cnpj) cnpj.value = resultado.cnpj || "";
                if (tipo) tipo.value = "Laboratório";
                if (telefone) telefone.value = resultado.telefone || "";

                if (cep) cep.value = resultado.endereco.cep || "";
                if (logradouro) logradouro.value = resultado.endereco.logradouro || "";
                if (numero) numero.value = resultado.endereco.numero || "";
                if (bairro) bairro.value = resultado.endereco.bairro || "";
                if (complemento) complemento.value = resultado.endereco.complemento || "";
                if (municipio) municipio.value = resultado.endereco.municipio || "";
                if (uf) uf.value = resultado.endereco.uf || "";
            }).catch((erro) => {
                console.log("Erro ao buscar unidade: ", erro)
        })
    });
});
