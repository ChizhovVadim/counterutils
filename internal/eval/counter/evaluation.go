package eval

import (
	. "github.com/ChizhovVadim/CounterGo/pkg/common"
	"github.com/ChizhovVadim/counterutils/internal/domain"
)

// 2024/03/12 10:43:31 Finished Epoch 10
// 2024/03/12 10:43:31 Current validation cost is: 0.025587
// 2024/03/12 10:43:31 Train finished
var w = []int{2195, 7543, 0, 0, 0, 0, 0, 0, 0, 0, 5989, 13050, 9150, 13239, 8851, 14386, 8336, 13675, 7185, 12756, 8333, 12462, 8896, 12127, 7799, 11719, 7794, 13355, 8333, 13969, 10950, 11799, 10531, 10854, 9359, 13983, 12245, 13322, 11499, 11872, 12142, 11828, 10382, 13985, 12554, 14564, 11484, 14445, 14177, 13577, 8029, 11634, 4229, 13204, 2431, 12604, 10716, 11254, 0, 0, 0, 0, 0, 0, 0, 0, 13515, 17122, 17714, 19814, 17566, 20030, 19862, 21273, 18074, 17569, 19180, 21676, 19001, 19604, 19515, 20668, 17333, 19092, 19793, 19927, 19675, 21339, 21214, 22203, 20301, 20314, 19705, 20748, 21122, 22482, 21551, 22340, 21018, 20562, 20825, 21292, 21204, 22404, 21837, 23970, 18728, 20979, 20945, 20780, 21177, 22603, 24828, 22000, 20813, 18436, 18025, 21258, 23783, 20225, 26454, 20412, 5781, 15478, 17097, 19856, 5286, 24149, 17831, 21090, 22157, 22886, 23962, 23858, 19707, 24372, 20655, 23511, 23028, 21825, 22897, 22337, 22823, 22823, 21452, 23651, 20918, 23340, 23583, 22624, 22159, 24213, 21886, 24004, 23850, 21697, 20340, 23684, 22238, 24823, 21894, 24862, 21536, 21978, 21824, 23755, 22881, 24510, 23066, 25341, 20934, 23507, 23652, 24276, 23901, 24383, 26397, 22934, 18981, 24139, 19502, 25599, 20978, 23812, 17424, 25815, 16464, 25199, 12209, 26988, 11802, 25451, 18026, 25052, 27296, 40678, 27807, 41785, 28484, 41261, 30088, 40250, 25079, 42582, 27096, 42001, 28483, 41623, 28969, 41296, 26519, 42963, 28033, 43465, 27893, 43516, 28434, 42455, 27164, 43836, 27037, 44652, 28376, 44499, 29942, 44218, 29103, 45100, 30782, 44756, 30685, 44609, 32643, 43459, 30559, 44717, 33065, 43623, 32522, 44165, 35502, 42444, 31481, 44049, 30094, 45176, 33565, 44371, 35309, 44524, 32805, 44054, 35599, 44175, 34012, 44746, 34765, 44754, 58333, 76065, 58471, 76309, 58648, 76400, 57310, 80808, 59262, 76025, 59061, 76962, 59791, 76046, 58966, 78853, 56768, 83289, 58075, 81722, 57945, 85065, 58019, 85242, 58035, 81997, 56760, 86225, 57157, 86579, 56420, 89555, 57975, 85593, 56246, 91145, 56533, 91161, 55612, 92792, 56061, 88815, 59101, 87721, 57103, 91674, 58862, 90116, 58057, 86911, 54150, 91192, 57251, 88586, 55118, 93241, 57145, 86771, 61037, 83855, 62369, 83184, 62296, 87791, 1777, -4410, 2216, -3237, -1749, -2850, -1390, -5722, 1005, -2044, -1501, 269, -2349, -740, -5743, 273, -2239, -294, -1844, 895, -3366, 1020, -3881, 1232, -1504, 244, -1090, 2277, 1188, 915, -62, 1294, 2410, -570, 8421, 460, 4643, 2270, 1171, 2714, 7968, -2166, 12230, 1868, 10215, 2732, 7260, 1959, 11229, -5126, 1125, 4344, 6286, 1768, 6925, 545, 27988, -18548, 3675, 2925, -822, 972, 3417, -1219, 20475, 18000, 21096, 24022, 22392, 28415, 23327, 30764, 24406, 30868, 24755, 31761, 25564, 32723, 26205, 31961, 26938, 30589, 21322, 20003, 22312, 24250, 23462, 27018, 24322, 28614, 24821, 30160, 25639, 31202, 25783, 31891, 26003, 32123, 25945, 32365, 26654, 32625, 26165, 32684, 27853, 31451, 26809, 32993, 30145, 29285, 20592, 43560, 23255, 42141, 25119, 44303, 26074, 47078, 26742, 47065, 27100, 49101, 27291, 50481, 27237, 50822, 27981, 50813, 28421, 51895, 28819, 51903, 28875, 53100, 28478, 52876, 29302, 53477, 29130, 53111, 0, 0, 48971, 61057, 56361, 89148, 58408, 87663, 58289, 87049, 58725, 87216, 58916, 90372, 59004, 91448, 59436, 92759, 59813, 93338, 59548, 93981, 59983, 95078, 59930, 96364, 60125, 95794, 59539, 96499, 59985, 96445, 59857, 96092, 60281, 95365, 61998, 93168, 61957, 92449, 64653, 90262, 66398, 87236, 69028, 86275, 69832, 84065, 72126, 79834, 72426, 78661, 75886, 77966, 87793, 68872, 0, 0, 447, -1485, 322, -1490, 2106, -1061, 753, -1162, 5503, -2312, -2831, 2014, -3482, -362, 0, 0, 2206, -1627, 627, -790, 481, 330, 303, 979, 2879, -257, 5491, -987, -3200, -299, 0, 0, 1613, 520, -2118, 1180, -1154, 388, -1589, -35, 2228, 785, 806, 2237, -2436, -411, 0, 0, 1131, 712, -891, 886, 421, -1194, 78, -382, -2223, -672, -8873, 2789, -2407, 465, 0, 0, -1367, 723, -999, 985, -2605, 779, -1578, -414, -3072, 1244, 3955, -378, -3855, 738, 0, 0, 2360, -341, -1826, 712, -947, -1102, -352, -1208, 2518, -658, 1602, 1566, -2699, -426, 0, 0, 2597, -601, 1796, -487, 687, -1376, -1646, -632, -2123, 663, -2355, 2392, -2478, -884, 0, 0, 1132, -2087, 2054, -618, 1459, -790, 959, 27, 2933, -1220, -1346, 3608, -1940, 369, 0, 0, 2845, -4569, 3247, -1374, 3740, -1242, 42, -479, 4975, -2926, -6913, 8150, -5728, 437, 0, 0, 3088, -1381, 2760, -2354, 1838, -1208, -1371, 224, 3338, 590, -1199, 6701, -4059, -748, 0, 0, 4487, 557, -1736, 1229, 567, -725, -1452, 108, -1836, 1711, 3229, 3115, -2860, -432, 0, 0, -2189, 4160, -1341, 1681, -1065, 830, -2167, 541, -4859, 450, -8948, 5121, -5155, 738, 0, 0, 1003, -471, 1020, 903, 131, 748, 159, -930, -6353, 2967, -5828, 6308, -2463, -231, 0, 0, 3229, -497, -2090, 1208, -2031, 652, -1306, 976, -4547, 2279, 7722, -1583, -5319, 370, 0, 0, 3137, -292, 1906, -921, -885, -727, -3208, 184, -4152, 2468, -10751, 6654, -3532, -629, 0, 0, 3301, -5531, 2779, -1987, 341, -623, 1939, -1007, 488, -119, 3661, 3910, -3492, 251, 0, 0, 0, 0, 0, 0, 0, 0, 1440, -809, 1186, -360, 59, 289, 355, 1115, 1246, -556, 484, 283, 1034, 830, 1523, 1163, 88, 1315, 2093, -12, 837, 1898, 1422, 1892, 2028, 1452, 575, 4478, 2154, 3296, 1998, 4813, -12267, 15745, 20748, 5874, 1963, 16671, 17912, 7016, 29656, 13230, 30170, 36593, -13083, 29824, 8714, 35018, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1589, 728, 3322, 512, 1991, 2285, 2231, 2332, 505, 848, 2895, 53, 1500, 1425, 1761, 1913, 818, 1055, 1005, 2066, 2867, 1897, 2889, 2830, 273, 5390, 3422, 7196, 6105, 7805, 9725, 7227, -2123, 18726, 20408, 13257, 22965, 14413, 15392, 15946, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1933, -662, -2615, 2164, -28, -890, 216, -247, 2761, -888, 2102, 5665, 0, 0, 0, 0, 85, 490, 269, 164, 377, 1270, 1266, 2013, 2393, 2458, 9569, 6948, 0, 0, 0, 0, -433, -229, -1372, 1592, -502, 503, 960, 295, 3340, -556, 4768, 2731, 0, 0, 0, 0, -370, 679, -803, 1536, -1368, 3150, -818, 5077, 1060, 8734, 4331, 17434, 0, 0, -4199, 3648, 161, 3082, 196, 1795, -834, 570, -996, -622, -1468, -595, 2352, -2220, 1928, -2134, -4384, 8068, -4347, 7556, -1810, 4259, -1977, 1768, -570, -1216, 1061, -1240, 4471, -2650, 3374, -2967, 2365, 11943, 2260, 12198, 3971, 5586, -1540, 2998, -893, 600, -796, -944, 3460, -2171, 4374, -2187, -5146, 24872, 20572, 15660, 29126, 4383, 7971, 4781, 1209, 3704, -1128, 2611, -2234, 3306, 4120, 1390, -4181, 43777, 2356, 29780, 7917, 22411, 13328, 12519, 17921, 6391, 14482, 3842, 5995, 5589, 1478, 7043, 6444, -2733, 7881, -478, 460, 1258, -682, 1407, 683, 686, -56, 1846, -2786, 2313, -5159, 2257, -327, -523, 3509, -329, 2456, 980, -168, 2141, -369, 4438, -2255, 6374, -2002, 5902, -3474, 5636, -1755, -2653, 5476, -1470, 1014, 669, 271, 4254, -345, 8053, -666, 10283, -823, 10117, -2810, 10448, -2554, -3946, 3498, -4497, 3331, 677, 2571, 6809, 1356, 10993, 1462, 14699, 2346, 13620, 1397, 14119, -4544, -7452, 970, -4591, 10017, -25, 6165, 9288, 6150, 15408, 8054, 18221, 2053, 19656, 6890, 21015, -60, -1591, -787, -577, 3351, -265, 1293, 1294, -495, -1185, -388, 140, -627, -731, -1096, -25, -339, -536, -1123, 164, 2373, -364, 370, 1638, 731, 1998, -6170, -5367, -2233, -3733, -3464, -5028, -5661, -4746, -2594, -5334, -8025, -1843, -5087, -3214, -1432, -4710, 2455, 1407}

