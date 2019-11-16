package common

import "strconv"

//ParseID2Str .
func ParseID2Str(id uint) (str string) {
	str = strconv.FormatUint(uint64(id), 10)
	return
}
