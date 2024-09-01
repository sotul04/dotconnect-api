package main

import (
	"dot-connect/board"
	"dot-connect/path"
	"dot-connect/solver"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type BoardRequest struct {
	Board [][]int `json:"board"`
}

type SolveResponse struct {
	Status    string  `json:"status"`
	Path      [][]int `json:"path"`
	Time      string  `json:"time"`
	NodeCount int     `json:"nodeCount"`
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API ready to use!")
	})

	r.POST("/solve", solveBoard)

	r.Run(port)

}

func solveBoard(c *gin.Context) {
	var request BoardRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	boardMatrix := request.Board
	rows := 0
	cols := 0
	for i := range boardMatrix {
		for j := range boardMatrix[i] {
			if boardMatrix[i][j] == 2 {
				rows = i
				cols = j
				break
			}
		}
	}

	board := board.NeWBoard(boardMatrix, board.NewSize(len(boardMatrix), len(boardMatrix[0])), path.New(rows, cols))
	solver := solver.New(board)
	start := time.Now()
	solver.Solve()
	elapsed := time.Since(start)

	if solver.Found {
		fmt.Println("Found")
		pathPoints := solver.Solution.ToPoints()
		c.JSON(http.StatusOK, SolveResponse{
			Status:    "success",
			Path:      pathPoints,
			Time:      elapsed.String(),
			NodeCount: solver.CounterNode,
		})
	} else {
		c.JSON(http.StatusOK, SolveResponse{
			Status:    "failed",
			Time:      elapsed.String(),
			NodeCount: solver.CounterNode,
		})
	}
}
