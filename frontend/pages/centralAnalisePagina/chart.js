import { pegarDadosQtdPacientes } from "../../api/centralAnaliseApi.js";

const graficoPizza = document.getElementById("graficoP");
const graficoBarra = document.getElementById("graficoB");

let chartPizza = null;
let chartBarra = null;

document.addEventListener("DOMContentLoaded", async () => {
    const funcao = document.getElementById("filtrarGrafico");

    funcao.addEventListener("change", () => {
        const novoValor = funcao.value;
        iniciarGraficoPizza(novoValor);
    })

    iniciarGraficoPizza(funcao.value);

    iniciarGraficoBarra("idade");

    const labelQtdPacientes = document.getElementById("labelQtdPacientes");
    const qtdPacientes = await iniciarQtdPacientes();
    labelQtdPacientes.textContent = qtdPacientes;
});

async function iniciarGraficoBarra(funcao) {
    const response = await pegarDadosQtdPacientes(funcao);

    if (chartBarra) chartBarra.destroy();
    
    chartBarra = new Chart(graficoBarra, {
        type: "bar",
        data: {
            labels: ["25 a 30", "30 a 40", "40 a 50", "50 a 60", "60 a 65"],
            datasets: [{
                label: "Pacientes",
                data: [response.qtd_25_a_30, response.qtd_30_a_40, response.qtd_40_a_50, response.qtd_50_a_60, response.qtd_60_a_65],
                backgroundColor: "#4e79a7"
            }]
        },
        options: {
            responsive: true,
            plugins: {
                legend: {display: false}
            }
        }
    });
}

async function iniciarGraficoPizza(funcao) {
    const response = await pegarDadosQtdPacientes(funcao);

    if (chartPizza) chartPizza.destroy();
    
    let labels = [];
    let data = [];

    if (response.branca !== undefined) {
        labels = ["Branca", "Preta", "Parda", "Amarela", "Indígena"];
        data = [
            response.branca || 0,
            response.preta || 0,
            response.parda || 0,
            response.amarela || 0,
            response.indigena || 0
        ];
    } 

    else if (response.analfabeta !== undefined) {
        labels = [
            "Analfabeta",
            "Fundamental Incompleto",
            "Fundamental Completo",
            "Médio Incompleto",
            "Médio Completo",
            "Superior Incompleto",
            "Superior Completo"
        ];
        data = [
            response.analfabeta || 0,
            response.fundamental_incompleto || 0,
            response.fundamental_completo || 0,
            response.medio_incompleto || 0,
            response.medio_completo || 0,
            response.superior_incompleto || 0,
            response.superior_completo || 0
        ];
    }


    console.log(data, labels)
    chartPizza = new Chart(graficoPizza, {
        type: "doughnut",
        data: {
            labels: labels,
            datasets: [{
                data: data,
                backgroundColor: ["#e15759", "#f28e2b", "#59a14f", "#4e79a7", "#edc949"]
            }]
        },
        options: {
            responsive: true,
            plugins: {
                legend: {
                position: 'right'
                }
            }
        }
    });
}

async function iniciarQtdPacientes() {
    const response = await pegarDadosQtdPacientes("padrao");
    if (response.quantidade_pacientes) {
        return response.quantidade_pacientes;
    }
}