package main

import (
	"context"
	"fmt"
)

type Key int
type Email string

func main() {
	const UserKey Key = 0
	const UserEmail Email = "gubaydullin_nurislam@gmail.com"

	withValueContext := context.WithValue(context.Background(), UserKey, UserEmail)

	if userEmail, ok := withValueContext.Value(UserKey).(Email); ok {
		fmt.Println("User Email:", userEmail)
		if userEmail != UserEmail {
			panic("values not same")
		}
	}

	if userEmail, ok := withValueContext.Value(UserKey + 1).(Email); ok {
		fmt.Println("User Email:", userEmail) // Never enter here, cause key not found
		if userEmail != UserEmail {
			panic("values not same")
		}
	}

	if userEmail, ok := withValueContext.Value(UserKey).(string); ok {
		fmt.Println("User Email:", userEmail) // Never enter here, cause value type mismatch
		if userEmail != string(UserEmail) {
			panic("values not same")
		}
	}
}
