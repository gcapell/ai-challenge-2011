package main

type PartRecord struct {
	part int
	ant *Ant
}

func (a Point) CouldInfluence (b Point ) bool {
	return a.CrowDistance2(b) <= INFLUENCERADIUS2
}

func (m *Map) nearbyEnemies(friendlies []*Ant) []Point {
	reply := make([]Point, 0)
	for _, e := range m.enemies {
		for _, a := range friendlies {
			if e.CouldInfluence(a.p) {
				reply = append(reply, e)
				break
			}
		}
	}
	return reply
}

// Partition all friendly ants into spheres of influence
func (m *Map) partitionFriendlies() [][]*Ant {
	assigned := make([]PartRecord, 0,  len(m.myAnts))

	nextPart := 1

	for _, a := range m.myAnts {

		// Assign initial partition
		partnum := nextPart
		nextPart += 1

		for i := range assigned {
			if a.p.CouldInfluence(assigned[i].ant.p) {
				// Join to existing partition
				partnum = assigned[i].part

				// Possibly stitch together some other partitions
				for j := range assigned[i+1:] {
					if a.p.CouldInfluence(assigned[j].ant.p) {
						assigned[j].part = partnum
					}
				}
				break
			}
		}
		assigned = append(assigned, PartRecord{partnum, a})
	}
	

	// Transform into slice of slices
	seen := make(map[int] [] *Ant)
	for _, a := range assigned {
		seen[a.part] = append(seen[a.part], a.ant)
	}
	
	reply := make([][]*Ant, 0, len(seen))
	for _, ants := range seen {
		reply = append(reply, ants)
	}
	return reply
}

