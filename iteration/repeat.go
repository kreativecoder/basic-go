//Package iteration based on this tutorial https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration
package iteration

const repeatCount = 5

//Repeat : accepts a character and repeats it 5 times
func Repeat(character string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
