package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person2 struct {
	Name   string
	Addr   string
	Job    string
	Reason string
}

func main() {

	var input = os.Args

	var person = []*Person2{
		{Name: "nama1", Addr: "addr1", Job: "job1", Reason: "reason1"},
		{Name: "nama2", Addr: "addr2", Job: "job2", Reason: "reason2"},
		{Name: "nama3", Addr: "addr3", Job: "job3", Reason: "reason3"},
		{Name: "nama4", Addr: "addr4", Job: "job4", Reason: "reason4"},
		{Name: "nama5", Addr: "addr5", Job: "job5", Reason: "reason5"},
		{Name: "nama6", Addr: "addr6", Job: "job6", Reason: "reason6"},
		{Name: "nama7", Addr: "addr7", Job: "job7", Reason: "reason7"},
		{Name: "nama8", Addr: "addr8", Job: "job8", Reason: "reason8"},
		{Name: "nama9", Addr: "addr9", Job: "job9", Reason: "reason9"},
		{Name: "nama10", Addr: "addr10", Job: "job10", Reason: "reason10"},
	}

	index, _ := strconv.Atoi(input[1])

	printDetails(index, person)
}

func printDetails(in int, p []*Person2) {
	fmt.Println("Index ke -", in)
	fmt.Println(p[in].Name)
	fmt.Println(p[in].Addr)
	fmt.Println(p[in].Job)
	fmt.Println(p[in].Reason)
}
