export function formatarCPF(cpf) {
    if (!cpf) return '';
    cpf = cpf.replace(/\D/g, ''); 
    if (cpf.length !== 11) return cpf; 
    return `${cpf.slice(0, 3)}.${cpf.slice(3, 6)}.${cpf.slice(6, 9)}-${cpf.slice(9)}`;
}

export function formatarTelefone(telefone) {
    if (!telefone) return '';
    telefone = telefone.replace(/\D/g, '');
    if (telefone.length !== 11) return telefone;
    return `(${telefone.slice(0, 2)}) ${telefone.slice(2, 7)}-${telefone.slice(7)}`;
}

export function formatarCNPJ(cnpj) {
    if (!cnpj) return '';
    cnpj = cnpj.replace(/\D/g, '');
    if (cnpj.length !== 14) return cnpj;
    return `${cnpj.slice(0, 2)}.${cnpj.slice(2, 5)}.${cnpj.slice(5, 8)}/${cnpj.slice(8, 12)}-${cnpj.slice(12)}`;
}

export function formatarCPFInput(event) {
    let valor = event.target.value
    valor = valor.replace(/\D/g, '')

    if (valor.length <= 3) {
        event.target.value = valor
    } else if (valor.length <= 6) {
        event.target.value = `${valor.slice(0, 3)}.${valor.slice(3)}`
    } else if (valor.length <= 9) {
        event.target.value = `${valor.slice(0, 3)}.${valor.slice(3, 6)}.${valor.slice(6)}`
    } else {
        event.target.value = `${valor.slice(0, 3)}.${valor.slice(3, 6)}.${valor.slice(6, 9)}-${valor.slice(9, 11)}`
    }
}

export function formatarCNPJInput(event) {
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

export function formatarTelefoneInput(event) {
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