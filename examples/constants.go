package examples

type RGB struct {
	R, G, B int
}

const (
	LogoPath      = "assets/logo.png"
	ServerUpImg   = "assets/server-up.png"
	ServerDownImg = "assets/server-down.png"
	leftX10       = 10
	leftX100      = 100
	MainFont      = "Helvetica"
	BaseFontSize  = 14
)

var (
	Amber         = RGB{231, 170, 64}
	Lime          = RGB{47, 241, 55}
	PureGreen     = RGB{0, 255, 0}
	Crimson       = RGB{220, 20, 60}
	DarkTurquoise = RGB{0, 206, 209}
	ColorDefault  = RGB{0, 0, 0}
	White         = RGB{255, 255, 255}
)

var ServerIpS = []string{
	"10.12.4.1",
	"10.12.4.2",
	"172.18.2.5",
	"172.18.2.6",
	"172.18.3.10",
	"192.168.17.1",
	"192.168.17.2",
	"10.45.1.100",
	"10.45.1.101",
	"10.45.2.50",
	"172.31.0.1",
	"172.31.0.2",
	"192.168.88.10",
	"192.168.88.11",
	"10.200.1.1",
	"10.200.2.1",
	"172.25.10.1",
	"172.25.10.2",
}
