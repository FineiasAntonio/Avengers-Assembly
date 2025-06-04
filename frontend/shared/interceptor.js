import { pegarTokenUsuario } from "./gerenciador-permissoes.js"

const originalFetch = window.fetch;

window.fetch = async function (input, init = {}) {

    const headers = new Headers(init.headers || {})

    if (!headers.has('Authorization')) {
        const tokenUsuario = pegarTokenUsuario()
        headers.append('Authorization', `Bearer ${tokenUsuario}`)
    }

    init.headers = headers

    return originalFetch.call(this, input, init)
    .then(response => {
      if (response.status === 401) {
        console.log('nao autorizado kk')
      }
      return response;
    })
    .catch(error => {
      console.error('Request failed:', error);
      throw error;
    });
}

