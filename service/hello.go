package service
type HelloService struct{}
func (service *HelloService)SayHello(name string) string{
	return "Hi:"+name
}
