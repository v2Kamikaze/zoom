# Zoom

![ZIGO](https://arkiana.com/wp-content/uploads/2023/07/Zig-vs-Go.webp "Zig 🧡 Go")

## Era pra ser em zig, mas a vida e a falta de tempo me levou para a zona de conforto.

## Objetivo

Por enquanto é um web server com endpoints para processar imagens (png e jpg, até agora).
Futuramente, antes da entrega, vou adicionar uma interface simples (não tão bonita) para a fazer tudo de forma visual.

### Pré-requisitos:

- Go

### Clonar o Repositório:

Utilize o git para clonar este repositório em sua máquina local:

```bash
git clone https://github.com/v2Kamikaze/zoom.git
```

### Instalar as Dependências:

Acesse a pasta do projeto e execute:

```bash
go mod tidy
```

## Executando:

Para rodar o servidor:

```bash
go run ./cmd/run .
```

Para rodar concorrentemente todos os filtros nas imagens presentes em `assets`(bom para testar):

```bash
go run cmd/batch .
```

Para limpar as imagens geradas por `batch` em `assets`:

```bash
go run cmd/reset .
```

### Rotas implementadas até agora:

| Método | Endpoint                               | Parâmetros                                                                                      | Retorno                                                           |
| ------ | -------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------- |
| POST   | /api/effects/negative                  | query={}                                                                                        | Nova imagem com efeito negativo aplicado                          |
| POST   | /api/effects/sobel-x                   | query={}                                                                                        | Nova imagem com filtro Sobel no eixo X                            |
| POST   | /api/effects/sobel-y                   | query={}                                                                                        | Nova imagem com filtro Sobel no eixo Y                            |
| POST   | /api/effects/sobel-mag                 | query={}                                                                                        | Nova imagem com magnitude do filtro Sobel                         |
| POST   | /api/effects/gaussian                  | query={ks: Tamanho do kernel (padrão 3), s: Sigma (padrão 1.0)}                                 | Nova imagem com filtro Gaussiano aplicado                         |
| POST   | /api/effects/laplacian                 | query={ks: Tamanho do kernel (padrão 3)}                                                        | Nova imagem com filtro Laplaciano aplicado                        |
| POST   | /api/effects/mean                      | query={ks: Tamanho do kernel (padrão 3)}                                                        | Nova imagem com filtro de média aplicado                          |
| POST   | /api/effects/bin                       | query={t: Limiar de binarização (padrão 128)}                                                   | Nova imagem binarizada                                            |
| POST   | /api/effects/gamma                     | query={g: Valor do gamma (padrão 2.0), c: Constante (padrão 1.0)}                               | Nova imagem com correção gama aplicada                            |
| POST   | /api/effects/high-boost                | query={ks: Tamanho do kernel (padrão 3), s: Sigma (padrão 1.0), k: Fator de boost (padrão 1.5)} | Nova imagem com filtro High Boost aplicado                        |
| POST   | /api/effects/sharpening                | query={ks: Tamanho do kernel (padrão 3)}                                                        | Nova imagem com filtro de nitidez aplicado                        |
| POST   | /api/transform/scale/bilinear          | query={x: Escala X (padrão 1), y: Escala Y (padrão 1)}                                          | Nova imagem escalada com interpolação bilinear                    |
| POST   | /api/transform/scale/nearest-neighbor  | query={x: Escala X (padrão 1), y: Escala Y (padrão 1)}                                          | Nova imagem escalada com interpolação por vizinho mais próximo    |
| POST   | /api/transform/rotate/bilinear         | query={a: Ângulo de rotação (padrão 0)}                                                         | Nova imagem rotacionada com interpolação bilinear                 |
| POST   | /api/transform/rotate/nearest-neighbor | query={a: Ângulo de rotação (padrão 0)}                                                         | Nova imagem rotacionada com interpolação por vizinho mais próximo |
