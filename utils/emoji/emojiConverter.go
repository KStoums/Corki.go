package emoji

func ConvertEmoji(emoji string) string {
	var result []byte
	for _, v := range []byte(emoji) {
		result = append(result, v)
	}
	return string(result[:])
}
