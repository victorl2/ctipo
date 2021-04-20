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
   add constraint UK_oc81nc6ig2dvu5ydf5kvak78d unique (prontuario);