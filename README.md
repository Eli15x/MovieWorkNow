# Movie Work Now

Movie Work Now é uma rede social para trabalho relativo a área de cinema, englobando:
Atuação ,roteirização ,filmagem (CameraMan) ,fotografia ,direção de arte, produtores ,Editores....


## Run
```sh
`go run main.go`
available at: http://localhost:1323
```

### Rotas

#### Check Valid Login
``` sh 
curl -X POST 'http://localhost:1323/profile/valid' 
 -H 'Content-Type: application/json'
 -d '{"email":"my_email","password":"my_password"}'
```

#### pegar Profile
``` sh 
curl -X GET 'http://localhost:1323/profile/id/:id' 
--header 'id: XXXXX' 
```

#### pegar Profile Companie
``` sh 
curl -X GET 'http://localhost:1323/profileCompanie/id/:id' 
--header 'id: XXXXX' 
```

#### Criar Profile
``` sh 
curl -X GET 'http://localhost:1323/profile/name/:name/email/:email/password/:password' 
--header 'name: XXXXX' 
--header 'email: XXXXX' 
--header 'password: XXXXX' 
```

#### Criar Profile Companie
``` sh 
curl -X GET 'http://localhost:1323/profileCompanie/name/:name/email/:email/password/:password' 
--header 'name: XXXXX'
--header 'email: XXXXX' 
--header 'password: XXXXX'
```

#### Inserir informações de trabalho no perfil Profile 
``` sh 
curl -X GET 'http://localhost:1323/profileCompanie/id/:id/job/:job/message/:message' \
--header 'id: XXXXX' 
--header 'job: X-X-XX-X' //sendo X um número, que corresponde ao id do job
--header 'message: XXXXX' 
```

#### Inserir informações de trabalho no perfil Profile Companie
``` sh 
curl -X GET 'http://localhost:1323/profileCompanie/companieId/:companieId/job/:job/message/:message' 
--header 'id: XXXXX' 
--header 'job: X-X-XX-X' //sendo X um número, que corresponde ao id do job
--header 'message: XXXXX' 

```