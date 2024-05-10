package main

import (
	"bufio"
	"fmt"
	"github.com/nam9nine/SSI_Project/config"
	"github.com/nam9nine/SSI_Project/pkg/actors"
	"log"
	"os"
)

func main() {
	hldr := new(actors.Holder)
	iss := new(actors.Issuer)

	// 설정 파일 로드
	cfg, err := config.LoadConfig("config/config.toml")
	if err != nil {
		log.Fatalf("설정 파일 로드 실패: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	var input rune

	for {
		fmt.Println("\n다음 중에서 선택해주세요:")
		fmt.Println("  (d) DID 및 DID Document 등록")
		fmt.Println("  (r) DID Resolver 실행")
		fmt.Println("  (v) VC 요청")
		fmt.Println("  (q) 프로그램 종료")
		fmt.Print("입력: ")

		input, _, err = reader.ReadRune()
		if err != nil {
			log.Fatalf("입력 읽기 실패: %v", err)
		}

		// 입력에 따른 조치
		switch input {
		case 'q':
			log.Println("프로그램을 종료합니다.")
			os.Exit(0)
		case 'd':
			log.Println("DID와 DID Document를 등록합니다.")
			hldr.InitHolder(cfg)
			log.Println("Holder DID 생성, DID Document 생성 및 VDR에 DID Document 등록 완료.")
			iss.InitIssuer(cfg)
			log.Println("Issuer DID 생성, DID Document 생성 및 VDR에 DID Document 등록 완료.")
		case 'r':
			log.Println("DID Resolver 실행합니다.")
			res, err := hldr.ResolveHolderDID()
			if err != nil {
				log.Printf("DID Resolver 실행 오류: %v", err)
				continue
			}
			log.Println("DID Resolver 성공")
			log.Println("--------DID Document----------")
			log.Println(res.DidDoc)
		case 'v':
			log.Println("Issuer VC 요청을 보냅니다 ")
			res, err := iss.RequestVC(hldr.DID.Did)

			if err != nil {
				panic(err)

			}

			log.Println("vc : ", res.VC)

		default:
			fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
		}

		// 개행 문자 처리
		if input != '\n' {
			_, _, err = reader.ReadRune() // 개행 문자 무시
			if err != nil {
				log.Fatalf("개행 문자 처리 실패: %v", err)
			}
		}
	}
}
