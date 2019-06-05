package videogeneration

// VideoInputObject : input object for the service
type VideoInputObject struct {
	pdfID       string
	audiofileID string
	mapping     map[string]int
}

// Image : To catch the images from external api
type Image struct {
	// custom implementation depending on some details
}
