// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Products struct {
// 	ID         int
// 	Name       string
// 	Pricelists Pricelists
// }

// type Pricelists struct {
// 	ID       int
// 	Category string
// 	Brand    string
// 	Price    int
// }

// var jsonSample = `
// [
//   {
//     "ID": 1,
//     "Name": "Pena Hitam",
//     "Price": 0,
//     "Pricelists": {
//       "ID": 1,
//       "Category": "pena",
//       "Brand": "bagus"
//     }
//   },
//   {
//     "ID": 2,
//     "Name": "Pena Hitam",
//     "Price": 0,
//     "Pricelists": {
//       "ID": 1,
//       "Category": "pena",
//       "Brand": "laris"
//     }
//   }
// ]
// `

// func main() {
// 	mapArr := getMaps()
// 	if bte, e := json.Marshal(mapArr); nil == e {
// 		fmt.Println(string(bte))
// 	}
// }

// func getMaps() map[string]map[string]string {
// 	var products []Products
// 	output := map[string]map[string]string{}
// 	if e := json.Unmarshal([]byte(jsonSample), &products); nil != e {
// 		panic(e)
// 	}
// 	for i := range products {
// 		productToMapItem(products[i], output)
// 	}

// 	return output
// }

// func productToMapItem(product Products, output map[string]map[string]string) {
// 	// CODE HERE
// 	// {"pena":{"bagus":"Pena Hitam","laris":"Pena Hitam"}}

// 	temp := map[string]map[string]string{
// 		product.Pricelists.Category: {
// 			product.Pricelists.Brand: product.Name,
// 		},
// 	}

// 	if output[product.Pricelists.Category] == nil {
// 		output[product.Pricelists.Category] = temp[product.Pricelists.Category]
// 	} else {
// 		for k, v := range temp[product.Pricelists.Category] {
// 			output[product.Pricelists.Category][k] = v
// 		}
// 	}
// }
