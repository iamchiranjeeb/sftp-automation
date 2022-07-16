package structs



type JsonStructure struct {
	localPath string `json:"localPath"`
	remotePath string `json:"remotePath"`
	user string `json:"user"`
	ip string `json:"ip"`
	port int `json:"port"`
	files []string `json:"files"`
	operation string `json:"operation"`
	ifCoommand bool `json:"execCommand"`
	command string `json:"command"`
}