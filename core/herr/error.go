package herr

type Err interface {
	Code() int
	Msg() string
	Data() interface{}
	Err() error
	String() string

	IsNil() bool
	NotNil() bool
	WithData(data interface{}) Err
	WithMsg(msg string) Err
	WithError(err error) Err

	Equal(err Err) bool
}
