package main
import "fmt"
func main() {
     
    var h, p, f, d int
    fmt.Scanln(&h, &p, &f, &d)

    if(h > 15 || p > 15 || f > 15){
        fmt.Println("error")
    } 

    var distH, distP int
    
    if d == 1 {
        distH = h - f
        distP = p - f
    } else {
        distH = f - h
        distP = f - p
    }

    if distH < 0{
        distH += 16
    }

    if distP < 0 {
        distP += 16
    }

    if distH < distP {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
    
}
