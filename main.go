package main

import "fmt"


type LoginError struct { 
	Username string
	Message string
}

func (e *LoginError) Error() string {
	return fmt.Sprintf("Login error for user '%s': %s",e.Username , e.Message)
}

func login(username,password string) error{
	if username != "admin" || password != "12345"{	
		return &LoginError{Username: username , Message: "Invalid credentials"}
	}
	return nil
}


func main() {
	// Success
	err := login("123","456")


	if err != nil {
		switch e := err.(type){
		case *LoginError:
			fmt.Println("Customer error" ,e)

		default:
		fmt.Println("Generic Error",e)
		}
		return 
	}
	fmt.Println("Login Sucess")


}