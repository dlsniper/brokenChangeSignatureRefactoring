package bpkg

type Storer interface {
	DeleteVol(volume string, forceDelete bool) error
}


type XlStorageDiskIDCheck struct {
	storage Storer
}

func (p *XlStorageDiskIDCheck) DeleteVol(volume string, forceDelete bool) (err error) {
	return p.storage.DeleteVol(volume, forceDelete)
}