CREATE TABLE client (	
	client_id int not null auto_increment,
	first_name varchar(255) not null,
    last_name varchar(255) not null,
    birth_date date not null,
    province varchar(255) not null,
    city varchar(255) not null,
    plan_id int not null,
    PRIMARY KEY (client_id),
	FOREIGN KEY (plan_id) REFERENCES plan(plan_id)
);

CREATE TABLE plan (
	plan_id int not null auto_increment,
    speed int not null,
    price float not null,
    discount int not null,
    PRIMARY KEY (plan_id)
);


INSERT INTO plan (speed, price, discount) VALUES
(100, 99.90, 10),
(200, 129.90, 15),
(300, 149.90, 20),
(500, 199.90, 25),
(1000, 299.90, 30);

INSERT INTO client (first_name, last_name, birth_date, province, city, plan_id) VALUES
('Lucas', 'Silva', '1992-05-14', 'São Paulo', 'Campinas', 1),
('Ana', 'Souza', '1987-11-23', 'Rio de Janeiro', 'Niterói', 3),
('Carlos', 'Oliveira', '1990-02-08', 'Minas Gerais', 'Belo Horizonte', 2),
('Juliana', 'Pereira', '2000-07-19', 'Paraná', 'Curitiba', 1),
('Felipe', 'Rodrigues', '1985-09-03', 'Bahia', 'Salvador', 4),
('Mariana', 'Costa', '1993-03-27', 'Rio Grande do Sul', 'Porto Alegre', 5),
('Bruno', 'Martins', '1991-12-01', 'Pernambuco', 'Recife', 3),
('Camila', 'Almeida', '1996-08-11', 'São Paulo', 'São Paulo', 2),
('Renato', 'Ferreira', '1989-04-17', 'Santa Catarina', 'Florianópolis', 4),
('Patrícia', 'Lima', '2002-10-06', 'Goiás', 'Goiânia', 5);


SELECT c.first_name, c.last_name, p.price, p.speed FROM client c INNER JOIN plan p ON c.plan_id = p.plan_id;
SELECT c.city, p.price FROM client c INNER JOIN plan p ON c.plan_id = p.plan_id;
SELECT c.city, p.price FROM client c INNER JOIN plan p ON c.plan_id = p.plan_id WHERE p.price > 100;
SELECT speed, discount FROM plan ORDER BY DISCOUNT DESC;
SELECT c.province, p.speed FROM client c INNER JOIN plan p ON c.plan_id = p.plan_id;
SELECT c.province, p.speed FROM client c INNER JOIN plan p ON c.plan_id = p.plan_id ORDER BY speed DESC;
SELECT * FROM client c INNER JOIN plan p ON c.plan_id = p.plan_id WHERE c.province = 'São Paulo';
SELECT * FROM client c INNER JOIN plan p ON c.plan_id = p.plan_id WHERE c.birth_date BETWEEN '1990-01-01' AND '2000-12-31';
SELECT * FROM client c INNER JOIN plan p ON c.plan_id = p.plan_id WHERE c.province = 'São Paulo' AND price > 100;
SELECT * FROM client c INNER JOIN plan p ON c.plan_id = p.plan_id WHERE c.province = 'São Paulo' ORDER BY discount DESC