# **NBA API**

## Prerequisites

- Golang 1.17
- Mysql

## Step to reproduce

- Clone this repository into your device
- Run your SQL Server
- Create a new database with `CREATE DATABASE Basketball;`
- Run this project with `go run main.go`

## API Reference

#### Get all team

```http
  GET /team
```

#### Get detail team

```http
  GET /team/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int`    | **Required**. Id of item to fetch |

#### Save new team

```http
  POST /team/save
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `team`      | `string`    | **Required**. Name of item to post  |
| `division`      | `string`    | **Required**. Division of item to post  |

#### Edit team

```http
  PUT /team/save/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int`    | **Required**. Id of item to edit  |
| `team`      | `string`    | **Required**. Name of item to edit  |
| `division`      | `string`    | **Required**. Division of item to edit  |

#### Delete team

```http
  DELETE /team/delete/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int`    | **Required**. Id of item to delete  |

