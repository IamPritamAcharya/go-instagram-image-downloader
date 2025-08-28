
# 📸 Instagram Downloader (Go + Gin)

A simple Go service powered by **Gin** that extracts **direct media URLs** (image or video) from Instagram posts.  
Paste an Instagram post URL, and get the actual `.jpg` or `.mp4` link in return.

---

## Live Demo

Deployed on Render:

**Endpoint:**

```

https://go-instagram-image-downloader.onrender.com/download?url=

https://go-instagram-image-downloader.onrender.com/download?url=\<INSTAGRAM_POST_URL>

````

**Example:**

```bash
curl "https://go-instagram-image-downloader.onrender.com/download?url=https://www.instagram.com/p/DN3FbNQWJCF/"
````

Response:

```json
{
  "media_url": "https://instagram.fna.fbcdn.net/..."
}
```

---

## 🚀 Features

* Retrieve direct media links from Instagram posts (**images only**).

---

## 📂 Project Structure

```
instagram-downloader/
│── main.go                 
│── go.mod
│── go.sum
│── handlers/               
│    └── download.go
│── services/               
│    └── instagram.go
│── docs/
│    └── README.md
```

---

## 🛠 Getting Started Locally

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/instagram-downloader.git
cd instagram-downloader
```

### 2. Install dependencies

```bash
go mod tidy
```

Dependencies:

* [`github.com/gin-gonic/gin`](https://github.com/gin-gonic/gin)
* [`github.com/PuerkitoBio/goquery`](https://github.com/PuerkitoBio/goquery)

### 3. Run the server

```bash
go run main.go
```

Server runs at:

```
http://localhost:8080
```

### 4. Test locally

```bash
curl "http://localhost:8080/download?url=https://www.instagram.com/p/DN3FbNQWJCF/"
```

Response:

```json
{
  "media_url": "https://instagram.fna.fbcdn.net/...jpg"
}
```

---

## ⚠️ Notes & Limitations

* Only works with **public Instagram posts**.
* Instagram changes its frontend frequently → scraper may require updates.
* **Use responsibly**: downloading or redistributing content without permission may violate Instagram’s **Terms of Service**.
