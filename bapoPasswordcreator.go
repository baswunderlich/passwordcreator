package bapoPass

var chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!?\\\"+-:;.[]()'="

func GeneratePassword(lengthOfPassword int, seed string, plattform string) string {
	result := ""
	var seedNum int64 = 0
	seedNumTable := make([]int64, 2)
	for i := 0; i < len(seed); i++ {
		seedNum = seedNum ^ int64(seed[i])
		seedNumTable = append(seedNumTable, int64(seedNum))
	}
	for i, v := range seedNumTable {
		seedNumTable[i] = v ^ seedNum
	}

	var plattformNum int64 = 0
	plattformTable := make([]int64, 2)
	for i := 0; i < len(plattform); i++ {
		plattformNum = plattformNum ^ int64(plattform[i])
		plattformTable = append(plattformTable, int64(plattform[i]))
	}
	for i := 0; i < lengthOfPassword; i++ {
		erg := (seedNum ^ int64(plattformTable[i%len(plattformTable)]) ^ plattformNum) % int64(len(chars))
		//fmt.Println(seedNum ^ int(plattformTable[i%len(plattformTable)]))
		result += string(chars[erg])
		plattformNum = seedNum ^ plattformTable[(i+1)%len(plattformTable)]
		seedNum = (erg + 1) ^ seedNumTable[(i+1)%len(seedNumTable)]
		seedNumTable = append(seedNumTable, seedNum)
	}

	return result
}
