package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	Owner    string
	account  string
	balance  int
	password string
}

var errorCodes = map[int]string{
	-2: "Not matched password",
	-1: "Not found account",
	0:  "Can't withdraw your accounts",
	1:  "OK",
}

func errorMsgs(code int) error {
	return errors.New(errorCodes[code])
}

func verifyAccount(a *Account) {
	unMatchCount := 0
	var owner string
	fmt.Print("Enter your owner name :")
	fmt.Scanln(&owner)
	var err error
	if owner != a.Owner {
		err = errorMsgs(-1)
	}
	if err != nil {
		panic(err)
	}

	defer func() {
		for {
			var pass string
			fmt.Print("Enter your password :")
			fmt.Scanln(&pass)
			if pass != a.password {
				unMatchCount++
				if unMatchCount == 3 {
					err = errorMsgs(-2)
				}
				if err != nil {
					panic(err)
				}
				fmt.Println("Password is not matched. Please try again.")
				continue
			}
			if pass == a.password {
				break
			}
		}
	}()
}

func NewAccount() *Account {
	var newOwner string
	fmt.Print("Enter your new owner name :")
	fmt.Scanln(&newOwner)
	account := Account{Owner: newOwner, balance: 0}
	var pass string
	var confirmPass string
	for {
		fmt.Print("Enter your password")
		fmt.Scanln(&pass)
		fmt.Print("Enter your confirm password")
		fmt.Scanln(&confirmPass)

		if pass == confirmPass {
			account.password = pass
			break
		}
		fmt.Println("Password is not matched. Please try again.")
	}
	fmt.Println("Hello!, ", account.Owner)

	account.account = "123-456-7890"
	return &account
}

func (a *Account) Inquiry() string {
	verifyAccount(a)
	return fmt.Sprintf("OWNER : %s, BALANCE : %d", a.Owner, a.balance)
}

func (a *Account) Deposit() {
	var amount int
	fmt.Print("Enter your deposit amount :")
	fmt.Scanln(&amount)
	a.balance += amount
	fmt.Println("Deposit is completed.")
	fmt.Println("Your balance is ", a.balance)
}

func (a *Account) Withdraw() {
	verifyAccount(a)
	var amount int
	fmt.Print("Enter your withdraw amount :")
	fmt.Scanln(&amount)

	var err error
	if amount > a.balance {
		err = errorMsgs(0)
	}
	if err != nil {
		panic(err)
	}

	defer func() {
		a.balance -= amount
		fmt.Println("Withdraw is completed.")
		fmt.Println("Your balance is ", a.balance)
	}()
}

func (a *Account) ChangeOwner() {
	verifyAccount(a)
	var newOwner string
	fmt.Print("Enter your new owner name :")
	fmt.Scanln(&newOwner)
}

func (a *Account) ChangePassword() {
	verifyAccount(a)
	var newPass string
	fmt.Print("Enter your new password :")
	fmt.Scanln(&newPass)
	var confirmPass string
	fmt.Print("Enter your confirm password :")
	fmt.Scanln(&confirmPass)

	var err error
	if newPass != confirmPass {
		err = errorMsgs(-2)
	}
	if err != nil {
		panic(err)
	}
	defer func() {
		a.password = newPass
		fmt.Println("Change password is completed.")
	}()
}
