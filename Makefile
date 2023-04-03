build:
	docker build -o one-lab .
run: build
	docker run -d --name task1 -p 9090:9090 one-lab