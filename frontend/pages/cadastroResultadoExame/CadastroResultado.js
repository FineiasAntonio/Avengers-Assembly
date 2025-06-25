import { listarLaboratorio } from '../../api/unidadeApi.js';
import { pegarUnidadeUsuario } from '../../shared/gerenciador-permissoes.js';
import { CadastrarResultadoExame } from '../../api/cadastroApi.js';

class ResultadoExameForm {

    protocoloExame = '';
    laboratorio = {};
    recebidoEm = '';

    constructor() {
        this.initializeEventListeners();
        this.setupOutrasNeoplasiasToggle();

        const cnes = pegarUnidadeUsuario();
        listarLaboratorio(cnes).then(data => {
            this.laboratorio = data;

            document.getElementById('nomeLaboratorio').textContent = `${this.laboratorio.nome} - ${this.laboratorio.cnes}`; 
        });

        const urlParams = new URLSearchParams(window.location.search);
        this.protocoloExame = urlParams.get('protocolo');
        document.getElementById('numeroExame').textContent = this.protocoloExame;

        this.recebidoEm = new Date().toISOString().split('T')[0];
        document.getElementById('recebidoEm').value = this.recebidoEm;
    }

    initializeEventListeners() {
        document.getElementById('btnSalvarEmitir')?.addEventListener('click', () => {
            this.handleSubmit();
        });
    }

    setupOutrasNeoplasiasToggle() {
        const outrasNeoplasiasCheckbox = document.getElementById('outras-neoplasias-checkbox');
        const outrasNeoplasiasInput = document.getElementById('outras-neoplasias-input');
        
        if (outrasNeoplasiasCheckbox && outrasNeoplasiasInput) {
            outrasNeoplasiasCheckbox.addEventListener('change', () => {
                outrasNeoplasiasInput.style.display = outrasNeoplasiasCheckbox.checked ? 'inline' : 'none';
            });
        }
    }

    pegarValorInput(campo) {
        const element = document.querySelector(`[data-field="${campo}"]`);
        if (!element) return null;

        if (element.type === 'checkbox') {
            return element.checked;
        }
        return element.value;
    }

    pegarValorGrupos(grupo) {
        const checkboxes = document.querySelectorAll(`[data-group="${grupo}"]`);
        const selectedValues = [];
        
        checkboxes.forEach(checkbox => {
            if (checkbox.checked) {
                selectedValues.push(checkbox.dataset.value);
            }
        });
        
        return selectedValues;
    }

    pegarValorUnico(grupo) {
        const checkboxes = document.querySelectorAll(`[data-group="${grupo}"]`);
        
        for (const checkbox of checkboxes) {
            if (checkbox.checked) {
                return checkbox.dataset.value;
            }
        }
        
        return null;
    }

    generateFormJSON() {
        return {
            protocolo_exame: this.protocoloExame,
            laboratorio: this.laboratorio.cnes,
            data_exame_recebido: new Date(this.recebidoEm).toISOString(),

            avaliacao_pre_analitica: {
                rejeicao_amostra: this.pegarValorGrupos('amostraRejeitada'),
                epitelios_representados: this.pegarValorGrupos('epiteliosRepresentados')
            },

            adequabilidade_material: {
                satisfatoria: this.pegarValorInput('adequacaoSatisfatoria'),
                insatisfatoria: [
                    ...this.pegarValorGrupos('adequacaoInsatisfatoria'),
                    ...(this.pegarValorInput('adequacaoInsatisfatoriaOutros') ? [this.pegarValorInput('adequacaoInsatisfatoriaOutros')] : [])
                ]
            },

            diagnostico_descritivo: {
                dentro_limites_normalidade: this.pegarValorInput('dentroLimitesNormalidade'),
                alteracoes_benignas: [
                    ...this.pegarValorGrupos('alteracoesBenignas'),
                    ...(this.pegarValorInput('alteracoesBenignasOutras') ? [this.pegarValorInput('alteracoesBenignasOutras')] : [])
                ]
            },

            microbiologia: {
                microorganismos: [
                    ...this.pegarValorGrupos('microorganismos'),
                    ...(this.pegarValorInput('microorganismosOutros') ? [this.pegarValorInput('microorganismosOutros')] : [])
                ],
                celulas_atipicas_escamosas: this.pegarValorUnico('celulasAtipicasEscamosas'),
                celulas_atipicas_glandulares: this.pegarValorUnico('celulasAtipicasGlandulares'),
                celulas_atipicas_origem_indefinida: this.pegarValorUnico('celulasAtipicasOrigemIndefinida'),

                atipias_escamosas: this.pegarValorGrupos('atipiasEscamosas'),
                atipias_glandulares: this.pegarValorGrupos('atipiasGlandulares'),
                
            },

            outras_neoplasias_malignas: this.pegarValorInput('outrasNeoplasias'),
            presenca_celulas_endometriais: this.pegarValorInput('presencaCelulasEndometriais'),

            observacoes: this.pegarValorInput('observacoes'),
            screening_citotecnico: this.pegarValorInput('screeningCitotecnico'),
            registro_responsavel: this.pegarValorInput('responsavelRegistro'),
            
        };
    }

    async handleSubmit() {
        const formData = this.generateFormJSON();
        console.log(formData);
        
        CadastrarResultadoExame(formData);
    }
}

document.addEventListener('DOMContentLoaded', () => {
    window.resultadoExameForm = new ResultadoExameForm();
});
