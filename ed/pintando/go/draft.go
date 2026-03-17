package main
import "fmt"
import "math"
func main() {

    var n1, n2, n3 float64
    fmt.Scanln(&n1)
    fmt.Scanln(&n2)
    fmt.Scanln(&n3)

    p := (n1 + n2 + n3)/2.0
    area := math.Sqrt(p*(p - n1)*(p - n2)*(p - n3))

    fmt.Printf("%.2f\n", area)

}