const (
	totalPhase        = 24
	scaleFactorNormal = 128
)

type EvaluationService struct {
	WeightList
	phase              int
	passed             uint64
	pieceCount         [COLOUR_NB][PIECE_NB]int
	attacked           [COLOUR_NB]uint64
	attackedBy2        [COLOUR_NB]uint64
	attackedBy         [COLOUR_NB][PIECE_NB]uint64
	pawnAttacksBy2     [COLOUR_NB]uint64
	kingSq             [COLOUR_NB]int
	kingAttackersCount [COLOUR_NB]int
	kingpawnTable      []kingPawnEntry
}

type kingPawnEntry struct {
	wpawns, bpawns uint64
	wking, bking   int
	score          Score
	passed         uint64
}

func NewEvaluationService() *EvaluationService {
	var e = &EvaluationService{
		kingpawnTable: make([]kingPawnEntry, 1<<16),
	}
	e.WeightList.init(featureSize)
	e.WeightList.InitWeights(w)
	return e
}

func (e *EvaluationService) Evaluate(p *Position) int {
	//init
	for pt := Pawn; pt <= King; pt++ {
		e.pieceCount[SideWhite][pt] = 0
		e.pieceCount[SideBlack][pt] = 0
		e.attackedBy[SideWhite][pt] = 0
		e.attackedBy[SideBlack][pt] = 0
	}

	e.passed = 0

	e.kingAttackersCount[SideWhite] = 0
	e.kingAttackersCount[SideBlack] = 0

	e.kingSq[SideWhite] = FirstOne(p.Kings & p.White)
	e.kingSq[SideBlack] = FirstOne(p.Kings & p.Black)

	e.pieceCount[SideWhite][Pawn] = PopCount(p.Pawns & p.White)
	e.pieceCount[SideBlack][Pawn] = PopCount(p.Pawns & p.Black)

	e.attackedBy[SideWhite][Pawn] = AllWhitePawnAttacks(p.Pawns & p.White)
	e.attackedBy[SideBlack][Pawn] = AllBlackPawnAttacks(p.Pawns & p.Black)

	e.pawnAttacksBy2[SideWhite] = UpLeft(p.Pawns&p.White) & UpRight(p.Pawns&p.White)
	e.pawnAttacksBy2[SideBlack] = DownLeft(p.Pawns&p.Black) & DownRight(p.Pawns&p.Black)

	e.attackedBy[SideWhite][King] = KingAttacks[e.kingSq[SideWhite]]
	e.attackedBy[SideBlack][King] = KingAttacks[e.kingSq[SideBlack]]

	e.attacked[SideWhite] = e.attackedBy[SideWhite][King]
	e.attacked[SideBlack] = e.attackedBy[SideBlack][King]

	e.attackedBy2[SideWhite] = e.attackedBy[SideWhite][Pawn] & e.attacked[SideWhite]
	e.attacked[SideWhite] |= e.attackedBy[SideWhite][Pawn]

	e.attackedBy2[SideBlack] = e.attackedBy[SideBlack][Pawn] & e.attacked[SideBlack]
	e.attacked[SideBlack] |= e.attackedBy[SideBlack][Pawn]

	//eval
	var score Score

	var pawnKingKey = murmurMix(p.Pawns&p.White,
		murmurMix(p.Pawns&p.Black,
			murmurMix(p.Kings&p.White,
				p.Kings&p.Black)))
	var pke = &e.kingpawnTable[pawnKingKey%uint64(len(e.kingpawnTable))]
	if e.tuning ||
		!(pke.wpawns == p.Pawns&p.White &&
			pke.bpawns == p.Pawns&p.Black &&
			pke.wking == e.kingSq[SideWhite] &&
			pke.bking == e.kingSq[SideBlack]) {

		var pawnKingScore = e.evalPawnsAndKings(p, SideWhite) + e.evalPawnsAndKings(p, SideBlack)
		score += pawnKingScore

		pke.wpawns = p.Pawns & p.White
		pke.bpawns = p.Pawns & p.Black
		pke.wking = e.kingSq[SideWhite]
		pke.bking = e.kingSq[SideBlack]
		pke.score = pawnKingScore
		pke.passed = e.passed
	} else {
		score += pke.score
		e.passed = pke.passed
	}

	score += e.evalFirstPass(p, SideWhite) + e.evalFirstPass(p, SideBlack)
	score += e.evalSecondPass(p, SideWhite) + e.evalSecondPass(p, SideBlack)

	score += e.Value(fMinorBehindPawn,
		PopCount((p.Knights|p.Bishops)&p.White&Down(p.Pawns))-
			PopCount((p.Knights|p.Bishops)&p.Black&Up(p.Pawns)))

	score += e.Value(fMinorProtected,
		PopCount((p.Knights|p.Bishops)&p.White&e.attackedBy[SideWhite][Pawn])-
			PopCount((p.Knights|p.Bishops)&p.Black&e.attackedBy[SideBlack][Pawn]))

	//score += e.Value(fPawnValue, e.pieceCount[SideWhite][Pawn]-e.pieceCount[SideBlack][Pawn])
	//score += e.Value(fKnightValue, e.pieceCount[SideWhite][Knight]-e.pieceCount[SideBlack][Knight])
	//score += e.Value(fBishopValue, e.pieceCount[SideWhite][Bishop]-e.pieceCount[SideBlack][Bishop])
	//score += e.Value(fRookValue, e.pieceCount[SideWhite][Rook]-e.pieceCount[SideBlack][Rook])
	//score += e.Value(fQueenValue, e.pieceCount[SideWhite][Queen]-e.pieceCount[SideBlack][Queen])

	if p.WhiteMove {
		score += e.Value(fTempo, 1)
	} else {
		score += e.Value(fTempo, -1)
	}

	//mix score
	var phase = e.pieceCount[SideWhite][Knight] + e.pieceCount[SideBlack][Knight] +
		e.pieceCount[SideWhite][Bishop] + e.pieceCount[SideBlack][Bishop] +
		2*(e.pieceCount[SideWhite][Rook]+e.pieceCount[SideBlack][Rook]) +
		4*(e.pieceCount[SideWhite][Queen]+e.pieceCount[SideBlack][Queen])
	if phase > totalPhase {
		phase = totalPhase
	}
	e.phase = phase

	var result = (score.Mg()*phase + score.Eg()*(totalPhase-phase)) / totalPhase

	var strongSide int
	if result > 0 {
		strongSide = SideWhite
	} else {
		strongSide = SideBlack
	}
	result = result * e.computeFactor(strongSide, p) / scaleFactorNormal

	result /= 100

	if !p.WhiteMove {
		result = -result
	}

	return result
}

