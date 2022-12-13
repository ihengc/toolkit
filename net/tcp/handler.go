package tcp

/*
* @author: Heng ChenChi
* @date: 2022/12/13 0013 12:07
* @version: 1.0
* @description:
**/

type Handler interface {
	Handle(RequestReader, ResponseWriter)
}

type defaultHandler struct {
}

func (d *defaultHandler) Handle(reader RequestReader, writer ResponseWriter) {
	//TODO implement me
	panic("implement me")
}
