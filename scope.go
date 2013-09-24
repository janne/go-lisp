package lisp

type Env map[string]interface{}

type Scope struct {
	envs []*Env
}

func NewScope() *Scope {
	scope := &Scope{}
	scope.envs = make([]*Env, 0, 10)
	return scope
}

func (s *Scope) Env() *Env {
	if len(s.envs) > 0 {
		return s.envs[len(s.envs)-1]
	}
	return nil
}

func (s *Scope) AddEnv() *Env {
	env := make(Env)
	s.envs = append(s.envs, &env)
	return &env
}

func (s *Scope) DropEnv() *Env {
	s.envs = s.envs[:len(s.envs)-1]
	return s.Env()
}

func (s *Scope) Set(key string, value interface{}) interface{} {
	env := *s.Env()
	env[key] = value
	return value
}

func (s *Scope) Get(key string) (val interface{}, ok bool) {
	env := *s.Env()
	val, ok = env[key]
	return
}
