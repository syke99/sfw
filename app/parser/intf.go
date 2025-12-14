package parser

type Parser interface {
	ParseSecrets() error
}
