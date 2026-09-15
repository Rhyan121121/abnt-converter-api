package abnt

import (
	"abnt-converter-api/internal/crossref"
	"testing"
)

func TestAuthorsFormatter(t *testing.T) {
	tests := []struct {
		name     string
		metadata crossref.ArticleMetadata
		expected string
	}{
		{
			name: "Lista de autores vazia",
			metadata: crossref.ArticleMetadata{
				Authors: []crossref.Author{},
			},
			expected: "AUTOR DESCONHECIDO",
		},
		{
			name: "Autor unico com sobrenome em caixa alta",
			metadata: crossref.ArticleMetadata{
				Authors: []crossref.Author{
					{Given: "Alan", Family: "Turing"},
				},
			},
			expected: "TURING, Alan",
		},
		{
			name: "Dois autores separados por ponto e virgula",
			metadata: crossref.ArticleMetadata{
				Authors: []crossref.Author{
					{Given: "Alan", Family: "Turing"},
					{Given: "Ada", Family: "Lovelace"},
				},
			},
			expected: "TURING, Alan; LOVELACE, Ada",
		},
		{
			name: "Tres autores (caso limite de listagem completa)",
			metadata: crossref.ArticleMetadata{
				Authors: []crossref.Author{
					{Given: "Alan", Family: "Turing"},
					{Given: "Ada", Family: "Lovelace"},
					{Given: "Grace", Family: "Hopper"},
				},
			},
			expected: "TURING, Alan; LOVELACE, Ada; HOPPER, Grace",
		},
		{
			name: "Quatro ou mais autores (deve exibir primeiro autor com et al.)",
			metadata: crossref.ArticleMetadata{
				Authors: []crossref.Author{
					{Given: "Alan", Family: "Turing"},
					{Given: "Ada", Family: "Lovelace"},
					{Given: "Grace", Family: "Hopper"},
					{Given: "Dennis", Family: "Ritchie"},
				},
			},
			expected: "TURING, Alan et al.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := authorsFormatter(tt.metadata)

			if got != tt.expected {
				t.Errorf("authorsFormatter() = %q, esperado %q", got, tt.expected)
			}
		})
	}
}

func TestTitleFormatter(t *testing.T) {
	tests := []struct {
		name     string
		metadata crossref.ArticleMetadata
		expected string
	}{
		{
			name: "Titulo vazio",
			metadata: crossref.ArticleMetadata{
				Title: []string{},
			},
			expected: "Sem Titulo",
		},
		{
			name: "Titulo com espacos em branco nas extremidades",
			metadata: crossref.ArticleMetadata{
				Title: []string{"   Estruturas de Dados e Algoritmos.   \t"},
			},
			expected: "Estruturas de Dados e Algoritmos.", //
		},
		{
			name: "Titulo normal sem espacos adicionais",
			metadata: crossref.ArticleMetadata{
				Title: []string{"Estruturas de Dados e Algoritmos."},
			},
			expected: "Estruturas de Dados e Algoritmos.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := titleFormatter(tt.metadata)

			if got != tt.expected {
				t.Errorf("titleFormatter() = %q, esperado %q", got, tt.expected)
			}
		})
	}
}

func TestDateFormatter(t *testing.T) {
	tests := []struct {
		name     string
		metadata crossref.ArticleMetadata
		expected string
	}{
		{
			name: "DateParts vazio deve retornar fallback s.d.",
			metadata: crossref.ArticleMetadata{
				Issued: crossref.IssuedDate{
					DateParts: [][]int{},
				},
			},
			expected: "s.d.",
		},
		{
			name: "Data completa com ano mes e dia deve extrair apenas o ano",
			metadata: crossref.ArticleMetadata{
				Issued: crossref.IssuedDate{
					DateParts: [][]int{{2016, 8, 25}},
				},
			},
			expected: "2016",
		},
		{
			name: "Data contendo somente o ano numerico",
			metadata: crossref.ArticleMetadata{
				Issued: crossref.IssuedDate{
					DateParts: [][]int{{2023}},
				},
			},
			expected: "2023",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dateFormatter(tt.metadata)

			if got != tt.expected {
				t.Errorf("dateFormatter() = %q, esperado %q", got, tt.expected)
			}
		})
	}
}

