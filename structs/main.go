package main

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Grayson",
		contact: contact{
			email:   "alexander.grayson717@gmail.com",
			zipCode: 94000,
		},
	}

	alex.updateName("Alexander")
	alex.print()
}
