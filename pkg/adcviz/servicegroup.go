package adcviz

import (
	"bytes"
	"context"
	"strings"

	"github.com/goccy/go-graphviz"
)

type LbVserverServiceGroupBinding struct {
	ServiceGroupName string
	Order            float64
}

func (g LbVserverServiceGroupBinding) CreateGraph(gv *graphviz.Graphviz) (*graphviz.Graph, error) {
	var err error

	var gr *graphviz.Graph
	if gr, err = gv.Graph(graphviz.WithName("cluster_" + g.ServiceGroupName)); err != nil {
		return nil, err
	}
	gr.SetLabel("Service Group")
	// gr.SetBackgroundColor("#")

	if err = g.CreateSubGraph(gr); err != nil {
		return nil, err
	}
	return gr, nil
}

func (g LbVserverServiceGroupBinding) CreateSubGraph(gr *graphviz.Graph) error {
	var err error
	var svg *graphviz.Node
	if svg, err = gr.CreateNodeByName("node_" + g.ServiceGroupName); err != nil {
		return err
	}
	svg.SetLabel("S")
	svg.SetShape(graphviz.RectangleShape)

	// var lbvDetailsGraph *graphviz.Graph
	// if lbvDetailsGraph, err = gr.CreateSubGraphByName("clusterDetails"); err != nil {
	// 	return err
	// }
	// defer lbvDetailsGraph.Close()
	// lbvDetailsGraph.SetBackgroundColor("#00FF00")
	//
	// var lbvDetailsNode *graphviz.Node
	// if lbvDetailsNode, err = lbvDetailsGraph.CreateNodeByName(g.Name + "-details"); err != nil {
	// 	return err
	// }
	// lbvDetailsNode.SetShape(graphviz.RectangleShape)
	// lbvDetailsNode.SetLabel(g.Details())
	// lbvDetailsNode.SetColor("#FF0000")

	// var lbvEdgeLbvDetails *graphviz.Edge
	// if lbvEdgeLbvDetails, err = gr.CreateEdgeByName("details", lbv, lbvDetailsNode); err != nil {
	// 	return err
	// }
	// lbvEdgeLbvDetails.SetLabel("Details")
	return nil
}

func (g LbVserverServiceGroupBinding) Details() string {
	var sb = new(strings.Builder)

	// sb.WriteString(fmt.Sprintf("%s: %s\n", "IP Address", g.IpAddress))
	// sb.WriteString(fmt.Sprintf("%s: %d\n", "Port", g.Port))
	// sb.WriteString(fmt.Sprintf("%s: %s\n", "Type", g.Type))
	// sb.WriteString("\n")
	// sb.WriteString(fmt.Sprintf("%s: %s\n", "Listen Policy", g.ListenPolicy))
	// sb.WriteString(fmt.Sprintf("%s: %d\n", "Listen Priority", int(g.ListenPriority)))
	// sb.WriteString("\n")
	// sb.WriteString(fmt.Sprintf("%s: %s\n", "Method", g.LbMethod))
	return sb.String()
}

func (g LbVserverServiceGroupBinding) String() (string, error) {
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

	if err = gv.RenderFilename(ctx, gr, graphviz.PNG, g.ServiceGroupName+".png"); err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err = gv.Render(ctx, gr, "dot", &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (g LbVserverServiceGroupBinding) SaveAsDot(filename string) error {
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

func (g LbVserverServiceGroupBinding) SaveAsJpg(filename string) error {
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

func (g LbVserverServiceGroupBinding) SaveAsPng(filename string) error {
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

func (g LbVserverServiceGroupBinding) SaveAsSvg(filename string) error {
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
