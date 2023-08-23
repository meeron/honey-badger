package api

func getMetaByContentType(contentType string) byte {
	switch contentType {
	case "text/plain":
		return 110
	case "application/json":
		return 120
	case "application/xml":
		return 121
	case "image/jpeg":
		return 130
	case "image/png":
		return 131
	default:
		return 0
	}
}

func getContentTypeByMeta(meta byte) string {
	switch meta {
	case 110:
		return "text/plain"
	case 120:
		return "application/json"
	case 121:
		return "application/xml"
	case 130:
		return "image/jpeg"
	case 131:
		return "mage/png"
	default:
		return "application/octet-stream"
	}
}
