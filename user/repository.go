package user

import "strings"

var users = []User{
	{Id: 1, Name: "Can kurt", Age: 30},
	{Id: 2, Name: "Banana king", Age: 25},
}

func GetUsers(search string) []User {
	filteredUsers := make([]User, 0)
	if search != "" {
		for _, u := range users {
			if strings.Contains(u.Name, search) {
				filteredUsers = append(filteredUsers, u)
			}
		}
	} else {
		filteredUsers = make([]User, len(users))
		copy(filteredUsers, users)
	}
	return filteredUsers
}

func GetUserById(id int) User {
	user := User{}
	for _, u := range users {
		if u.Id == id {
			user = u
			break
		}
	}
	return user
}

func AddUser(user User) User {
	user.Id = len(users) + 1
	users = append(users, user)
	return user
}

func UpdateUser(user User) User {
	for i, u := range users {
		if u.Id == user.Id {
			users[i] = user
			break
		}
	}
	return user
}

func DeleteUser(id int) {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
}
