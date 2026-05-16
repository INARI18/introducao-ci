package greeting

import "fmt"

func Message(name string) string {
	return fmt.Sprintf("Olá, %s!", name)
}
