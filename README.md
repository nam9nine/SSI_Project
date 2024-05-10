# SSI 기반 대학교 MT(멤버십 트레이닝) 참여 인증 시스템


## Server Configuration

| Server Type | IP Address  | Port  | Database Path            |
|-------------|-------------|-------|--------------------------|
| Registrar   | 127.0.0.1   | 50051 | N/A                      |
| Resolver    | 127.0.0.1   | 50052 | N/A                      |
| Holder      | 127.0.0.1   | 50053 | `./internal/db/holder`   |
| Issuer      | 127.0.0.1   | 50054 | `./internal/db/issuer`   |
| Verifier    | 127.0.0.1   | 50055 | `./internal/db/verifier` |

각 서버는 특정 포트에서 작동하며 해당되는 경우 전용 데이터베이스 경로를 가질 수 있습니다.


## 의존성 주입
```shell
go build
```
## registrar server 생성
```shell
go run cmd/server/vdr/registrar/main.go
```
## resolver server 생성
```shell
go run cmd/server/vdr/resolver/main.go
```

## Issuer server 생성
```shell
go run cmd/server/issuer.go
```

## Verifier server 생성
```shell
go run cmd/server/verifier.go
```

## main client 생성
```shell
go run cmd/client/holder/main.go
```

## 실행 시나리오
1. (d)DID 및 DID Document 등록 : Holder의 DID

