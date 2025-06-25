import { notificar } from "../../shared/notificacao.js"
import "../../shared/interceptor.js"
import { listarUnidade, listarLaboratorio } from "../../api/unidadeApi.js"
import * as formatador from "../../shared/formatador.js"
import { cadastrarFuncionario } from "../../api/usuarioApi.js"
import { pegarPermissaoUsuario, pegarUnidadeUsuario, pegarTipoUnidadeUsuario } from "../../shared/gerenciador-permissoes.js"

document.addEventListener('DOMContentLoaded', function () {
    const cpfInput = document.getElementById('cpf')
    const telefoneInput = document.getElementById('telefone')
    const emailInput = document.getElementById('email')
    const registroInput = document.getElementById('registro')
    const permissaoInput = document.getElementById('permissao')
    const cnesUnidadeInput = document.getElementById('cnesUnidade')
    const tipoUnidadeInput = document.getElementById('tipoUnidade')

    const cadastroFuncionario = document.getElementById('cadastroFuncionario')

    const permissaoUsuarioLogado = pegarPermissaoUsuario()
    const cnesUsuarioLogado = pegarUnidadeUsuario()
    const tipoUnidadeUsuarioLogado = pegarTipoUnidadeUsuario()
    const isAdministrador = permissaoUsuarioLogado === 'ADMINISTRADOR'
    const isGestor = permissaoUsuarioLogado === 'GESTAO'

    configurarInterfacePorPermissao()

    if (cpfInput) {
        cpfInput.addEventListener('input', formatador.formatarCPFInput)
    }
    if (telefoneInput) {
        telefoneInput.addEventListener('input', formatador.formatarTelefoneInput)
    }
    if (cnesUnidadeInput) {
        cnesUnidadeInput.addEventListener('blur', buscarUnidadePorCNES)
    }

    if (permissaoInput) {
        permissaoInput.addEventListener('change', selecionarTipoUnidadeAutomaticamente)
        permissaoInput.addEventListener('change', controlarVisibilidadeRegistro)
    }
    if (tipoUnidadeInput) {
        tipoUnidadeInput.addEventListener('change', validarCompatibilidadePermissao)
    }

    function configurarInterfacePorPermissao() {
        if (isAdministrador) {
            filtrarOpcoesPermissao()
            controlarVisibilidadeRegistro()
            return
        }

        if (isGestor) {
            configurarUnidadeGestor()
            filtrarOpcoesPermissao()
            controlarVisibilidadeRegistro()
            return
        }

        notificar('Você não tem permissão para cadastrar funcionários', 'error')
        desabilitarFormulario()
    }

    function configurarUnidadeGestor() {
        tipoUnidadeInput.disabled = true
        cnesUnidadeInput.disabled = true
        
        if (cnesUnidadeInput) {
            cnesUnidadeInput.removeEventListener('blur', buscarUnidadePorCNES)
        }
        if (permissaoInput) {
            permissaoInput.removeEventListener('change', selecionarTipoUnidadeAutomaticamente)
        }
        if (tipoUnidadeInput) {
            tipoUnidadeInput.removeEventListener('change', validarCompatibilidadePermissao)
        }
        
        buscarUnidadeGestor()
    }

    async function buscarUnidadeGestor() {
        if (!cnesUsuarioLogado) {
            notificar('CNES da unidade não encontrado', 'error')
            return
        }

        try {
            const tipoUnidade = pegarTipoUnidadeUsuario()
            const unidade = tipoUnidade === 'unidade' ? await listarUnidade(cnesUsuarioLogado) : await listarLaboratorio(cnesUsuarioLogado)

            if (unidade) {
                tipoUnidadeInput.value = tipoUnidade
                cnesUnidadeInput.value = cnesUsuarioLogado
                exibirDadosUnidade(unidade)
            } else {
                notificar('Unidade do gestor não encontrada', 'error')
                desabilitarFormulario()
            }
        } catch (error) {
            console.error('Erro ao buscar unidade do gestor:', error)
            notificar('Erro ao buscar informações da unidade', 'error')
            desabilitarFormulario()
        }
    }

    function desabilitarFormulario() {
        const campos = [cpfInput, telefoneInput, emailInput, registroInput, permissaoInput, cnesUnidadeInput, tipoUnidadeInput]
        campos.forEach(campo => {
            if (campo) {
                campo.disabled = true
            }
        })
        
        const submitButton = document.querySelector('.acoes button')
        if (submitButton) {
            submitButton.disabled = true
            submitButton.textContent = 'Sem permissão'
        }
    }

    function selecionarTipoUnidadeAutomaticamente() {
        const permissao = permissaoInput.value
        
        if (isGestor) {
            return
        }
        
        limparMensagensErro()
        limparEstilosErro()
        
        cnesUnidadeInput.value = ''
        const dadosUnidadeDiv = document.getElementById('dadosUnidade')
        if (dadosUnidadeDiv) {
            dadosUnidadeDiv.style.display = 'none'
        }

        switch (permissao) {
            case 'ACESSO_ATENDIMENTO':
            case 'ACESSO_EXAMES':
                tipoUnidadeInput.value = 'unidade'
                tipoUnidadeInput.disabled = true
                break
            case 'ACESSO_LABORATORIO':
                tipoUnidadeInput.value = 'laboratorio'
                tipoUnidadeInput.disabled = true
                break
            case 'GESTAO':
                if (isAdministrador) {
                    tipoUnidadeInput.disabled = false
                } else {
                    tipoUnidadeInput.value = tipoUnidadeUsuarioLogado
                    tipoUnidadeInput.disabled = true
                }
                break
            case 'ADMINISTRADOR':
                if (isAdministrador) {
                    tipoUnidadeInput.disabled = false
                }
                break
            default:
                tipoUnidadeInput.value = ''
                tipoUnidadeInput.disabled = false
        }
        validarCompatibilidadePermissao()
    }

    function limparCamposIncompatíveis() {
        const permissao = permissaoInput.value
        const tipoUnidade = tipoUnidadeInput.value

        if ((permissao === 'ACESSO_LABORATORIO' && tipoUnidade === 'unidade') ||
            (permissao === 'ACESSO_ATENDIMENTO' && tipoUnidade === 'laboratorio') ||
            (permissao === 'ACESSO_EXAMES' && tipoUnidade === 'laboratorio')) {
            
            cnesUnidadeInput.value = ''
            const dadosUnidadeDiv = document.getElementById('dadosUnidade')
            if (dadosUnidadeDiv) {
                dadosUnidadeDiv.style.display = 'none'
            }
        }
    }

    function validarCompatibilidadePermissao() {
        const permissao = permissaoInput.value
        const tipoUnidade = tipoUnidadeInput.value
        const cnesUnidade = cnesUnidadeInput.value.trim()

        limparMensagensErro()
        limparEstilosErro()

        if (!permissao || !tipoUnidade) {
            return true
        }

        if (isGestor) {
            if (cnesUnidade !== cnesUsuarioLogado) {
                exibirMensagemErro('Como gestor, você só pode cadastrar funcionários na sua própria unidade')
                return false
            }
            return true
        }

        let erro = null
        let camposComErro = []

        if (permissao === 'ACESSO_LABORATORIO' && tipoUnidade === 'unidade') {
            erro = 'Usuários com permissão de laboratório não podem ser cadastrados em unidades de saúde'
            camposComErro = [permissaoInput, tipoUnidadeInput]
        }

        else if (permissao === 'ACESSO_ATENDIMENTO' && tipoUnidade === 'laboratorio') {
            erro = 'Usuários com acesso de atendimento não podem ser cadastrados em laboratórios'
            camposComErro = [permissaoInput, tipoUnidadeInput]
        }

        else if (permissao === 'ACESSO_EXAMES' && tipoUnidade === 'laboratorio') {
            erro = 'Usuários com acesso a exames não podem ser cadastrados em laboratórios'
            camposComErro = [permissaoInput, tipoUnidadeInput]
        }

        if (erro) {
            exibirMensagemErro(erro)
            destacarCamposComErro(camposComErro)
            limparCamposIncompatíveis()
            return false
        }

        return true
    }

    function exibirMensagemErro(mensagem) {
        let erroDiv = document.getElementById('erroValidacao')
        if (!erroDiv) {
            erroDiv = document.createElement('div')
            erroDiv.id = 'erroValidacao'
            erroDiv.className = 'erro-validacao'
            
            const fieldsetProfissional = document.querySelector('fieldset:nth-of-type(2)')
            fieldsetProfissional.parentNode.insertBefore(erroDiv, fieldsetProfissional.nextSibling)
        }
        
        erroDiv.textContent = mensagem
        erroDiv.style.display = 'flex'
    }

    function limparMensagensErro() {
        const erroDiv = document.getElementById('erroValidacao')
        if (erroDiv) {
            erroDiv.style.display = 'none'
        }
    }

    function destacarCamposComErro(campos) {
        campos.forEach(campo => {
            if (campo) {
                campo.classList.add('campo-erro')
            }
        })
    }

    function limparEstilosErro() {
        const campos = [permissaoInput, tipoUnidadeInput, cnesUnidadeInput]
        campos.forEach(campo => {
            if (campo) {
                campo.classList.remove('campo-erro')
            }
        })
    }

    async function buscarUnidadePorCNES(event) {
        const cnes = event.target.value.trim()
        const dadosUnidadeDiv = document.getElementById('dadosUnidade')
        
        if (!cnes) {
            dadosUnidadeDiv.style.display = 'none'
            return
        }

        if (isGestor && cnes !== cnesUsuarioLogado) {
            notificar('Como gestor, você só pode cadastrar funcionários na sua própria unidade', 'error')
            cnesUnidadeInput.value = cnesUsuarioLogado
            return
        }

        if (!validarCompatibilidadePermissao()) {
            return
        }

        try {
            const unidade = document.getElementById('tipoUnidade').value === 'unidade' ? await listarUnidade(cnes) : await listarLaboratorio(cnes)
            if (unidade) {
                 exibirDadosUnidade(unidade)
            } else {
                notificar('Unidade não encontrada', 'error')
                dadosUnidadeDiv.style.display = 'none'
            }            
        } catch (error) {
            console.error('Erro ao buscar unidade:', error)
            notificar('Erro ao buscar unidade', 'error')
            dadosUnidadeDiv.style.display = 'none'
        }
    }

    function exibirDadosUnidade(unidade) {
        const dadosUnidadeDiv = document.getElementById('dadosUnidade')
        const nomeUnidade = document.getElementById('nomeUnidade')
        const cnpjUnidade = document.getElementById('cnpjUnidade')
        const telefoneUnidade = document.getElementById('telefoneUnidade')
        const enderecoUnidade = document.getElementById('enderecoUnidade')

        nomeUnidade.textContent = unidade.nome || '-'
        cnpjUnidade.textContent = formatador.formatarCNPJ(unidade.cnpj) || '-'
        telefoneUnidade.textContent = formatador.formatarTelefone(unidade.telefone) || '-'
        
        if (unidade.endereco) {
            const end = unidade.endereco
            const enderecoCompleto = `${end.logradouro}, ${end.numero} - ${end.bairro}, ${end.municipio} - ${end.uf}, ${end.cep}`
            enderecoUnidade.textContent = enderecoCompleto
        } else {
            enderecoUnidade.textContent = '-'
        }

        dadosUnidadeDiv.style.display = 'block'
    }

    function filtrarOpcoesPermissao() {
        if (!permissaoInput) return

        const placeholder = permissaoInput.querySelector('option[value=""]')
        permissaoInput.innerHTML = ''
        if (placeholder) {
            permissaoInput.appendChild(placeholder)
        }

        if (isAdministrador) {
            adicionarOpcao('ACESSO_ATENDIMENTO', 'Atendimento')
            adicionarOpcao('ACESSO_EXAMES', 'Exames')
            adicionarOpcao('ACESSO_LABORATORIO', 'Laboratório')
            adicionarOpcao('GESTAO', 'Gestão')
            adicionarOpcao('ADMINISTRADOR', 'Administrador')
        } else if (isGestor) {
            if (tipoUnidadeUsuarioLogado === 'unidade') {
                adicionarOpcao('ACESSO_ATENDIMENTO', 'Atendimento')
                adicionarOpcao('ACESSO_EXAMES', 'Exames')
                adicionarOpcao('GESTAO', 'Gestão')
            } else if (tipoUnidadeUsuarioLogado === 'laboratorio') {
                adicionarOpcao('ACESSO_LABORATORIO', 'Laboratório')
                adicionarOpcao('GESTAO', 'Gestão')
            }
        }
    }

    function adicionarOpcao(value, text) {
        const option = document.createElement('option')
        option.value = value
        option.textContent = text
        permissaoInput.appendChild(option)
    }

    function controlarVisibilidadeRegistro() {
        const permissao = permissaoInput.value
        const registroRow = document.getElementById('registroRow')
        
        if (!registroRow) return
        
        const precisaRegistro = permissao === 'ACESSO_EXAMES' || permissao === 'ACESSO_LABORATORIO'
        
        if (precisaRegistro) {
            registroRow.style.display = 'flex'
            registroInput.required = true
        } else {
            registroRow.style.display = 'none'
            registroInput.required = false
            registroInput.value = ''
        }
    }

    cadastroFuncionario.addEventListener('submit', async function (event) {
        event.preventDefault()

        if (!validarCompatibilidadePermissao()) {
            notificar('Existem incompatibilidades entre permissão e tipo de unidade', 'error')
            return
        }

        if (isGestor) {
            const cnesDigitado = cnesUnidadeInput.value.trim()
            if (cnesDigitado !== cnesUsuarioLogado) {
                notificar('Como gestor, você só pode cadastrar funcionários na sua própria unidade', 'error')
                return
            }
        }

        const dados = new FormData(cadastroFuncionario)
        const valores = Object.fromEntries(dados.entries())

        const cnesUnidade = cnesUnidadeInput.value.trim()

        const requisicaoCadastro = {
            registro: valores.registro,
            nome: valores.nome,
            cpf: valores.cpf.replace(/\D/g, ''),
            email: valores.email,
            telefone: valores.telefone.replace(/\D/g, ''),
            permissao: valores.permissao,
        }

        if (document.getElementById('tipoUnidade').value === 'unidade') {
            requisicaoCadastro.unidade_saude_cnes = cnesUnidade
        } else {
            requisicaoCadastro.laboratorio_cnes = cnesUnidade
        }

        try {
            cadastrarFuncionario(requisicaoCadastro)
        } catch (error) {
            console.error('Erro ao cadastrar funcionário:', error)
            alert('Erro ao cadastrar funcionário: ' + error.message)
        }
    })
})

