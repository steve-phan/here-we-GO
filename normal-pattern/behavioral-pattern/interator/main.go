package main

import (
	"fmt"

	i "bookable24.de/main/normal-pattern/behavioral-pattern/interator/interator"
)

func main() {

	user1 := &i.User{
		Name: "Tedy",
		Age:  10,
	}
	user2 := &i.User{
		Name: "Luna",
		Age:  8,
	}
	user3 := &i.User{
		Name: "LiLy",
		Age:  8,
	}

	userCollection := &i.UserCollection{
		Users: []*i.User{user1, user2, user3},
	}

	userIterator := userCollection.CreateInterator()
	fmt.Printf("There are %v users\n", len(userCollection.Users))
	for userIterator.HasNext() {
		user := userIterator.GetNext()
		fmt.Printf(" User's name is %v and user's age is %v\n", user.Name, user.Age)
	}

}
