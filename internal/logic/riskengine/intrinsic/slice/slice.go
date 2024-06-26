package slice

// /
type NumberList []int64

func (s *NumberList) Len() int {
	return len(*s)
}
func (s *NumberList) Del(id int64) {
	for i, v := range *s {
		if v == id {
			*s = append((*s)[:i], (*s)[i+1:]...)
			return
		}
	}
}
func (s *NumberList) Add(id int64) {
	if !s.Exist(id) {
		*s = append(*s, id)
	}
}
func (s *NumberList) Exist(id int64) bool {
	for _, v := range *s {
		if v == id {
			return true
		}
	}
	return false
}
func (s *NumberList) Append(id int64) *NumberList {
	*s = append(*s, id)
	return s
}

func (s *NumberList) New(id int64) *NumberList {
	*s = []int64{id}
	return s
}
