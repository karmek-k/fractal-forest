# Fractal Forest Generator

A Go web application that generates beautiful fractal trees using SVG.

## Features

- Interactive web interface
- Randomly generated fractal trees
- Colorful tree variations
- Real-time forest generation
- Health check endpoint

## Local Development

1. Install Go 1.21 or later
2. Clone the repository
3. Run the application:
```bash
go run fractal_forest.go
```
4. Visit http://localhost:8080

## Deployment to DigitalOcean App Platform

1. Create a new app in DigitalOcean App Platform
2. Connect your GitHub repository
3. Select the Dockerfile deployment method
4. Configure the following:
   - HTTP Port: 8080
   - Health Check Path: /health
   - Instance Count: 1 (or more for scaling)
   - Instance Size: Basic (or your preferred size)

## Environment Variables

- `PORT`: The port the application will listen on (default: 8080)

## API Endpoints

- `/`: Main web interface
- `/forest`: Generates SVG forest
- `/health`: Health check endpoint

## Building and Running with Docker

```bash
# Build the image
docker build -t fractal-forest .

# Run the container
docker run -p 8080:8080 fractal-forest
``` 