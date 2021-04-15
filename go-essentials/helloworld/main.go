/* if the package name is main then when the code is built an executable file will be made. If the package name is anything else then it can become a library, reusable code source, or a dependency not an executable */
package main

/* import is used to bring in an external package and use the code tools/functions that are within that external package.
The fmt library implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler. */
import "fmt"

/* body of the application where all the code and functions are written and run. Every package named main needs a primary function named main  */
func main() {
	fmt.Println("Hello World!")
}
