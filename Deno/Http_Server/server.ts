import { Application, Router } from "https://deno.land/x/oak/mod.ts";

const BASE_PATH = Deno.env.get("BASE_PATH") || "./";
const HOST = Deno.env.get("HOST") || "";
const PORT = Number(Deno.env.get("PORT")) || 8000;
const app = new Application();
const router = new Router();

app.use(router.routes());

router.get("/hello", async (ctx) => {
  ctx.response.body = "Hello World!";
});

router.get("/", async (ctx) => {
  //list files
  const files = await Deno.readDir(BASE_PATH);
  let fileNames: Object = {};
  for await (const file of files) {
    // get path of file
    // get current path
    const cwd_path = Deno.cwd();
    const path: string = cwd_path + "/" + file.name;
    fileNames[file.name] = `http://localhost:${PORT}${path}`;
  }
  ctx.response.body = fileNames;
});

// get files
router.get("/(.*)", async (ctx) => {
  const path = ctx.request.url.pathname;
  const file = await Deno.open(path);
  ctx.response.body = file;
});

await app.listen({ host: HOST, port: PORT });
