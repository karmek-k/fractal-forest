package main

import (
	"math"
	"testing"
)

func TestPointCreation(t *testing.T) {
	p := Point{x: 100, y: 200}
	if p.x != 100 || p.y != 200 {
		t.Errorf("Point creation failed. Expected (100, 200), got (%.2f, %.2f)", p.x, p.y)
	}
}

func TestTreeCreation(t *testing.T) {
	start := Point{x: 100, y: 200}
	tree := Tree{
		start:  start,
		angle:  math.Pi / 2,
		length: 50.0,
		depth:  3,
	}

	if tree.start.x != start.x || tree.start.y != start.y {
		t.Errorf("Tree start point mismatch. Expected (%.2f, %.2f), got (%.2f, %.2f)",
			start.x, start.y, tree.start.x, tree.start.y)
	}
	if tree.angle != math.Pi/2 {
		t.Errorf("Tree angle mismatch. Expected %.2f, got %.2f", math.Pi/2, tree.angle)
	}
	if tree.length != 50.0 {
		t.Errorf("Tree length mismatch. Expected 50.0, got %.2f", tree.length)
	}
	if tree.depth != 3 {
		t.Errorf("Tree depth mismatch. Expected 3, got %d", tree.depth)
	}
}

func TestGenerateTree(t *testing.T) {
	tree := Tree{
		start:  Point{x: 400, y: 300},
		angle:  -math.Pi / 2, // Pointing upward
		length: 100.0,
		depth:  2,
	}

	svg := generateTree(tree)
	if svg == "" {
		t.Error("generateTree returned empty string")
	}

	// Test that the SVG contains the expected number of path commands
	// For depth 2, we should have 1 main branch + 2 first-level branches
	expectedPaths := 3
	pathCount := 0
	for i := 0; i < len(svg); i++ {
		if i+1 < len(svg) && svg[i:i+2] == "M " {
			pathCount++
		}
	}

	if pathCount != expectedPaths {
		t.Errorf("Expected %d paths, got %d", expectedPaths, pathCount)
	}
}

func TestGenerateForest(t *testing.T) {
	forest := generateForest()
	if forest == "" {
		t.Error("generateForest returned empty string")
	}

	// Test that the SVG contains the basic structure
	if len(forest) < 100 {
		t.Error("Generated forest SVG is too short")
	}

	// Test that the SVG contains the skyblue background
	if !contains(forest, "fill=\"skyblue\"") {
		t.Error("Forest SVG missing skyblue background")
	}

	// Test that the SVG contains path elements
	if !contains(forest, "<path") {
		t.Error("Forest SVG missing path elements")
	}
}

func TestTreeBranching(t *testing.T) {
	tree := Tree{
		start:  Point{x: 400, y: 300},
		angle:  -math.Pi / 2,
		length: 100.0,
		depth:  1,
	}

	svg := generateTree(tree)

	// For depth 1, we should have exactly one branch
	expectedPaths := 1
	pathCount := 0
	for i := 0; i < len(svg); i++ {
		if i+1 < len(svg) && svg[i:i+2] == "M " {
			pathCount++
		}
	}

	if pathCount != expectedPaths {
		t.Errorf("Expected %d paths for depth 1, got %d", expectedPaths, pathCount)
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[0:len(substr)] == substr || len(s) > len(substr) && contains(s[1:], substr)
}
