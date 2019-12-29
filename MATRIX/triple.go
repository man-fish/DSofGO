package MATRIX

import "strconv"

type Triple struct {
	Row,Col,Data	int
}

func NewTriple(row,col,data int) *Triple {
	if row >= 0 && col >= 0 {
		return &Triple{row,col,data}
	}
	return nil
}

func (tp Triple) ToSymmetry() Triple {
	return Triple{tp.Col,tp.Row,tp.Data}
}

func (tp Triple) IsEqual(a,b interface{}) bool {
	if a.(Triple).Row == b.(Triple).Row && a.(Triple).Col == b.(Triple).Col {
		return true
	}
	return false
}

func (tp *Triple) SubRow () {
	tp.Row = tp.Row-1
}

func (tp *Triple) SubCol () {
	tp.Col = tp.Col-1
}

func (tp Triple) CompareTo(a,b interface{}) int {
	if a.(Triple).Row == b.(Triple).Row && a.(Triple).Col == b.(Triple).Col {
		if a.(Triple).Data > b.(Triple).Data {
			return 1
		}else if a.(Triple).Data < b.(Triple).Data {
			return -1
		}else {
			return 0
		}
	}else{
		panic("cant not compare")
	}

}

func (tp *Triple) String() string {
	return  "(" + strconv.Itoa(tp.Row) + "," +  strconv.Itoa(tp.Col) + "," +  strconv.Itoa(tp.Data) + ")"
}
