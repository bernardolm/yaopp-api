package main

type Scale int

const (
	Simple        Scale = 0
	Linear        Scale = 1
	PlanningPoker Scale = 2
	Fibonacci     Scale = 3
)

func (scale Scale) String() string {
	names := [...]string{
		"Simple",
		"Linear",
		"Planning poker",
		"Fibonacci",
	}

	if scale < Simple || scale > Fibonacci {
		return "Unknown"
	}

	return names[scale]
}

func main() {
	FibonacciScale := []int{0, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	LinearScale := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	PlanningPokerScale := []int{0, 1, 2, 3, 5, 8, 13, 20, 40, 100}
	SimpleScale := []int{0, 1, 2}
}
