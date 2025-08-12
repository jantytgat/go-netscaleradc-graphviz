package adcviz

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/goccy/go-graphviz"
)

type LbVserver struct {
	Name      string
	IpAddress string
	Port      int
	Type      string
}

func (g LbVserver) Print() string {
	var err error
	var graphv *graphviz.Graphviz
	ctx := context.Background()

	if graphv, err = graphviz.New(ctx); err != nil {
		panic(err)
	}

	var graph *graphviz.Graph
	if graph, err = graphv.Graph(); err != nil {
		panic(err)
	}

	defer func() {
		var err2 error
		if err2 = graph.Close(); err2 != nil {
			panic(err2)
		}
		if err2 = graphv.Close(); err2 != nil {
			panic(err2)
		}
	}()

	var lbvs *graphviz.Node
	if lbvs, err = graph.CreateNodeByName(g.Name); err != nil {
		panic(err)
	}
	fmt.Println("LbVserver:", lbvs.Label())

	var svg *graphviz.Node
	if svg, err = graph.CreateNodeByName("SVG"); err != nil {
		panic(err)
	}

	fmt.Println("SVG:", svg.Label())

	var subnode *graphviz.Node
	if subnode, err = graph.CreateSubNode(svg); err != nil {
		panic(err)
	}
	subnode.MainSub().SetNode(lbvs)
	fmt.Println("SubNode:", subnode.Label())

	// m, err := graph.CreateNodeByName("m")
	// if err != nil { panic(err) }
	//
	// e, err := graph.CreateEdgeByName("e", n, m)
	// if err != nil { panic(err) }
	// e.SetLabel("e")

	if err = graphv.RenderFilename(ctx, graph, graphviz.PNG, g.Name+".png"); err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err = graphv.Render(ctx, graph, "dot", &buf); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func (g LbVserver) SaveAsSvg(filename string) error {
	return nil
}

func (g LbVserver) SaveAsPng(filename string) error {
	return nil
}
