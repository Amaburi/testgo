package main

import (
	"fmt"
	"time"

	"rsc.io/quote"
)

type Murid struct {
	nama        string
	umur        int
	alamat      string
	hobi        string
	universitas string
}

func Muridd() {
	var person1 Murid
	person1.nama = "Akame"
	person1.umur = 20
	person1.alamat = "Bandung"
	person1.hobi = "Makan"
	person1.universitas = "Universitas telkom"

	PrintMurid(person1)
	fmt.Printf("\x1b[34m----------\x1b[0m\n")
	var person2 Murid
	person2.nama = "Taylor"
	person2.umur = 22
	person2.alamat = "Bojongsoang"
	person2.hobi = "Nyanyi"
	person2.universitas = "Universitas telkom"
	PrintMurid(person2)
}
func PrintMurid(pers Murid) {
	fmt.Println("Nama: ", pers.nama)
	fmt.Println("Umur : ", pers.umur)
	fmt.Println("Univ: ", pers.universitas)
	fmt.Println("Alamat: ", pers.alamat)

}
func datadiri(name string, mahasiswa bool, yearborn int, nim int) {
	currentTime := time.Now()
	fmt.Println("Nama:", name, "Mahasiswa:", mahasiswa, "Berapa tahun hidup:", currentTime.Year()-yearborn, "Nim : ", nim)
}
func switchh() {
	var day = 0
	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Day is not set")
	}
}
func repeat() {
	var name = []string{"Nama", "Masuk", "Masuk", "Masuuuk"}
	// normal loop
	for i := 0; i < len(name); i++ {
		fmt.Println("\n", name[i])
	}
	// using range
	// for idx, name := range name {
	// 	fmt.Println("\n", idx, name)
	// }
}
func slice() {
	// len Use len when you want to work with the actual number of elements in the array or slice. This ensures that you only access valid elements.
	my_variable := []string{"satu", "dua", "tiga"}
	fmt.Println("Initial length value :", len(my_variable))
	// what [1:3] means is [start:end] it just counts the 0-2 index and doesnt include the end index which i sindex 3
	slicep := my_variable[0:2]
	lenn := len(slicep)
	fmt.Println("Sliced Value :", slicep)
	fmt.Println("Length value :", lenn)
	// checking the full capacity using "cap"
	fmt.Println("Capacity :", cap(my_variable))
	fmt.Println("Capacity :", cap(slicep))

	fmt.Printf("\x1b[34m-------Append-------\x1b[0m\n")
	// append
	my_variable = append(my_variable, "empat", "lima", "enam")
	fmt.Println("After append Value :", my_variable)
	fmt.Println("Length value :", len(my_variable))
	fmt.Println("Capacity :", cap(my_variable))
	sliceppp := my_variable[2:5]
	fmt.Println(sliceppp)
	sliceppp = append(sliceppp, "enam", "tujuh", "delapan", "sembilan", "sepuluh")

	fmt.Printf("\x1b[34m-------Copy-------\x1b[0m\n")
	// Copy
	neededWords := sliceppp[:len(sliceppp)-5]
	wordstoCopy := make([]string, len(neededWords))
	copy(wordstoCopy, neededWords)
	fmt.Printf("numbersCopy = %v\n", wordstoCopy)
	fmt.Printf("length = %d\n", len(wordstoCopy))
	fmt.Printf("capacity = %d\n", cap(wordstoCopy))
	// cap Use cap when you want to work with the capacity of the slice, which represents the maximum number of elements it can hold without resizing its underlying array. Using cap for slicing or iterating may lead to unexpected behavior if you access elements beyond the actual data stored in the slice.
	sliice := make([]int, 0, 20)
	capacity := cap(sliice)
	fmt.Println(capacity)
}
func mapp() {
	// we use this line of code if the map has a value instead using make(map[string]string) use this if you have empty map
	var map1 = map[string]string{"Brand": "Ford", "Product": "Car", "Year release": "1990"}
	fmt.Println(map1)

	fmt.Printf("\x1b[34m------Checking the key and value in the map--------\x1b[0m\n")
	// how to check specific values of the map val stands for value and ok stands for the key
	val1, ok1 := map1["Brand"]
	val2, ok2 := map1["color"]
	_, ok3 := map1["Product"]
	val4, _ := map1["Product"]
	fmt.Println(val1, ":", ok1)
	fmt.Println(val2, ok2)
	// Checking for only its key
	fmt.Println(ok3)
	// Checking for only its value
	fmt.Println(val4)

	// using map references
	fmt.Printf("\x1b[34m------Example of using reference in map--------\x1b[0m\n")
	fmt.Println("initial value :", map1)
	map3 := map1
	map3["Brand"] = "Honda"
	map3["Product"] = "Motorcycle"
	map3["Year release"] = "2005"
	fmt.Println("after using reference :", map3)

	fmt.Printf("\x1b[34m------Another Example of using map--------\x1b[0m\n")
	// empty array then we add using map2["name variable"] = "name value" but i recommend using make(map[keyType]valueType) instead of map([keyType]valueType{value})
	var map2 = make(map[string]string)
	// this is how we access the map value if the there's no value in the map
	map2["Brand"] = "Ferarri"
	map2["Product"] = "Car"
	map2["Year release"] = "2000"
	fmt.Println(map2)
	fmt.Println(map2["Brand"])
	// this is how we can delete a value from the map
	delete(map2, "Brand")
	fmt.Println(map2)
	// this is how we update a value in the map just override it.
	map2["Year release"] = "2024"
	map2["Brand"] = "Ferrari"
	fmt.Println(map2)

	fmt.Printf("\x1b[34m------Using iterate in map--------\x1b[0m\n")
	// Iterate over map k stands for key and v for value you can customize the it based on your preferences
	for k, v := range map2 {
		fmt.Printf("%v : %v ", k, v)
	}
	// using this as a space
	fmt.Println()
	// iretate in order
	var b []string
	b = append(b, "Year release", "Brand", "Product")
	for _, element := range b { // loop with the defined order
		fmt.Printf("%v : %v, ", element, map2[element])
	}
}

func main() {

	var boo bool = false
	fmt.Println(boo)
	fmt.Printf("\x1b[31m---------Custom Parameter---------\x1b[0m\n")
	datadiri("hokuto", true, 1999, 103012300216)
	age := 1
	fmt.Println(age)
	fmt.Println(quote.Opt())
	fmt.Printf("\x1b[31m---------Switch---------\x1b[0m\n")
	switchh()
	fmt.Printf("\x1b[31m---------Looping---------\x1b[0m\n")
	repeat()
	fmt.Printf("\x1b[31m---------Struct---------\x1b[0m\n")
	Muridd()
	fmt.Printf("\x1b[31m---------Slice---------\x1b[0m\n")
	slice()
	fmt.Printf("\x1b[31m-------Map-------\x1b[0m\n")
	mapp()
}
