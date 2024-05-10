package crdt

func comparePositions(a, b Position) int {
	if a.SiteID < b.SiteID {
		return -1
	} else if a.SiteID > b.SiteID {
		return 1
	}

	if a.Counter < b.Counter {
		return -1
	} else if a.Counter > b.Counter {
		return 1
	}

	for i := range a.LastView {
		if i >= len(b.LastView) {
			return 1
		}
		if a.LastView[i] < b.LastView[i] {
			return -1
		} else if a.LastView[i] > b.LastView[i] {
			return 1
		}
	}

	if len(a.LastView) < len(b.LastView) {
		return -1
	}

	return 0
}

func insertElement(list []Element, elem Element) []Element {
	i := 0
	for i < len(list) {
		if comparePositions(elem.Position, list[i].Position) < 0 {
			break
		}
		i++
	}
	return append(list[:i], append([]Element{elem}, list[i:]...)...)
}

func removeElement(list []Element, elemID string) []Element {
	for i, elem := range list {
		if elem.ID == elemID {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func generatePosition(ol *OrderedListCRDT, siteID uint) Position {
	counter := ol.IncrementCounter(siteID)
	ol.UpdateLastView(siteID)

	return Position{
		SiteID:   siteID,
		Counter:  counter,
		LastView: ol.Elements[len(ol.Elements)-1].Position.LastView,
	}
}
