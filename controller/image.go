package controller

// ImageController is a controller of docker image
type ImageController struct {
}

// Image creates an instance of ImageController
func Image() (c *ImageController) {
	return &ImageController{
	}
}