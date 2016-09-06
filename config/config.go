// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	Testbeat TestbeatConfig
}

type TestbeatConfig struct {
	Period string `config:"period"`
	Command string `config:"command"`
}
