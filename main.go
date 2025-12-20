package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/russross/blackfriday/v2"
)

type Article struct {
	Title   string
	Date    string
	Slug    string
	Content template.HTML
	Summary string
}

type SocialLink struct {
	Name string
	URL  string
}

type PageData struct {
	Title       string
	Articles    []Article
	Skills      []string
	SocialLinks []SocialLink
}

func main() {
	// Create output directory
	os.MkdirAll("dist", 0755)
	os.MkdirAll("dist/articles", 0755)

	// Generate site
	generateHomePage()
	generateArticles()
	copyAssets()

	fmt.Println("Portfolio site generated in 'dist' directory!")
}

func generateHomePage() {
	skills := []string{
		"Go/Golang", "Microservices", "Docker", "Kubernetes",
		"REST APIs", "gRPC", "Kafka", "Redis",
		"AWS", "Git", "OpenTelemtry", "CI/CD", "Cassandra",
		"Performance Engineering", "C/C++, Python",
	}

	articles := getArticles()

	socialLinks := []SocialLink{
		{Name: "GitHub", URL: "https://github.com/HemanthKumarKola1"},
		{Name: "LeetCode", URL: "https://leetcode.com/u/Go_hemanth"},
		{Name: "LinkedIn", URL: "https://www.linkedin.com/in/hemanth-kumar-kola-03415b193"},
	}

	data := PageData{
		Title:       "Hemanth Kola",
		Articles:    articles,
		Skills:      skills,
		SocialLinks: socialLinks,
	}

	tmpl := template.Must(template.New("index").Parse(indexTemplate))
	file, _ := os.Create("dist/index.html")
	defer file.Close()
	tmpl.Execute(file, data)
}

func generateArticles() {
	articles := getArticles()
	tmpl := template.Must(template.New("article").Parse(articleTemplate))

	for _, article := range articles {
		file, _ := os.Create(fmt.Sprintf("dist/articles/%s.html", article.Slug))
		tmpl.Execute(file, article)
		file.Close()
	}
}

func getArticles() []Article {
	var articles []Article

	filepath.WalkDir("content", func(path string, d fs.DirEntry, err error) error {
		if err != nil || !strings.HasSuffix(path, ".md") {
			return nil
		}

		content, _ := os.ReadFile(path)
		html := blackfriday.Run(content)

		filename := strings.TrimSuffix(filepath.Base(path), ".md")
		title := strings.ReplaceAll(filename, "-", " ")
		title = strings.Title(title)

		articles = append(articles, Article{
			Title:   title,
			Date:    time.Now().Format("Jan 2, 2006"),
			Slug:    filename,
			Content: template.HTML(html),
			Summary: extractSummary(string(content)),
		})
		return nil
	})

	return articles
}

func extractSummary(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			if len(line) > 150 {
				return line[:150] + "..."
			}
			return line
		}
	}
	return "Read more..."
}

func copyAssets() {
	// Copy profile image
	os.MkdirAll("dist/images", 0755)
	if data, err := os.ReadFile("Hemanth-Mongo.jpeg"); err == nil {
		os.WriteFile("dist/images/profile.jpeg", data, 0644)
	}

	// Copy award images
	if data, err := os.ReadFile("GEN_AI.jpeg"); err == nil {
		os.WriteFile("dist/images/gen-ai-award.jpeg", data, 0644)
	}
	if data, err := os.ReadFile("OTEL_SPOT.jpeg"); err == nil {
		os.WriteFile("dist/images/otel-spot-award.jpeg", data, 0644)
	}

	css := `
* { margin: 0; padding: 0; box-sizing: border-box; }
body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; line-height: 1.6; color: #333; }
.container { max-width: 1200px; margin: 0 auto; padding: 0 20px; }
header { background: #2c3e50; color: white; padding: 2rem 0; }
.hero { display: flex; align-items: center; gap: 2rem; }
.profile-img { width: 150px; height: 150px; border-radius: 50%; object-fit: cover; border: 4px solid white; }
.hero-text { flex: 1; }
.hero h1 { font-size: 3rem; margin-bottom: 0.5rem; }
.hero p { font-size: 1.2rem; opacity: 0.9; }
.skills { background: #f8f9fa; padding: 3rem 0; }
.skills-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 1rem; margin-top: 2rem; }
.skill-card { background: white; padding: 1rem; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); text-align: center; }
.articles { padding: 3rem 0; }
.articles-grid { display: grid; gap: 2rem; margin-top: 2rem; }
.article-card { border: 1px solid #ddd; border-radius: 8px; padding: 1.5rem; }
.article-card h3 { color: #2c3e50; margin-bottom: 0.5rem; }
.article-card .date { color: #666; font-size: 0.9rem; margin-bottom: 1rem; }
.btn { display: inline-block; background: #3498db; color: white; padding: 0.5rem 1rem; text-decoration: none; border-radius: 4px; margin-top: 1rem; }
.btn:hover { background: #2980b9; }
.social-links { display: flex; justify-content: center; gap: 2rem; margin-bottom: 1rem; }
.social-links a { color: white; text-decoration: none; padding: 0.5rem 1rem; border: 1px solid rgba(255,255,255,0.3); border-radius: 4px; transition: background 0.3s; }
.social-links a:hover { background: rgba(255,255,255,0.1); }
.awards { background: #f8f9fa; padding: 3rem 0; }
.toggle-btn { background: #3498db; color: white; border: none; padding: 0.75rem 1.5rem; border-radius: 4px; cursor: pointer; margin-top: 1rem; }
.toggle-btn:hover { background: #2980b9; }
.awards-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(300px, 1fr)); gap: 2rem; margin-top: 2rem; }
.award-card { background: white; padding: 1rem; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); text-align: center; }
.award-img { width: 100%; max-width: 400px; height: auto; border-radius: 4px; }
footer { background: #2c3e50; color: white; text-align: center; padding: 2rem 0; margin-top: 3rem; }
@media (max-width: 768px) { .hero { flex-direction: column; text-align: center; } .hero h1 { font-size: 2rem; } .skills-grid { grid-template-columns: repeat(auto-fit, minmax(150px, 1fr)); } }
`
	os.WriteFile("dist/style.css", []byte(css), 0644)
}

const indexTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <header>
        <div class="container">
            <div class="hero">
                <img src="images/profile.jpeg" alt="Hemanth Kumar Kola" class="profile-img">
                <div class="hero-text">
                    <h1>Hemanth Kumar Kola</h1>
                    <p>Go Engineer & Backend Developer</p>
                    <p style="font-size: 1rem; margin-top: 0.5rem;">Bengaluru | +91 7337283959 | hemanthkumarkola1@gmail.com</p>
                </div>
            </div>
        </div>
    </header>

    <section class="skills">
        <div class="container">
            <h2>Technical Skills</h2>
            <div class="skills-grid">
                {{range .Skills}}
                <div class="skill-card">
                    <strong>{{.}}</strong>
                </div>
                {{end}}
            </div>
        </div>
    </section>

    <section class="awards">
        <div class="container">
            <h2>Awards & Recognition</h2>
            <button class="toggle-btn" onclick="toggleAwards()">Show Awards</button>
            <div class="awards-grid" id="awards-grid" style="display: none;">
                <div class="award-card">
                    <img src="images/gen-ai-award.jpeg" alt="Gen AI Award" class="award-img">
                </div>
                <div class="award-card">
                    <img src="images/otel-spot-award.jpeg" alt="OpenTelemetry Spot Award" class="award-img">
                </div>
            </div>
        </div>
    </section>

    <section class="articles">
        <div class="container">
            <h2>Latest Articles</h2>
            <div class="articles-grid">
                {{range .Articles}}
                <article class="article-card">
                    <h3>{{.Title}}</h3>
                    <div class="date">{{.Date}}</div>
                    <p>{{.Summary}}</p>
                    <a href="articles/{{.Slug}}.html" class="btn">Read More</a>
                </article>
                {{end}}
            </div>
        </div>
    </section>

    <footer>
        <div class="container">
            <div class="social-links">
                {{range .SocialLinks}}
                <a href="{{.URL}}" target="_blank">{{.Name}}</a>
                {{end}}
            </div>
            <p>&copy; 2024 Hemanth Kumar Kola. Built with Go.</p>
        </div>
    </footer>

    <script>
    function toggleAwards() {
        const grid = document.getElementById('awards-grid');
        const btn = document.querySelector('.toggle-btn');
        if (grid.style.display === 'none') {
            grid.style.display = 'grid';
            btn.textContent = 'Hide Awards';
        } else {
            grid.style.display = 'none';
            btn.textContent = 'Show Awards';
        }
    }
    </script>
</body>
</html>`

const articleTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="../style.css">
</head>
<body>
    <header>
        <div class="container">
            <h1><a href="../index.html" style="color: white; text-decoration: none;">← Back to Portfolio</a></h1>
        </div>
    </header>

    <main class="container" style="padding: 3rem 0;">
        <article>
            <h1>{{.Title}}</h1>
            <div class="date" style="margin-bottom: 2rem;">{{.Date}}</div>
            <div class="content">{{.Content}}</div>
        </article>
    </main>

    <footer>
        <div class="container">
            <p>&copy; 2024 Hemanth Kumar Kola. Built with Go.</p>
        </div>
    </footer>
</body>
</html>`
