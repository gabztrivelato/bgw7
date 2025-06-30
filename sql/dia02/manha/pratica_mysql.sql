
/* Exibir o título e o nome do gênero de todas as séries. 
SELECT s.title, g.name FROM series s INNER JOIN genres g ON s.genre_id = g.id;
*/

/* Mostre o título dos episódios, o nome e o sobrenome dos atores que trabalham em cada episódio.
SELECT e.title,
 a.first_name,
 a.last_name 
FROM actors a 
INNER JOIN actor_episode 
ON actor_episode.actor_id = a.id 
INNER JOIN episodes e 
ON actor_episode.episode_id = e.Id; */

/* Mostre o título de todas as séries e o número total de temporadas de cada série.
SELECT s.title, 
COUNT(*) AS total 
FROM series s 
INNER JOIN seasons 
ON s.id = seasons.serie_id 
GROUP BY s.title;
*/

/*
Mostre o nome de todos os gêneros e o número total de filmes de cada gênero, desde que seja maior ou igual a 3.
SELECT g.name, 
COUNT(*) as total 
FROM movies m 
INNER JOIN genres g 
ON m.genre_id = g.id 
GROUP BY g.name 
HAVING total >= 3;
*/

/*Mostre apenas o nome e o sobrenome dos atores que trabalharam em todos os filmes de Guerra nas Estrelas e não os repita. */

SELECT a.first_name,
 a.last_name 
FROM actors a 
INNER JOIN actor_movie 
ON a.id = actor_movie.actor_id 
INNER JOIN movies m
ON m.id = actor_movie.movie_id
WHERE m.title LIKE '%La Guerra de las galaxias%'
GROUP BY a.first_name, a.last_name;
*/
