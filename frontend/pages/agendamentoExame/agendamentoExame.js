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
            if (response) {
                renderizarHorariosDisponiveis(response)
            }
        })

    })
})

function renderizarHorariosDisponiveis(horariosOcupados) {
    const container_root = document.getElementById("root")
    container_root.innerHTML = ""
    container_root.classList.remove("invisible")
    container_root.classList.add("visible")

    const container = document.createElement("fieldset")

    container_root.appendChild(container)

    const horariosPossiveis = []
    for (let hora = 8; hora <= 18; hora += 2) {
        horariosPossiveis.push(`${hora.toString().padStart(2, '0')}:00`)
    }

    const ocupadosPorProfissional = {}
    horariosOcupados.forEach(item => {
        item.data = item.data.split("Z")[0]
        if (!ocupadosPorProfissional[item.profissional]) {
            ocupadosPorProfissional[item.profissional] = []
        }
        ocupadosPorProfissional[item.profissional].push(item.data)
    })

    for (const [profissional, horariosOcupados] of Object.entries(ocupadosPorProfissional)) {
        const profissionalDiv = document.createElement("div")
        profissionalDiv.className = "profissional-horarios"

        const titulo = document.createElement("label")
        titulo.textContent = profissional
        profissionalDiv.appendChild(titulo)

        const listaHorarios = document.createElement("div")
        listaHorarios.className = "horarios-container"

        horariosPossiveis.forEach(horario => {
            const estaOcupado = horariosOcupados.some(ocupado => {
                return new Date(ocupado).getHours() === parseInt(horario.split(':')[0])
            })


            const botaoHorario = document.createElement("div")
            botaoHorario.className = "horario"
            botaoHorario.textContent = horario
            if (!estaOcupado) {
                botaoHorario.addEventListener('click', () => {
                    selecionarHorario(profissional, horario)
                })
            } else {
                botaoHorario.className += " horario-ocupado"
            }
            listaHorarios.appendChild(botaoHorario)

        })

        profissionalDiv.appendChild(listaHorarios)
        container.appendChild(profissionalDiv)
    }
}