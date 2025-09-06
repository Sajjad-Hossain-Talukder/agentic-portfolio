build: 
	docker build -t chatbot-backend -f docker/Dockerfile .

run: 
	docker run -d -p 8082:8082 --name chatbot-back chatbot-backend

stop: 
	docker stop chatbot-back

logs: 
	docker logs chatbot-back

remove: 
	docker rm chatbot-back

run-server: 
	go run src/main.go

.PHONY: create-docker-container run-docker run-server