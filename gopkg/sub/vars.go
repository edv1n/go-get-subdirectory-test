package gopkg

import "fmt"

var (
	MetaTagContent = fmt.Sprintf(`<meta name="go-import" content="%s git %s %s"></head>`, Module, Repo, Subdir)
)
