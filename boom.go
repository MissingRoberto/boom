package boom

type Boom struct {
}

type Manifest struct {
	Jobs []Job `yaml:"jobs"`
}

type Job struct {
	Name      string
	Instances string
}
