package main

type Fragment interface {
	Exec(transInfo *TransInfo) error
}

type GetPodAction struct {
}

func (g GetPodAction) Exec(transInfo *TransInfo) error {
	//var fragment Fragment = new(GetPodAction)
	//var fragment Fragment = &GetPodAction{}
	//var fragment Fragment = GetPodAction{}
	return nil
}
