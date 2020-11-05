package middlewares

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUserFromCAS(t *testing.T) {
	session := "eyJDQVNfVVNFUk5BTUUiOiJaWlplbm8iLCJyZWYiOiJodHRwcyUzQS8vZ2l0aHViLmNvbS8lMjMvIiwic2VydmljZSI6InRlc3QifQ.X6Ob5Q.98uQ2q96Jvq1Tkqb2Q7-oFEuAEc"

	fmt.Println(UserFromCAS(session))
	b, _ := json.Marshal(UserFromCAS(session))
	fmt.Println(string(b))
	//fmt.Printf("%+v", UserFromCAS(session))
}
