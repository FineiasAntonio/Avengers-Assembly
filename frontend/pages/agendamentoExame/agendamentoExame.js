import { consultarHorarioOcupado } from "../../api/agendamentoApi.js"
import { pegarUnidadeUsuario } from "../../shared/gerenciador-permissoes.js"
import { listarUnidade } from "../../api/unidadeApi.js"

document.addEventListener("DOMContentLoaded", () => {
    const inputCartaoSus = document.getElementById("cartaoSUS")
    const dataInput = document.getElementById("data")
    const cnesUnidadeUsuario = pegarUnidadeUsuario()

    listarUnidade(cnesUnidadeUsuario).then(resultado => {
        document.getElementById('ubsPaciente').textContent = resultado.nome
        document.getElementById('municipioPaciente').textContent = `${resultado.endereco.logradouro}, ${resultado.endereco.numero}, ${resultado.endereco.complemento}, ${resultado.endereco.municipio}/${resultado.endereco.uf}`
    })


    inputCartaoSus.addEventListener('blur', () => {

    })

    dataInput.addEventListener('change', (e) => {
        const data = e.target.value

        consultarHorarioOcupado(cnesUnidadeUsuario, data).then(response => {
            renderizarHorariosDisponiveis(response)
        })

    })
})

function renderizarHorariosDisponiveis(horariosPorProfissional) {
    const containerRoot = document.getElementById("root");
    containerRoot.innerHTML = "";
    containerRoot.classList.remove("invisible");
    containerRoot.classList.add("visible");

    const container = document.createElement("fieldset");
    containerRoot.appendChild(container);

    if (!horariosPorProfissional || typeof horariosPorProfissional !== 'object') {
        const mensagem = document.createElement("p");
        mensagem.textContent = "Nenhum dado de horários disponível.";
        container.appendChild(mensagem);
        return;
    }

    const horariosPossiveis = [];
    for (let hora = 8; hora <= 18; hora += 2) {
        horariosPossiveis.push(`${hora.toString().padStart(2, '0')}:00`);
    }

    const profissionais = Object.keys(horariosPorProfissional);
    if (profissionais.length === 0) {
        const mensagem = document.createElement("p");
        mensagem.textContent = "Nenhum profissional disponível para esta data.";
        container.appendChild(mensagem);
        return;
    }

    profissionais.forEach(profissional => {
        const profissionalDiv = document.createElement("div");
        profissionalDiv.className = "profissional-horarios";

        const titulo = document.createElement("h3");
        titulo.textContent = profissional;
        profissionalDiv.appendChild(titulo);

        const listaHorarios = document.createElement("div");
        listaHorarios.className = "horarios-container";

        const horariosOcupados = horariosPorProfissional[profissional] || [];

        horariosPossiveis.forEach(horario => {
            const hora = parseInt(horario.split(':')[0]);

            const estaOcupado = horariosOcupados.some(horarioOcupado => {
                const dataHora = new Date(horarioOcupado);
                const mesmaHora = dataHora.getUTCHours() === hora;
                return mesmaHora;
            });

            const botaoHorario = document.createElement("button");
            botaoHorario.className = "horario";
            botaoHorario.textContent = horario;

            if (estaOcupado) {
                botaoHorario.disabled = true;
                botaoHorario.classList.add("horario-ocupado");
            } else {
                botaoHorario.addEventListener('click', () => {
                    //TODO: Agendar o exame com o horário selecionado
                });
            }

            listaHorarios.appendChild(botaoHorario);
        });

        profissionalDiv.appendChild(listaHorarios);
        container.appendChild(profissionalDiv);
    });
}