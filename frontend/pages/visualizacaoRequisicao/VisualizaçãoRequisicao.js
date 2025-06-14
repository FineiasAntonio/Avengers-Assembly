import { pegarUnidadeUsuario } from "../../shared/gerenciador-permissoes.js";
import { listarUnidade } from "../../api/unidadeApi.js";
import { listarPaciente } from "../../api/pacienteApi.js";

buscaPaciente.addEventListener('blur', (event) => {

    if (!event.target.value) {
        return;
    }

    listarPaciente(event.target.value).then(resultado => {
        cartaoSus.value = resultado.cartao_sus || '';
        nomePaciente.value = resultado.nome || '';
        idadePaciente.value = resultado.idade || '';
        dataNascimento.value = resultado.data_nascimento || '';
        nomeMae.value = resultado.nome_mae || '';
        enderecoPaciente.value = `${resultado.endereco.logradouro}, ${resultado.endereco.numero}, ${resultado.endereco.complemento}, ${resultado.endereco.bairro}` || '';
        telefonePaciente.value = resultado.telefone || '';
        nacionalidadePaciente.value = resultado.nacionalidade || '';
        racaCorPaciente.value = resultado.raca || '';
        escolaridadePaciente.value = resultado.escolaridade || '';
    }).catch(error => {
        console.error("Erro ao buscar paciente:", error);
    });
});

window.addEventListener('DOMContentLoaded', () => {

    const unidadeCnes = document.getElementById("unidadeCnes")
    const unidadeNome = document.getElementById("unidadeNome")
    const unidadeEndereco = document.getElementById("unidadeEndereco")

    const buscaPaciente = document.getElementById("buscaPaciente");

    const cpfPaciente = document.getElementById("cpfPaciente");
    const cartaoSus = document.getElementById("cartaoSus");
    const nomePaciente = document.getElementById("nomePaciente");
    const idadePaciente = document.getElementById("idadePaciente");
    const dataNascimento = document.getElementById("dataNascimento");
    const nomeMae = document.getElementById("nomeMae");
    const enderecoPaciente = document.getElementById("enderecoPaciente");
    const telefonePaciente = document.getElementById("telefonePaciente");
    const nacionalidadePaciente = document.getElementById("nacionalidadePaciente");
    const racaCorPaciente = document.getElementById("racaCorPaciente");
    const escolaridadePaciente = document.getElementById("escolaridadePaciente");

    buscaPaciente.addEventListener('blur', (event) => {
        if (!event.target.value) {
            return;
        }

        listarPaciente(event.target.value).then(resultado => {
            cpfPaciente.textContent += `${resultado.cpf.slice(0, 3)}.${resultado.cpf.slice(3, 6)}.${resultado.cpf.slice(6, 9)}-${resultado.cpf.slice(9, 11)}` || '';
            cartaoSus.textContent += resultado.cartao_sus || '';
            nomePaciente.textContent += resultado.nome || '';
            idadePaciente.textContent += `${resultado.idade} anos` || '';
            dataNascimento.textContent += new Date(resultado.data_nascimento).toLocaleDateString('pt-BR') || '';
            nomeMae.textContent += resultado.nome_mae || '';
            enderecoPaciente.textContent += `${resultado.endereco.logradouro}, ${resultado.endereco.numero}, ${resultado.endereco.complemento}, ${resultado.endereco.bairro}, ${resultado.endereco.municipio} - ${resultado.endereco.uf}, ${resultado.endereco.cep}` || '';
            telefonePaciente.textContent += `(${resultado.telefone.slice(0, 2)}) ${resultado.telefone.slice(2, 7)}-${resultado.telefone.slice(7, 11)}` || '';
            nacionalidadePaciente.textContent += resultado.nacionalidade || '';
            racaCorPaciente.textContent += resultado.raca || '';
            escolaridadePaciente.textContent += resultado.escolaridade || '';
        }).catch(error => {
            console.error("Erro ao buscar paciente:", error);
        });
    });

    listarUnidade(pegarUnidadeUsuario()).then(resultado => {
        unidadeCnes.textContent = resultado.cnes
        unidadeNome.textContent = resultado.nome
        unidadeEndereco.textContent = `${resultado.endereco.logradouro}, ${resultado.endereco.numero} - ${resultado.endereco.bairro}, ${resultado.endereco.municipio} - ${resultado.endereco.uf}, ${resultado.endereco.cep}`
    })
});