func (e *EvaluationService) computeFactor(strongSide int, p *Position) int {
	var result = scaleFactorNormal

	result = result * (200 - p.Rule50) / 200

	var strongSidePawns = e.pieceCount[strongSide][Pawn]
	var x = 8 - strongSidePawns
	result = result * (128 - x*x) / 128

	/*const (
		QueenSideBB = FileAMask | FileBMask | FileCMask | FileDMask
		KingSideBB  = FileEMask | FileFMask | FileGMask | FileHMask
	)
	if p.Colours(strongSide)&p.Pawns&QueenSideBB == 0 ||
		p.Colours(strongSide)&p.Pawns&KingSideBB == 0 {
		result = result * 85 / 100
	}*/

	if strongSidePawns == 0 {
		var strongMinors = e.pieceCount[strongSide][Knight] + e.pieceCount[strongSide][Bishop]
		var strongMajors = e.pieceCount[strongSide][Rook] + 2*e.pieceCount[strongSide][Queen]

		var weakSide = strongSide ^ 1
		var weakMinors = e.pieceCount[weakSide][Knight] + e.pieceCount[weakSide][Bishop]
		var weakMajors = e.pieceCount[weakSide][Rook] + 2*e.pieceCount[weakSide][Queen]

		var balance = 4*(strongMinors-weakMinors) + 6*(strongMajors-weakMajors)

		if strongMajors == 0 && strongMinors <= 1 {
			return scaleFactorNormal / 16
		} else if balance <= 4 {
			return scaleFactorNormal / 4
		}
	}

	if e.pieceCount[SideWhite][Bishop] == 1 &&
		e.pieceCount[SideBlack][Bishop] == 1 &&
		onlyOne(p.Bishops&darkSquares) {
		if p.Knights|p.Rooks|p.Queens == 0 {
			result = result * 1 / 2
		}
	}

	return result
}

