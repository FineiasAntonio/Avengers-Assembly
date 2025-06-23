function validarCPF(cpf) {
    cpf = cpf.replace(/\D/g, '')
    
    if (cpf.length !== 11) return false
    
    if (/^(\d)\1{10}$/.test(cpf)) return false
    
    let soma = 0
    for (let i = 0; i < 9; i++) {
        soma += parseInt(cpf.charAt(i)) * (10 - i)
    }
    let resto = 11 - (soma % 11)
    let digito1 = resto < 2 ? 0 : resto
    
    if (parseInt(cpf.charAt(9)) !== digito1) return false
    
    soma = 0
    for (let i = 0; i < 10; i++) {
        soma += parseInt(cpf.charAt(i)) * (11 - i)
    }
    resto = 11 - (soma % 11)
    let digito2 = resto < 2 ? 0 : resto
    
    return parseInt(cpf.charAt(10)) === digito2
}

function validarEmail(email) {
    const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    return regex.test(email)
}