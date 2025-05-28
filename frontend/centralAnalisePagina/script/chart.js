const graficoPizza = document.getElementById("graficoP");
const graficoBarra = document.getElementById("graficoB");

new Chart(graficoPizza, {
    type: "doughnut",
    data: {
        labels: ["Alta Prioridade", "Média Prioridade", "Baixa Prioridade"],
        datasets: [{
            data: [60, 25, 15],
            backgroundColor: ["#e15759", "#f28e2b", "#59a14f"]
        }]
    },
    options: {
        responsive: true,
        plugins: {
            legend: {
            position: 'bottom'
            }
        }
    }
});

new Chart(graficoBarra, {
    type: "bar",
    data: {
        labels: ["25 a 30", "30 a 40", "40 a 50", "50 a 60", "60 a 65"],
        datasets: [{
            label: "Pacientes",
            data: [35, 50, 40, 60, 55],
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