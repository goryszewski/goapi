package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type User struct {
	ID int
	FirstName string
	LastName string
}

var (
	users []*User
	nextID = 1
)

func GetUsers() []*User {
	client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        DB:		  0,  // use default DB
    })
	ctx := context.Background()
	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("log:",val)
	return users
}

func AddUser(u User) (User , error) {
	if u.ID != 0 {
		return User{} , errors.New("NEw User must not include id or it must be set to zero")
	}
	u.ID=nextID
	nextID++
	users = append(users, &u)
	return u,nil
}

func GetUserByID(id int) (User,error) {
	for _, u := range users{
		if u.ID == id {
			return *u,nil
		}
	}
	return User{},fmt.Errorf("User not found %v",id)
}

func UpdateUser(u User) (User, error){
	for i , candidate := range users{
		if u.ID == candidate.ID {
			users[i] = &u
			return u,nil
		}
	}
	return User{}, fmt.Errorf("User not found %v",u.ID)

}

func RemoveUserByID(id int) error {
	for i,u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]... )
			return nil
		}
	}
	return fmt.Errorf("Uset not found %v",id)
}