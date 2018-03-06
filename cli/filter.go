package cli

//Filter reporting filter (use package name, or package name request/event prefix

//DefaultFilter create default filter
func DefaultFilter() map[string]bool {
	var result = make(map[string]bool)
	result["dsunit"] = true
	result["vc"] = true
	result["sdk"] = true
	result["workflow.print"] = true
	result["build"] = true
	result["deploy"] = true
	result["exec.stdin"] = true
	result["exec.stdout"] = true
	return result
}

//WildcardFilter
func WildcardFilter() map[string]bool {
	return map[string]bool{
		"*": true,
	}
}

