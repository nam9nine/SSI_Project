# SSI 기반 대학교 MT(멤버십 트레이닝) 참여 인증 시스템


**순서대로 실행**
## registrar 서버 생성
```shell
go run cmd/server/vdr/registrar/main.go
```
## resolver 서버 생성
```shell
go run cmd/server/vdr/resolver/main.go
```
## holder client 생성
```shell
go run cmd/client/holder/main.go
```

