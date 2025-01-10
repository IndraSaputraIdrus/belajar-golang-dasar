package main

import "fmt"

func main() {
  var name string // wajib menyebutkan tipe data

  // bisa juga tanpa menyebutkan, contoh
  // var name = "indra saputra idrus"

  name = "Indra saputra idrus"
  fmt.Println(name)

  name = "Muh rafa idrus"
  fmt.Println(name)

  game := "Free fire" // shorthand deklarasi variable
  fmt.Println(game)

  game = "Mobile legend"
  fmt.Println(game)

  // multiple variable

  var (
    firstName = "Indra"
    middleName = "Saputra"
    lastName = "Idrus"
  )

  fmt.Println(firstName)
  fmt.Println(middleName)
  fmt.Println(lastName)
}
