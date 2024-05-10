package crdt

type OrderedListCRDT struct {
	Elements []Element
}

func (ol *OrderedListCRDT) Apply(op Operation) {
	ol.Elements = op.Apply(ol.Elements)
}

func (ol *OrderedListCRDT) IncrementCounter(siteID uint) uint64 {
	for i, elem := range ol.Elements {
		if elem.Position.SiteID == siteID {
			ol.Elements[i].Position.Counter++
			return ol.Elements[i].Position.Counter
		}
	}
	// If the SiteID is not found, create a new element with Counter = 1
	counter := uint64(1)
	ol.Elements = append(ol.Elements, Element{Position: Position{SiteID: siteID, Counter: counter}})
	return counter
}

func (ol *OrderedListCRDT) UpdateLastView(siteID uint) {
	found := false
	lastView := make([]uint64, len(ol.Elements))
	for i, elem := range ol.Elements {
		lastView[i] = elem.Position.Counter
		if elem.Position.SiteID == siteID {
			found = true
		}
	}
	if found {
		for i, elem := range ol.Elements {
			if elem.Position.SiteID == siteID {
				ol.Elements[i].Position.LastView = lastView
				break
			}
		}
	} else {
		// If the SiteID is not found, create a new element with the LastView
		ol.Elements = append(ol.Elements, Element{Position: Position{SiteID: siteID, LastView: lastView}})
	}
}
