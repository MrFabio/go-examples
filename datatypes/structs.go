package datatypes

import "fmt"

// customer represents a customer with basic information
type customer struct {
	name    string
	address string
	bal     float64
}

// getCustomerInfo prints customer information
func getCustomerInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}

// newCustomerAddress updates a customer's address
// Takes a pointer to customer to modify the original struct
func newCustomerAddress(c *customer, address string) {
	c.address = address
}

// rectangle represents a geometric rectangle
type rectangle struct {
	length, height float64
}

// Area calculates the area of a rectangle
func (r rectangle) Area() float64 {
	return r.length * r.height
}

// contact represents a person's contact information
type contact struct {
	firstName, lastName string
	phone               string
}

// business represents a business with contact information
// Uses composition by embedding the contact struct
type business struct {
	name, address string
	contact
}

// info prints the business contact information
func (b business) info() {
	fmt.Printf("Contact at %s is %s %s\n", b.name, b.contact.firstName, b.contact.lastName)
}

// StructsExamples demonstrates various struct operations
func StructsExamples() {
	pl("Structs Examples")

	// Create a customer using field assignment
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "112 Main St"
	tS.bal = 123.56

	getCustomerInfo(tS)

	// Update address using a pointer receiver
	newCustomerAddress(&tS, "21 South St")

	pl("After address change:", tS.address)

	// Create a customer using struct literal (anonymous)
	sS2 := customer{"Sue Jones", "123 Second St", 100.0}
	pl("Anonymous Struct:", sS2.name)

	// Create a customer using named fields
	sS := customer{
		name:    "Sue Jones",
		address: "123 Second St",
		bal:     100.0,
	}
	pl("Named Struct:", sS.name)

	getCustomerInfo(sS)

	// Create and use a rectangle
	rect1 := rectangle{10, 20}
	// Accessing the Area method of the rectangle struct
	pl("Area of rectangle:", rect1.Area())

	// Create a contact using named fields
	contact1 := contact{
		firstName: "John",
		lastName:  "Doe",
		phone:     "1234567890",
	}

	// Create a business with embedded contact
	business1 := business{
		name:    "ABC Plumbing",
		address: "456 North St",
		contact: contact1,
	}

	// Call the info method to display business contact information
	business1.info()
}
