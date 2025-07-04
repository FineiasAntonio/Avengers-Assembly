CREATE TABLE IF NOT EXISTS endereco (
    endereco_id VARCHAR(10) PRIMARY KEY,
    logradouro VARCHAR(50) NOT NULL,
    numero VARCHAR (10) NOT NULL,
    complemento VARCHAR(20),
    bairro VARCHAR(20),
    codmunicipio INTEGER,
    municipio VARCHAR(20) NOT NULL,
    uf VARCHAR(2) NOT NULL,
    cep VARCHAR(8) NOT NULL,
    pontoreferencia VARCHAR(30)
);

CREATE TABLE IF NOT EXISTS unidade_saude (
    cnes VARCHAR(15) PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    cnpj VARCHAR(14) UNIQUE NOT NULL,
    endereco VARCHAR(10) UNIQUE NOT NULL,
    telefone VARCHAR(13) NOT NULL,

    FOREIGN KEY (endereco) REFERENCES endereco(endereco_id)
);

CREATE TABLE IF NOT EXISTS laboratorio (
    cnes VARCHAR(15) PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    cnpj VARCHAR(14) UNIQUE NOT NULL,
    endereco VARCHAR(10) UNIQUE NOT NULL,
    telefone VARCHAR(13) NOT NULL,

    FOREIGN KEY (endereco) REFERENCES endereco(endereco_id)
);

CREATE TABLE IF NOT EXISTS usuario (
    registro VARCHAR(10) PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    cpf VARCHAR(11) UNIQUE NOT NULL,
    email VARCHAR(50),
    telefone VARCHAR(13) NOT NULL,
    senha VARCHAR(255) NOT NULL,
    unidadesaude VARCHAR(20),
    laboratorio VARCHAR(20), 
    permissao VARCHAR (20) NOT NULL,
    primeiroacesso BOOLEAN NOT NULL,

    FOREIGN KEY (unidadesaude) REFERENCES unidade_saude(cnes),
    FOREIGN KEY (laboratorio) REFERENCES laboratorio(cnes)
);

CREATE TABLE IF NOT EXISTS paciente (
    cartaosus VARCHAR(10) PRIMARY KEY,
    prontuario VARCHAR(30) UNIQUE NOT NULL,
    nome VARCHAR(50) NOT NULL,
    nomemae VARCHAR(50),
    cpf VARCHAR(11) UNIQUE NOT NULL,
    datanascimento TIMESTAMP NOT NULL,
    idade INTEGER NOT NULL,
    raca  VARCHAR(10),
    nacionalidade VARCHAR(20) NOT NULL,
    escolaridade VARCHAR(30),
    telefone VARCHAR(13) NOT NULL,
    endereco VARCHAR(10) NOT NULL,
    senha VARCHAR(255) NOT NULL,
    primeiroacesso BOOLEAN NOT NULL,

    FOREIGN KEY (endereco) REFERENCES endereco(endereco_id)
);

CREATE TABLE IF NOT EXISTS agendamento_exame (
    protocolo VARCHAR(10) PRIMARY KEY,
    unidade VARCHAR(15) NOT NULL,
    paciente VARCHAR(10) NOT NULL,
    profissional VARCHAR(10) NOT NULL,
    data TIMESTAMP NOT NULL,

    FOREIGN KEY (unidade) REFERENCES unidade_saude(cnes),
    FOREIGN KEY (paciente) REFERENCES paciente(cartaosus),
    FOREIGN KEY (profissional) REFERENCES usuario(registro)
);

CREATE TABLE IF NOT EXISTS requisicao_exame (
    protocolo VARCHAR(30) PRIMARY KEY,
    paciente VARCHAR(10) NOT NULL,
    motivoexame VARCHAR(255),
    fezexamepreventivo BOOLEAN,
    anoultimoexame VARCHAR(4),
    usadiu VARCHAR(10),
    estagravida VARCHAR(10),
    usaanticoncepcional VARCHAR(10),
    usahormoniomenopausa VARCHAR(10),
    fezradioterapia VARCHAR(10),
    dataultimamenstruacao TIMESTAMP,
    sangramentoaposrelacoes VARCHAR(10),
    sangramentoaposmenopausa VARCHAR(10),
    inspecaocolo VARCHAR(10),
    sinaisdst BOOLEAN,
    datacoleta TIMESTAMP NOT NULL,
    responsavel VARCHAR(10) NOT NULL,
    resultado VARCHAR(10),
    status VARCHAR(20) NOT NULL,

    FOREIGN KEY (paciente) REFERENCES paciente(cartaosus),
    FOREIGN KEY (responsavel) REFERENCES usuario(registro)
);