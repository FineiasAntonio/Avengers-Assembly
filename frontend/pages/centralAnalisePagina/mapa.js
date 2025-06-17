let mapa = null;

document.addEventListener("DOMContentLoaded", () => {
    const visualizarMapa = document.getElementById("visualizarMapa");
    const overlay = document.getElementById("overlay");
    const fechar = document.getElementById("fecharMapaTela");

    visualizarMapa.addEventListener('click', async () => {
        overlay.style.display = 'flex';
        mapa = await iniciarMapa();

        if (mapa) {
            setTimeout(() => {mapa.invalidateSize()}, 200);
        };
    });

    fechar.addEventListener('click', async () => {
        overlay.style.display = 'none'; 

        if (mapa) {
            mapa.remove();
            mapa = null;
        }
    });
});

async function pegarCoordenadas(cidade) {
    try {
        const response = await fetch(`https://nominatim.openstreetmap.org/search?q=${cidade}&format=json`);
        const resposta = await response.json();

        if (resposta.length > 0) {
            const latitude = resposta[0].lat;
            const longitude = resposta[0].lon;
            return [latitude, longitude];
        }
        else {
            throw new Error("Cidade nao encontrada!");
        }
    }
    catch (error) {
        console.log(error);
        window.alert("Algo deu errado ao receber Cordenadas!")
    }
}

async function iniciarMapa() {
    const cidadeLatLgn = await pegarCoordenadas('Goiânia')
   
    if (cidadeLatLgn) {
        mapa = L.map('tela').setView(cidadeLatLgn, 11);

        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution:
            'Map data &copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        }).addTo(mapa);

        const dadosRegioes = [
            { nome: "Setor Sul", coords: [-16.704, -49.263], quantidade: 80 },
            { nome: "Setor Aeroporto", coords: [-16.723, -49.233], quantidade: 40 },
            { nome: "Setor Central", coords: [-16.686, -49.256], quantidade: 60 },
            { nome: "Setor Universitário", coords: [-16.711, -49.292], quantidade: 50 }
        ];

        dadosRegioes.forEach(regiao => {
            L.circle(regiao.coords, {
                color: '#4caf50',
                fillColor: '#4caf50',
                fillOpacity: 0.5,
                radius: regiao.quantidade * 20
            }).addTo(mapa)
                .bindPopup(`<strong>${regiao.nome}</strong><br>${regiao.quantidade} pessoas cadastradas`);
        });
        return mapa;
    }
    return null;
}