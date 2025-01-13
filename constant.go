package main

import "fmt"

func main() {
  const firstName string = "Indra"
  const lastName = "Rafa"

  fmt.Println(firstName)
  fmt.Println(lastName)

  const (
    game = "freefire"
    anime = "naruto"
  )

  fmt.Println(game)
  fmt.Println(anime)
}
