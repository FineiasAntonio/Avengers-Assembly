import { listarPaciente } from "../../api/pacienteApi.js";

window.addEventListener("DOMContentLoaded", () => {
    const nome = document.getElementById("nome");
    const apelido = document.getElementById("apelido");
    const dataNascimento = document.getElementById("dataNascimento");
    const idade = document.getElementById("idade");
    const cpf = document.getElementById("cpf");
    const cartaoSus = document.getElementById("cartaoSus");
    const telefone = document.getElementById("telefone");

    const cep = document.getElementById("cepEndereco");
    const logradouro = document.getElementById("logradouro");
    const numero = document.getElementById("numero");
    const bairro = document.getElementById("bairro");
    const complemento = document.getElementById("complemento");
    const municipio = document.getElementById("municipio");
    const uf = document.getElementById("uf");

    const nomeMae = document.getElementById("nomeMae");
    const nacionalidade = document.getElementById("nacionalidade");
    const raca = document.getElementById("raca_cor");
    const escolaridade = document.getElementById("escolaridade");

    const cartao_sus = new URLSearchParams(window.location.search).get("cartao_sus");

    listarPaciente(cartao_sus).then((resultado) => {
        if (nome) nome.value = resultado.nome || "";
        if (apelido) apelido.value = resultado.apelido || "";
        if (dataNascimento) dataNascimento.value = new Date(resultado.data_nascimento).toLocaleDateString();
        if (idade) idade.value = resultado.idade || "";
        if (cpf) cpf.value = resultado.cpf || "";
        if (cartaoSus) cartaoSus.value = resultado.cartao_sus || "";
        if (telefone) telefone.value = resultado.telefone || "";

        if (cep) cep.value = resultado.endereco?.cep || "";
        if (logradouro) logradouro.value = resultado.endereco?.logradouro || "";
        if (numero) numero.value = resultado.endereco?.numero || "";
        if (bairro) bairro.value = resultado.endereco?.bairro || "";
        if (complemento) complemento.value = resultado.endereco?.complemento || "";
        if (municipio) municipio.value = resultado.endereco?.municipio || "";
        if (uf) uf.value = resultado.endereco?.uf || "";

        if (nomeMae) nomeMae.value = resultado.nome_mae || "";
        if (nacionalidade) nacionalidade.value = resultado.nacionalidade || "";
        if (raca) raca.value = resultado.raca || "";
        if (escolaridade) escolaridade.value = resultado.escolaridade || "";
    }).catch((erro) => {
        console.error("Erro ao buscar paciente:", erro);
    });
});
