package container

type Binding struct {
  Concrete any
  Shared bool
}

type Container struct {
  Instances map[string] any
  Bindings map[string] Binding
}

/* initialize container data */
func (c *Container) Init(){
  c.Instances = make(map[string]any)
  c.Bindings = make(map[string] Binding)
}

/* Register a binding with container. */
func (c *Container) Bind(abstract string, concrete any , shared  bool){
  c.Bindings[abstract] = Binding{concrete, shared}
}

/* Register shard binding in container. */
func (c *Container) Singleton(abstract string, concrete any){
  c.Bindings[abstract] = Binding{concrete, true}
}

func (c *Container) Build(abstract string, params any) any{
  concrete:= c.Bindings[abstract].Concrete;

  // check if the concrete is functio
  if(c.Bindings[abstract].Shared){
    c.Instances[abstract] = concrete;
  }
  return concrete;
}

func (c *Container) Make(abstract string, params interface{})any{
  if val, ok:= c.Instances[abstract]; ok{
    return val
  }

  return c.Build(abstract, params)
}

