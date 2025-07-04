package shortener

import "log"

var base62 = []byte("1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

func MakeShort(id int64) string {
	var buf [10]byte
	i := len(buf)

	for id > 0 {
		i--
		buf[i] = base62[id%62]
		id /= 62
	}
	log.Println("Сокращенная ссылка: ", string(buf[i:]))
	return string(buf[i:])
}