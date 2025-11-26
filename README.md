# shorty – Pure Go URL Shortener CLI

Like bit.ly, but fully offline. Zero Redis. Zero external services.

### Features
- `shorty create <url>` → https://s.ly/9p558L
- Same URL → same short code (smart deduplication)
- `shorty get <code>` → opens in browser + counts click
- `shorty stats <code>` → shows analytics
- Single `shorty.db` file
- Base62 encoding (real shortener style)

### Download & Run

**Windows**  
Download → [shorty.exe](https://github.com/sabari7320/cli-url-shortner/releases/tag/v1.0.0)  
Double-click or run in terminal.

**Linux & macOS** (run from source – 5 seconds)
```bash
git clone https://github.com/yourname/shorty.git
cd shorty
go run main.go create https://github.com
go run main.go get 9p558L
go run main.go stats 9p558L
