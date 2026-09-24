# forge-demo

A tiny Go web app used to test [Forge](https://forge.luccinmasirika.com), a self-hosted deployment platform.

Every push to `main` is built and deployed to https://forge-demo.apps.luccinmasirika.com.

```bash
go run .              # http://localhost:8080
docker build -t forge-demo . && docker run -p 8080:8080 forge-demo
```
