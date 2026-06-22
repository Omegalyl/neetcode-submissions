type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encoded string = ""
	if len(strs) == 0 {
		return ""
	}
	for i := range strs {
		byteRP := ""
		for j := range strs[i] {
			byteRP += fmt.Sprintf("%03d", strs[i][j])
		}
		encoded  += byteRP + ","
	}
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return nil
	}
	var decoded []string = []string{}
	strs := strings.Split(encoded, ",")
	strs = strs[:len(strs)-1]

	for i := range strs {
		encodedStr := strs[i]
		decodedBytes := []byte{}
		if len(encodedStr) == 0  || encodedStr == "" {
			decoded = append(decoded,  "")
			continue
		}
		for j:= 0; j < len(encodedStr); j += 3 {
			chunk := encodedStr[j: j+3]
			val, _ := strconv.ParseUint(chunk, 10, 8)
			decodedBytes = append(decodedBytes, byte(val))
		}
		decoded = append(decoded,  string(decodedBytes))
	}
	fmt.Println(decoded)
	return decoded
}
