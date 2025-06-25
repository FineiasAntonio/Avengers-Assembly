import "../../api/cadastroApi.js"
import { cadastrarUnidade } from "../../api/unidadeApi.js"

document.addEventListener('DOMContentLoaded', function () {
    const cnpjInput = document.getElementById('cnpj')
    const telefoneInput = document.getElementById('telefone')
    const cepInput = document.getElementById('cepEndereco')
    const logradouroInput = document.getElementById('logradouro')
    const bairroInput = document.getElementById('bairro')
    const municipioInput = document.getElementById('municipio')
    const ufInput = document.getElementById('uf')
    const tipoUnidadeInput = document.getElementById('tipoUnidade')

    const cadastroUnidade = document.getElementById('cadastroUnidade')

    if (cnpjInput) {
        cnpjInput.addEventListener('input', formatarCNPJ)
    }
    if (telefoneInput) {
        telefoneInput.addEventListener('input', formatarTelefone)
    }
    if (cepInput) {
        cepInput.addEventListener('input', formatarCEPInput);
        cepInput.addEventListener('blur', buscarEnderecoPorCEP);
    }

    async function buscarEnderecoPorCEP(event) {
        const cep = event.target.value.replace(/\D/g, '');

        function limparCamposEndereco() {
            if (logradouroInput) logradouroInput.value = '';
            if (bairroInput) bairroInput.value = '';
            if (municipioInput) municipioInput.value = '';
            if (ufInput) ufInput.value = '';

            logradouroInput.removeAttribute("disabled")
            bairroInput.removeAttribute("disabled")
            municipioInput.removeAttribute("disabled")
            ufInput.removeAttribute("disabled")
        }

        if (cep.length === 8) {
            logradouroInput.value = "Buscando...";
            bairroInput.value = "Buscando...";
            municipioInput.value = "Buscando...";
            ufInput.value = "Buscando...";

            const response = await fetch(`https://viacep.com.br/ws/${cep}/json/`);
            if (!response.ok) {
                limparCamposEndereco();
            }
            const data = await response.json();

            if (data.erro) {
                alert('CEP não encontrado. Verifique o CEP digitado.');
                limparCamposEndereco();
            } else {
                logradouroInput.value = data.logradouro || '';
                bairroInput.value = data.bairro || '';
                municipioInput.value = data.localidade || '';
                ufInput.value = data.uf || '';
            }
        } else {
            limparCamposEndereco()
        }
    }

    cadastroUnidade.addEventListener('submit', async function (event) {
        event.preventDefault()

        const dados = new FormData(cadastroUnidade)
        const valores = Object.fromEntries(dados.entries())

        valores.cnpj = valores.cnpj.replace(/\D/g, '')
        valores.telefone = valores.telefone.replace(/\D/g, '')
        valores.cep = valores.cep.replace(/\D/g, '')

        valores.endereco = {
            logradouro: valores.logradouro,
            numero: valores.numero,
            complemento: valores.complemento,
            bairro: valores.bairro,
            municipio: valores.municipio,
            uf: valores.uf,
            cep: valores.cep,
        }
        
        cadastrarUnidade(valores, tipoUnidadeInput.value)
    })
})
    
function formatarCNPJ(event) {
    let valor = event.target.value
    valor = valor.replace(/\D/g, '')

    if (valor.length <= 2) {
        event.target.value = valor
    } else if (valor.length <= 5) {
        event.target.value = `${valor.slice(0, 2)}.${valor.slice(2)}`
    } else if (valor.length <= 8) {
        event.target.value = `${valor.slice(0, 2)}.${valor.slice(2, 5)}.${valor.slice(5)}`
    } else if (valor.length <= 12) {
        event.target.value = `${valor.slice(0, 2)}.${valor.slice(2, 5)}.${valor.slice(5, 8)}/${valor.slice(8)}`
    } else {
        event.target.value = `${valor.slice(0, 2)}.${valor.slice(2, 5)}.${valor.slice(5, 8)}/${valor.slice(8, 12)}-${valor.slice(12, 14)}`
    }
}

function formatarTelefone(event) {
    let valor = event.target.value
    valor = valor.replace(/\D/g, '')

    if (valor.length === 0) {
        event.target.value = ''
        return
    }

    if (valor.length <= 2) {
        event.target.value = `(${valor}`
    } else if (valor.length <= 6 && valor.length > 2) {
        event.target.value = `(${valor.slice(0, 2)}) ${valor.slice(2)}`
    } else if (valor.length <= 10 && valor.length > 6) {
        event.target.value = `(${valor.slice(0, 2)}) ${valor.slice(2, 6)}-${valor.slice(6, 10)}`
    } else if (valor.length === 11 && valor.length > 6) {
        event.target.value = `(${valor.slice(0, 2)}) ${valor.slice(2, 7)}-${valor.slice(7, 11)}`
    } else if (valor.length > 11) {
        event.target.value = `(${valor.slice(0, 2)}) ${valor.slice(2, 7)}-${valor.slice(7, 11)}`
    }
}

function formatarCEPInput(event) {
    let valor = event.target.value
    valor = valor.replace(/\D/g, '')

    if (valor.length > 5) {
        valor = valor.replace(/^(\d{5})(\d)/, '$1-$2')
    }
    event.target.value = valor.slice(0, 9)
}
