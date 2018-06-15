package controller

// VolumeController is a controller of docker volume
type VolumeController struct {
}

// Volume creates an instance of VolumeController
func Volume() (c *VolumeController) {
	return &VolumeController{
	}
}
