CREATE TABLE IF NOT EXISTS clientes (
    id smallint PRIMARY KEY,
    nome VARCHAR(23),
    limite integer,
    saldo integer
);

CREATE TABLE IF NOT EXISTS transactions (
    clientId smallint references clientes(id),
    valor integer,
    tipo char(1),
    descricao varchar(10),
    realizada_em TIMESTAMP WITH TIME ZONE DEFAULT timezone('America/Sao_Paulo'::text, now())
);

INSERT INTO clientes (id, nome, limite, saldo)
VALUES
    (1, 'o barato sai caro', 1000 * 100, 0),
    (2, 'zan corp ltda', 800 * 100, 0),
    (3, 'les cruders', 10000 * 100, 0),
    (4, 'padaria joia de cocaia', 100000 * 100, 0),
    (5, 'kid mais', 5000 * 100, 0);


CREATE INDEX IF NOT EXISTS idx_clientes_id ON clientes (id);
CREATE INDEX IF NOT EXISTS idx_transactions_realizada_em ON transactions (realizada_em);