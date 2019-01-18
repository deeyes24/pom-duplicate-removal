package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName1 := os.Args[1]
	fileName2 := os.Args[2]

	//Load content from fileName1 and fileName2 as pom1 and pom2

	fmt.Println("Found files ", fileName1, fileName2)

	pomFile1, err := os.Open(fileName1)
	if err != nil {
		fmt.Println("Unable to read fileName ", fileName1)
		panic(2)
	}

	pomFile2, err := os.Open(fileName2)
	if err != nil {
		fmt.Println("Unable to read fileName ", fileName2)
		panic(2)
	}
	defer pomFile1.Close()
	defer pomFile2.Close()

	byteValue1, err := ioutil.ReadAll(pomFile1)
	var deps Project
	var deps2 Project

	err = xml.Unmarshal(byteValue1, &deps)

	if err != nil {
		fmt.Println("Unable to marshal the file ", fileName1)
	}

	byteValue2, err := ioutil.ReadAll(pomFile2)

	err = xml.Unmarshal(byteValue2, &deps2)

	unique1 := SliceUniqMap(deps.Dependencies.Dependencies)

	unique2 := SliceUniqMap(deps.Dependencies.Dependencies)

	fmt.Println("Dependencies found in ", fileName1, "but not in ", fileName2)
	printDiff(unique1, unique2)

	fmt.Println("Dependencies found in ", fileName2, "but not in ", fileName1)
	printDiff(unique2, unique1)

}

func printDiff(u1 []Dependency, u2 []Dependency) {
	seen1 := make(map[Dependency]struct{}, len(u1))
	seen2 := make(map[Dependency]struct{}, len(u2))
}
