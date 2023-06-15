package excepts

import (
	"log"
	"os"
)

func ExceptOs() {
	f, err := os.Open("filename.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	println(f.Name())
}

// func ExceptFunc() {
//   _, err := otherFunc()
//   switch err.(type) {
//   default:
//     println("ok")
//   case MyError:
//     log.Print("Log my error")
//   case error:
//     log.Fatal(err.Error())
//   }
// }

// func otherFunc() {
// 	panic("unimplemented")
// }
