package internal

import "github.com/matrixorigin/wecom-group-bot/internal/utils/poolutils"

func GenURL(url string, paths ...string) string {
	builder := poolutils.GetStringBuilder()
	defer poolutils.PutStringBuilder(builder)
	builder.WriteString(url)
	for _, path := range paths {
		builder.WriteString(path)
	}
	return builder.String()
}
