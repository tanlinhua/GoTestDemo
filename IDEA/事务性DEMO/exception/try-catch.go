package exception

type Block struct {
	Try     func()
	Catch   func(interface{})
	Finally func()
}

func (t Block) Do() {
	if t.Finally != nil {
		defer t.Finally()
	}
	if t.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				t.Catch(r)
			}
		}()
	}
	t.Try()
}
