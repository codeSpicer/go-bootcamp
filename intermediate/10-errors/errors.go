package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math error: less than zero")
	}
	return math.Sqrt(x), nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: data is empty")
	}
	return nil
}

func main() {

	// res1, err := sqrt(2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res1)
	// res2, err := sqrt(-2)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res2)
	// }

	data := []byte{}
	// fmt.Println(process(data))
	// err = process(data)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	err2 := eprocess(data)
	if err2 != nil {
		fmt.Println(err2)
	}

}

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("my error: %s", e.message)
}

func eprocess(data []byte) error {
	return &myError{message: "Error: error"}

}
