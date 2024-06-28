package template

func List() ([]string, error) {
	lists := make([]string, 3)
	lists[0] = "template1"
	lists[1] = "template2"
	lists[2] = "template3"

	return lists, nil
}
