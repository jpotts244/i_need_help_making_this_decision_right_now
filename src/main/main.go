package main

import "fmt"
import "os"
import "bufio"
import "time"
import "github.com/fatih/color"
import "math/rand"

func main() {
  rand.Seed(time.Now().Unix())
  input := make([]string, 0)

  c := color.New(color.FgCyan)
  c.Println("Enter first  choice: ")
  scan := bufio.NewScanner(os.Stdin)
  scan.Scan()

  s := color.New(color.FgBlue)
  s.Println("Enter second choice: ")
  second_scan := bufio.NewScanner(os.Stdin)
  second_scan.Scan()

  t := color.New(color.FgMagenta)
  t.Println("Enter third  choice: ")
  third_scan := bufio.NewScanner(os.Stdin)
  third_scan.Scan()

  // array[0] = scan.Text()
  // array[1] = second_scan.Text()
  // array[2] = third_scan.Text()

  input = append(input, 
    scan.Text(),
    second_scan.Text(),
    third_scan.Text(),
    )

  n := rand.Int() % len(input)

  message := fmt.Sprint("And the winner is....", input[n])
  fmt.Println(message)

}
