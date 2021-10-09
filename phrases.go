package main

var phrases = []string{
	"%m becomes compatible with %b. %p riots.",
	"%p forks %m into %m, causing %p to rage quit modding.",
	"%m's community splits due to %p, much to the dismay of %p.",
	"%m is not compatible with %b. %p is now using %m.",
	"%m is now compatible with %m. %p commits insanity.",
	"%p sees %m as a massive improvement over %m, and %p goes crazy.",
	"%p rewrites %m in %l, causing massive outrage from %p.",
	"%p makes a big change to %m, making half of the community switch to %m.",
	"%p gets banned by Discord. The %m community is in shambles.",
	"%p moves to %m, uprooting a whole community.",
}

var replacers = map[string][]string{
	"%p": people,
	"%m": mods,
	"%b": branches,
	"%l": languages,
}
