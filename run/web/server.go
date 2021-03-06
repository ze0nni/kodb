package web

import (
	"log"

	"github.com/ze0nni/kodb/run/web/msg"

	"github.com/ze0nni/kodb/internal/engine"
	"github.com/ze0nni/kodb/internal/types"
)

func newServer(engine engine.Engine) *serverInstance {
	return &serverInstance{
		engine:               engine,
		clientConnectedCh:    make(chan *clientConnection),
		clientDisconnectedCh: make(chan *clientConnection),
		clients:              make(map[ClientID]*clientConnection),

		msgQueue: make(chan ClientMsg),

		msgGetSchemaCh:      make(chan msgGetSchema),
		msgGetLibraryRowsCh: make(chan msgGetLibraryRows),
		msgNewRowCh:         make(chan msgNewRow),
		msgDeleteRowCh:      make(chan msgDeleteRow),
		msgUpdateValueCh:    make(chan msgUpdateValue),

		msgAddLibraryCh: make(chan msgAddLibrary),
	}
}

type msgGetSchema = struct{ ClientId ClientID }
type msgGetLibraryRows = struct {
	ClientID    ClientID
	LibraryName engine.LibraryName
}

type msgDeleteRow = struct {
	ClientID    ClientID
	LibraryName engine.LibraryName
	RowID       engine.RowID
}
type msgUpdateValue = struct {
	ClientID    ClientID
	LibraryName engine.LibraryName
	RowID       engine.RowID
	FieldID     engine.FieldID
	value       string
}
type msgAddLibrary = struct {
	ClientID    ClientID
	LibraryName engine.LibraryName
}

type serverInstance struct {
	engine engine.Engine

	clientConnectedCh    chan *clientConnection
	clientDisconnectedCh chan *clientConnection
	clients              map[ClientID]*clientConnection

	msgQueue            chan ClientMsg
	msgGetSchemaCh      chan msgGetSchema
	msgGetLibraryRowsCh chan msgGetLibraryRows
	msgNewRowCh         chan msgNewRow
	msgDeleteRowCh      chan msgDeleteRow
	msgUpdateValueCh    chan msgUpdateValue

	msgAddLibraryCh chan msgAddLibrary
}

// ClientConnected
func (server *serverInstance) ClientConnected(client *clientConnection) {
	server.clientConnectedCh <- client
}

func (server *serverInstance) clientConnected(client *clientConnection) {
	if _, ok := server.clients[client.id]; false == ok {
		server.clients[client.id] = client
	} else {
		log.Panicf("Client [%d] already conected", client.id)
	}
}

// ClientDisconnected
func (server *serverInstance) ClientDisconnected(client *clientConnection) {
	server.clientDisconnectedCh <- client
}

func (server *serverInstance) clientDisconnected(client *clientConnection) {
	if storedClient, ok := server.clients[client.id]; ok && storedClient == client {
		delete(server.clients, client.id)
	} else {
		log.Panicf("Can't disconnec for client [%d]", client.id)
	}
}

// Perform

func (server *serverInstance) Perform(msg ClientMsg, err error) {
	if nil != err {
		log.Print(err)
		return
	}
	if nil == msg {
		panic("msg can't be nil")
	}
	server.msgQueue <- msg
}

func (server *serverInstance) perform(msg ClientMsg) {
	err := msg.Perform(server)
	if nil != err {
		log.Printf("Error when perform %s: %s", msg, err)
	}
}

func (server *serverInstance) GetSchema(clientId ClientID) {
	server.msgGetSchemaCh <- msgGetSchema{clientId}
}

func (server *serverInstance) getSchema(m msgGetSchema) {
	if client, ok := server.clients[m.ClientId]; ok {
		client.SetSchema(msg.SetSchemaMsgFromEngine(server.engine))
	}

}

// GetLibraryRows
func (server *serverInstance) GetLibraryRows(clientID ClientID, libraryName string) {
	server.msgGetLibraryRowsCh <- msgGetLibraryRows{clientID, engine.LibraryName(libraryName)}
}

