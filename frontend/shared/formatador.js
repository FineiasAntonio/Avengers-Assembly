export function formatarCPF(cpf) {
    if (!cpf) return '';
    cpf = cpf.replace(/\D/g, ''); 
    if (cpf.length !== 11) return cpf; 
    return `${cpf.slice(0, 3)}.${cpf.slice(3, 6)}.${cpf.slice(6, 9)}-${cpf.slice(9)}`;
}