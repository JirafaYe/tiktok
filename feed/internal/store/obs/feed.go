package obs

func GetVideoPrefix() string {
	return "http://" + C.Address + "/videos/"
}

func GetImagePrefix() string {
	return "http://" + C.Address + "/images/"
}
