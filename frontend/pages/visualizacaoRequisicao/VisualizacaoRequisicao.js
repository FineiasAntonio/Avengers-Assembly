import { pegarUnidadeUsuario } from "../../shared/gerenciador-permissoes.js";
import { listarUnidade } from "../../api/unidadeApi.js";
import { ListarRequisicaoExame } from "../../api/cadastroApi.js";

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
    const nomeResponsavel = document.getElementById("nomeResponsavel");

    const protocolo = new URLSearchParams(window.location.search).get("protocolo");

    ListarRequisicaoExame(protocolo).then((resultado) => {
        cpfPaciente.textContent += resultado.paciente.cpf;
        cartaoSus.textContent += resultado.paciente.cartao_sus;
        nomePaciente.textContent += resultado.paciente.nome;
        idadePaciente.textContent += resultado.paciente.idade;
        dataNascimento.textContent += new Date(resultado.paciente.data_nascimento).toLocaleDateString();
        nomeMae.textContent += resultado.paciente.nome_mae || "Não informado";

        const endereco = resultado.paciente.endereco;
        enderecoPaciente.textContent += `${endereco.logradouro}, ${endereco.numero}${endereco.complemento ? ', ' + endereco.complemento : ''} - ${endereco.bairro}, ${endereco.municipio} - ${endereco.uf}, CEP: ${endereco.cep}`;

        telefonePaciente.textContent += resultado.paciente.telefone;
        nacionalidadePaciente.textContent += resultado.paciente.nacionalidade || "Não informado";
        racaCorPaciente.textContent += resultado.paciente.raca || "Não informado";
        escolaridadePaciente.textContent += resultado.paciente.escolaridade || "Não informado";

        motivoExame.textContent = resultado.motivo_exame;
        fezExame.textContent = resultado.fez_exame_preventivo ? "Sim" : "Não";
        anoUltimoExame.textContent = resultado.ano_ultimo_exame || "Não informado";
        usaDIU.textContent = resultado.usa_diu;
        estaGravida.textContent = resultado.esta_gravida;
        usaAnticoncepcional.textContent = resultado.usa_anticoncepcional;
        usaHormonio.textContent = resultado.usa_hormonio_menopausa || "Não informado";
        fezRadioterapia.textContent = resultado.fez_radioterapia;
        teveSangramentoRelacoes.textContent = resultado.sangramento_apos_relacoes || "Não informado";
        teveSangramentoMenopausa.textContent = resultado.sangramento_apos_menopausa || "Não informado";
        ultimaMenstruacao.textContent = resultado.data_ultima_menstruacao ?
            new Date(resultado.data_ultima_menstruacao).toLocaleDateString() : "Não informado";

        inspecaoColo.textContent = resultado.inspecao_colo;
        sinaisDST.textContent = resultado.sinais_dst ? "Sim" : "Não";
        dataColeta.textContent = new Date(resultado.data_coleta).toLocaleDateString();
        registroResponsavel.textContent = resultado.responsavel_registro;
        nomeResponsavel.textContent = resultado.responsavel.nome;
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
