module github.com/eyedeekay/samtracker

go 1.13

require (
	github.com/eyedeekay/sam-forwarder v0.32.1-0.20191021175341-aa52be69ffe8
	github.com/vvampirius/retracker v0.0.0-20171226134001-fdbec17ad537
	github.com/zeebo/bencode v1.0.0 // indirect
)

replace github.com/vvampirius/retracker v0.0.0-20171226134001-fdbec17ad537 => github.com/eyedeekay/retracker v0.0.0-20191208024817-1068d9dccb6d
