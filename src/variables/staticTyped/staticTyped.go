package main

import "fmt"

func threeTwoInt() int32{
	var k int32;
	k = 30;
	return k
}

func usixFourInt() uint64{
	var k uint64;
	k = 30;
	return k
}

func complexSixFour() complex64{
	var t complex64;
	t = 3-96i;
	return t;
}

func main(){
	var k int32;
	k = 30
	fmt.Println(k)
	fmt.Println("Complex Six Four ::", complexSixFour())
}