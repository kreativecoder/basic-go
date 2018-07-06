//Package iteration based on this tutorial https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration
package iteration

//Repeat : accepts a character and repeats it "count" times
func Repeat(character string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}