func (e *EvaluationService) evalPawnsAndKings(p *Position, side int) Score {
	var s Score
	var x uint64
	var sq int

	var sign int
	if side == SideWhite {
		sign = 1
	} else {
		sign = -1
	}
	var US = side
	var THEM = side ^ 1
	var friendly = p.Colours(US)
	var enemy = p.Colours(THEM)

	for x = p.Pawns & friendly; x != 0; x &= x - 1 {
		sq = FirstOne(x)
		s += e.addPst32(fPawnPST, US, sq, sign)
		if PawnAttacksNew(THEM, sq)&friendly&p.Pawns != 0 {
			s += e.addPst32(fPawnProtected, US, sq, sign)
		}
		if adjacentFilesMask[File(sq)]&rankMasks[Rank(sq)]&friendly&p.Pawns != 0 {
			s += e.addPst32(fPawnDuo, US, sq, sign)
		}
		if adjacentFilesMask[File(sq)]&friendly&p.Pawns == 0 {
			s += e.Value(fPawnIsolated, sign)
		}
		if FileMask[File(sq)]&^SquareMask[sq]&friendly&p.Pawns != 0 {
			s += e.Value(fPawnDoubled, sign)
		}
		if pawnPassedMask[US][sq]&enemy&p.Pawns == 0 {
			e.passed |= SquareMask[sq]
		}
	}
	{
		sq = FirstOne(p.Kings & friendly)
		s += e.addPst32(fKingPST, US, sq, sign)
		var file = limit(File(sq), FileB, FileG)
		var mask = friendly & p.Pawns & forwardRanksMasks[US][Rank(sq)]
		for f := file - 1; f <= file+1; f++ {
			var ours = FileMask[f] & mask
			var ourDist int
			if ours == 0 {
				ourDist = 7
			} else {
				ourDist = relativeRankOf(US, backmost(US, ours))

				/*ourDist = Rank(sq) - Rank(backmost(US, ours))
				if ourDist < 0 {
					ourDist = -ourDist
				}*/
			}
			var index = boolToInt(f == File(sq))*64 + f*8 + ourDist
			s += e.Value(fKingShelter+index, sign)
		}
	}
	return s
}

