INSERT INTO movies (
	title,
    rating,
    awards,
    release_date,
    length,
    genre_id
) 
VALUES(
	'Mission: Impossible',
    8.9,
    2,
    '1996-04-13',
    190,
    4
);

INSERT INTO genres(
	name,
    ranking,
    active
)
VALUES(
	'Dorama',
    13,
    1
);

UPDATE movies
SET genre_id = 14
WHERE id = 22;

UPDATE actors 
SET favorite_movie = 22
WHERE id = 1;

CREATE TEMPORARY TABLE movies_copy AS
SELECT * FROM movies;

DELETE FROM movies_copy
WHERE awards < 5;

SELECT * FROM genres 
INNER JOIN movies
ON movies.genre_id = genres.id;

SELECT * FROM actors
INNER JOIN movies
ON movies.id = actors.favorite_movie_id
WHERE movies.awards > 3;

CREATE INDEX index_nome ON movies (title);

/*
No banco de dados de movies, você notaria uma melhora significativa com a criação de índices? Analise e justifique sua resposta.
Sim, pois ao usar filtros que buscaria pelo nome, teria uma busca mais rápida e efetiva.
Em qual outra tabela você criaria um índice e por quê? Justifique sua resposta.
Criaria na coluna titulo na tabela de episodios, pois também ajudaria em um filtro que fosse buscar os episodios deixando mais rápido a busca.
*/
