# go-json

Clone the repository.

```sh
git clone https://github.com/chrisbradleydev/go-json.git .
```

Build Docker image.

```sh
docker build -t go-json .
```

Run Docker container.

```sh
docker run -v $(pwd):/app -v /app/tmp go-json
```
