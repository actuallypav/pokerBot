package main

import (
	"fmt"

	proto "github.com/Dan-Sa/poker-lib/pb"
	lib "github.com/Dan-Sa/poker-lib/shared"
)

type SimpleBot struct{}

const (
	handStrengthWeak = iota
	handStrengthPair
	handStrengthTwoPair
	handStrengthThree
	handStrengthStraight
	handStrengthFlush
	handStrengthFullHouse
	handStrengthFour
	handStrengthStrFlush
	handStrengthRoyFlush
)

func (b *SimpleBot) Action(players []*lib.Player, holeCards, board []*lib.Card, pots []*lib.Pot, betSoFar, betToPlayerThisRound, minRaise uint) (proto.Action, uint) {
	// Evaluate hand strength based on hole cards and board cards
	handStrength := handStrengthWeak

	handStrength = EvaluateHand(holeCards, board)

	//TODO: call below function to work out a number
	//TODO: between 0-1 to see how much of a risk it is to bet
	//opponentStrenght := OpponentLogic(players, holeCards, board, pots, betSoFar, betToPlayerThisRound, minRaise)

	if handStrength != handStrengthWeak {
		return proto.Action_Bet, minRaise
	} else {
		return proto.Action_Fold, 0//fold for a weak hand
	}

	// Implement logic to analyze opponents' actions and adjust strategy

	// Develop a betting strategy considering pot odds, position, and hand strength

	// Introduce bluffing logic based on opponents' behavior and hand strength

	// Calculate pot odds and determine whether to call, raise, or fold

	// Consider the bot's position at the table and adjust strategy accordingly

}
func EvaluateHand(holeCards, board []*lib.Card) int {
	//TODO: add deeper logic - we don't want to go through all of them every time
	//check if a flag has been raised, eg check for handStrengthPair etc.

	//create a map to store the frequency of each rank (board/hand)
	rankFrequency := make(map[proto.Rank]int)

	//count the frequency of ranks in hole cards
	for _, card := range holeCards {
		rankFrequency[card.Rank]++
	}

	//count the frequency of ranks in board cards
	for _, card := range board {
		rankFrequency[card.Rank]++
	}

	//check for royal flush
	//check for straight flush
	//check for four of a kind
	//check for full house
	//check for flush
	//check for straight
	//check for three of a kind
	//check for two pair
	//check for pair

	//check if any rank has a frequency of 2, indicating a pair
	for _, frequency := range rankFrequency {
		if frequency == 2 {
			return handStrengthPair // pair found
		}
	}

	return handStrengthWeak // weak odds or high card
}

func OpponentLogic(players []*lib.Player, holeCards, board []*lib.Card, pots []*lib.Pot, betSoFar, betToPlayerThisRound, minRaise uint) float64 {
	return handStrengthPair
}

