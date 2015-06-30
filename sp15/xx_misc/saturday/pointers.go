package main

type User struct {
    Name string
}

func main() {
    u := &User{Name: "Leto"}
    println(u)
    println(&u)
    println(u.Name)
    Modify(&u)
    println(u.Name)

    println("---------")

    u2 := &User{Name: "Leto"}
    println(u2.Name)
    Modify2(u2)
    println(u2.Name)
}

func Modify(u **User) {
    *u = &User{Name: "Paul"}
}

func Modify2(u *User) {
    u = &User{Name: "Paul"}
}