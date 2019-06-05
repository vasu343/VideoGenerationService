package videogeneration

import (
	"errors"
	"fmt"
)

type VideoGenerateService struct {
	dao Dao // database layer object
}

// GetImage : struct for the payload sent to the getImage service
type GetImage struct {
	PdfID  string `json:"pdfId"`
	PageNo string `json:"pageNumber"`
}

// GetResp : Function responsible for calling external service to get images
func GetImage(url string,Page No string,pdfId string) (Image,error) {
	request := GetImage{PageNo : PageNo,PdfID : pdfId}
	buf, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return Image{},err
	} else {
		defer resp.Body.Close()

		// Image Struct to catch the custom images from the api
		var image Image
		if err := json.NewDecoder(resp.Body).Decode(&image); err != nil {
			log.Println(err)
			return Image{},err
		}
		return image,err
	}
}

// GenerateVideo : Takes in 3 params and returns the string id of the saved video in database or returns an error 
func (svc VideoGenerateService) GenerateVideo(pdfId string, audiofileId string, mapping map[string]int) (string,error){
	GetUrlforImage := config.GetUrl()// url for getting the image
	var images []Image
	for PageNo, duration := range mapping {			
		image,err := GetImage(url,PageNo,pdfId)
		if err != nil {
			log.Println(err)
			return "",err
		}
		images = append(images,image)
	}
	uuid,err := svc.dao.SaveVideo(images)
	return uuid,err
}