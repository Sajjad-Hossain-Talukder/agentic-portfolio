build: 
	docker build -t chatbot .

run: 
	docker run -d -p 5173:80 --name chatbot-ui chatbot

stop: 
	docker stop chatbot-ui

logs: 
	docker logs chatbot-ui

remove: 
	docker rm chatbot-ui

run-local: 
	yarn dev

.PHONY: build 

