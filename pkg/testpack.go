package testpackage

import "fmt"

//Concat provides only one result by two values
func Concat(value1, value2 string) string {

	return fmt.Sprintf("%s%s", value1, value2)

}
