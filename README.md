# CRUD 블로그 만들기 
-  블로그 목록 표시, 특정 블로그 글 읽기, 블로그 생성, 수정, 삭제 기능을 포함
-  간단한 블로그 어플리케이션이지만 CRUD 기능을 통해 다양한 작업을 수행
-  https://www.notion.so/shinjangwoon/W-Golang-React-js-b14cad61acca4e24989f11fcd2eadbe8?pvs=4

## 구성
- Front: React.js
- Backend: Golang
   - Golang에서는 Gin, gin-bigo, Iris, Echo, Fiber 등 여러 프레임워크가 있다. 특히 Fiber는 가벼우면서도 빠르기 때문에 Fiber로 진행
- DB: MySQL
   - GORM(ORM)을 통해 go_fiber에서 DB와 연결

### GORM
- GORM은 MySQL, Postgres, SQLite 등 다양한 데이터베이스 지원하며, 설치는 간단하여 'go get'을 통해 설치 가능
- MVC 아키텍처로 코드베이스 구성, controller(비즈니스 로직), models(데이터베이스 테이블 struct) 등 디렉토리 생성
- 데이터베이스 처리를 위해 database.go 파일을 생성하고 GORM 및 해당 드라이버 import하여 작업 진행.
