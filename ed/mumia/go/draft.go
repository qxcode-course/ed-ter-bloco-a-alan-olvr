package main
import "fmt"
func main() {

    var name string
    var age int
    fmt.Scanln(&name)
    fmt.Scanln(&age)

    if age < 12 {
        fmt.Println(name + " eh crianca")
    } else {
        if age >= 12 && age < 18{
            fmt.Println(name + " eh jovem")
        } else {
            if age >= 18 && age < 65{
                fmt.Println(name + " eh adulto")
            } else {
                if age >= 65 && age < 1000{
                    fmt.Println(name + " eh idoso")
                } else {
                    fmt.Println(name + " eh mumia")
                }
            }
        }
    }
}
