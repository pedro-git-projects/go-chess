package main

import (
	"fmt"

	"github.com/pedro-git-projects/projeto-integrado-frontend/blogapi/pkg/encoder"
)

func createPaths() []string {
	paths := []string{}
	for i := 0; i < 9; i++ {
		paths = append(paths, fmt.Sprintf("blogapi/static/images/img%d.png", i))
	}
	return paths
}

func encodePaths(paths []string) []string {
	result, err := encoder.ImgstoBase64(paths)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	return result
}

var content = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed neque quam, lacinia eget ultricies vel, rhoncus sit amet dolor. Morbi eu ultrices leo, sit amet gravida lacus. In imperdiet lacus mauris, non dignissim mauris porttitor sed. Nullam vehicula nunc lacus, eget rhoncus velit auctor sit amet. Duis ipsum ipsum, ultricies non varius et, sagittis non ex. Donec neque nulla, semper malesuada commodo non, efficitur ac justo. Integer sed tempor erat. Ut venenatis mauris luctus tortor placerat, vitae cursus quam tristique. Mauris erat erat, rhoncus eget suscipit id, varius eu arcu. In ante lectus, mollis a aliquet ac, varius et tortor. Donec pellentesque at eros eget vehicula. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Nam quis enim suscipit, efficitur sapien nec, mattis mi. Proin at faucibus lorem, sed vestibulum est.

Curabitur neque justo, dignissim elementum justo nec, sagittis tincidunt augue. Pellentesque sit amet pellentesque orci. Phasellus pulvinar euismod ullamcorper. Donec vitae commodo dui. Nullam convallis diam non efficitur tristique. Curabitur fermentum eu quam quis feugiat. Nullam auctor ex sapien. Sed vitae augue maximus ex volutpat tempor ut vitae enim. Etiam eros turpis, sollicitudin at diam ut, porttitor blandit diam.`

func createPosts(result []string) []Post {
	posts := []Post{
		{ID: 1, Title: "How to Move the Pieces", Body: "Learn how all of the pieces move and how to set up a chessboard!", Image: result[0], Content: content},
		{ID: 2, Title: "Playing the Game", Body: "Learn the rules of the game and start playing today!", Image: result[1], Content: content},
		{ID: 3, Title: "Opening Principles", Body: "Learn the correct way to play the first moves of the game! ", Image: result[2], Content: content},
		{ID: 4, Title: "Winning the Game", Body: "Learn different ways to checkmate your opponent and win!", Image: result[3], Content: content},
		{ID: 5, Title: "Capturing Pieces", Body: "Learn these tactics to capture more of your opponent's pieces!", Image: result[4], Content: content},
		{ID: 6, Title: "Finding Checkmate", Body: "Start attacking weak kings and deliver checkmates!", Image: result[5], Content: content},
		{ID: 7, Title: "Intro to Book Openings", Body: "Learn the main ideas of the most popular and classic openings! ", Image: result[6], Content: content},
		{ID: 8, Title: "Make the Most of Your Pieces", Body: "Learn how to maximize the strength of your pieces! ", Image: result[7], Content: content},
		{ID: 9, Title: "Understanding Endgames", Body: "Stop avoiding king and pawn endgames and start winning them! ", Image: result[8], Content: content},
	}
	return posts
}

var paths []string = createPaths()

var result = encodePaths(paths)

var posts = createPosts(result)
