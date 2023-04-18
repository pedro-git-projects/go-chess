package utils_test

import (
	"fmt"
	"testing"

	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/utils"
)

func TestCoordFromString(t *testing.T) {
	c, err := utils.CoordFromStr("a1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)

}
