package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"
)

// Point represents a 2D point
type Point struct {
	x, y float64
}

// Tree represents a fractal tree
type Tree struct {
	start  Point
	angle  float64
	length float64
	depth  int
	color  string
}

// Tree colors in natural shades
var treeColors = []string{
	"#2d5a27", // dark green
	"#3a7d32", // forest green
	"#4a8c3f", // medium green
	"#5c9c4c", // light green
	"#6dac59", // pale green
	"#7ebc66", // mint green
	"#8fcc73", // sage green
	"#a0dc80", // lime green
}

// generateTree recursively generates SVG path commands for a fractal tree
func generateTree(tree Tree) string {
	if tree.depth <= 0 {
		return ""
	}

	// Calculate end point of current branch
	endX := tree.start.x + tree.length*math.Cos(tree.angle)
	endY := tree.start.y - tree.length*math.Sin(tree.angle)
	end := Point{endX, endY}

	// Create SVG line for current branch
	path := fmt.Sprintf("M %.2f %.2f L %.2f %.2f ",
		tree.start.x, tree.start.y, end.x, end.y)

	// Generate left and right branches
	leftTree := Tree{
		start:  end,
		angle:  tree.angle + math.Pi/4,
		length: tree.length * 0.7,
		depth:  tree.depth - 1,
		color:  tree.color,
	}

	rightTree := Tree{
		start:  end,
		angle:  tree.angle - math.Pi/4,
		length: tree.length * 0.7,
		depth:  tree.depth - 1,
		color:  tree.color,
	}

	return path + generateTree(leftTree) + generateTree(rightTree)
}

func generateForest() string {
	var svgContent string
	svgContent = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<svg width="800" height="600" xmlns="http://www.w3.org/2000/svg">
<rect width="100%" height="100%" fill="skyblue"/>
`

	// Generate multiple trees
	numTrees := 5
	for i := 0; i < numTrees; i++ {
		// Random position for each tree
		x := rand.Float64()*700 + 50
		y := 600.0 // Start from bottom of canvas

		// Random angle variation - now pointing downward
		angle := math.Pi/2 + (rand.Float64()-0.5)*0.5

		// Select random color for the tree
		color := treeColors[rand.Intn(len(treeColors))]

		// Create tree
		tree := Tree{
			start:  Point{x, y},
			angle:  angle,
			length: 100.0,
			depth:  8,
			color:  color,
		}

		// Generate and write tree path with its color
		path := generateTree(tree)
		svgContent += fmt.Sprintf(`<path d="%s" stroke="%s" stroke-width="2" fill="none"/>`, path, tree.color)
	}

	svgContent += "\n</svg>"
	return svgContent
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Fractal Forest</title>
    <style>
        body {
            margin: 0;
            padding: 20px;
            background-color: #f0f0f0;
            font-family: Arial, sans-serif;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
            text-align: center;
        }
        h1 {
            color: #2c3e50;
        }
        .forest-container {
            background-color: white;
            padding: 20px;
            border-radius: 10px;
            box-shadow: 0 0 10px rgba(0,0,0,0.1);
            margin-top: 20px;
        }
        button {
            background-color: #3498db;
            color: white;
            border: none;
            padding: 10px 20px;
            border-radius: 5px;
            cursor: pointer;
            font-size: 16px;
            margin: 10px;
        }
        button:hover {
            background-color: #2980b9;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Fractal Forest</h1>
        <div class="forest-container">
            <div id="forest"></div>
            <button onclick="refreshForest()">Generate New Forest</button>
        </div>
    </div>
    <script>
        function refreshForest() {
            fetch('/forest')
                .then(response => response.text())
                .then(svg => {
                    document.getElementById('forest').innerHTML = svg;
                });
        }
        // Load initial forest
        refreshForest();
    </script>
</body>
</html>`
	fmt.Fprint(w, html)
}

func forestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprint(w, generateForest())
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Set up routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/forest", forestHandler)

	// Start server
	fmt.Println("Server starting at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
