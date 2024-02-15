DROP TABLE IF EXISTS clientes;
DROP TABLE IF EXISTS transactions;

CREATE TABLE IF NOT EXISTS clientes (
    id serial PRIMARY KEY,
    nome VARCHAR(50),
    limite INT,
    saldo INT
);

CREATE INDEX IF NOT EXISTS idx_clientes_id ON clientes (id);

INSERT INTO clientes (nome, limite, saldo)
VALUES
    ('o barato sai caro', 1000 * 100, 0),
    ('zan corp ltda', 800 * 100, 0),
    ('les cruders', 10000 * 100, 0),
    ('padaria joia de cocaia', 100000 * 100, 0),
    ('kid mais', 5000 * 100, 0);

CREATE TABLE IF NOT EXISTS transactions (
    id serial PRIMARY KEY,
    clientId INT,
    valor INT,
    tipo varchar(1),
    descricao varchar(10),
    realizada_em TIMESTAMP WITH TIME ZONE DEFAULT timezone('America/Sao_Paulo'::text, now())
);

CREATE INDEX IF NOT EXISTS idx_transactions_clientId ON transactions (clientId);