# Http Web Server in Deno (TypeScript) with oak

This is a simple web server written in Deno (TypeScript) using oak that serves
static files from a directory.

## Usage

```bash
deno run --allow-net --allow-env --allow-read server.ts
```

## HTTP Endpoints

#### Hello World

```http
GET /hello
```

#### Get list of files in `BASE_PATH`

```http
GET /
```

#### Get file and download it

```http
GET /${path}
```

| Parameter | Type     | Description                               |
| :-------- | :------- | :---------------------------------------- |
| `path`    | `string` | **Required**. Path of the file to provide |
