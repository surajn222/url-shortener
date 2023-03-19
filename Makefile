build:
	-echo "Docker build"
	-docker build . -t surajn222/url-shortener:latest

push:
	-docker push surajn222/url-shortener:latest

run: build push
	-echo "Docker run"
	-cd docker && docker-compose down
	-cd docker && docker-compose rm -f
	-cd docker && docker-compose pull   
	-cd docker && docker-compose up -d

down:
	-cd docker && docker-compose down


make run-latest: build push run
