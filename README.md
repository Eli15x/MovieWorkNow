# Movie Work Now

Movie Work Now é uma plataforma social para trabalho relativo a área de cinema, englobando:
Atuação ,roteirização ,filmagem (CameraMan) ,fotografia ,direção de arte, produtores ,Editores....


## Run
```sh
`go run main.go`
available at: http://localhost:8888
```

### Rotas
```sh
// pegar Profile
curl -X GET 'http://localhost:8888/profile/id/:id' \
--header 'ID: XXXXX' \

```