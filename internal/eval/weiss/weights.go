package eval

import (
	. "github.com/ChizhovVadim/CounterGo/pkg/common"
)

type Weights struct {
	PawnValue         Score
	KnightValue       Score
	BishopValue       Score
	RookValue         Score
	QueenValue        Score
	BishopPair        Score
	KnightMobility    [9]Score
	BishopMobility    [14]Score
	RookMobility      [15]Score
	QueenMobility     [28]Score
	PawnDoubled       Score
	PawnSupport       Score
	PawnOpen          Score
	PawnPhalanx       [RANK_NB]Score
	PawnIsolated      Score
	PassedPawn        [RANK_NB]Score
	PassedDefended    [RANK_NB]Score
	PassedBlocked     [RANK_NB]Score
	PassedDistUs      [FILE_NB]Score
	PassedDistThem    Score
	PassedRookBack    Score
	KingAtkPawn       Score
	NBBehindPawn      Score
	RookFile          [2]Score
	SafetyAttackPower [PIECE_NB]int
	SafetyCheckPower  [PIECE_NB]int
	KingLineDanger    [28]Score
	ThreatByPawn      Score
	ThreatByPawnPush  Score
	PSQT              [COLOUR_NB][PIECE_NB][64]Score `json:"-"`
}

func (w *Weights) init() {
	w.PawnValue = S(104, 205)
	w.KnightValue = S(408, 625)
	w.BishopValue = S(413, 653)
	w.RookValue = S(558, 1102)
	w.QueenValue = S(1479, 1945)
	w.BishopPair = S(25, 124)

	w.KnightMobility = [...]Score{
		S(-35, -142), S(-28, -19), S(-7, 46), S(3, 81), S(13, 92), S(16, 111), S(23, 111), S(32, 106),
		S(46, 77)}
	w.BishopMobility = [...]Score{
		S(-40, -93), S(-21, -54), S(-9, 6), S(-2, 44), S(8, 61), S(17, 89), S(22, 107), S(22, 114),
		S(20, 126), S(26, 126), S(31, 125), S(55, 107), S(51, 119), S(103, 71)}
	w.RookMobility = [...]Score{
		S(-106, -146), S(-22, 4), S(-5, 70), S(-3, 77), S(-3, 113), S(2, 131), S(0, 151), S(7, 155),
		S(12, 162), S(20, 168), S(28, 174), S(30, 178), S(31, 182), S(45, 172), S(86, 140)}
	w.QueenMobility = [...]Score{
		S(-63, -48), S(-96, -54), S(-91, -107), S(-21, -129), S(3, -58), S(-1, 65), S(-1, 135), S(2, 177),
		S(5, 204), S(9, 224), S(11, 241), S(14, 256), S(17, 260), S(17, 269), S(18, 276), S(19, 282),
		S(17, 290), S(15, 295), S(11, 302), S(9, 310), S(17, 303), S(13, 307), S(38, 287), S(49, 268),
		S(116, 213), S(130, 187), S(143, 163), S(123, 160)}

	w.initPSQT()

	w.PawnDoubled = S(-13, -42)
	w.PawnIsolated = S(-8, -19)
	w.PawnSupport = S(20, 11)
	w.PawnOpen = S(-14, -17)
	w.PawnPhalanx = [RANK_NB]Score{
		S(0, 0), S(8, -3), S(19, 7), S(24, 33),
		S(60, 126), S(171, 248), S(169, 318), S(0, 0),
	}

	w.PassedPawn = [RANK_NB]Score{
		S(0, 0), S(-15, 36), S(-13, 43), S(-70, 128),
		S(-13, 161), S(107, 200), S(278, 233), S(0, 0),
	}
	w.PassedDefended = [RANK_NB]Score{
		S(0, 0), S(0, 0), S(5, -14), S(-2, -14),
		S(4, 15), S(49, 64), S(161, 68), S(0, 0),
	}
	w.PassedBlocked = [RANK_NB]Score{
		S(0, 0), S(0, 0), S(0, 0), S(-1, -23), S(4, -34), S(9, -93), S(-28, -121), S(0, 0),
	}
	w.PassedDistUs = [RANK_NB]Score{
		S(0, 0), S(0, 0), S(0, 0), S(16, -29), S(10, -37), S(-8, -35), S(-13, -27), S(0, 0),
	}
	w.PassedDistThem = S(-3, 19)
	w.PassedRookBack = S(11, 23)

	w.KingAtkPawn = S(3, 54)
	w.NBBehindPawn = S(7, 42)
	w.RookFile = [2]Score{S(10, 18), S(29, 20)}

	w.SafetyAttackPower = [PIECE_NB]int{Knight: 35, Bishop: 20, Rook: 40, Queen: 80}
	w.SafetyCheckPower = [PIECE_NB]int{Knight: 100, Bishop: 35, Rook: 65, Queen: 65}

	w.KingLineDanger = [28]Score{
		S(0, 0), S(0, 0), S(0, 0), S(-11, 41),
		S(-27, 41), S(-31, 29), S(-30, 25), S(-37, 34),
		S(-41, 33), S(-54, 39), S(-50, 34), S(-65, 43),
		S(-68, 42), S(-79, 43), S(-88, 45), S(-86, 43),
		S(-98, 44), S(-100, 37), S(-107, 33), S(-112, 29),
		S(-119, 27), S(-133, 18), S(-131, 9), S(-157, -3),
		S(-140, -18), S(-132, -33), S(-130, -32), S(-135, -34),
	}

	w.ThreatByPawn = S(-52, -73)    // weiss positive
	w.ThreatByPawnPush = S(-18, -7) // weiss positive
}

