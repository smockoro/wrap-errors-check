package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func returnError() error {
	err := fmt.Errorf("sample")
	if err != nil {
		return err
	}

	err = fmt.Errorf("sample1111")

	if err != nil {
		return errors.Wrap(err, "faild")
	}

	return nil
}

func returnError2() (string, error) {
	err := fmt.Errorf("sample")
	if err != nil {
		return "ng", err
	}

	err = fmt.Errorf("sample1111")

	if err != nil {
		return "ng", errors.Wrap(err, "faild")
	}

	return "ok", nil
}

func main() {
	fmt.Println("hello")
}
