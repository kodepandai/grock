package container

type Binding struct {
  concrete interface {}
  shared bool
}

type Container struct {
  Instances map[string] interface {}
  Bindings map[string] Binding
}
