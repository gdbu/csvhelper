package csvhelper

func splitOnChar(bs []byte, b byte) (spl [][]byte) {
	var (
		index       int
		escapeState bool
		quoteState  bool
	)

	for i, char := range bs {
		if escapeState {
			// We are currently in an escaped state for this character
			// Escaped state only lasts for one character, set back to false
			escapeState = false
			// This character was escaped, continue
			continue
		}

		switch char {
		case '"':
			// We encounted a double quote, inverse the quoted state
			quoteState = !quoteState
		case '\\':
			// We encounted a backslash, set the escape state to true
			escapeState = true
		case b:
			if quoteState {
				// We cannot split on during an active quote state, continue
				continue
			}

			// Append the part to the split slice
			spl = append(spl, bs[index:i])
			// Update the index
			index = i + 1
		}
	}

	if index < len(bs)-1 {
		spl = append(spl, bs[index:])
	}

	return
}
