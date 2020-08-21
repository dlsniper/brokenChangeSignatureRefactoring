package opkg

import (
	"net/http"

	"awesomeProject35/bpkg"
)

type StorageRESTServer struct {
	storage *bpkg.XlStorageDiskIDCheck
}

// DeleteVolumeHandler - delete a volume.
func (s *StorageRESTServer) DeleteVolHandler(w http.ResponseWriter, r *http.Request) {
	var volume string
	var forceDelete bool
	err := s.storage.DeleteVol(volume, forceDelete)
	if err != nil {
		//
	}
}
