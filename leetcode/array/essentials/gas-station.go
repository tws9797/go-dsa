package essentials

// https://leetcode.com/problems/gas-station/
/*
	Input: gas = [1,2,3,4,5],
	cost = [3,4,5,1,2]
	Explanation:
	Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
	Travel to station 4. Your tank = 4 - 1 + 5 = 8
	Travel to station 0. Your tank = 8 - 2 + 1 = 7
	Travel to station 1. Your tank = 7 - 3 + 2 = 6
	Travel to station 2. Your tank = 6 - 4 + 3 = 5
	Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
	Therefore, return 3 as the starting index.
		Output: 3
*/

func CanCompleteCircuit(gas []int, cost []int) int {

	/*
		Scenario 1:
		Start at station i is impossible if gas[i] - cost[i] < 0. Set the starting station to i+1.

		Scenario 2:
		Start at station i and gas[i] - cost[i] > 0. Continue to the next station.
		If one of the station cause the currTank to be < 0, means there is no interval stations (including starting station of current loop) that is able to finish the trip.
		Hence set the starting station to i+1 as well

		Scenario 3:
		If currTank is equal or more than 0 until the end of the lenStation, check the totalTank.
		The totalTank must be equal or greater than 0 to have at least one starting point
	*/

	totalTank := 0
	currTank := 0
	lenStation := len(gas)
	startingStation := 0

	for i := 0; i < lenStation; i++ {

		// Total tank is to calculate the cost required for the previous path
		// For e.g, to determine station 3 is the starting station, have to calculate the path from stations 0 to 2, which is total tank.
		totalTank += gas[i] - cost[i]

		// Curr tank is to determine the starting point
		currTank += gas[i] - cost[i]

		// If one couldn't get here,
		if currTank < 0 {

			// Start with an empty tank.
			currTank = 0

			// Pick up the next station as the starting one.
			// One thing to note is that If there exists a solution, it is guaranteed to be unique
			startingStation = i + 1
		}
	}

	if totalTank < 0 {
		return -1
	}

	return startingStation
}
