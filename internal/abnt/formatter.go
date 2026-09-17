package abnt

import (
	"abnt-converter-api/internal/crossref"
	"strconv"
	"strings"
)

func authorsFormatter(metadata crossref.ArticleMetadata) string {
	builder := strings.Builder{}
	qtdAuthors := len(metadata.Authors)

	if qtdAuthors == 0 {
		return "AUTOR DESCONHECIDO"
	} else if qtdAuthors <= 3 {
		for i, author := range metadata.Authors {
			if i > 0 {
				builder.WriteString("; ")
			}
			builder.WriteString(strings.ToUpper(author.Family))
			builder.WriteString(", ")
			builder.WriteString(author.Given)
		}
		builder.WriteString(". ")
		return builder.String()
	}

	author := metadata.Authors[0]
	builder.WriteString(strings.ToUpper(author.Family))
	builder.WriteString(", ")
	builder.WriteString(author.Given)
	builder.WriteString(" et al.")
	return builder.String()
}
func magazineVerify(metadata crossref.ArticleMetadata) string {
	if len(metadata.ContainerTitle) == 0 {
		return ""
	}

	return metadata.ContainerTitle[0]
}

func formatField(prefix, value string) string {
	return prefix + " " + value
}

func journalFormatter(metadata crossref.ArticleMetadata) string {
	builder := strings.Builder{}
	magazine := magazineVerify(metadata)
	if magazine != "" {
		builder.WriteString(magazine)
	}

	var details []string

	if metadata.Volume != "" {
		details = append(details, formatField("v.", metadata.Volume))
	}

	if metadata.Issue != "" {
		details = append(details, formatField("n.", metadata.Issue))
	}

	if metadata.Page != "" {
		details = append(details, formatField("p.", metadata.Page))
	}

	year := dateFormatter(metadata)
	if year != "" {
		details = append(details, year)
	}

	if len(details) > 0 {
		if builder.Len() != 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(strings.Join(details, ", "))
	}

	if builder.Len() != 0 {
		builder.WriteString(".")
	}
	return builder.String()
}

func dateFormatter(metadata crossref.ArticleMetadata) string {
	if len(metadata.Issued.DateParts) == 0 {
		return "s.d."
	}

	return strconv.Itoa(metadata.Issued.DateParts[0][0])
}

func titleFormatter(metadata crossref.ArticleMetadata) string {
	if len(metadata.Title) == 0 {
		return "Sem Titulo"
	}

	return strings.TrimSpace(metadata.Title[0])

}

func Formatter(metadata crossref.ArticleMetadata) string {
	builder := strings.Builder{}

	authors := authorsFormatter(metadata)
	title := titleFormatter(metadata)
	details := journalFormatter(metadata)

	builder.WriteString(authors)
	builder.WriteString(" ")
	builder.WriteString(title)
	builder.WriteString(". ")
	builder.WriteString(details)
	builder.WriteString(" ")
	builder.WriteString("[")
	builder.WriteString(metadata.DOI)
	builder.WriteString("]")

	return builder.String()
}
