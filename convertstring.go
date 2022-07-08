//converting strings
package convertstring

import "strconv"

func StrToInt(s string) (int) {
    num, _ := strconv.Atoi(s)
    
    return num
}

func StrToFloat(s string) (float64) {
    num, _ := strconv.ParseFloat(s, 64)
    
    return num
}

