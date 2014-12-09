package controllers

func NewContainers() *ContainersController {
	return &ContainersController{}
}

func NewImages() *ImagesController {
	return &ImagesController{}
}

func NewApps() *AppsController {
	return &AppsController{}
}
