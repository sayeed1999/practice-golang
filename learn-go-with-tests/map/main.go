package dictionary

// First make a custom type Dictionary from the map of string:string
type Dictionary map[string]string

// Then make a custom Search method on it
// Caution: For map, don't use d *Dictionary! Why?
func (d Dictionary) Search(key string) string {
	found := d[key]
	return found
}
