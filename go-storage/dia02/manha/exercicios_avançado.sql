
CREATE SCHEMA biblioteca;

CREATE TABLE estudante (
	id_leitor INT NOT NULL AUTO_INCREMENT,
    nome VARCHAR(50) NOT NULL,
    sobrenome VARCHAR(100) NOT NULL,
    endereco VARCHAR(255) NOT NULL,
    carreira VARCHAR(50) NOT NULL,
	idade INT NOT NULL,
    PRIMARY KEY (id_leitor)
);

CREATE TABLE livro(
	id_livro INT NOT NULL AUTO_INCREMENT,
    titulo VARCHAR(100) NOT NULL,
    editora VARCHAR(100) NOT NULL,
    genero VARCHAR(50) NOT NULL,
    PRIMARY KEY (id_livro)
);

CREATE TABLE autor(
	id_autor INT NOT NULL AUTO_INCREMENT,
    nome VARCHAR(100) NOT NULL,
    nacionalidade varchar(100) NOT NULL,
    PRIMARY KEY (id_autor)
);

CREATE TABLE livro_autor(
	id_autor INT NOT NULL,
    id_livro INT NOT NULL,
    FOREIGN KEY (id_autor) REFERENCES autor(id_autor),
    FOREIGN KEY (id_livro) REFERENCES livro(id_livro)
);

CREATE TABLE emprestimo (
	id_leitor INT NOT NULL,
    id_livro INT NOT NULL,
    data_emprestimo DATE NOT NULL,
    data_devolucao DATE NOT NULL,
    devolvido BOOL NOT NULL,
    FOREIGN KEY (id_leitor) REFERENCES estudante(id_leitor),
    FOREIGN KEY (id_livro) REFERENCES livro(id_livro)
);

INSERT INTO autor (id_autor, nome, nacionalidade) VALUES
(1, 'Gabriel García Márquez', 'Colombiana'),
(2, 'Isabel Allende', 'Chilena'),
(3, 'Jorge Luis Borges', 'Argentina'),
(4, 'Mario Vargas Llosa', 'Peruana'),
(5, 'Paulo Coelho', 'Brasileira');

INSERT INTO livro (id_livro, titulo, editora, genero) VALUES
(1, 'Cien Años de Soledad', 'Sudamericana', 'Literatura'),
(2, 'La Casa de los Espíritus', 'Plaza & Janés', 'Literatura'),
(3, 'El Aleph', 'Emecé', 'Cuentos'),
(4, 'Conversación en La Catedral', 'Alfaguara', 'Novela'),
(5, 'El Alquimista', 'Rocco', 'Autoayuda');


INSERT INTO livro_autor (id_autor, id_livro) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5);

INSERT INTO estudante (id_leitor, nome, sobrenome, endereco, carreira, idade) VALUES
(1, 'Juan', 'Pérez', 'Calle 123', 'Ingeniería', 21),
(2, 'María', 'Gómez', 'Avenida 456', 'Derecho', 22),
(3, 'Luis', 'Martínez', 'Calle 789', 'Medicina', 23),
(4, 'Ana', 'Rodríguez', 'Boulevard 101', 'Arquitectura', 20),
(5, 'Pedro', 'López', 'Pasaje 202', 'Psicología', 24);

INSERT INTO emprestimo (id_leitor, id_livro, data_emprestimo, data_devolucao, devolvido) VALUES
(1, 1, '2025-06-01', '2025-06-15', 1),
(2, 2, '2025-06-02', '2025-06-16', 0),
(3, 3, '2025-06-03', '2025-06-17', 1),
(4, 4, '2025-06-04', '2025-06-18', 0),
(5, 5, '2025-06-05', '2025-06-19', 0);

/*
1. Enumera os dados dos autores
SELECT 
    ROW_NUMBER() OVER (ORDER BY id_autor) AS numero,
    id_autor,
    nome,
    nacionalidade
FROM 
    autor;

*/

/*
2. Indica o nome e a idade dos alunos
SELECT nome,
	idade
FROM estudante
*/

/*
3. Que alunos pertencem ao curso de informática?
SELECT * 
FROM estudante
WHERE carrera = 'Informática';
*/

/*
4. Quais são os autores de nacionalidade francesa ou italiana?
SELECT * FROM autor WHERE nacionalidade = 'francesa' OR nacionalidade = 'italiana'
*/

/*
5. Quais são os livros que não são da área da internet
SELECT * FROM livros WHERE genero <> 'internet'
*/

/*
6. Enumera os livros publicados pela Salamandra
SELECT
    ROW_NUMBER() OVER (ORDER BY id_livro) AS numero,
    id_livro,
    titulo,
    editora,
    genero
FROM 
    livro
WHERE editora = 'Salamandra';
*/

/*
7. Enumera os nomes dos alunos cuja idade é superior à média
SELECT
	ROW_NUMBER() OVER (ORDER BY nome) AS numero,
    nome, 
    sobrenome,
    idade
FROM estudante
WHERE idade > (SELECT AVG(idade) FROM estudante)
*/

/*
8. Enumere os nomes dos alunos cujo sobrenome começa com a letra G
SELECT 
	ROW_NUMBER() OVER (ORDER BY nome) AS numero,
    nome,
    sobrenome,
    idade
FROM estudante
WHERE sobrenome LIKE 'G%';
*/

/*
9. Faz uma lista de autores do livro "O universo: guia de viagem"
SELECT
    nome
FROM autor
INNER JOIN livro_autor
ON autor.id_autor = livro_autor.id_autor
INNER JOIN livro
ON livro_autor.id_livro = livro.id_livro
WHERE livro.titulo = 'O Universo: Guia de Viagem';
*/

/
/*
10. Que livros foi emprestado ao leitor "Filippo Galli"
SELECT 
	l.titulo,
    es.nome
FROM livro l
INNER JOIN emprestimo e
ON l.id_livro = e.id_livro
INNER JOIN estudante es
ON e.id_leitor = es.id_leitor
WHERE es.nome = 'Filippo' AND es.sobrenome = 'Galli';
*/

/*
11. Indica o nome do aluno mais novo
SELECT nome
FROM estudante
WHERE idade = (SELECT min(idade) FROM estudantes)
*/

/*
12. Enumera os nomes dos alunos que foi emprestado os livros da base de dados
SELECT 
	ROW_NUMBER() OVER(ORDER BY e.nome),
    nome
FROM estudante
INNER JOIN emprestimo
ON estudante.id_leitor = emprestimo.id_leitor
INNER JOIN livro
ON emprestimo.id_livro = livro.id_livro
WHERE genero = 'Base de Dados'
*/

/*
13. Enumero os livros que pertecem a autora J.K. Rowling
SELECT 
     ROW_NUMBER() OVER(ORDER BY id_livro),
     titulo
FROM livro
INNER JOIN livro_autor
ON livro.id_livro = livro_autor.id_livro
INNER JOIN autor
ON livro_autor.id_autor = autor.id_autor
WHERE autor.nome = 'J.K. Rowling'
*/

/*
14. Enumera os titulos dos livros que deviam ser devolvidos em 16/07/2021
SELECT
	ROW_NUMBER() OVER(ORDER BY titulo),
    titulo
FROM livro
INNER JOIN emprestimo
ON livro.id_livro = emprestimo.id_livro
WHERE data_devolucao = '2021-07-16'
*/









