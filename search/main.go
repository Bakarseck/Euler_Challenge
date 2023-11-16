package main

import "fmt"
import "time"

func search(liste []int, elementRecherche int) int {
    debut := 0
    fin := len(liste) - 1

    for debut <= fin {
        milieu := (debut + fin) / 2

        if liste[milieu] == elementRecherche {
            return milieu
        } else if liste[milieu] < elementRecherche {
            debut = milieu + 1
        } else {
            fin = milieu - 1
        }
    }

    return -1
}


func ls(a []int, n int, c chan int) {
    for i, v := range a {
        if v == n {
            c <- i
            return
        }
    }
}
var c = 10000

func main() {
    a := []int{}
    n := 1000000000
    for i := 0; i <= n;i++ {
        a = append(a, i)
    }

	start := time.Now()
    fmt.Println(search(a, n))
    fmt.Println(time.Since(start))
    
    start := time.Now()
    cha := make(chan int)
    x := 0
    for i := 1; i <= c; i++ {
        go ls(a[x:((len(a)/c)*i)], n, cha)
    }
    
    go func() {
        <- cha
        close(cha)
    }()
    fmt.Println(time.Since(start))
}