func (w *Weights) initPSQT() {
	var psqt [PIECE_NB][SQUARE_NB]Score

	psqt[Pawn] = [SQUARE_NB]Score{
		S(0, 0), S(0, 0), S(0, 0), S(0, 0), S(0, 0), S(0, 0), S(0, 0), S(0, 0),
		S(57, 3), S(51, 32), S(32, 64), S(73, 35), S(82, 40), S(88, 27), S(-40, 85), S(-52, 53),
		S(2, 74), S(16, 77), S(33, 53), S(39, 28), S(62, 31), S(133, 32), S(91, 60), S(33, 66),
		S(-11, 40), S(-11, 19), S(-10, 6), S(4, -25), S(23, -15), S(39, -17), S(10, 8), S(2, 16),
		S(-18, 19), S(-27, 13), S(-13, -12), S(-15, -16), S(-3, -16), S(4, -14), S(-4, -8), S(-4, -6),
		S(-25, 4), S(-35, -3), S(-30, -3), S(-25, -9), S(-15, 1), S(-7, 4), S(2, -16), S(-5, -15),
		S(-13, 10), S(-14, 8), S(-21, 12), S(-11, 10), S(-17, 31), S(19, 17), S(32, -4), S(3, -24),
		S(0, 0), S(0, 0), S(0, 0), S(0, 0), S(0, 0), S(0, 0), S(0, 0), S(0, 0)}

	psqt[Knight] = [SQUARE_NB]Score{
		S(-207, -71), S(-112, -12), S(-145, 34), S(-53, 8), S(-1, 16), S(-128, 43), S(-79, -11), S(-158, -115),
		S(0, -26), S(0, -1), S(51, -12), S(57, 27), S(65, 10), S(79, -32), S(-12, 2), S(16, -41),
		S(-21, -13), S(26, 6), S(36, 53), S(55, 53), S(98, 31), S(80, 36), S(45, -6), S(3, -19),
		S(9, -3), S(25, 18), S(48, 56), S(43, 73), S(29, 71), S(69, 48), S(30, 25), S(39, -3),
		S(0, 2), S(18, 12), S(23, 56), S(27, 58), S(25, 59), S(36, 54), S(54, 13), S(31, 13),
		S(-22, -44), S(-5, -5), S(-1, 15), S(15, 39), S(16, 37), S(8, 14), S(11, 1), S(-1, -17),
		S(-34, -37), S(-40, -3), S(-19, -16), S(-5, 10), S(-10, 6), S(-10, -16), S(-22, -15), S(-6, -6),
		S(-88, -61), S(-15, -52), S(-40, -19), S(-7, 7), S(0, 6), S(-5, -26), S(-13, -22), S(-54, -37)}

	psqt[Bishop] = [SQUARE_NB]Score{
		S(-40, 51), S(-55, 44), S(-133, 58), S(-125, 63), S(-128, 57), S(-150, 49), S(-25, 25), S(-39, 24),
		S(-21, 23), S(23, 29), S(8, 28), S(-23, 36), S(7, 23), S(-7, 33), S(-18, 37), S(-49, 31),
		S(4, 25), S(31, 24), S(57, 24), S(39, 18), S(49, 22), S(57, 33), S(27, 33), S(12, 18),
		S(-5, 17), S(40, 19), S(30, 20), S(58, 39), S(44, 36), S(38, 27), S(45, 19), S(-1, 23),
		S(15, -7), S(21, 2), S(23, 29), S(39, 30), S(34, 31), S(25, 21), S(25, 10), S(40, -8),
		S(12, -8), S(32, 11), S(22, 14), S(16, 21), S(21, 25), S(27, 13), S(35, 1), S(37, -8),
		S(25, -13), S(22, -31), S(26, -12), S(2, 5), S(3, 6), S(10, -13), S(29, -30), S(29, -53),
		S(34, -17), S(40, -2), S(15, 7), S(-4, 5), S(9, 6), S(12, -2), S(22, -8), S(33, -24)}

	psqt[Rook] = [SQUARE_NB]Score{
		S(36, 67), S(36, 78), S(-3, 98), S(4, 88), S(14, 88), S(15, 88), S(34, 83), S(50, 75),
		S(-1, 74), S(-17, 84), S(11, 82), S(18, 85), S(8, 83), S(19, 52), S(4, 62), S(27, 51),
		S(-3, 68), S(50, 45), S(26, 64), S(53, 46), S(70, 33), S(65, 44), S(89, 18), S(30, 46),
		S(-8, 64), S(14, 58), S(25, 60), S(47, 48), S(41, 44), S(46, 36), S(44, 32), S(25, 43),
		S(-23, 43), S(-21, 60), S(-20, 58), S(-12, 48), S(-16, 44), S(-20, 45), S(5, 36), S(-12, 30),
		S(-30, 20), S(-22, 34), S(-30, 29), S(-23, 21), S(-22, 21), S(-22, 11), S(7, 6), S(-18, 5),
		S(-49, 21), S(-26, 19), S(-17, 22), S(-13, 14), S(-9, 7), S(-18, 0), S(-7, -6), S(-44, 12),
		S(-23, 26), S(-22, 25), S(-20, 28), S(-10, 12), S(-11, 10), S(-9, 19), S(-3, 14), S(-19, 10)}

	psqt[Queen] = [SQUARE_NB]Score{
		S(-32, 87), S(-10, 101), S(-5, 128), S(10, 135), S(7, 157), S(45, 146), S(43, 134), S(31, 125),
		S(-18, 57), S(-60, 99), S(-29, 105), S(-84, 199), S(-80, 239), S(-15, 179), S(-43, 165), S(16, 150),
		S(-15, 44), S(0, 33), S(-10, 79), S(-15, 122), S(7, 164), S(42, 167), S(63, 126), S(14, 149),
		S(3, 17), S(2, 60), S(-11, 70), S(-15, 130), S(-12, 167), S(0, 182), S(40, 166), S(18, 133),
		S(12, -3), S(10, 45), S(1, 50), S(-6, 95), S(-8, 103), S(11, 87), S(27, 72), S(39, 76),
		S(6, -26), S(19, 3), S(9, 27), S(4, 21), S(8, 24), S(8, 27), S(35, -6), S(28, -13),
		S(10, -37), S(15, -28), S(20, -40), S(15, -3), S(16, -13), S(17, -95), S(29, -125), S(23, -78),
		S(2, -44), S(-6, -47), S(-2, -42), S(7, -44), S(6, -47), S(-16, -53), S(-1, -83), S(12, -57)}

	psqt[King] = [SQUARE_NB]Score{
		S(-24, -67), S(25, 18), S(11, 45), S(9, 83), S(0, 59), S(25, 71), S(43, 88), S(15, -58),
		S(-6, 29), S(54, 110), S(55, 112), S(78, 99), S(81, 100), S(89, 115), S(95, 132), S(42, 45),
		S(35, 64), S(133, 110), S(120, 132), S(92, 137), S(128, 131), S(162, 138), S(147, 124), S(45, 64),
		S(38, 67), S(109, 98), S(103, 138), S(65, 158), S(78, 154), S(128, 130), S(118, 103), S(2, 72),
		S(29, 30), S(122, 74), S(129, 113), S(45, 151), S(90, 140), S(125, 115), S(123, 81), S(-36, 61),
		S(37, 30), S(119, 56), S(93, 87), S(71, 109), S(92, 105), S(79, 99), S(91, 68), S(7, 52),
		S(80, 33), S(80, 53), S(61, 72), S(5, 88), S(22, 87), S(34, 86), S(72, 60), S(63, 27),
		S(40, -41), S(90, 12), S(55, 41), S(-35, 34), S(29, 5), S(-30, 60), S(70, 12), S(47, -51)}

	for pieceType := Pawn; pieceType <= King; pieceType++ {
		for sq := 0; sq < 64; sq++ {
			var val = psqt[pieceType][sq]
			w.PSQT[SideWhite][pieceType][FlipSquare(sq)] = val
			w.PSQT[SideBlack][pieceType][sq] = val
		}
	}
}