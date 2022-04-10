# AluraChallenge
#alurachallengeback2

Project created in Go, using GIN to expose /receitas and /despesas endpoints.
Needs MySQL database to work and env.yaml with [username, password, name, host, port] db credentials.


build image docker
````shell
docker build -t juliazuin/alurachallenge .
````

docker run iterative
````shell
docker run -it -p 8080:8080 juliazuin/alurachallenge
````

docker run background
````shell
docker run -d  -p 8080:8080 juliazuin/alurachallenge
````