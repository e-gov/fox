package fox

type Fox struct {
	Name 		string	`json:"name"`
	Parents 	[]string `json:"parents"`
	Uuid		string 	`json:"uuid"`
}

// Compare compares two foxes
func Compare(a, b Fox) bool{
	if &a == &b {
		return true
	}
	
	if a.Name != b.Name || a.Uuid != b.Uuid{
		return false
	}
	
	if len(a.Parents) != len(b.Parents){
		return false
	}

	for i := range a.Parents{
		if a.Parents[i] != b.Parents[i] {
			return false
		}
	}
	
	return true
}