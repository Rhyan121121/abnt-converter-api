# ABNT Converter 📚✨

Ferramenta CLI desenvolvida em Go para automatizar a conversão de citações acadêmicas brutas e identificadores DOI diretamente no padrão **ABNT NBR 6023**, com busca automática de metadados na API da Crossref e envio imediato para a área de transferência.

---

## 💡 Sobre o Projeto

O projeto nasceu de uma necessidade prática: **facilitar a rotina de pesquisa acadêmica da minha namorada**.

Ao consultar artigos científicos em bases biomédicas como o **PubMed**, as referências vêm comumente padronizadas no formato APA. Ter que formatar manualmente cada uma das referências para as normas brasileiras exigidas pela faculdade e periódicos (ABNT NBR 6023) é um processo repetitivo, demorado e sujeito a erros de pontuação.

Esta aplicação resolve esse atrito: basta colar a citação bruta copiada da página do artigo. O programa identifica e isola o DOI, busca os metadados bibliográficos atualizados na **Crossref**, aplica as regras da ABNT e **já envia a referência formatada diretamente para o clipboard (`Ctrl + V`)**.

---

## 🧪 Exemplo de Teste

### 1. Entrada (Citação no formato APA copiada do PubMed)
```text
Broutier, L., Andersson-Rolf, A., Hindley, C. J., Boj, S. F., Clevers, H., Koo, B. K., & Huch, M. (2016). Culture and establishment of self-renewing human and mouse adult liver and pancreas 3D organoids and their genetic manipulation. Nature protocols, 11(9), 1724–1743. [https://doi.org/10.1038/nprot.2016.097](https://doi.org/10.1038/nprot.2016.097)
```

### 2. Saída (Referência ABNT NBR 6023 gerada e enviada para o Clipboard)
```text
BROUTIER, L. et al. Culture and establishment of self-renewing human and mouse adult liver and pancreas 3D organoids and their genetic manipulation. Nature Protocols, v. 11, n. 9, p. 1724-1743, 2016. DOI: [https://doi.org/10.1038/nprot.2016.097](https://doi.org/10.1038/nprot.2016.097).
```

---

## 🚀 Funcionalidades

- **Extração Resiliente de DOI:** Isola automaticamente o padrão `10.xxxx/...` mesmo quando a citação contém nomes de múltiplos autores, títulos longos, pontuações e URLs.
- **Integração com a Crossref:** Consulta metadados completos (autores, periódico, volume, fascículo, páginas, ano e link do DOI).
- **Formatação ABNT Rigorosa:**
    - Sobrenomes de autores em caixa alta (`SOBRENOME, Prenome/Inicial`).
    - Aplicação automática de *et al.* quando houver mais de 3 autores.
    - Destaque do título da publicação/periódico.
    - Formatação padronizada de fascículo (`v.`, `n.`, `p.`, `ano`).
- **Interface TUI Interativa:** Formulário de entrada no terminal com componentes visuais usando [Huh?](https://github.com/charmbracelet/huh) e [Lip Gloss](https://github.com/charmbracelet/lipgloss).
- **Clipboard Integrado:** Cópia instantânea para o sistema operacional logo após a geração da referência.

---

## 📂 Estrutura do Repositório

```text
.
├── cmd/
│   └── main.go                 # Entrada CLI, formulário Huh e cópia para o clipboard
├── internal/
│   ├── abnt/
│   │   ├── formatter.go        # Regras de formatação da norma ABNT NBR 6023
│   │   └── formatter_test.go   # Testes unitários com casos reais de citação
│   └── crossref/
│       ├── client.go           # Cliente HTTP, extração de DOI e escape de URL (%2F)
│       └── models.go           # Structs para deserialização do JSON da Crossref
├── go.mod                      # Módulo e dependências do projeto
├── go.sum                      # Checksums das dependências
└── README.md                   # Documentação do projeto
```

---

## ⚙️ Pré-requisitos

- **Go:** Versão 1.22 ou superior instalada.
- **Suporte a Clipboard no Sistema Operacional (Linux):**
    - Wayland: pacote `wl-clipboard`
    - X11: pacotes `xclip` ou `xsel`
      *(No Windows e macOS o clipboard funciona nativamente sem pacotes adicionais)*

---

## 🛠️ Instalação e Uso

1. **Clone o repositório:**
   ```bash
   git clone [https://github.com/seu-usuario/abnt-converter-api.git](https://github.com/seu-usuario/abnt-converter-api.git)
   cd abnt-converter-api
   ```

2. **Instale as dependências:**
   ```bash
   go mod tidy
   ```

3. **Execute os testes:**
   ```bash
   go test ./...
   ```

4. **Inicie o programa:**
   ```bash
   go run cmd/main.go
   ```

5. **(Opcional) Gere o executável binário:**
   ```bash
   go build -o abnt cmd/main.go
   ./abnt
   ```

---

## 🗺️ Próximos Passos (Roadmap)

- [ ] **API HTTP em Go:** Expor endpoints REST para permitir o consumo por outros clientes.
- [ ] **Interface Web Responsiva:** Criar front-end limpo para uso direto pelo navegador.
- [ ] **Deploy no Netlify:** Publicar a versão web de forma contínua e acessível publicamente.
