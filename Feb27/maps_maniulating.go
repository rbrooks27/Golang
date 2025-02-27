package main
import ("fmt")

func main(){
	// Let_me_in()
	// UpDating()
	// get_deleted()
	// I_Want_That_One()
	// refer()
	// interate()
	// specific_order()
}

/*
Accessing Map Elements:

		Syntax:
			value = map_name[key]
	*/
		func Let_me_in(){
			var a = make(map[string]string)
			a["college"] = "Wooster"
			a["major"] = "Chem & CS"
			a["year"] = "2027"
		
			// Accessing a element
			fmt.Printf(a["college"])
		} // Output: Wooster



/*

Updating and Adding Map Elements:

		Syntax:
			map_name[key] = value
	*/
		//Example
		func UpDating(){
			var b = make(map[string]string)
			b["Company"] = "Apple"
			b["Position"] = "Software Engineer"
			b["Level"] = "I"
		  
			fmt.Println("Before Update:", b)
		  
			b["Level"] = "II" // Updating an element
			b["Name"] = "Ray" // Adding an element
		  
			fmt.Println("After Updating and Adding:", b)
		}
		/*
		Output:
			Before Update: map[Company:Apple Level:I Position:Software Engineer]
			After Updating and Adding: map[Company:Apple Level:II Name:Ray Position:Software Engineer]

		*/
/*

Removing Elements from Map
		- Removing elements is done using the delete() function.

	Syntax: 
		delete(map_name, key)
	*/
			func get_deleted(){
				var c = make(map[string]string)
				c["Company"] = "Dell"
				c["Position"] = "Data Engineer"
				c["Level"] = "IV"
				c["Name"] = "Mike"

				fmt.Println("Call,", c , " into my office.")

				delete(c,"Level")

				fmt.Println("Now that levels are gone, ", c , " leave my office.")


			}
				/*
				Output:
							Call, map[Company:Dell Level:IV Name:Mike Position:Data Engineer]  into my office.
							Now that levels are gone,  map[Company:Dell Name:Mike Position:Data Engineer]  leave my office.
				*/
/*

Checking for Specific Elements in a Map: 
	- You can check if a certain key exists in a map using:

	Syntax:
			val, ok :=map_name[key]
			(If you only want to check the existence of a certain key, you can use the blank identifier (_) in place of val.)
	*/
	// Example:
			func I_Want_That_One(){
				var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}

				val1, ok1 := a["brand"] // Checking for existing key and its value
				val2, ok2 := a["color"] // Checking for non-existing key and its value
				val3, ok3 := a["day"]   // Checking for existing key and its value
				_, ok4 := a["model"]    // Only checking for existing key and not its value
			  
				fmt.Println(val1, ok1)
				fmt.Println(val2, ok2)
				fmt.Println(val3, ok3)
				fmt.Println(ok4)
			}

			/*
			Output: 
				Ford true
				false
				true
				true
			*/


			



/*

Maps Are References:
	- Maps are references to hash tables.
	- If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.
	*/
	// Example:
			func refer(){
				var e = map[string]string{"genre":"comedy", "movie":"50 First Dates", "year":"2005"}

				f := e

				// Exactly the same because it is the same map
				fmt.Println(e)
				fmt.Println(f)

				// Even through I make a change to f since it is referring it will change e
				f["year"] = "2004"
				fmt.Println("After change to f: ")

				fmt.Println(e)
				fmt.Println(f)

				/*
				Output:
				map[genre:comedy movie:50 First Dates year:2005]
				map[genre:comedy movie:50 First Dates year:2005]
				After change to f: 
				map[genre:comedy movie:50 First Dates year:2004]
				map[genre:comedy movie:50 First Dates year:2004]
				*/

			}


/*

Iterating Over Maps
	- You can use range to iterate over maps.
	*/
	// Example:
			func iterate(){
				p := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

				for k, v := range p {
					fmt.Printf("%v : %v, ", k, v)
				}
			}
			/*
			Output:
				one : 1, two : 2, three : 3, four : 4,
			*/

/*

Iterating over Maps in a Specific Order:
	- Maps are unordered data structures
	- If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order.
	*/
	// Example
			func specific_order(){
				a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

				var b []string             // defining the order
				b = append(b, "one", "two", "three", "four")

				for k, v := range a {        // loop with no order
					fmt.Printf("%v : %v, ", k, v)
				}

				fmt.Println()

				for _, element := range b {  // loop with the defined order
					fmt.Printf("%v : %v, ", element, a[element])
  				}
				}
				/*
				Output:
				two : 2, three : 3, four : 4, one : 1, 
				one : 1, two : 2, three : 3, four : 4,
				*/