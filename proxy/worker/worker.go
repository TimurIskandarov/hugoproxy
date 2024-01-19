package worker

import (
	"log"
	"os"
	"test/worker/binary"
	"test/worker/counter"
	"test/worker/graph"
	"time"

	"github.com/mohae/deepcopy"
)

func Tasks() {
	t := time.NewTicker(5 * time.Second)
	var b byte = 0

	var totalNodes int = 4
	avl := binary.GenerateTree(totalNodes)
	bin := deepcopy.Copy(avl).(*binary.AVLTree)

	for {
		select {
		case tick := <-t.C:
			contentCounter := counter.GetCounterPage(tick, b)
			err := os.WriteFile("/app/static/tasks/_index.md", []byte(contentCounter), 0644)
			if err != nil {
				log.Println(err)
			}
			b++

			binary.SetRandomNode(avl, bin)
			contentAVL := binary.GetAVLPage(avl, bin)
			err = os.WriteFile("/app/static/tasks/binary.md", []byte(contentAVL), 0644)
			if err != nil {
				log.Println(err)
			}

			totalNodes++
			if totalNodes == 100 {
				totalNodes = 4
			}

			contentGraph := graph.GetGraphPage()
			err = os.WriteFile("/app/static/tasks/graph.md", []byte(contentGraph), 0644)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
