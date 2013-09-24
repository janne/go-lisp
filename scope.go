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

func (s *Scope) Create(key string, value interface{}) interface{} {
	env := *s.Env()
	env[key] = value
	return value
}

func (s *Scope) Set(key string, value interface{}) interface{} {
	for i := len(s.envs) - 1; i >= 0; i-- {
		env := *s.envs[i]
		if _, ok := env[key]; ok {
			env[key] = value
			return value
		}
	}
	return s.Create(key, value)
}

func (s *Scope) Get(key string) (val interface{}, ok bool) {
	for i := len(s.envs) - 1; i >= 0; i-- {
		env := *s.envs[i]
		if val, ok = env[key]; ok {
			break
		}
	}
	return
}
