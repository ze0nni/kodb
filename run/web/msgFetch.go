package web

type msgFetch struct {
	clientID ClientID
}

func (m *msgFetch) Perform(srv *serverInstance) error {
	client := srv.clients[m.clientID]
	if nil == client {
		return nil
	}

	resp, err := rspFetch(srv.engine)
	if nil != err {
		return err
	}
	client.Send(resp)

	return nil
}
