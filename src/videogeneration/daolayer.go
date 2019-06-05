package videogeneration

type Dao struct {
	config.DatabaseDao
}

 // SaveVideo : Saves video file to a file store and returns a stringId as a reference to the file
func (dao *Dao) SaveVideo(images []Image,pdfId string,audiofileID string) (string, error) {
	// Here we can have the concrete implementation to save the video file. The video file can be saved in something like a distributed file storage 
	// and we can also maintain a reference to the video like this in cassandra or redis :- VideoreferenceId PdfId AudiofileID
	// The above mapping can be used for querying a video by pdf or audio id.
}
