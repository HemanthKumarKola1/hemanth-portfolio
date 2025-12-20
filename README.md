# Go Engineer Portfolio

A static site generator built in Go for showcasing your skills and technical articles.

## Features

- 📝 Markdown-based article system
- 🎨 Responsive design
- ⚡ Fast static site generation
- 🚀 Easy deployment

## Quick Start

1. **Build the site:**

   ```bash
   ./build.sh
   ```

2. **Preview locally:**
   Open `docs/index.html` in your browser

## Adding Articles

1. Create `.md` files in the `content/` directory
2. Run `./build.sh` to regenerate the site

## Deployment Options

### GitHub Pages (Free)

1. Push to GitHub repository
2. Enable GitHub Pages in repository settings
3. Upload `docs/` contents

### Netlify (Free)

1. Connect your GitHub repository
2. Set build command: `go run main.go`
3. Set publish directory: `docs`

### Vercel (Free)

1. Import your GitHub repository
2. Framework preset: Other
3. Build command: `go run main.go`
4. Output directory: `docs`

## Customization

- Edit skills in `main.go`
- Modify templates for design changes
- Update CSS in the `copyAssets()` function

## Project Structure

```
portfolio/
├── main.go           # Site generator
├── content/          # Markdown articles
├── docs/            # Generated site
├── build.sh         # Build script
└── go.mod           # Go dependencies
```
