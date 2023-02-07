package obs

func GetVideoPrefix() string {
	return C.Address + "/videos/"
}

func GetImagePrefix() string {
	return C.Address + "/images/"
}
