import { pegarUnidadeUsuario } from "../../shared/gerenciador-permissoes.js";
import { listarUnidade } from "../../api/unidadeApi.js";
import { listarPaciente } from "../../api/pacienteApi.js";
import { listarUsuario } from "../../api/usuarioApi.js";
import { CadastroRequisicao } from "../../api/cadastroApi.js";
import { mostarNotificacao } from "../../shared/notificacao.js";

window.addEventListener("DOMContentLoaded", () => {
    const unidadeCnes = document.getElementById("unidadeCnes");
    const unidadeNome = document.getElementById("unidadeNome");
    const unidadeEndereco = document.getElementById("unidadeEndereco");

    const buscaPaciente = document.getElementById("buscaPaciente");
    const submitRequisicao = document.getElementById("submitRequisicao");

    const cpfPaciente = document.getElementById("cpfPaciente");
    const cartaoSus = document.getElementById("cartaoSus");
    const nomePaciente = document.getElementById("nomePaciente");
    const idadePaciente = document.getElementById("idadePaciente");
    const dataNascimento = document.getElementById("dataNascimento");
    const nomeMae = document.getElementById("nomeMae");
    const enderecoPaciente = document.getElementById("enderecoPaciente");
    const telefonePaciente = document.getElementById("telefonePaciente");
    const nacionalidadePaciente = document.getElementById(
        "nacionalidadePaciente"
    );
    const racaCorPaciente = document.getElementById("racaCorPaciente");
    const escolaridadePaciente = document.getElementById("escolaridadePaciente");

    const motivoExame = document.getElementById("motivoExame");
    const fezExame = document.getElementById("fezExame");
    const anoUltimoExame = document.getElementById("anoUltimoExame");
    anoUltimoExame.addEventListener("input", (event) => {
        if (event.target.value.length > 4) {
            event.target.value = event.target.value.slice(0, 4);
        }
    })
    const usaDIU = document.getElementById("usaDIU");
    const estaGravida = document.getElementById("estaGravida");
    const usaAnticoncepcional = document.getElementById("usaAnticoncepcional");
    const usaHormonio = document.getElementById("usaHormonio");
    const fezRadioterapia = document.getElementById("fezRadioterapia");
    const teveSangramentoRelacoes = document.getElementById(
        "teveSangramentoRelacoes"
    );
    const teveSangramentoMenopausa = document.getElementById(
        "teveSangramentoMenopausa"
    );
    const ultimaMenstruacao = document.getElementById("ultimaMenstruação");

    const inspecaoColo = document.getElementById("inspecaoColo");
    const sinaisDST = document.getElementById("sinaisDST");
    const dataColeta = document.getElementById("dataColeta");
    const registroResponsavel = document.getElementById("registroResponsavel");
    const nomeProfissional = document.getElementById("nomeProfissional");

    buscaPaciente.addEventListener("blur", (event) => {
        if (!event.target.value) {
            return;
        }

        listarPaciente(event.target.value)
            .then((resultado) => {
                cpfPaciente.textContent +=
                    `${resultado.cpf.slice(0, 3)}.${resultado.cpf.slice(
                        3,
                        6
                    )}.${resultado.cpf.slice(6, 9)}-${resultado.cpf.slice(9, 11)}` || "";
                cartaoSus.textContent += resultado.cartao_sus || "";
                nomePaciente.textContent += resultado.nome || "";
                idadePaciente.textContent += `${resultado.idade} anos` || "";
                dataNascimento.textContent +=
                    new Date(resultado.data_nascimento).toLocaleDateString("pt-BR") || "";
                nomeMae.textContent += resultado.nome_mae || "";
                enderecoPaciente.textContent +=
                    `${resultado.endereco.logradouro}, ${resultado.endereco.numero}, ${resultado.endereco.complemento}, ${resultado.endereco.bairro}, ${resultado.endereco.municipio} - ${resultado.endereco.uf}, ${resultado.endereco.cep}` ||
                    "";
                telefonePaciente.textContent +=
                    `(${resultado.telefone.slice(0, 2)}) ${resultado.telefone.slice(
                        2,
                        7
                    )}-${resultado.telefone.slice(7, 11)}` || "";
                nacionalidadePaciente.textContent += resultado.nacionalidade || "";
                racaCorPaciente.textContent += resultado.raca || "";
                escolaridadePaciente.textContent += resultado.escolaridade || "";

                buscaPaciente.disabled = true;
            })
            .catch((error) => {
                console.error("Erro ao buscar paciente:", error);
            });
    });

    submitRequisicao.addEventListener("click", () => {

        if (buscaPaciente.value.trim() === "") {
            mostarNotificacao("Por favor, informe o CPF ou Cartão SUS do paciente.", "warning", 3000);
            return;
        }
        if (registroResponsavel.value.trim() === "") {
            mostarNotificacao("Por favor, informe o registro do profissional responsável.", "warning", 3000);
            return;
        }

        const requisicao = {
            paciente_id: cartaoSus.textContent.split(":")[1].trim(),
            motivo_exame: motivoExame.value,
            fez_exame_preventivo: fezExame.checked,
            ano_ultimo_exame: anoUltimoExame.value,
            usa_diu: usaDIU.value,
            esta_gravida: estaGravida.value,
            usa_anticoncepcional: usaAnticoncepcional.value,
            usa_hormonio_menopausa: usaHormonio.value,
            fez_radioterapia: fezRadioterapia.value,
            data_ultima_menstruacao: new Date(ultimaMenstruacao.value).toISOString(),
            sangramento_apos_relacoes: teveSangramentoRelacoes.value,
            sangramento_apos_menopausa: teveSangramentoMenopausa.value,
            inspecao_colo: inspecaoColo.value,
            sinais_dst: sinaisDST.checked,
            data_coleta: new Date(dataColeta.value).toISOString(),
            responsavel_registro: registroResponsavel.value,
        };

        enviarRequisicaoExame(requisicao);
    });

    registroResponsavel.addEventListener("blur", (event) => {
        const registro = event.target.value.trim();

        if (registro.length > 0) {
            listarUsuario(registro)
                .then((profissional) => {
                    if (profissional) {
                        nomeProfissional.textContent = `Profissional: ${profissional.nome}`;
                        nomeProfissional.style.color = "#2c3e50";
                    } else {
                        nomeProfissional.textContent = "Profissional não encontrado";
                        nomeProfissional.style.color = "#e74c3c";
                    }
                })
                .catch((error) => {
                    console.error("Erro ao buscar profissional:", error);
                    nomeProfissional.textContent = "Erro ao buscar profissional";
                    nomeProfissional.style.color = "#e74c3c";
                });
        } else {
            nomeProfissional.textContent = "";
        }
    });

    listarUnidade(pegarUnidadeUsuario()).then((resultado) => {
        unidadeCnes.textContent = resultado.cnes;
        unidadeNome.textContent = resultado.nome;
        unidadeEndereco.textContent = `${resultado.endereco.logradouro}, ${resultado.endereco.numero} - ${resultado.endereco.bairro}, ${resultado.endereco.municipio} - ${resultado.endereco.uf}, ${resultado.endereco.cep}`;
    });

});

function enviarRequisicaoExame(requisicao) {
    CadastroRequisicao(requisicao);
}
