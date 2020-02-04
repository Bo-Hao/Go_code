package main 

import(
	"fmt"
)

func main(){
	//use string as key
	x := make(map[string]int)
	x["go"] = 10
	fmt.Println(x, len(x))

	//use integer as key
	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y, len(y))
	//delete(y, 1) 

	elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"

	fmt.Println(elements, len(elements))
	
	//if key is not exist, map can still call it. It consider "empty" as empty.
	fmt.Println(elements["empty"])
	name, ok := elements["empty"]	//we can see this as two return value.
	fmt.Println(name, ok)	

	//short representation
	/* elements := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Fluorine",
		"Ne": "Neon",
	} */
	
	//json is known as nested structure, and so do go can archieve it.
	e := map[string]map[string]string{
        "H": map[string]string{
            "name":"Hydrogen",
            "state":"gas",
        },
        "He": map[string]string{
            "name":"Helium",
            "state":"gas",
        },
        "Li": map[string]string{
            "name":"Lithium",
            "state":"solid",
        },
        "Be": map[string]string{
            "name":"Beryllium",
            "state":"solid",
        },
        "B":  map[string]string{
            "name":"Boron",
            "state":"solid",
        },
        "C":  map[string]string{
            "name":"Carbon",
            "state":"solid",
        },
        "N":  map[string]string{
            "name":"Nitrogen",
            "state":"gas",
        },
        "O":  map[string]string{
            "name":"Oxygen",
            "state":"gas",
        },
        "F":  map[string]string{
            "name":"Fluorine",
            "state":"gas",
        },
        "Ne":  map[string]string{
            "name":"Neon",
            "state":"gas",
        },
    }
    if el, ok := e["Li"]; ok {
        fmt.Println(el["name"], el["state"])
    }

}