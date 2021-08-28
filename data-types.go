package main


import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    
    // Declare second integer, double, and String variables.
    var a uint64 = 0
    var b float64 = 0.0
    var t string = "is the best place to learn and practice coding!"
    
    // Read and save an integer, double, and String to your variables.
    if scanner.Scan() {
        za, _ := strconv.ParseInt(scanner.Text(), 10, 64)
        a = uint64(za)        
    }
    if scanner.Scan() {
        f, _ := strconv.ParseFloat(scanner.Text(), 64)
        b = f      
    }
    if scanner.Scan() {
        t = scanner.Text()
    }
    
    // Print the sum of both integer variables on a new line.
    fmt.Println(i+a)
    
    // Print the sum of the double variables on a new line.
     zz := float64(int((b+d) * 100)) / 100
    
    // Concatenate and print the String variables on a new line
    fmt.Printf("%.1f\n",zz)
    
    // The 's' variable above should be printed first.
    fmt.Println(s + t)
}
