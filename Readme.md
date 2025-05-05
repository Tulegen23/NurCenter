README.md

NurCenter Microservices

A microservices-based productivity application with User and Productivity services.

Setup





Install Dependencies:

go mod tidy



Setup PostgreSQL:





Create databases: nurcenter_users and nurcenter_productivity.



Update .env with your credentials.



Run Migrations:

migrate -path migrations -database "postgresql://postgres:yourpassword@localhost:5432/nurcenter_users?sslmode=disable" up
migrate -path migrations -database "postgresql://postgres:yourpassword@localhost:5432/nurcenter_productivity?sslmode=disable" up



Run with Docker:

docker-compose up --build



Run Tests:

go test ./tests



Test with Postman:





Import postman/nurcenter_microservices_postman_collection.json.



Register and login to get a JWT token, then use it for Productivity Service endpoints.

Services





User Service: http://localhost:8081 (register, login, user profile)



Productivity Service: http://localhost:8080 (todos, notes, habits, goals, finances, reminders, dashboard)