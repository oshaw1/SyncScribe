package crdt

func (op InsertOperation) Apply(list []Element) []Element {
	position := generatePosition(&OrderedListCRDT{Elements: list}, op.SiteID)
	op.Element.Position = position
	op.Counter = position.Counter
	return insertElement(list, op.Element)
}

func (op InsertOperation) GetSiteID() uint {
	return op.SiteID
}

func (op InsertOperation) GetCounter() uint64 {
	return op.Counter
}

func (op DeleteOperation) Apply(list []Element) []Element {
	// Remove the element from the list based on the ElementID
	return removeElement(list, op.ElementID)
}

func (op DeleteOperation) GetSiteID() uint {
	return op.SiteID
}

func (op DeleteOperation) GetCounter() uint64 {
	return op.Counter
}
