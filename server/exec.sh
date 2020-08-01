function run() {
    go run cmd/api/main.go
}

function deploy() {
    gcloud app deploy
}

$1
