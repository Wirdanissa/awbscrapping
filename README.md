# AWB Scrapping API (Go)

API untuk scrape informasi status resi (AWB) dan mengembalikan JSON.

## Endpoint

- `GET /awb/{awb_number}`

## Cara Menjalankan (Docker)

```bash
# Build Docker image
docker build -t awbscrapping .

# Jalankan container
docker run -p 8080:8080 awbscrapping
```
