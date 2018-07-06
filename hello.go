//following tutorial on https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world
package main

const helloPrefix = "Hello, "

//Hello : function to learn unit testing
func Hello(name string) string {
	if name == "" {
		name = "Testing"
	}

	return helloPrefix + name
}
