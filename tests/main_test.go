package tests

import (
	"fmt"
	"go-avanzado/utils"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	fmt.Println("Inicializando tests")

	code := m.Run()

	os.Exit(code)
}
func BenchmarkSum(b *testing.B) {

	for i := 0; i < b.N; i++ {

		utils.Sum(2, 3)
	}
}
