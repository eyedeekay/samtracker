package samtracker

import (
	"log"

	"github.com/eyedeekay/sam-forwarder/interface"
	"github.com/eyedeekay/sam-forwarder/tcp"

	Core "github.com/vvampirius/retracker/core"
	"github.com/vvampirius/retracker/core/common"
)

//SamTracker is a structure which automatically configured the forwarding of
//a local service to i2p over the SAM API.
type SamTracker struct {
	*samforwarder.SAMForwarder
	Server   *Core.Core
	config   *common.Config
	ServeDir string
	up       bool
}

var err error

func (f *SamTracker) GetType() string {
	return "samtracker"
}

func (f *SamTracker) ServeParent() {
	log.Println("Starting eepsite server", f.Base32())
	if err = f.SAMForwarder.Serve(); err != nil {
		f.Cleanup()
	}
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *SamTracker) Serve() error {
	go f.ServeParent()
	if f.Up() {
		log.Println("Starting web server", f.Target())
		f.Server = Core.New(f.config)
	}
	return nil
}

func (f *SamTracker) Up() bool {
	return f.up
}

//Close shuts the whole thing down.
func (f *SamTracker) Close() error {
	return f.SAMForwarder.Close()
}

func (s *SamTracker) Load() (samtunnel.SAMTunnel, error) {
	if !s.up {
		log.Println("Started putting tunnel up")
	}
	f, e := s.SAMForwarder.Load()
	if e != nil {
		return nil, e
	}
	s.SAMForwarder = f.(*samforwarder.SAMForwarder)
	s.up = true
	log.Println("Finished putting tunnel up")
	return s, nil
}

//NewSamTracker makes a new SAM forwarder with default options, accepts host:port arguments
func NewSamTracker(host, port string) (*SamTracker, error) {
	return NewSamTrackerFromOptions(SetHost(host), SetPort(port))
}

//NewSamTrackerFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewSamTrackerFromOptions(opts ...func(*SamTracker) error) (*SamTracker, error) {
	var s SamTracker
	s.SAMForwarder = &samforwarder.SAMForwarder{}
	s.Server = &Core.Core{}
	log.Println("Initializing eephttpd")
	s.config = &common.Config{}
	s.config.Age = 180
	s.config.Debug = false
	s.config.XRealIP = false
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	s.SAMForwarder.Config().SaveFile = true
	l, e := s.Load()
	s.config.Listen = s.Target()
	if e != nil {
		return nil, e
	}
	return l.(*SamTracker), nil
}
