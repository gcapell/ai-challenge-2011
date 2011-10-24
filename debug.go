package main

//String returns an ascii diagram of the map.
func (m *Map) String() string {
	str := ""
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			loc := m.FromRowCol(row, col)
			s := m.ItemAt(loc)
			str += string([]byte{s}) + " "
		}
		str += "\n"
	}
	return str
}

func (m *Map) ItemAt(loc Location) byte {
	s := &m.squares[loc]
	if s.isWater {
		return '%'	// water
	}
	ant, found := m.Ants[loc]
	if found {
		return byte(ant) + 'a'
	}
	if !s.wasSeen {
		return '.'	// unknown
	}
	_, found = m.Food[loc]
	if found {
		return '*'	// food
	}
	return ' '	// land
}

