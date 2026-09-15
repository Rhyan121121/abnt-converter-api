# ABNT Converter

Ferramenta em Go para converter citações acadêmicas e DOIs em referências formatadas no padrão **ABNT NBR 6023** automaticamente, consultando os metadados diretamente na API da Crossref.

---

## Funcionalidades

- **Extração Automática:** Isola o identificador DOI a partir de citações brutas (ex: formato APA).
- **Consulta à Crossref:** Busca automática de autores, título, periódico, volume, fascículo, páginas e ano de publicação.
- **Formatação ABNT NBR 6023:** Aplica regras de autoria (caixa alta, iniciais, *et al.*), destaque tipográfico do periódico e links resolúveis de DOI.
- **Interface de Terminal:** Entrada interativa via formulário estilizado com [Huh?](https://github.com/charmbracelet/huh) e [Lip Gloss](https://github.com/charmbracelet/lipgloss).
- **Cópia Automática:** Envia a referência final diretamente para a área de transferência do sistema.

---

## Estrutura do Projeto

```text
.
├── cmd/
│   └── main.go                 # Ponto de entrada e interface CLI
├── internal/
│   ├── abnt/
│   │   ├── formatter.go        # Lógica de formatação ABNT NBR 6023
│   │   └── formatter_test.go   # Testes unitários do formatador
│   └── crossref/
│       ├── client.go           # Cliente HTTP e sanitização de DOI
│       └── models.go           # Structs para decodificação da API
├── go.mod
└── go.sum


```

## Próximos Passos (Roadmap)

- [ ] **API Backend em Go:** Criar endpoints REST/HTTP expondo a busca e a formatação ABNT.
- [ ] **Front-End Web:** Desenvolver interface amigável e responsiva para navegadores.
- [ ] **Deploy & Hospedagem:** Publicar o front-end na Netlify integrado à API.
- [ ] **Conversão em Lote:** Permitir colar múltiplas citações/DOIs de uma só vez.
- [ ] **Novos Tipos de Documentos:** Suporte a livros, teses e anais de congressos.