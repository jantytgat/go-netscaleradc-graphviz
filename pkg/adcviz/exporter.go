package adcviz

import (
	"context"
	"os"

	"github.com/goccy/go-graphviz"
)

type Exporter interface {
	String() (string, error)
	SaveAsDot(filename string) error
	SaveAsJpg(filename string) error
	SaveAsPng(filename string) error
	SaveAsSvg(filename string) error
}

type Grapher interface {
	CreateGraph(gv *graphviz.Graphviz) (*graphviz.Graph, error)
	CreateSubGraph(gr *graphviz.Graph) error
}

type Detailer interface {
	Details() string
}

func renderToFile(ctx context.Context, gv *graphviz.Graphviz, gr *graphviz.Graph, format graphviz.Format, filename string) error {
	var err error
	var file *os.File
	if file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666); err != nil {
		return err
	}
	defer file.Close()

	if err = gv.Render(context.Background(), gr, format, file); err != nil {
		return err
	}
	return nil
}
