package gql

var (
	conn = newConnecter()
)

func Connect(opts ...ClientOption) {
	conn.Connect(opts...)
}

func ConnectAt(u string, opts ...ClientOption) {
	conn.ConnectAt(u, opts...)
}

func Disconnect() {
	conn.Disconnect()
}

func Run(req *Request) (resp Response, err error) {
	resp, err = conn.Run(req)
	return
}

func RunWithSpecialResp(req *Request, resp interface{}) (err error) {
	err = conn.RunWithSpecialResp(req, resp)
	return
}
