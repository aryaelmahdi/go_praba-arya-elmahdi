package main

//user harusnya menjadi User
type user struct {
	//tipe data tidak sesuai, harusnya string
	id       int
	username int
	password int
}

type userservice struct {
	//t tidak deskriptif, bisa diganti dengan users
	t []user
}

//tidak menggunakan pointer *, nama method harusnya menggunakan camel case
func (u userservice) getallusers() []user {
	return u.t
}

//tidak menggunakan pointer *, nama method harusnya menggunakan camel case
func (u userservice) getuserbyid(id int) user {
	//r tidak deskriptif, bisa diganti dengan user
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}
	//diganti User mengikuti struct
	return user{}
}