func testBot(bot *SimpleBot) {
	// define different scenarios to test
	scenarios := []struct {
		players     []*lib.Player
		holeCards   []*lib.Card
		boardCards  []*lib.Card
		pots        []*lib.Pot
		betSoFar    uint
		betToPlayer uint
		minRaise    uint
	}{
		// Scenario 1: Initial scenario with no actions
		{
			players: []*lib.Player{
				{Guid: "1", Name: "Player1", Bank: 1000},
				{Guid: "2", Name: "Player2", Bank: 1000},
			},
			holeCards: []*lib.Card{
				{Rank: proto.Rank_Queen, Suit: proto.Suit_Spades},
				{Rank: proto.Rank_Jack, Suit: proto.Suit_Hearts}, // pair of queens
			},
			boardCards: []*lib.Card{
				{Rank: proto.Rank_Two, Suit: proto.Suit_Dimonds},
				{Rank: proto.Rank_Queen, Suit: proto.Suit_Clubs}, // pair of queens on the board
				{Rank: proto.Rank_Ten, Suit: proto.Suit_Spades},
			},
			pots: []*lib.Pot{
				{Size: 100, Players: []string{"Player1", "Player2"}},
			},
			betSoFar:    50,
			betToPlayer: 100,
			minRaise:    200,
		},
		// // Scenario 2: Player raises in the second round
		// {
		// 	players: []*lib.Player{
		// 		{Guid: "1", Name: "Player1", Bank: 1000},
		// 		{Guid: "2", Name: "Player2", Bank: 1000},
		// 	},
		// 	holeCards: []*lib.Card{
		// 		{Rank: proto.Rank_Queen, Suit: proto.Suit_Spades},
		// 		{Rank: proto.Rank_Jack, Suit: proto.Suit_Hearts}, //pair of queens
		// 	},
		// 	boardCards: []*lib.Card{
		// 		{Rank: proto.Rank_Two, Suit: proto.Suit_Dimonds},
		// 		{Rank: proto.Rank_Queen, Suit: proto.Suit_Clubs}, //pair of queens on the board
		// 		{Rank: proto.Rank_Ten, Suit: proto.Suit_Spades},
		// 		{Rank: proto.Rank_Four, Suit: proto.Suit_Dimonds}, //card shows a four
		// 	},
		// 	pots: []*lib.Pot{
		// 		{Size: 100, Players: []string{"Player1", "Player2"}},
		// 	},
		// 	betSoFar:    50,
		// 	betToPlayer: 100,
		// 	minRaise:    200,
		// },
		// // Scenario 3: All other players fold in the third round
		// {
		// 	players: []*lib.Player{
		// 		{Guid: "1", Name: "Player1", Bank: 1000},
		// 		{Guid: "2", Name: "Player2", Bank: 1000},
		// 	},
		// 	holeCards: []*lib.Card{
		// 		{Rank: proto.Rank_Queen, Suit: proto.Suit_Spades},
		// 		{Rank: proto.Rank_Jack, Suit: proto.Suit_Hearts}, // pair of queens
		// 	},
		// 	boardCards: []*lib.Card{
		// 		{Rank: proto.Rank_Two, Suit: proto.Suit_Dimonds},
		// 		{Rank: proto.Rank_Queen, Suit: proto.Suit_Clubs}, // pair of queens on the board
		// 		{Rank: proto.Rank_Ten, Suit: proto.Suit_Spades},
		// 		{Rank: proto.Rank_Four, Suit: proto.Suit_Dimonds}, // final card reveals a 6
		// 		{Rank: proto.Rank_Six, Suit: proto.Suit_Spades},   // final card reveals a 4
		// 	},
		// 	pots: []*lib.Pot{
		// 		{Size: 100, Players: []string{"Player1", "Player2"}},
		// 	},
		// 	betSoFar:    50,
		// 	betToPlayer: 100,
		// 	minRaise:    200,
		// },
	}

	// Test each scenario
	for i, scenario := range scenarios {
		fmt.Printf("Scenario %d\n", i+1)
		action, amount := bot.Action(
			scenario.players,     //players in scenario
			scenario.holeCards,   //cards the bot is holding
			scenario.boardCards,  //board cards
			scenario.pots,        //pots is a list (0 is the total) - other could refer to other players
			scenario.betSoFar,    //total amount of bets so far
			scenario.betToPlayer, //how many chips to contribute to stay in the game
			scenario.minRaise,    //the min raise amount in this scenario
		)
		fmt.Println("Current Pot Size:", scenario.pots[0].Size)
		fmt.Println("Bot action:", action)
		fmt.Println("Amount:", amount)
		fmt.Println("------------")

		// //simulate multiple rounds
		// for round := 2; round <= 5; round++ {
		// 	//update board cards and pot size for the next round
		// 	scenario.boardCards = append(scenario.boardCards, &lib.Card{Rank: proto.Rank_Ace, Suit: proto.Suit_Spades}) // add a new card to the board
		// 	scenario.pots[0].Size += 50                                                                                 // increase pot size for the next round

		// 	fmt.Printf("Round %d\n", round)
		// 	fmt.Println("Current Pot Size:", scenario.pots[0].Size)
		// 	action, amount := bot.Action(
		// 		scenario.players,
		// 		scenario.holeCards,
		// 		scenario.boardCards,
		// 		scenario.pots,
		// 		scenario.betSoFar,
		// 		scenario.betToPlayer,
		// 		scenario.minRaise,
		// 	)
		// 	fmt.Println("Bot action:", action)
		// 	fmt.Println("Amount:", amount)
		// 	fmt.Println("------------")

		// 	if action == proto.Action_Bet {
		// 		scenario.pots[0].Size += amount //increase pot size by the bet amount
		// 	}
		}
	}
}

func main() {
	//create a simple bot instance
	bot := &SimpleBot{}

	//run the testing function
	testBot(bot)
}
