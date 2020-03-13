package convert


// 将十进制转换为62进制   0-9a-zA-Z 六十二进制
func TransTo62(id int64)string{
	// 1 -- > 1
	// 10-- > a
	// 61-- > Z
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var shortUrl []byte
	for{
		var result byte
		number := id % 62
		result = charset[number]
		var tmp []byte
		tmp = append(tmp,result)
		shortUrl = append(tmp,shortUrl...)
		id = id / 62
		if id == 0{
			break
		}
	}
	return string(shortUrl)
}