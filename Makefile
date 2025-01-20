docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

migration-down:
	migrate -source "file://database/migrations" -database "mysql://admin:1234@tcp(127.0.0.1:3306)/be103" down

migration-up:
	migrate -source "file://database/migrations" -database "mysql://admin:1234@tcp(127.0.0.1:3306)/be103" up

	//migrate -source "file://database/migrations" -database "mysql://root:my-secret-pw@tcp(127.0.0.1:3306)/be103" up