start: up
    go run main.go
up: 
    docker-compose up --build -d

down: 
    docker-compose down

db-init:
    docker exec -i mysql-db mysql -u root -proot fantasy_products < ./docs/db/mysql/database.sql
