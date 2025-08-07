package find_the_maximum_number_of_cruits_collected

type Direction int16

const (
	Direction_Up Direction = iota
	Direction_Down
	Direction_Left
	Direction_Right

	Direction_Up_Right
	Direction_Up_Left

	Direction_Down_Right
	Direction_Down_Left
)

type Player struct {
	AvailableDirections map[Direction]struct{}
}

// 0,0; 0,1; 0,2;
// 1,0; 1,1; 1,2;
// 2,0; 2,1; 2,2;

var (
	// start: 0,0 (top-left)
	pOne = Player{AvailableDirections: map[Direction]struct{}{
		Direction_Down_Right: {},
		Direction_Right:      {},
		Direction_Down:       {},
	}}
	// start: 0,n-1 (bottom-left)
	pTwo = Player{AvailableDirections: map[Direction]struct{}{
		Direction_Up_Right:   {},
		Direction_Right:      {},
		Direction_Down_Right: {},
	}}
	// start: n-1,0 (top-right)
	pThree = Player{AvailableDirections: map[Direction]struct{}{
		Direction_Down_Left:  {},
		Direction_Down:       {},
		Direction_Down_Right: {},
	}}
)

func maxCollectedFruits(fruits [][]int) int {

}
