function verificarAutenticacaoValida() {
    const token = localStorage.getItem('token');

    if (!token) {
        window.location.replace('/auth/LoginPagina.html');
    }
}

verificarAutenticacaoValida()