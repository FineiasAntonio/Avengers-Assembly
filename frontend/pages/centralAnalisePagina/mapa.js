const visualizarMapa = document.getElementById("visualizarMapa");
const overlay = document.getElementById("overlay");
const fechar = document.getElementById("fecharMapaTela");

visualizarMapa.addEventListener('click', () => {
    overlay.style.display = 'flex';
});

fechar.addEventListener('click', () => {
    overlay.style.display = 'none';
});

overlay.addEventListener('click', () => {
    overlay.style.display = 'none';
});

