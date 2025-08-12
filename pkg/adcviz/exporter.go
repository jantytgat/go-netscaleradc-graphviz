package adcviz

type Exporter interface {
	Print() string
	SaveAsSvg(filename string) error
	SaveAsPng(filename string) error
}
