package main

import "fmt"

var courseName []string = []string{"JAVA", "Python"}

func main() {
	courseName = append(courseName, "Golang")
	courseName = append(courseName, "C#")

	courseName = append(courseName, "JavaScript", "PHP")

	fmt.Println(courseName)

	courseWeb := courseName[4:6]
	fmt.Println(courseWeb)

	courseWeb = courseName[:4]
	fmt.Println(courseWeb)
}