func (e *EvaluationService) evalFirstPass(p *Position, side int) Score {
	var s Score
	var x, attacks uint64
	var sq int

	var allPieces = p.AllPieces()
	var sign int
	if side == SideWhite {
		sign = 1
	} else {
		sign = -1
	}
	var US = side
	var THEM = side ^ 1
	var friendly = p.Colours(US)
	var enemy = p.Colours(THEM)

	var mobilityArea = ^(p.Pawns&friendly | e.attackedBy[THEM][Pawn])
	var kingArea = kingAreaMasks[THEM][e.kingSq[THEM]] &^ e.pawnAttacksBy2[THEM]
	for x = p.Knights & friendly; x != 0; x &= x - 1 {
		sq = FirstOne(x)
		e.pieceCount[US][Knight]++
		s += e.addPst32(fKnightPST, US, sq, sign)
		attacks = KnightAttacks[sq]
		s += e.Value(fKnightMobility+PopCount(attacks&mobilityArea), sign)
		e.attackedBy2[US] |= e.attacked[US] & attacks
		e.attacked[US] |= attacks
		e.attackedBy[US][Knight] |= attacks
		if attacks&kingArea != 0 {
			e.kingAttackersCount[THEM]++
		}
		if outpostSquares[side]&SquareMask[sq] != 0 &&
			outpostSquareMasks[US][sq]&enemy&p.Pawns == 0 {
			s += e.Value(fKnightOutpost, sign)
		}
	}
	for x = p.Bishops & friendly; x != 0; x &= x - 1 {
		sq = FirstOne(x)
		e.pieceCount[US][Bishop]++
		s += e.addPst32(fBishopPST, US, sq, sign)
		attacks = BishopAttacks(sq, allPieces)
		s += e.Value(fBishopMobility+PopCount(attacks&mobilityArea), sign)
		e.attackedBy2[US] |= e.attacked[US] & attacks
		e.attacked[US] |= attacks
		e.attackedBy[US][Bishop] |= attacks
		if attacks&kingArea != 0 {
			e.kingAttackersCount[THEM]++
		}
		if side == SideWhite {
			s += e.Value(fBishopRammedPawns, sign*PopCount(
				sameColorSquares(sq)&p.Pawns&p.White&Down(p.Pawns&p.Black)))
		} else {
			s += e.Value(fBishopRammedPawns, sign*PopCount(
				sameColorSquares(sq)&p.Pawns&p.Black&Up(p.Pawns&p.White)))
		}
	}
	for x = p.Rooks & friendly; x != 0; x &= x - 1 {
		sq = FirstOne(x)
		e.pieceCount[US][Rook]++
		s += e.addPst32(fRookPST, US, sq, sign)
		attacks = RookAttacks(sq, allPieces&^(friendly&p.Rooks))
		s += e.Value(fRookMobility+PopCount(attacks&mobilityArea), sign)
		e.attackedBy2[US] |= e.attacked[US] & attacks
		e.attacked[US] |= attacks
		e.attackedBy[US][Rook] |= attacks
		if attacks&kingArea != 0 {
			e.kingAttackersCount[THEM]++
		}
		var mask = FileMask[File(sq)]
		if (mask & friendly & p.Pawns) == 0 {
			if (mask & p.Pawns) == 0 {
				s += e.Value(fRookOpen, sign)
			} else {
				s += e.Value(fRookSemiopen, sign)
			}
		}
	}
	for x = p.Queens & friendly; x != 0; x &= x - 1 {
		sq = FirstOne(x)
		e.pieceCount[US][Queen]++
		s += e.addPst32(fQueenPST, US, sq, sign)
		attacks = QueenAttacks(sq, allPieces)
		s += e.Value(fQueenMobility+PopCount(attacks&mobilityArea), sign)
		e.attackedBy2[US] |= e.attacked[US] & attacks
		e.attacked[US] |= attacks
		e.attackedBy[US][Queen] |= attacks
		if attacks&kingArea != 0 {
			e.kingAttackersCount[THEM]++
		}
	}
	if e.pieceCount[US][Bishop] >= 2 {
		s += e.Value(fBishopPair, sign)
	}
	return s
}

