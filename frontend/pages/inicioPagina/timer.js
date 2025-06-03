import { pegarNomeUsuario } from "../../shared/gerenciador-permissoes.js";

document.addEventListener("DOMContentLoaded", () => {
    const relogio = document.getElementById("relogio");
    const data = document.getElementById("data");

    const saudacao = document.getElementById("saudacao");

    function atualizarRelogio() {
        const agora = new Date();
        const horas = String(agora.getHours()).padStart(2, '0');
        const minutos = String(agora.getMinutes()).padStart(2, '0');
        const segundos = String(agora.getSeconds()).padStart(2, '0');

        relogio.textContent = `${horas}:${minutos}:${segundos}`;
    }

    function atualizarData() {
        const agora = new Date();
        const opcoes = { year: 'numeric', month: 'long', day: 'numeric' };
        data.textContent = agora.toLocaleDateString('pt-BR', opcoes);
    }

    function atualizarSaudacao() {
        const agora = new Date();
        const horas = agora.getHours();

        if (horas < 12) {
            saudacao.textContent = "Bom dia,";
        } else if (horas < 18) {
            saudacao.textContent = "Boa tarde,";
        } else {
            saudacao.textContent = "Boa noite,";
        }

        const usuario = pegarNomeUsuario();
        saudacao.textContent += ` ${usuario ? usuario : "Fulano"}!`;
    }

    setInterval(() => {
        atualizarRelogio();
    }, 1000);

    atualizarData();
    atualizarSaudacao();

})