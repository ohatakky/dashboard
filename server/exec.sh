function run() {
    watchexec -r go run cmd/api/main.go
}

function deploy() {
    gcloud app deploy
}

$1
