package csvhelper

func splitOnChar(bs []byte, b byte) (spl [][]byte) {
	var (
		index       int
		escapeState bool
		quoteState  bool
	)

	for i, char := range bs {
		if escapeState {
			escapeState = false
			continue
		}

		switch char {
		case '"':
			quoteState = !quoteState
		case '\\':
			escapeState = true
		case b:
			if quoteState {
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
