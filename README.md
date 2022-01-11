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
curl -X GET 'http://localhost:8888/profile/id/:id' 
--header 'id: XXXXX' 

// pegar Profile Companie
curl -X GET 'http://localhost:8888/profileCompanie/id/:id' 
--header 'id: XXXXX' 

// Criar Profile
curl -X GET 'http://localhost:8888/profile/name/:name/email/:email/password/:password' 
--header 'name: XXXXX' 
--header 'email: XXXXX' 
--header 'password: XXXXX' 


// Criar Profile Companie
curl -X GET 'http://localhost:8888/profileCompanie/name/:name/email/:email/password/:password' 
--header 'name: XXXXX'
--header 'email: XXXXX' 
--header 'password: XXXXX'

// Inserir informações de trabalho no perfil Profile 
curl -X GET 'http://localhost:8888/profileCompanie/id/:id/job/:job/message/:message' \
--header 'id: XXXXX' 
--header 'job: X-X-XX-X' //sendo X um número, que corresponde ao id do job
--header 'message: XXXXX' 


// Inserir informações de trabalho no perfil Profile Companie
curl -X GET 'http://localhost:8888/profileCompanie/companieId/:companieId/job/:job/message/:message' 
--header 'id: XXXXX' 
--header 'job: X-X-XX-X' //sendo X um número, que corresponde ao id do job
--header 'message: XXXXX' 


```