func (server *serverInstance) getLibraryRows(m msgGetLibraryRows) {
	if client, ok := server.clients[m.ClientID]; ok {
		client.SetLibraryRows(msg.SetLibraryRowsMsgFromEngine(
			m.LibraryName,
			server.engine,
		))
	}
}

// NewRow
func (server *serverInstance) NewRow(m msgNewRow) {
	server.msgNewRowCh <- m
}

func (server *serverInstance) newRow(m msgNewRow) {
	l, err := server.engine.Library(m.LibraryName)
	if nil != err {
		log.Printf("Error when <newRow>: %s", err)
		return
	}

	id, err := l.NewRow()
	if nil != err {
		log.Printf("Error when <newRow>: %s", err)
		return
	}
	if m.HasParent {
		l.UpdateValue(id, engine.FieldID("parentLibrary"), m.ParentLibraryName.ToString())
		l.UpdateValue(id, engine.FieldID("parentRow"), m.ParentRowID.ToString())
		l.UpdateValue(id, engine.FieldID("parentColumn"), m.ParentFieldID.String())
	}
}

// DeleteRow
func (server *serverInstance) DeleteRow(clientID ClientID, libraryName string, rowID string) {
	server.msgDeleteRowCh <- msgDeleteRow{clientID, engine.LibraryName(libraryName), engine.RowID(rowID)}
}

func (server *serverInstance) deleteRow(m msgDeleteRow) {
	l, err := server.engine.Library(m.LibraryName)
	if nil != err {
		log.Printf("Error when <deleteRow>: %s", err)
	}
	err = l.DeleteRow(m.RowID)
	if nil != err {
		log.Printf("Error when <deleteRow>: %s", err)
	}
}

func (server *serverInstance) UpdateValue(clientID ClientID, libraryName, rowID, fieldID, value string) {
	server.msgUpdateValueCh <- msgUpdateValue{
		clientID,
		engine.LibraryName(libraryName),
		engine.RowID(rowID),
		engine.FieldID(fieldID),
		value,
	}
}

func (server *serverInstance) updateValue(m msgUpdateValue) {
	l, err := server.engine.Library(m.LibraryName)
	if nil != err {
		log.Printf("Error when <updateValue>: %s", err)
	}
	err = l.UpdateValue(m.RowID, m.FieldID, m.value)
	if nil != err {
		log.Printf("Error when <updateValue>: %s", err)
	}
}

//

func (server *serverInstance) AddLibrary(clientID ClientID, libraryName string) {
	server.msgAddLibraryCh <- msgAddLibrary{
		clientID,
		engine.LibraryName(libraryName),
	}
}

func (server *serverInstance) addLibrary(m msgAddLibrary) {
	//TODO: pass library name
	_, err := server.engine.AddLibrary(m.LibraryName, types.TypeName(""))
	if nil != err {
		log.Printf("Error when <addLibrary>: %s", err)
	}
}

//listen
func (server *serverInstance) listen() {
	listener := &serverListener{server}

	listenerHandle := server.engine.Listen(listener)
	defer listenerHandle()

	typesHandle := server.engine.Types().Listen(listener)
	defer typesHandle()

	for {
		select {
		case client := <-server.clientConnectedCh:
			server.clientConnected(client)
		case client := <-server.clientDisconnectedCh:
			server.clientDisconnected(client)

		case msg := <-server.msgQueue:
			server.perform(msg)
		case msg := <-server.msgGetSchemaCh:
			server.getSchema(msg)
		case msg := <-server.msgGetLibraryRowsCh:
			server.getLibraryRows(msg)
		case msg := <-server.msgNewRowCh:
			server.newRow(msg)
		case msg := <-server.msgDeleteRowCh:
			server.deleteRow(msg)
		case msg := <-server.msgUpdateValueCh:
			server.updateValue(msg)
		case msg := <-server.msgAddLibraryCh:
			server.addLibrary(msg)
		}
	}
}
