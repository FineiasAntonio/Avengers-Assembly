const horarios = [
  "08:00", "08:30", "09:00", "09:30",
  "10:00", "10:30", "11:00", "11:30",
  "13:00", "13:30", "14:00", "14:30"
];

const agendadosPorData = {
  "2025-06-25": ["08:30", "10:30", "13:00"],
  "2025-06-26": ["09:00", "11:00", "14:00"]
};

document.addEventListener("DOMContentLoaded", () => {
  const inputData        = document.getElementById("data");
  const disponiveisWrapper = document.getElementById("horariosDisponiveis");
  const listaDisponiveis = document.getElementById("listaDisponiveis");

  disponiveisWrapper.classList.add("invisible");

  const atualizarDisponiveis = (dataSelecionada) => {
    if (!dataSelecionada) {
      disponiveisWrapper.classList.add("invisible");
      listaDisponiveis.innerHTML = "";
      return;
    }

    const agendados = agendadosPorData[dataSelecionada] || [];
    const livres = horarios.filter(h => !agendados.includes(h));

    listaDisponiveis.innerHTML = "";
    livres.forEach(h => {
      const li = document.createElement("li");
      li.textContent = h;
      listaDisponiveis.appendChild(li);
    });

    disponiveisWrapper.classList.remove("invisible");
  };

  inputData.addEventListener("change", () => {
    atualizarDisponiveis(inputData.value);
  });
});
