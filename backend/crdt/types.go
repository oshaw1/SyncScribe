package crdt

type Element struct {
	ID       string
	Content  string
	Position Position
}

type Position struct {
	SiteID   uint     `json:"SiteID"`
	Counter  uint64   `json:"Counter"`
	LastView []uint64 `json:"LastView,omitempty"`
}

type Operation interface {
	Apply(list []Element) []Element
	GetSiteID() uint
	GetCounter() uint64
}

type InsertOperation struct {
	Element Element
	SiteID  uint
	Counter uint64
}

type DeleteOperation struct {
	ElementID string
	SiteID    uint
	Counter   uint64
}
