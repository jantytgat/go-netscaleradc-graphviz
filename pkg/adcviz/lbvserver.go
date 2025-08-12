package adcviz

type LbVserver struct {
	Name      string
	IpAddress string
	Port      int
	Type      string
}

func (g LbVserver) Print() string {
	return g.Name
}

func (g LbVserver) SaveAsSvg(filename string) error {
	return nil
}

func (g LbVserver) SaveAsPng(filename string) error {
	return nil
}
