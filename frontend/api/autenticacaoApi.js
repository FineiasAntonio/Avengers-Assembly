const API_ENDERECO = "http://localhost:8080/api/"

export async function loginRequisicao(credencial, senha) {
    const credenciais = { credencial, senha }

    const response = await fetch(API_ENDERECO + "auth/login", {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(credenciais)
    })

    if (!response.ok) {
            const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status}` }));
            throw new Error(errorData.message);
    }

    const token = await response.json()

    console.log("Autenticado com sucesso, token: " + token)

    localStorage.setItem("token", token)
}