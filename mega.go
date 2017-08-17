package main

import (
    //"bufio"
    "fmt"
    //"os"
)

func check(err error) {
    if err != nil {
        //panic(err)
    }
}

/* Pega o valor da última posição e coloca em matrix[pos] (jogo que deve ser removido)
* Ou seja, é feito uma duplicação de jogo
* E a última posição da matrix deve ser descartada.
*/
func remove(matrix [][]int, pos int, row []int) [][]int {
    matrix[pos] = matrix[len(matrix)-1]
    return matrix[:len(matrix)-1]
}



func main(){
    var i, j, k, l, m, n uint8
    /*i := 0
    j := 0
    k := 0
    l := 0
    m := 0
    n := 0*/
    cont := 0

    /*f, err := os.Create("C:/Users/luisr/Desktop/Result.txt")
    check(err)
    defer f.Close()
    w := bufio.NewWriter(f)*/

    matrix := make([][]uint8, 50063860) // One row per unit of y.
    // Loop over the rows, allocating the slice for each row.
    for p := range matrix {
	   matrix[p] = make([]uint8, 6)
    }

    //matrix[50000][3] = 1
    //fmt.Println("posicao [50000][3]", matrix[50000][3])


    for i=1; i<=60; i++{
        for j=i+1; j<=60; j++{
            for k=j+1; k<=60; k++{
                for l=k+1; l<=60; l++{
                    for m=l+1; m<=60; m++{
                        for n=m+1; n<=60; n++{
                            matrix[cont][0] = i
                            matrix[cont][1] = j
                            matrix[cont][2] = k
                            matrix[cont][3] = l
                            matrix[cont][4] = m
                            matrix[cont][5] = n
                            cont++;
                        }
                    }
                }
            }
        }
    }

    fmt.Println("cont: ", cont)

    /*_, err = fmt.Fprintf(w, "%#v", matrix)
    check(err)
    w.Flush()*/


}

func removeSequencesOf3(){

}

func removeVerticalSequences(){

}

func removeDiagonalSequencesOfRight(){

}

func removeDiagonalSequencesOfLeft(){

}

func removeGamesOfLessThen3Quadrants(){

}