func (e *EvaluationService) addPst32(index, side, sq, sign int) Score {
	return e.Value(index+relativeSq32(side, sq), sign)
}

func (e *EvaluationService) addPst12(index, side, sq, sign int) Score {
	return e.Value(index+file4(sq), sign) +
		e.Value(index+4+relativeRankOf(side, sq), sign)
}

var kingAttackWeight = [...]int{2, 4, 8, 12, 13, 14, 15, 16}

func (e *EvaluationService) evalSecondPass(p *Position, side int) Score {
	var s Score
	var x uint64
	var sq int
	var sign int
	var forward int
	if side == SideWhite {
		sign = 1
		forward = 8
	} else {
		sign = -1
		forward = -8
	}
	var allPieces = p.AllPieces()
	var US = side
	var friendly = p.Colours(US)
	var THEM = side ^ 1

	for x = e.passed & friendly; x != 0; x &= x - 1 {
		sq = FirstOne(x)
		if forwardFileMasks[US][sq]&^SquareMask[sq]&p.Pawns != 0 {
			continue
		}
		var rank = relativeRankOf(US, sq)
		var keySq = sq + forward
		var bitboard = SquareMask[keySq]
		var canAdvance = allPieces&bitboard == 0
		var safeAdvance = e.attacked[THEM]&bitboard == 0
		var index = 0
		if canAdvance {
			index |= 1
		}
		if safeAdvance {
			index |= 2
		}
		s += e.Value(fPawnPassed+index*8+rank, sign)
		//if forwardFileMasks[US][sq]&^SquareMask[sq]&p.Pawns == 0 {
		//s += e.Value(fPassedFriendlyDistance+rank, sign*distanceBetween[keySq][e.kingSq[US]])
		//s += e.Value(fPassedEnemyDistance+rank, sign*distanceBetween[keySq][e.kingSq[THEM]])
		//}

		var r = Max(0, relativeRankOf(side, sq)-Rank3)
		s += e.Value(fPassedFriendlyDistance+8*r+distanceBetween[keySq][e.kingSq[US]], sign)
		s += e.Value(fPassedEnemyDistance+8*r+distanceBetween[keySq][e.kingSq[THEM]], sign)
	}

	{
		//king safety

		var kingSq = e.kingSq[US]
		var kingArea = kingAreaMasks[US][kingSq]

		var weak = e.attacked[THEM] &
			^e.attackedBy2[US] &
			(^e.attacked[US] | e.attackedBy[US][Queen] | e.attackedBy[US][King])

		var safe = ^p.Colours(THEM) &
			(^e.attacked[US] | (weak & e.attackedBy2[THEM]))

		var occupied = p.AllPieces()
		var knightThreats = KnightAttacks[kingSq]
		var bishopThreats = BishopAttacks(kingSq, occupied)
		var rookThreats = RookAttacks(kingSq, occupied)
		var queenThreats = bishopThreats | rookThreats

		var knightChecks = knightThreats & safe & e.attackedBy[THEM][Knight]
		var bishopChecks = bishopThreats & safe & e.attackedBy[THEM][Bishop]
		var rookChecks = rookThreats & safe & e.attackedBy[THEM][Rook]
		var queenChecks = queenThreats & safe & e.attackedBy[THEM][Queen]

		var val = sign * kingAttackWeight[Min(len(kingAttackWeight)-1, e.kingAttackersCount[US])]

		s += e.Value(fSafetyWeakSquares, val*PopCount(weak&kingArea))
		s += e.Value(fSafetySafeQueenCheck, val*PopCount(queenChecks))
		s += e.Value(fSafetySafeRookCheck, val*PopCount(rookChecks))
		s += e.Value(fSafetySafeBishopCheck, val*PopCount(bishopChecks))
		s += e.Value(fSafetySafeKnightCheck, val*PopCount(knightChecks))
	}

	{
		// threats

		var minors = friendly & (p.Knights | p.Bishops)
		var rooks = friendly & p.Rooks
		var queens = friendly & p.Queens

		var attacksByPawns = e.attackedBy[THEM][Pawn]
		var attacksByMinors = e.attackedBy[THEM][Knight] | e.attackedBy[THEM][Bishop]
		var attacksByMajors = e.attackedBy[THEM][Rook] | e.attackedBy[THEM][Queen]

		var poorlyDefended = (e.attacked[THEM] & ^e.attacked[US]) |
			(e.attackedBy2[THEM] & ^e.attackedBy2[US] & ^e.attackedBy[US][Pawn])

		s += e.Value(fThreatWeakPawn, sign*PopCount(friendly&p.Pawns & ^attacksByPawns & poorlyDefended))
		s += e.Value(fThreatMinorAttackedByPawn, sign*PopCount(minors&attacksByPawns))
		s += e.Value(fThreatMinorAttackedByMinor, sign*PopCount(minors&attacksByMinors))
		s += e.Value(fThreatMinorAttackedByMajor, sign*PopCount(minors&poorlyDefended&attacksByMajors))
		s += e.Value(fThreatRookAttackedByLesser, sign*PopCount(rooks&(attacksByPawns|attacksByMinors)))
		s += e.Value(fThreatMinorAttackedByKing, sign*PopCount(minors&poorlyDefended&e.attackedBy[THEM][King]))
		s += e.Value(fThreatRookAttackedByKing, sign*PopCount(rooks&poorlyDefended&e.attackedBy[THEM][King]))
		s += e.Value(fThreatQueenAttackedByOne, sign*PopCount(queens&e.attacked[THEM]))
	}

	return s
}

func (e *EvaluationService) Size() int {
	return featureSize
}

func (e *EvaluationService) ComputeFeatures(pos *Position) domain.TuneEntry {
	e.tuning = true
	for i := range e.values {
		e.values[i] = 0
	}
	e.Evaluate(pos)
	var result = domain.TuneEntry{
		Features:         e.WeightList.Features(),
		MgPhase:          float32(e.phase) / totalPhase,
		WhiteStrongScale: float32(e.computeFactor(SideWhite, pos)) / scaleFactorNormal,
		BlackStrongScale: float32(e.computeFactor(SideBlack, pos)) / scaleFactorNormal,
	}
	result.EgPhase = 1 - result.MgPhase
	return result
}
