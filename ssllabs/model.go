package ssllabs

var CSV_LOCATION = "/tmp/ssl-scan.csv"

var CHAIN_ISSUES = map[int]string{
    0: "none",
    1: "unused",
    2: "incomplete chain",
    4: "chain contains unrelated or duplicate certificates",
    8: "the certificates form a chain (trusted or not) but incorrect order",
    16: "contains a self-signed root certificate",
    32: "the certificates form a chain but cannot be validated",
}

var FORWARD_SECRECY = map[int]string{
    1: "With some browsers WEAK",
    2: "With modern browsers",
    4: "Yes (with most browsers) ROBUST",
}

var RC4 = []string{"Support RC4", "RC4 with modern protocols", "RC4 Only"}

var PROTOCOLS = []string{
	"TLS 1.3", "TLS 1.2", "TLS 1.1", "TLS 1.0", "SSL 3.0 INSECURE", "SSL 2.0 INSECURE",
}

var PROTOCOL_IDS = map[string]int{
	"TLS 1.3": 772,
	"TLS 1.2": 771,
	"TLS 1.1": 770,
	"TLS 1.0": 769,
	"SSL 3.0 INSECURE": 768,
	"SSL 2.0 INSECURE": 2,
}

var VULNERABLES = []string{
    "Vuln Beast", "Vuln Drown", "Vuln Heartbleed", "Vuln FREAK",
    "Vuln openSsl Ccs", "Vuln openSSL LuckyMinus20", "Vuln POODLE", "Vuln POODLE TLS",
}

var VUL_TEST_RESULTS = map[int]string{
	-3: "Timeout" ,
	-2: "TLS not supported",
	-1: "Test Failed",
	0: "Unknown",
	1: "Not vulnerable",
	2: "Vulnerable",
}

type APIInfo struct {
	EngineVersion        string
	CriteriaVersion      string
	MaxAssessments       int
	CurrentAssessments   int
	NewAssessmentCoolOff int
	Messages             []string
}

type Report struct {
	Host            string
	Port            int
	Protocol        string
	IsPublic        bool
	Status          string
	StatusMessage   string
	StartTime       int
	TestTime        int
	EngineVersion   string
	CriteriaVersion string
	CacheExpiryTime int
	Endpoints       []Endpoint
	CertHostnames   []string
	rawJSON         string
}

type Cert struct {
	Subject              string
	CommonNames          []string
	AltNames             []string
	NotBefore            int64
	NotAfter             int64
	IssuerSubject        string
	SigAlg               string
	IssuerLabel          string
	RevocationInfo       int
	CrlURIs              []string
	OcspURIs             []string
	RevocationStatus     int
	CrlRevocationStatus  int
	OcspRevocationStatus int
	Sgc                  int
	ValidationType       string
	Issues               int
	Sct                  bool
	MustStaple           int
}

type Endpoint struct {
	IpAddress            string
	ServerName           string
	StatusMessage        string
	StatusDetailsMessage string
	Grade                string
	GradeTrustIgnored    string
	HasWarnings          bool
	IsExceptional        bool
	Progress             int
	Duration             int
	Eta                  int
	Delegation           int 
	Details              EndpointDetails
}

type Chain struct {
	Certs  []ChainCert
	Issues int
}

type ChainCert struct {
	Subject              string
	Label                string
	NotBefore            int64
	NotAfter             int64
	IssuerSubject        string
	IssuerLabel          string
	SigAlg               string
	Issues               int
	KeyAlg               string
	KeySize              int
	KeyStrength          int
	RevocationStatus     int
	CrlRevocationStatus  int
	OcspRevocationStatus int
	Raw                  string
}

type EndpointDetails struct {
	HostStartTime                  int
	Protocols                      []Protocol
	Cert                           Cert
	Chain                          Chain
	ServerSignature                string
	PrefixDelegation               bool 
	NonPrefixDelegation            bool
	VulnBeast                      bool
	RenegSupport                   int
	SessionResumption              int
	CompressionMethods             int
	SupportsNpn                    bool
	NpnProtocols                   string
	SessionTickets                 int
	OcspStapling                   bool
	StaplingRevocationStatus       int
	StaplingRevocationErrorMessage string
	SniRequired                    bool
	HttpStatusCode                 int
	HttpForwarding                 string
	ForwardSecrecy                 int
	SupportsRc4                    bool
	Rc4WithModern                  bool
	Rc4Only                        bool
	Heartbleed                     bool
	Heartbeat                      bool
	OpenSslCcs                     int
	OpenSSLLuckyMinus20            int
	Poodle                         bool
	PoodleTls                      int
	FallbackScsv                   bool
	Freak                          bool
	HasSct                         int
	DhPrimes                       []string
	DhUsesKnownPrimes              int
	DhYsReuse                      bool
	Logjam                         bool
	ChaCha20Preference             bool
	DrownErrors                    bool
	DrownVulnerable                bool
}

type Protocol struct {
	Id               int
	Name             string
	Version          string
	V2SuitesDisabled bool
	ErrorMessage     bool
	Q                int
}
