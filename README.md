# Projeto de Consulta de APIs

Este projeto foi criado para consultar as seguintes APIs:

- [OpenWeather](https://openweathermap.org/)
- [JSONPlaceholder](https://jsonplaceholder.typicode.com/)

## Requisitos

Antes de começar, certifique-se de ter:

- **Go** instalado (versão **1.23.4**).
Caso não tenha, você pode baixá-lo em https://go.dev/.
- Uma **API Key** do OpenWeather.
Caso não tenha, você pode baixá-lo em https://openweathermap.org/

## Instalação

### 1. Configurar variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto e adicione sua chave da API do OpenWeather:

```ini
OPEN_WEATHER_API_KEY=SuaChaveAqui
```
Crie outro arquivo em `/tests/tests.env` e adicione a chave da API
```ini
OPEN_WEATHER_API_KEY=SuaChaveAqui
```

### 2. Baixar as dependências

Na raiz do projeto execute

```bash
go mod tidy
```

## Executar a aplicação

### 1. Rodar todos os testes

Na raiz do projeto execute

```bash
go test ./...
```

### 2. Rodar a aplicação

Na raiz do projeto execute

```bash
go run .
```

## Rotas

### 1. Usuario

### Pegar as informações do usuario

- **Método**: GET
- **Rota**: `/users/userId` 
- **Parâmetros**: `userId` (um int de 1 até 10)
 
#### Exemplo de Retorno

```json
{
    "id": 1,
    "name": "Leanne Graham",
    "username": "Bret",
    "email": "Sincere@april.biz",
    "address": {
        "street": "Kulas Light",
        "suite": "Apt. 556",
        "city": "Gwenborough",
        "zipcode": "92998-3874",
        "geo": {
            "lat": "-37.3159",
            "lng": "81.1496"
        }
    },
    "phone": "1-770-736-8031 x56442",
    "website": "hildegard.org",
    "company": {
        "name": "Romaguera-Crona",
        "catchPhrase": "Multi-layered client-server neural-net",
        "bs": "harness real-time e-markets"
    }
}
```

### Criar um usuario

- **Método**: POST
- **Rota**: `/users/create`

#### Exemplo de JSON aceito:

```json
{
    "name": "test name",
    "nickname": "test nickname",
    "email": "test@test.com"
}
```

#### Exemplo de Retorno:

```json
{
    "user": {
        "id": 11,
        "name": "test name",
        "nickname": "test nickname",
        "email": "test@test.com"
    },
    "link": "https://jsonplaceholder.typicode.com/users/11"
}
```

**Aviso:** A criação de usuário não cria um usuário de verdade. Os dados são apenas fictícios e retornados como exemplo, pois estamos utilizando a API de teste **JSONPlaceholder**.

### 2. Clima

### Pegar as informações do clima de uma cidade

- **Método**: GET
- **Rota**: `/weathers/{cityName}`
- **Parâmetro**: `cityName` (Nome da cidade)

#### Exemplo de Retorno:

```json
{
    "coord": {
        "lon": 10.307,
        "lat": 45.9484
    },
    "weather": [
        {
            "id": 804,
            "main": "Clouds",
            "description": "nublado",
            "icon": "04n"
        }
    ],
    "base": "stations",
    "main": {
        "temp": 11.29,
        "feels_like": 9.99,
        "temp_min": 9.78,
        "temp_max": 12.28,
        "pressure": 1011,
        "humidity": 58,
        "sea_level": 1011,
        "grnd_level": 851
    },
    "visibility": 10000,
    "wind": {
        "speed": 1.34,
        "deg": 234
    },
    "clouds": {
        "all": 100
    },
    "dt": 1741545314,
    "sys": {
        "type": 2,
        "id": 2011703,
        "country": "IT",
        "sunrise": 1741498981,
        "sunset": 1741540549
    },
    "timezone": 3600,
    "id": 3174478,
    "name": "Campogrande",
    "cod": 200
}
```