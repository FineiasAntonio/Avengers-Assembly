
export function deslogar() {
    localStorage.removeItem('token')
    verificarAutenticacaoValida()    
}