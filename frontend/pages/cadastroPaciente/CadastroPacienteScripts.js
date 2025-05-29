document.addEventListener('DOMContentLoaded', function () {

    const cpfInput = document.getElementById('cpf');
    const telefoneInput = document.getElementById('telefone');
    const dataNascimentoInput = document.getElementById('dataNascimento');
    const cepInput = document.getElementById('cepEndereco');
    const logradouroInput = document.getElementById('logradouro');
    const bairroInput = document.getElementById('bairro');
    const municipioInput = document.getElementById('municipio');
    const ufInput = document.getElementById('uf');

    if (cpfInput) {
        cpfInput.addEventListener('input', formatarCPF);
    }
    if (telefoneInput) {
        telefoneInput.addEventListener('input', formatarTelefone);
    }
    if (dataNascimentoInput) {
        dataNascimentoInput.addEventListener('input', formatarData);
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

                logradouroInput.disabled = !!data.logradouro;
                bairroInput.disabled = !!data.bairro;
                municipioInput.disabled = !!data.localidade;
                ufInput.disabled = !!data.uf;
            }
        } else {
            limparCamposEndereco()
        }
    }
})

function formatarCPF(event) {
    let valor = event.target.value;
    valor = valor.replace(/\D/g, '');

    if (valor.length <= 3) {
        event.target.value = valor;
    } else if (valor.length <= 6) {
        event.target.value = `${valor.slice(0, 3)}.${valor.slice(3)}`;
    } else if (valor.length <= 9) {
        event.target.value = `${valor.slice(0, 3)}.${valor.slice(3, 6)}.${valor.slice(6)}`;
    } else {
        event.target.value = `${valor.slice(0, 3)}.${valor.slice(3, 6)}.${valor.slice(6, 9)}-${valor.slice(9, 11)}`;
    }
}

function formatarTelefone(event) {
    let valor = event.target.value;
    valor = valor.replace(/\D/g, '');

    if (valor.length === 0) {
        event.target.value = '';
        return;
    }

    if (valor.length <= 2) {
        event.target.value = `(${valor}`;
    } else if (valor.length <= 6 && valor.length > 2) {
        event.target.value = `(${valor.slice(0, 2)}) ${valor.slice(2)}`;
    } else if (valor.length <= 10 && valor.length > 6) {
        event.target.value = `(${valor.slice(0, 2)}) ${valor.slice(2, 6)}-${valor.slice(6, 10)}`;
    } else if (valor.length === 11 && valor.length > 6) {
        event.target.value = `(${valor.slice(0, 2)}) ${valor.slice(2, 7)}-${valor.slice(7, 11)}`;
    } else if (valor.length > 11) {
        event.target.value = `(${valor.slice(0, 2)}) ${valor.slice(2, 7)}-${valor.slice(7, 11)}`;
    }
}

function formatarData(event) {
    let valor = event.target.value;
    valor = valor.replace(/\D/g, '');

    if (valor.length <= 2) {
        event.target.value = valor;
    } else if (valor.length <= 4) {
        event.target.value = `${valor.slice(0, 2)}/${valor.slice(2)}`;
    } else {
        event.target.value = `${valor.slice(0, 2)}/${valor.slice(2, 4)}/${valor.slice(4, 8)}`;
    }
}

function formatarCEPInput(event) {
    let valor = event.target.value;
    valor = valor.replace(/\D/g, '');

    if (valor.length > 5) {
        valor = valor.replace(/^(\d{5})(\d)/, '$1-$2');
    }
    event.target.value = valor.slice(0, 9);
}


