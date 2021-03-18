package processor

// Record stores log information.
type Record struct {
	Request             Request             `json:"request"`
	UpstreamURI         string              `json:"upstream_uri"`
	Response            Response            `json:"response"`
	AuthenticatedEntity AuthenticatedEntity `json:"authenticated_entity"`
	Route               Route               `json:"route"`
	Service             Service             `json:"service"`
	Latencies           Latencies           `json:"latencies"`
	ClientIP            string              `json:"client_ip"`
	StartedAt           int64               `json:"started_at"`
}

type AuthenticatedEntity struct {
	ConsumerID struct {
		UUID string `json:"uuid"`
	} `json:"consumer_id"`
}

type Latencies struct {
	Proxy   int64 `json:"proxy"`
	Gateway int64 `json:"gateway"`
	Request int64 `json:"request"`
}

type Request struct {
	Method      string         `json:"method"`
	URI         string         `json:"uri"`
	URL         string         `json:"url"`
	Size        int            `json:"size"`
	Querystring []string       `json:"querystring"`
	Headers     RequestHeaders `json:"headers"`
}

type RequestHeaders struct {
	Accept    string `json:"accept"`
	Host      string `json:"host"`
	UserAgent string `json:"user-agent"`
}

type Response struct {
	Status  int64           `json:"status"`
	Size    int             `json:"size"`
	Headers ResponseHeaders `json:"headers"`
}

type ResponseHeaders struct {
	ContentLength                 string `json:"Content-Length"`
	Via                           string `json:"via"`
	Connection                    string `json:"Connection"`
	AccessControlAllowCredentials string `json:"access-control-allow-credentials"`
	ContentType                   string `json:"Content-Type"`
	Server                        string `json:"server"`
	AccessControlAllowOrigin      string `json:"access-control-allow-origin"`
}

type Route struct {
	CreatedAt     int64        `json:"created_at"`
	Hosts         string       `json:"hosts"`
	ID            string       `json:"id"`
	Methods       []string     `json:"methods"`
	Paths         []string     `json:"paths"`
	PreserveHost  bool         `json:"preserve_host"`
	Protocols     []string     `json:"protocols"`
	RegexPriority int64        `json:"regex_priority"`
	Service       RouteService `json:"service"`
	StripPath     bool         `json:"strip_path"`
	UpdatedAt     int64        `json:"updated_at"`
}

type RouteService struct {
	ID string `json:"id"`
}

type Service struct {
	ConnectTimeout int64  `json:"connect_timeout"`
	CreatedAt      int64  `json:"created_at"`
	Host           string `json:"host"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	Path           string `json:"path"`
	Port           int64  `json:"port"`
	Protocol       string `json:"protocol"`
	ReadTimeout    int64  `json:"read_timeout"`
	Retries        int64  `json:"retries"`
	UpdatedAt      int64  `json:"updated_at"`
	WriteTimeout   int64  `json:"write_timeout"`
}
