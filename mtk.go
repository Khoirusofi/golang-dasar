package main

func main() {
	var a = 10
	var b = 20
	var d = 30
	var e = a*a - d + b
	var c = a + b
	println(c)
	println(e)

	// Output
	// 30
	// 90
	var i = 10
	i += 10 // i = i + 10
	println(i)
	// Output
	// 20

	var j = 20
	j -= 10 // j = j - 10
	println(j)
	// Output
	// 10

	var k = 30
	k *= 10 // k = k * 10
	println(k)
	// Output
	// 300

	var s = 1
	s++ // s = s + 1
	println(s)
	// Output
	// 2

	var t = 21
	t-- // t = t - 1
	println(t)
	// Output
	// 20

}