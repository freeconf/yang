package node

type EventMulticast struct {
	//A Events
	//B Events
}

//func (multi *EventMulticast) AddListener(l *Listener) {
//	multi.A.AddListener(l)
//	multi.B.AddListener(l)
//}
//
//func (multi *EventMulticast) RemoveListener(l *Listener) {
//	multi.A.RemoveListener(l)
//	multi.B.RemoveListener(l)
//}
//
//func (multi *EventMulticast) Fire(path *Path, e Event) (err error) {
//	if err = multi.A.Fire(path, e); err == nil {
//		err = multi.B.Fire(path, e)
//	}
//	return
//}
