package main

// "fmt"
// "go_tutorial/accounts"
// "io/ioutil"
// "net/http"

// func banking() {
// 	var account *accounts.Account

// 	for {
// 		var number int
// 		fmt.Println("Create new account: 1")
// 		fmt.Println("Inquiry: 2")
// 		fmt.Println("Deposit: 3")
// 		fmt.Println("Withdraw: 4")
// 		fmt.Println("Change owner: 5")
// 		fmt.Println("Change password: 6")
// 		fmt.Println("Exit: -1")

// 		switch number {
// 		case -1:
// 			fmt.Println("Account task is end! Bye!")
// 		case 1:
// 			account = accounts.NewAccount()
// 		case 2:
// 			account.Inquiry()
// 		case 3:
// 			account.Deposit()
// 		case 4:
// 			account.Withdraw()
// 		case 5:
// 			account.ChangeOwner()
// 		case 6:
// 			account.ChangePassword()
// 		default:
// 			fmt.Println("Please enter a valid number.")
// 		}

// 		if number == -1 {
// 			break
// 		}
// 	}
// }

func main() {
	// resp, err := http.Get("http://csharp.news")

	// if err != nil {
	// 	panic(err)
	// }

	// defer resp.Body.Close()

	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", string(data))
	// banking()
}
