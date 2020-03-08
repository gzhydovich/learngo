package main

type placeholder [5]string

var	zero = placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

var	one = placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

var	two = placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

var	three = placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

var	four = placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

var	five = placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

var	six = placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

var	seven = placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

var	eight = placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

var	nine = placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	// Dots ...
var	dots = placeholder{
	" ",
	"█",
	" ",
	"█",
	" ",
}

var	aa = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"█ █",
}

var	ll = placeholder{
	"█  ",
	"█  ",
	"█  ",
	"█  ",
	"███",
}

var	rr = placeholder{
	"██ ",
	"█ █",
	"██ ",
	"█ █",
	"█ █",
}

var	mm = placeholder{
	"█ █",
	"███",
	"█ █",
	"█ █",
	"█ █",
}

var	ex = placeholder{
	" █ ",
	" █ ",
	" █ ",
	"   ",
	" █ ",
}

var	dot = placeholder{
	" ",
	" ",
	" ",
	" ",
	"█",
}

var alarm = [10]placeholder{
	aa, ll, aa, rr, mm, ex,
}


var	digits = [...]placeholder{
		zero,
		one,
		two,
		three,
		four,
		five,
		six,
		seven,
		eight,
		nine,
	}