package request

const (
	categorizersPath = "/categorizers"
)

type Categorizers struct{}

func (c *Categorizers) Path() (string, error) {

	categorizersParameters, err := c.Encode()

	if err != nil {
		return "", err
	}

	return categorizersPath + categorizersParameters, nil
}

func (c *Categorizers) Validate() error {
	return nil
}

func (c *Categorizers) Encode() (string, error) {
	return "", nil
}
