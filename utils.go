package csvhelper

func splitOnChar(bs []byte, b byte) (spl [][]byte) {
	var (
		index      int
		escapeChar byte
	)

	for i, char := range bs {
		switch char {
		case '"', '\'':
			switch escapeChar {
			case 0:
				escapeChar = char
			case char:
				escapeChar = 0
			}

		case b:
			if escapeChar != 0 {
				continue
			}

			spl = append(spl, bs[index:i])
			index = i + 1
		}
	}

	if index < len(bs)-1 {
		spl = append(spl, bs[index:])
	}

	return
}
