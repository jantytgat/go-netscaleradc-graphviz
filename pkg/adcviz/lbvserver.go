package adcviz

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/goccy/go-graphviz"
)

type LbVserver struct {
	Name           string
	IpAddress      string
	Port           int
	Type           string
	ListenPolicy   string
	ListenPriority float64
	LbMethod       string
	ServiceGroups  []LbVserverServiceGroupBinding
}

func (g LbVserver) CreateGraph(gv *graphviz.Graphviz) (*graphviz.Graph, error) {
	var err error

	var gr *graphviz.Graph
	if gr, err = gv.Graph(graphviz.WithName("cluster_" + g.Name)); err != nil {
		return nil, err
	}
	gr.SetLabel("Load-Balancing Virtual Server")
	gr.SetLabelLocation(graphviz.TopLocation)
	gr.SetCompound(true)

	if err = g.CreateSubGraph(gr); err != nil {
		return nil, err
	}
	return gr, nil
}

func (g LbVserver) CreateSubGraph(gr *graphviz.Graph) error {
	var err error
	// var lbv *graphviz.Node
	// if lbv, err = gr.CreateNodeByName(g.Name); err != nil {
	// 	return err
	// }
	// lbv.SetLabel(g.Name)
	// lbv.SetShape(graphviz.RectangleShape)
	// lbv.SetColor("#0000FF")

	var lbvDetailsGraph *graphviz.Graph
	if lbvDetailsGraph, err = gr.CreateSubGraphByName("cluster_" + g.Name + "_details"); err != nil {
		return err
	}
	// defer lbvDetailsGraph.Close()
	lbvDetailsGraph.SetLabel(g.Name)
	lbvDetailsGraph.SetCompound(true)

	var lbvDetailsNode *graphviz.Node
	if lbvDetailsNode, err = lbvDetailsGraph.CreateNodeByName("node_" + g.Name + "_details"); err != nil {
		return err
	}
	lbvDetailsNode.SetShape(graphviz.RectangleShape)
	lbvDetailsNode.SetLabel(g.Details())
	lbvDetailsNode.SetColor("#000000")

	for _, svg := range g.ServiceGroups {
		var svgGraph *graphviz.Graph
		if svgGraph, err = gr.CreateSubGraphByName("cluster_" + svg.ServiceGroupName); err != nil {
			return err
		}
		svgGraph.SetLabel(svg.ServiceGroupName)
		svgGraph.SetCompound(true)

		if err = svg.CreateSubGraph(svgGraph); err != nil {
			return err
		}
		var svgNode *graphviz.Node
		if svgNode, err = svgGraph.NodeByName("node_" + svg.ServiceGroupName); err != nil {
			return err
		}
		var edgeGraph *graphviz.Edge
		if edgeGraph, err = gr.CreateEdgeByName("edge_"+g.Name+"_TO_"+svg.ServiceGroupName, lbvDetailsNode, svgNode); err != nil {
			return err
		}
		edgeGraph.SetColor("#654321")
		edgeGraph.SetLogicalTail("cluster_" + g.Name + "_details")
		edgeGraph.SetLogicalHead("cluster_" + svg.ServiceGroupName)
		edgeGraph.SetArrowHead(graphviz.NoneArrow)
		edgeGraph.SetLabel(fmt.Sprintf("Order: %d", int(svg.Order)))
		edgeGraph.SetLabel("\n")
	}
	return nil
}

func (g LbVserver) Details() string {
	var sb = new(strings.Builder)

	sb.WriteString(fmt.Sprintf("%s: %s\n", "IP Address", g.IpAddress))
	sb.WriteString(fmt.Sprintf("%s: %d\n", "Port", g.Port))
	sb.WriteString(fmt.Sprintf("%s: %s\n", "Type", g.Type))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("%s: %s\n", "Listen Policy", g.ListenPolicy))
	sb.WriteString(fmt.Sprintf("%s: %d\n", "Listen Priority", int(g.ListenPriority)))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("%s: %s\n", "Method", g.LbMethod))
	return sb.String()
}

func (g LbVserver) String() (string, error) {
	var err error
	var gv *graphviz.Graphviz
	ctx := context.Background()

	if gv, err = graphviz.New(ctx); err != nil {
		return "", err
	}
	defer gv.Close()

	var gr *graphviz.Graph
	if gr, err = g.CreateGraph(gv); err != nil {
		return "", err
	}
	defer gr.Close()

	if err = gv.RenderFilename(ctx, gr, graphviz.PNG, g.Name+".png"); err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err = gv.Render(ctx, gr, "dot", &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (g LbVserver) SaveAsDot(filename string) error {
	var err error
	var gv *graphviz.Graphviz
	ctx := context.Background()

	if gv, err = graphviz.New(ctx); err != nil {
		return err
	}
	defer gv.Close()

	var gr *graphviz.Graph
	if gr, err = g.CreateGraph(gv); err != nil {
		return err
	}
	defer gr.Close()

	return renderToFile(ctx, gv, gr, graphviz.XDOT, filename)
}

func (g LbVserver) SaveAsJpg(filename string) error {
	var err error
	var gv *graphviz.Graphviz
	ctx := context.Background()

	if gv, err = graphviz.New(ctx); err != nil {
		return err
	}
	defer gv.Close()

	var gr *graphviz.Graph
	if gr, err = g.CreateGraph(gv); err != nil {
		return err
	}
	defer gr.Close()

	if err = gv.RenderFilename(ctx, gr, graphviz.JPG, filename); err != nil {
		return err
	}

	return nil
}

func (g LbVserver) SaveAsPng(filename string) error {
	var err error
	var gv *graphviz.Graphviz
	ctx := context.Background()

	if gv, err = graphviz.New(ctx); err != nil {
		return err
	}
	defer gv.Close()

	var gr *graphviz.Graph
	if gr, err = g.CreateGraph(gv); err != nil {
		return err
	}
	defer gr.Close()

	if err = gv.RenderFilename(ctx, gr, graphviz.PNG, filename); err != nil {
		return err
	}

	return nil
}

func (g LbVserver) SaveAsSvg(filename string) error {
	var err error
	var gv *graphviz.Graphviz
	ctx := context.Background()

	if gv, err = graphviz.New(ctx); err != nil {
		return err
	}
	defer gv.Close()

	var gr *graphviz.Graph
	if gr, err = g.CreateGraph(gv); err != nil {
		return err
	}
	defer gr.Close()

	return renderToFile(ctx, gv, gr, graphviz.SVG, filename)
}
