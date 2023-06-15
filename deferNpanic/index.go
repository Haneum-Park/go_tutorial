package defernpanic

import "os"

func DeferBased() {
	f, err := os.Open("filename.txt")
	if err != nil {
		panic(err)
	}

	// ? 마지막에 파일 close 실행
	defer f.Close()

	// ? 파일 읽기
	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
}

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}

func PanicBased() {
	openFile("filename.txt")

	// ? openFile 안에서 panic이 실행되면 아래 println은 실행되지 않음
	println("Done")
}
