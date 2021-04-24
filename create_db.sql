create table Paciente (
    id int8 not null,
    cpf varchar(255),
    data_nascimento date not null,
    nome varchar(255) not null,
    prontuario varchar(255) not null,
    sexo int4,
    id_endereco int8,
    primary key (id)
);

alter table if exists Paciente
   add constraint UNIQUE_PRONTUARIO unique (prontuario);

create table Endereco (
    id int8 not null,
    bairro varchar(255),
    cep varchar(255),
    cidade varchar(255),
    complemento varchar(255),
    numero int4,
    ponto_referencia varchar(255),
    rua varchar(255),
    primary key (id)
);

alter table if exists Paciente
   add constraint FK_ENDERECO_DO_PACIENTE
   foreign key (id_endereco)
   references Endereco;