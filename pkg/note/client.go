package note

// curl "https://note.com/api/v2/creators/ohatakky/contents?kind=note&page=1" | jq .

const (
	Endpoint       = "https://note.com/api/v2"
	contentPathFmt = "creators/%s/contents?kind=note"
)
