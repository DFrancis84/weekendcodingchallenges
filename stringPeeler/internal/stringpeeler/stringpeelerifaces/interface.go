package stringpeelerifaces

type StringPeelerClient interface {
	StringPeeler(input string) (string, error)
}
