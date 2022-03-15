package bapoPass

var chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!?\\\"+-:;.[]()'="

func GeneratePassword(lengthOfPassword int, seed string, plattform string) string {
	result := ""
	var seedNum int = 0
	seedNumTable := make([]int, 2)
	for i := 0; i < len(seed); i++ {
		seedNum = seedNum ^ int(seed[i])
		seedNumTable = append(seedNumTable, int(seedNum))
	}
	for i, v := range seedNumTable {
		seedNumTable[i] = v ^ seedNum
	}

	plattformTable := make([]byte, 2)
	for i := 0; i < len(plattform); i++ {
		plattformTable = append(plattformTable, byte(plattform[i]))
	}
	for i := 0; i < lengthOfPassword; i++ {
		erg := (seedNum ^ int(plattformTable[i%len(plattformTable)])) % len(chars)
		//fmt.Println(seedNum ^ int(plattformTable[i%len(plattformTable)]))
		result += string(chars[erg])
		seedNum = (erg + 1) ^ seedNumTable[(i+1)%len(seedNumTable)]
		seedNumTable = append(seedNumTable, seedNum)
	}

	return result
}
