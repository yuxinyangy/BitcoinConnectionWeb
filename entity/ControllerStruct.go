package entity

type HASH struct {
	Hash string `form:"hash"`
}

type Heights struct {
	Height int `form:"height"`
}

type NewAddress struct {
	Label string `form:"label"`
	Type string `form:"type"`
}

type HashPs struct {
	Blocks int `form:"blocks"`
	Height int `form:"height"`
}

type DataAddress struct {
	Address string `form:"address"`
}
