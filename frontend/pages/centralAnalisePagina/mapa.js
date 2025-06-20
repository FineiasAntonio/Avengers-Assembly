import { pegarDadosQuantidadePacientesPorRegiao } from "../../api/centralAnaliseApi.js";

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
            const latitude = parseFloat(resposta[0].lat);
            const longitude = parseFloat(resposta[0].lon);

            const bbox = resposta[0].boundingbox.map(Number);

            const raio = raioDoBairro(bbox[0], bbox[2], bbox[1], bbox[3]);

            return { coords: [latitude, longitude], raio };
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

function raioDoBairro(lat1, lon1, lat2, lon2) {
    const R = 6371000;
    const toRad = x => x * Math.PI / 180;

    const dLat = toRad(lat2 - lat1);
    const aLat = Math.sin(dLat / 2) * Math.sin(dLat / 2);
    const cLat = 2 * Math.atan2(Math.sqrt(aLat), Math.sqrt(1 - aLat));
    const distanciaVertical = R * cLat;

    const latMedia = (lat1 + lat2) / 2;
    const dLon = toRad(lon2 - lon1);
    const aLon = Math.cos(toRad(latMedia)) * Math.cos(toRad(latMedia)) * Math.sin(dLon / 2) * Math.sin(dLon / 2);
    const cLon = 2 * Math.atan2(Math.sqrt(aLon), Math.sqrt(1 - aLon));
    const distanciaHorizontal = R * cLon;

    const area = distanciaVertical * distanciaHorizontal;

    const raio = Math.sqrt(area / Math.PI);

    return raio;
}


async function iniciarMapa() {
    const cidadeLatLgn = await pegarCoordenadas('Goiânia')
   
    if (cidadeLatLgn) {
        mapa = L.map('tela').setView(cidadeLatLgn.coords, 11);

        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution:
            'Map data &copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        }).addTo(mapa);

        const response = await pegarDadosQuantidadePacientesPorRegiao();

        const dadosRegioes = await Promise.all(
            response.map(async (item) => {
                const cordReg = await pegarCoordenadas(item.bairro+", Goiânia");
                return {
                    bairro: item.bairro,
                    coords: cordReg.coords,
                    raio: cordReg.raio,
                    quantidade: item.quantidade,
                }
        }))
        console.log(dadosRegioes)

        dadosRegioes.forEach(regiao => {
            L.circle(regiao.coords, {
                color: '#4caf50',
                fillColor: '#4caf50',
                fillOpacity: 0.5,
                radius: regiao.raio
            }).addTo(mapa)
                .bindPopup(`<strong>${regiao.bairro}</strong><br>${regiao.quantidade} pessoas cadastradas`);
        });
        return mapa;
    }
    return null;
}