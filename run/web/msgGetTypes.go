package web

type msgGetTypes struct {
	clientID ClientID
}

func (m *msgGetTypes) Perform(srv *serverInstance) error {
	client := srv.clients[m.clientID]
	if nil == client {
		return nil
	}

	resp, err := rspSetTypes(srv.engine.Types())
	if nil != err {
		return err
	}

	client.Send(resp)

	return nil
}
