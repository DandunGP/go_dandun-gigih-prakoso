/*

BEFORE

package main

type user struct {
	id       int
	username int
	password int
}

type userservice struct {
	t []user
}

func (u userservice) getallusers() []user {
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}
	return user{}
}

*/

/* AFTER */

package main

type User struct {
	id       int
	username int
	password int
}

type userService struct {
	t []User
}

func (u userService) getAllUser() []User {
	return u.t
}

func (u userService) getUserById(id int) User {
	for _, items := range u.t {
		if id == items.id {
			return items
		}
	}
	return User{}
}
