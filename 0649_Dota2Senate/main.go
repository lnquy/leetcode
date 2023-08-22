package main

func main() {
	predictPartyVictory("DDRRR")
}

const (
	dire    = byte('D') // 68
	radiant = byte('R') // 82
)

func predictPartyVictory(senate string) string {
	currBannedParty := dire
	if senate[0] == dire {
		currBannedParty = radiant
	}
	noOfCurrBannedSenates := 0
	votingRound := []byte(senate)
	totalBannedSenates := 0

	for len(votingRound) > 0 {
		currSenate := votingRound[0]
		votingRound = votingRound[1:] // Pop first in queue
		if currSenate != currBannedParty {
			noOfCurrBannedSenates++                       // Any senate will try to ban a senate from the other party
			votingRound = append(votingRound, currSenate) // Move this currSenate to the end of queue after voted
		} else {
			if noOfCurrBannedSenates > 0 { // This senate has been banned
				// Simply remove this senate from votingRound
				noOfCurrBannedSenates--
				totalBannedSenates++
			} else {
				// All banned senates from the same party has been removed
				// Now we can flip the banning party
				currBannedParty = dire
				if currSenate == dire {
					currBannedParty = radiant
				}
				noOfCurrBannedSenates = 1
				votingRound = append(votingRound, currSenate) // Move this currSenate to the end of queue after voted
			}
		}

		// Voting finished
		if len(votingRound) == 1 || noOfCurrBannedSenates >= len(senate) {
			break
		}
	}

	if votingRound[0] == dire {
		return "Dire"
	} else {
		return "Radiant"
	}
}
