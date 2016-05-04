package manifest

// Manifest represents a yaml file configuring the execution of the ratchet command
type Manifest []struct {
	Package  string `yaml:"package"`
	Ratchets []struct {
		Command   string   `yaml:"command"`
		Arguments []string `yaml:"args"`
	} `yaml:"ratchets"`
}
