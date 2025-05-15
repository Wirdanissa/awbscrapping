# AWB Scrapping API (Go)

API untuk scrape informasi status resi (AWB) dan mengembalikan JSON.

## Endpoint

- `GET /awb/{awb_number}`

# Successfull AWB Scrapping
Endpoint : http://localhost:8080/awb/3452440340005

# Invalid AWB Handling
Endping : http://localhost:8080/awb/00000

## Cara Menjalankan (Docker)

```bash
# Build Docker image
docker build -t awbscrapping .

# Jalankan container
docker run -p 8080:8080 awbscrapping
```
