package uniq

import "fmt"

type Options struct {
	Count      bool
	Duplicate  bool
	Unique     bool
	IgnoreCase bool
	SkipFields int
	SkipChars  int
}

func (o *Options) Validate() error {
	var shareFlags uint

	if o.Count {
		shareFlags++
	}

	if o.Duplicate {
		shareFlags++
	}

	if o.Unique {
		shareFlags++
	}

	if shareFlags > 1 {
		return fmt.Errorf("Options -c, -d, -u are not used together")
	}

	if o.SkipFields < 0 {
		return fmt.Errorf("-f num can not be negative")
	}

	if o.SkipChars < 0 {
		return fmt.Errorf("-c num can not be negative")
	}

	return nil
}
