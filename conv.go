package main

import (
	
)

// Encode integer id to 64base number in string
func Encode(id int32) string {
	arr := make([]byte, 0)
	for id > 0 {
		r := id % 64
		arr = append([]byte{to64Base(r)}, arr...)
		id = id >> 6
	}
	return string( arr )
}

//TODO:	Decode

// Convert remainder to 64base digit in byte
// Digit table:
//	number	character
//	 0 9	   0-9
//	10-35	   A-Z
//	36-61	   a-z
//	62-63	   -, .
func to64Base(num int32) byte {
	switch {
		case num >=0 && num <= 9:
			// '0'+num
			return byte(48+num)
		case num >=10 && num <= 35:
			// 'A'+num-10
			return byte(55+num)
		case num >=36 && num <= 61:
			// 'a'+num-36
			return byte(61+num)
		case num == 62:
			return '-'
		case num == 63:
			return '.'
		default:
			return '*'
	}
}


