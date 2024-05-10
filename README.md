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



## Client 및 Server생성
### registrar server 생성
```shell
go run cmd/server/vdr/registrar/main.go
```
### resolver server 생성
```shell
go run cmd/server/vdr/resolver/main.go
```

### Issuer server 생성
```shell
go run cmd/server/issuer.go
```

### Verifier server 생성
```shell
go run cmd/server/verifier.go
```

### main client 생성
```shell
go run cmd/client/holder/main.go
```
## main client에서 사용 가능한 명령어

| 명령어 | 설명                          |
|--------|-------------------------------|
| (d)    | DID 및 DID Document 등록      |
| (v)    | VC 요청                       |
| (p)    | VP 생성                       |
| (s)    | VP 전달 및 검증 요청          |
| (r)    | DID Resolver 실행             |
| (q)    | 프로그램 종료                 |

**입력:** _원하는 명령어를 입력하고 Enter 키를 누릅니다._

## 실행 시나리오

1. (d) DID 및 DID Document 등록 : Holder DID 및 DID Document 생성

2. (v) Holder -> Issuer VC 발급 요청 (Issuer DID 및 DID Document 생성)

3. (p) Holder VP 생성

4. (s) Holder -> Issuer VP 전달 및 검증 요청 (Verifier DID 및 DID Document 생성)

5. (r) Holder -> VDR(Resolver) Holder의 DID로 DID Document 요청

6. (q) 프로그램 종료
