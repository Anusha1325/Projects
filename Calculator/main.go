package main

import "github.com/Anusha1325/Projects/operations"

func main() {

	// Add
	operations.AddInt8(-102, 14)
	operations.AddInt16(25483, -12306)
	operations.AddInt32(341798237, -2147483483)
	operations.AddInt64(9213684637386453568, 3784687567427483)
	operations.AddUint8(12, 230)
	operations.AddUint16(23748, 38127)
	operations.AddUint32(3927486376, 38172832)
	operations.AddUint64(3298283477345786, 18437847348723473643)
	operations.AddFloat32(237.754, -3274.834)
	operations.AddFloat64(2347.9475, 334558.4876)

    // Subtract
	operations.SubInt8(-102, 14)
	operations.SubInt16(25483, -12306)
	operations.SubInt32(341798237, -2147483483)
	operations.SubInt64(9213684637386453568, 3784687567427483)
	operations.SubUint8(120, 230)
	operations.SubUint16(23748, 38127)
	operations.SubUint32(3927486376, 38172832)
	operations.SubUint64(3298283477345786465, 18446744073709551615)
	operations.SubFloat32(237.754, -3274.834)
	operations.SubFloat64(2347.9475, 33455875.4876)

}