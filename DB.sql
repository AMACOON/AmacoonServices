CREATE DATABASE dbcatclubsystem;


CREATE USER catclubsystem WITH PASSWORD '2010amacoon2010';



GRANT ALL PRIVILEGES ON DATABASE dbcatclubsystem TO catclubsystem;



-- Concede permissões de uso no esquema 'public'
GRANT USAGE ON SCHEMA public TO catclubsystem;

-- Concede todas as permissões no esquema 'public'
GRANT ALL PRIVILEGES ON SCHEMA public TO catclubsystem;

-- Concede permissões para criar, modificar e excluir tabelas no esquema 'public'
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO catclubsystem;

-- Para permitir que o usuário crie novas tabelas no futuro
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO catclubsystem;


--amacoon013  postgres   