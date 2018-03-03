package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

)
var Totn int
var tot []int
func main() {
	var outerloop int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	
	outerloop, err := strconv.Atoi(scanner.Text())
    if err != nil{
    	fmt.Println("outerloop error")
    }
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
    Totn = outerloop
	outerrec(outerloop)
}
func printans(n int, c int) int{
	if n ==0{
		return 0
	}
    fmt.Println(tot[c])
    return printans(n-1, c+1)
}
func outerrec(n int) int{
	//fmt.Println("outerrec")
	if n== 0{
        printans(Totn, 0)
    	return 0
    }
    var innerloop int
    var iArr []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	innerloop, err := strconv.Atoi(scanner.Text())
    if err != nil{
    	fmt.Println("innerloop error")
    }
    
   
    innerrec(innerloop, iArr)

    
    return outerrec(n-1)

}
func innerrec(n int, arr []int) int{
	//fmt.Println("innerrec")
	if n == 0{
		//fmt.Println((len(arr)-1))
		sumofs(arr, (len(arr)-1), 0)
		//fmt.Println("CHECK")
		return 0
	}
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil{
		return 0
	}
	
	arr = append(arr, i)
	return innerrec(n-1, arr)

}

func sumofs(arr []int, n int, sum int) int{
	//fmt.Println("sumofs")

	if n == -1{
		//fmt.Println("TOT APPEND")
		tot = append(tot, sum)
		return -1
	}
    if((arr[n] <= 0)&& n>=0){
    	return sumofs(arr, n-1, sum)
    }else{
    	//fmt.Println(arr[n])
    	sum+= (arr[n]*arr[n])
    }

    return sumofs(arr, n-1, sum)
     
}
