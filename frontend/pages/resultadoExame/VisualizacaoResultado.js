import { listarLaboratorio } from '../../api/unidadeApi.js';
import { pegarUnidadeUsuario } from '../../shared/gerenciador-permissoes.js';
import { EmitirResultadoExame, BuscarResultadoExame } from '../../api/cadastroApi.js';
import { ListarRequisicaoExame } from '../../api/cadastroApi.js';

class VisualizacaoResultadoExame {

    protocoloExame = '';
    resultadoExame = null;
    requisicaoExame = null;
    laboratorio = {};

    constructor() {
        this.initializeEventListeners();
        this.carregarDados();
    }

    async carregarDados() {
        const urlParams = new URLSearchParams(window.location.search);
        this.protocoloExame = urlParams.get('protocolo');
        
        if (!this.protocoloExame) {
            alert('Protocolo não fornecido');
            return;
        }

        document.getElementById('numeroExame').textContent = this.protocoloExame;

        // Carregar dados do laboratório
        const cnes = pegarUnidadeUsuario();
        try {
            this.laboratorio = await listarLaboratorio(cnes);
            document.getElementById('nomeLaboratorio').textContent = `${this.laboratorio.nome} - ${this.laboratorio.cnes}`;
        } catch (error) {
            console.error('Erro ao carregar laboratório:', error);
        }

        // Carregar resultado do exame e requisição
        await Promise.all([
            this.carregarResultadoExame(),
            this.carregarRequisicaoExame()
        ]);
    }

    async carregarResultadoExame() {
        try {
            this.resultadoExame = await BuscarResultadoExame(this.protocoloExame);
            this.exibirDados();
        } catch (error) {
            console.error('Erro ao carregar resultado:', error);
            // Não mostrar alerta aqui, pois pode não ter resultado ainda
        }
    }

    async carregarRequisicaoExame() {
        try {
            this.requisicaoExame = await ListarRequisicaoExame(this.protocoloExame);
            this.exibirStatus();
        } catch (error) {
            console.error('Erro ao carregar requisição:', error);
        }
    }

    initializeEventListeners() {
        document.getElementById('btnEmitir')?.addEventListener('click', () => {
            this.emitirLaudo();
        });
    }

    async emitirLaudo() {
        if (confirm('Deseja realmente emitir este laudo?')) {
            try {
                await EmitirResultadoExame(this.protocoloExame);
                alert('Laudo emitido com sucesso!');
                // Recarregar dados
                await Promise.all([
                    this.carregarResultadoExame(),
                    this.carregarRequisicaoExame()
                ]);
            } catch (error) {
                alert('Erro ao emitir laudo');
            }
        }
    }

    exibirStatus() {
        if (!this.requisicaoExame) return;

        document.getElementById('statusResultado').textContent = this.requisicaoExame.status || 'N/A';

        // Mostrar botão de emitir apenas se o status for SALVO
        if (this.requisicaoExame.status === 'SALVO') {
            document.getElementById('btnEmitir').style.display = 'inline-block';
        } else {
            document.getElementById('btnEmitir').style.display = 'none';
        }
    }

    exibirDados() {
        if (!this.resultadoExame) return;

        // Informações básicas
        if (this.resultadoExame.data_exame_recebido) {
            const data = new Date(this.resultadoExame.data_exame_recebido);
            document.getElementById('recebidoEm').textContent = data.toLocaleDateString('pt-BR');
        }

        if (this.resultadoExame.data_emissao_laudo) {
            const dataEmissao = new Date(this.resultadoExame.data_emissao_laudo);
            document.getElementById('dataEmissao').textContent = dataEmissao.toLocaleDateString('pt-BR');
            document.getElementById('dataEmissaoContainer').style.display = 'block';
        }

        // Avaliação Pré-Analítica
        this.exibirLista('amostraRejeitada', this.resultadoExame.avaliacao_pre_analitica?.rejeicao_amostra);
        this.exibirLista('epiteliosRepresentados', this.resultadoExame.avaliacao_pre_analitica?.epitelios_representados);

        // Adequação do Material
        if (this.resultadoExame.adequabilidade_material?.satisfatoria) {
            document.getElementById('adequacaoSatisfatoria').textContent = '✓ Satisfatória';
        } else {
            this.exibirLista('adequacaoInsatisfatoria', this.resultadoExame.adequabilidade_material?.insatisfatoria);
        }

        // Diagnóstico Descritivo
        if (this.resultadoExame.diagnostico_descritivo?.dentro_limites_normalidade) {
            document.getElementById('dentroLimitesNormalidade').textContent = '✓ Dentro dos Limites da Normalidade no material examinado';
        }
        this.exibirLista('alteracoesBenignas', this.resultadoExame.diagnostico_descritivo?.alteracoes_benignas);

        // Microbiologia
        this.exibirLista('microorganismos', this.resultadoExame.microbiologia?.microorganismos);
        
        this.exibirItem('celulasAtipicasEscamosas', this.resultadoExame.microbiologia?.celulas_atipicas_escamosas);
        this.exibirItem('celulasAtipicasGlandulares', this.resultadoExame.microbiologia?.celulas_atipicas_glandulares);
        this.exibirItem('celulasAtipicasOrigemIndefinida', this.resultadoExame.microbiologia?.celulas_atipicas_origem_indefinida);
        
        this.exibirLista('atipiasEscamosas', this.resultadoExame.microbiologia?.atipias_escamosas);
        this.exibirLista('atipiasGlandulares', this.resultadoExame.microbiologia?.atipias_glandulares);

        // Outras informações
        if (this.resultadoExame.outras_neoplasias_malignas) {
            document.getElementById('outrasNeoplasias').textContent = this.resultadoExame.outras_neoplasias_malignas;
        } else {
            document.getElementById('outrasNeoplasiasContainer').style.display = 'none';
        }

        if (this.resultadoExame.presenca_celulas_endometriais) {
            document.getElementById('celulasEndometriais').textContent = '✓ Presença de células endometriais (na pós-menopausa ou acima de 40 anos, fora do período menstrual)';
        } else {
            document.getElementById('celulasEndometriais').style.display = 'none';
        }

        document.getElementById('screeningCitotecnico').textContent = this.resultadoExame.screening_citotecnico || 'N/A';
        document.getElementById('responsavelRegistro').textContent = this.resultadoExame.registro_responsavel || 'N/A';

        if (this.resultadoExame.observacoes) {
            document.getElementById('observacoes').textContent = this.resultadoExame.observacoes;
        } else {
            document.getElementById('observacoesContainer').style.display = 'none';
        }
    }

    exibirLista(elementId, lista) {
        const element = document.getElementById(elementId);
        if (!element || !lista || lista.length === 0) {
            if (element) element.style.display = 'none';
            return;
        }

        element.innerHTML = '';
        lista.forEach(item => {
            const div = document.createElement('div');
            div.textContent = `• ${item}`;
            div.className = 'resultado-item';
            element.appendChild(div);
        });
        element.style.display = 'block';
    }

    exibirItem(elementId, valor) {
        const element = document.getElementById(elementId);
        if (!element || !valor) {
            if (element) element.style.display = 'none';
            return;
        }

        element.textContent = valor;
        element.style.display = 'block';
    }
}

document.addEventListener('DOMContentLoaded', () => {
    window.visualizacaoResultadoExame = new VisualizacaoResultadoExame();
}); 