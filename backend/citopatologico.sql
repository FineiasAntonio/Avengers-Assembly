CREATE TABLE Paciente (
    CartaoSUS VARCHAR(15) PRIMARY KEY NOT NULL,
    Prontuario VARCHAR(10),
    NomeCompleto VARCHAR(100) NOT NULL,
    NomeMae VARCHAR(100) NOT NULL,
    CPF VARCHAR(11),
    DataNascimento DATE NOT NULL,
    Idade INT,
    Raca VARCHAR(10),
    Nacionalidade VARCHAR(30),
    Escolaridade VARCHAR(50)
);

CREATE TABLE EnderecoPaciente (
    CartaoSUS VARCHAR(15) PRIMARY KEY,
    Logradouro VARCHAR(100),
    Numero VARCHAR(6),
    Complemento VARCHAR(50),
    CodMunicipio VARCHAR(7),
    Municipio VARCHAR(50),
    Bairro VARCHAR(50),
    UF CHAR(2),
    CEP VARCHAR(8),
    DDD VARCHAR(2),
    Telefone VARCHAR(9),
    PontoReferencia VARCHAR(100),

    FOREIGN KEY (CartaoSUS) REFERENCES Paciente(CartaoSUS)
);

CREATE TABLE Exame (
    Protocolo VARCHAR(14) PRIMARY KEY,
    CartaoSUS VARCHAR(15) NOT NULL,
    UF CHAR(2),
    CNES_UnidadeSaude VARCHAR(7),
    UnidadeSaude VARCHAR(100),
    Municipio VARCHAR(50),
    MotivoExame VARCHAR(15),
    FezExamePreventivo BOOLEAN NOT NULL,
    AnoUltimoExame VARCHAR(4),
    UsaDIU BOOLEAN NOT NULL,
    EstaGravida BOOLEAN NOT NULL,
    UsaAnticoncepcional BOOLEAN NOT NULL,
    UsaHormônioMenopausa BOOLEAN NOT NULL,
    FezRadioterapia BOOLEAN NOT NULL,
    DataUltimaMenstruacao DATE,
    SangramentoAposRelacoes BOOLEAN,
    SangramentoAposMenopausa BOOLEAN,
    InspecaoColo VARCHAR(20),
    SinaisDST BOOLEAN,
    DataColeta DATE NOT NULL,
    Responsavel VARCHAR(100) NOT NULL,
    
    FOREIGN KEY (CartaoSUS) REFERENCES Paciente(CartaoSUS)
);