func TestJournalFormatter(t *testing.T) {
	// Tabela de cenários cobrindo variações comuns de metadados em periódicos
	tests := []struct {
		name     string                   // Descrição contextual do caso de teste
		metadata crossref.ArticleMetadata // Dados simulados de entrada da Crossref
		expected string                   // Saída textual esperada segundo a ABNT
	}{
		{
			name: "Publicacao completa com revista, volume, fasciculo, pagina e ano",
			metadata: crossref.ArticleMetadata{
				ContainerTitle: []string{"Nature"},
				Volume:         "500",
				Issue:          "1",
				Page:           "45-50",
				Issued: crossref.IssuedDate{
					DateParts: [][]int{{2016}},
				},
			},
			expected: "Nature, v. 500, n. 1, p. 45-50, 2016.", // Formato pleno da ABNT
		},
		{
			name: "Publicacao sem fasciculo e sem paginas",
			metadata: crossref.ArticleMetadata{
				ContainerTitle: []string{"Science"},
				Volume:         "12",
				Issue:          "",
				Page:           "",
				Issued: crossref.IssuedDate{
					DateParts: [][]int{{2020}},
				},
			},
			expected: "Science, v. 12, 2020.", // Valida omissao limpa sem virgulas orfas
		},
		{
			name: "Publicacao continua apenas com identificador de pagina e ano",
			metadata: crossref.ArticleMetadata{
				ContainerTitle: []string{"PLoS ONE"},
				Volume:         "",
				Issue:          "",
				Page:           "e0123456",
				Issued: crossref.IssuedDate{
					DateParts: [][]int{{2022}},
				},
			},
			expected: "PLoS ONE, p. e0123456, 2022.", // Valida artigo eletronico sem volume e issue
		},
		{
			name: "Apenas nome do periodico e ano de publicacao",
			metadata: crossref.ArticleMetadata{
				ContainerTitle: []string{"The Lancet"},
				Volume:         "",
				Issue:          "",
				Page:           "",
				Issued: crossref.IssuedDate{
					DateParts: [][]int{{2018}},
				},
			},
			expected: "The Lancet, 2018.", // Junta diretamente revista com o ano
		},
	}

	// Executa a iteracao sobre todos os cenarios de teste da tabela
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Executa a funcao que monta o bloco de periodico
			got := journalFormatter(tt.metadata)

			// Valida se o texto formatado corresponde rigorosamente ao padrao ABNT
			if got != tt.expected {
				t.Errorf("journalFormatter() = %q, esperado %q", got, tt.expected)
			}
		})
	}
}

func TestFormatter(t *testing.T) {
	// Tabela de cenarios para validar a agregacao final de todos os blocos ABNT
	tests := []struct {
		name     string                   // Descricao contextual do caso de teste
		metadata crossref.ArticleMetadata // Metadados completos simulados da resposta Crossref
		expected string                   // String consolidada esperada com todas as pontuacoes
	}{
		{
			name: "Formatacao completa com autores, titulo, periodico e DOI",
			metadata: crossref.ArticleMetadata{
				Authors: []crossref.Author{
					{Given: "John", Family: "Doe"},
				},
				Title:          []string{"Machine learning in medicine"},
				ContainerTitle: []string{"Nature Medicine"},
				Volume:         "28",
				Issue:          "4",
				Page:           "100-105",
				DOI:            "10.1038/s41591-022-01234-5",
				Issued: crossref.IssuedDate{
					DateParts: [][]int{{2022}},
				},
			},
			// Espera todos os blocos separados corretamente por ponto e espaco, terminando com [DOI]
			expected: "DOE, John. Machine learning in medicine. Nature Medicine, v. 28, n. 4, p. 100-105, 2022. [10.1038/s41591-022-01234-5]",
		},
		{
			name: "Status diferente de ok deve retornar string vazia",
			metadata: crossref.ArticleMetadata{
				DOI: "10.1038/nature123",
			},
			expected: "", // Espera retorno vazio imediato pelo guard clause
		},
	}

	// Executa cada cenario de teste individualmente
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Invoca a funcao montadora principal
			got := Formatter(tt.metadata)

			// Valida se a referencia completa gerada bate com a expectativa
			if got != tt.expected {
				t.Errorf("\ngot:      %q\nexpected: %q", got, tt.expected)
			}
		})
	}
}
