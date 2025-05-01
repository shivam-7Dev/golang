package intermediate

import (
	"fmt"
	"math/rand"
	"time"
)

func RandNumberInro() {
	fmt.Println("+++++++this is randdome number intro+++++++++=")

	fmt.Println("without seed:", withoutSeed())
	fmt.Println("with seed:", randWithSeed())
	fmt.Println("radn in range 3,7:", randInRange(3, 7))
	randInloop()

}

func withoutSeed() int {
	randNumber := rand.Intn(100)
	return randNumber
}

func randWithSeed() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100)
}

func randInRange(min, max int) int {
	randnum := rand.Intn(max-min+1) + min
	return randnum
}

func randInloop() {
	for i := 1; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}
}
