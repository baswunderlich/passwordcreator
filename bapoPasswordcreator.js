chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!?\\\"+-:;.[]()'="

/*
lengthOfPassword should be a number
seed and plattform should be strings
*/
function generate_password(lengthOfPassword, seed, plattform) {
    result = ""
    seedNumTable = []
    seedNum = 0
    for (i = 0; i < seed.length; i++){
        seedNum = seedNum ^ seed[i].charCodeAt(0)
        seedNumTable.push(seedNum)
    }

    for (i = 0; i < seedNumTable.length; i++){
        seedNumTable[i] = seedNumTable[i] ^ seedNum
    }

    plattformTable = []
    for(i = 0; i < plattform.length; i++){
        plattformTable.push(plattform[i].charCodeAt(0))
    }
    for(i = 0; i < lengthOfPassword; i++){
        erg = (seedNum ^ plattformTable[i%plattformTable.length]) % chars.length
        result = result + chars[erg]
        seedNum = (erg+1) ^ seedNumTable[(i+1)%seedNumTable.length]
        seedNumTable.push(seedNum)
    }

    return result
}