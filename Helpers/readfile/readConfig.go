package readfile

import(
	"path"
)

func ReadFile(filePath string)([]map[string]interface{}){

	pathExt := path.Ext(filePath)

	var blankData []map[string]interface{}

	switch pathExt {
		
		case ".json":
			return ReadJson(filePath)
		
		case ".yaml":
			return ReadYml(filePath)

		case ".yml":
			return ReadYml(filePath)

	}

	return blankData
}