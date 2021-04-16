package model

type Cell struct {
	flags byte
}

var EmptyCell = Cell{}
var SingleCell = Cell{1}

func (cell Cell) Exists() bool {
	return cell != EmptyCell
}
