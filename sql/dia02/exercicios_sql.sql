
/*
Selecciona o nome, a posição e a localização dos departamentos onde os vendedores trabalham.
SELECT f.name,
	f.posto,
    d.localidad
FROM funcionarios f
INNER JOIN departamento d
ON f.depto_nro = d.depto_nro
*/

/*
Mostra os departamentos com mais de cinco empregados.
SELECT d.nombre_depto, COUNT(f.depto_nro) AS total FROM departamentos d INNER JOIN funcionarios f ON d.depto_nro = f.depto_nro GROUP BY f.depto_nro, d.nombre_depto HAVING total > 5
/*
SELECT f.nome,
	f.salario,
    d.nombre_depto
FROM funcionarios f
INNER JOIN departamentos d
ON f.depto_nro = d.depto_nro
WHERE f.posto = (SELECT posto FROM funcionarios WHERE nome = 'Mito' AND sobrenome = 'Barchuk');
*/

/*
Mostra os detalhes dos empregados que trabalham no departamento de contabilidade, ordenados por nome.

SELECT f.nome,
	f.sobrenome,
	f.posto,
    f.data_alta,
    f.salario,
    f.comissao
FROM funcionarios f 
INNER JOIN departamentos d
ON f.depto_nro = d.depto_nro
WHERE d.nombre_depto = 'Contabilidad'
ORDER BY f.nome;

/*
Mostra o nome do empregado com o salário mais baixo.

SELECT nome 
FROM funcionarios
WHERE salario = (SELECT min(salario) FROM funcionarios);
*/

/*
Mostra os detalhes do empregado com o salário mais alto no departamento de "Vendas".
SELECT 
	nome,
    sobrenome,
    posto,
    data_alta,
    salario,
    comissao
FROM funcionarios 
WHERE salario = (SELECT max(f.salario) FROM funcionarios f INNER JOIN departamentos d ON f.depto_nro = d.depto_nro WHERE d.nombre_depto = 'Vendas' )
*/


