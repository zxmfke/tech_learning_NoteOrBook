package main

func main() {
	mcdonaldStore := new(Mcdonald)
	mcdonaldStore.OrderHamburger("Big Mac")
	mcdonaldStore.OrderHamburger("Double Cheeseburger")

	mossStore := new(Moss)
	mossStore.OrderHamburger("Unworthy")
	mossStore.OrderHamburger("Big Mac")
}
