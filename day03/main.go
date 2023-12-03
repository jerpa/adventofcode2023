package main

import (
	c "github.com/jerpa/adventofcode2023/common"
	"github.com/jerpa/adventofcode2023/day03/part2"

	"time"
)

func main() {
	start := time.Now()

	//c.Print("Part1: ", part1.Part1(c.GetData(Data)))
	c.Print("Took: ", time.Since(start).String())
	start = time.Now()
	c.Print("Part2: ", part2.Part2(c.GetData(Data)))
	c.Print("Took: ", time.Since(start).String())
}

var Data = `
................................................965..583........389.................307.................512......................395.....387
........................#....374...382....250...*..........737*....*896.395...........*....................$.........................#......
..494.........532-...474......*.......#....*...................522......*..........%...........................%...+................269.....
.....*..#................506..143........375......77.....155...........400.518...64....773...718..797........694....972.603.....*...........
....479.795...............*..........800...........*.$.......264*636.......@..............&..*...*.......499...............*...5.20.........
515...................512.484...*....*...=......390...427...................................644.804.........*...@......-..532............28.
..........607...........&.....105...906...910.......@............979.969...........-..=.............462....414..101.361..........283$.......
..........*...781................................925..............=...*..........434..899....368.......*..................33...........*....
.......850.......=........559..249...................732.....430....132........................*.....................817-.=.........613.381.
....................157....*...&....978..............$......-...................*626.-.......297.............750............................
..........312...........606........*....136.............593.............638...........177.....................*....672.772....998.491=......
....450......*..156...............760.....=.......227......%..794...157........................919.......&....607..*...*....................
.......*...108...............968...................-..........*........*....80....669..........$..........854......6..572.....743....&......
......759.......649.............*832.......351.........433...418...997........*.....*.......46....187...........................*.....759...
............344.&..........&.........573......@...556.............*..........884...748..317../....+..........................236............
......#..................627.....334*..............*...933........774.534................*.......................................905........
...642.....105..........................100.......372.....$...........*..........*982..527....$837..406.948....*127......455........*.......
...........*.......383............293......*750.&.....774....630......51......826...........................361.........=.......993..822....
....*..576.346.876....@..907.953..@..............982..*.........+.................913*782...*115..565..526........779.......971.............
..20...*..........*.......*....%.............46-.......8................232................5.....................*....@........$........903.
.....$.............483.....420......115....$......287...............................873.......................131....818...............*....
......161....*362.....................-.266.........*......524...157.79....@.396.......*............848@.293............................988.
..........412...................................97..721.....#.......*...638../......305.......................+.................630*818.....
.194..................556....297.253...........*.........................................................@.....186..213..819................
................154.....$.......*....454......46.795.686.......=..204.........$...=............771.....171..........#................-456...
..................*.....................*........*.....-.....457.....*....638.466..579......................#..........109.....727..........
..................177......785.769...827..........500....168........479....*...........=....255..............74...........*845....*..375....
................-.........*.....*..........................*..............360..681......700.......................853...........781.....*...
..846..91......966........511.435.............245.......617.............*.....*....................................$.....810...........836..
.....*.................................%86......*..............121.......721..751......386*.......476*60...%195......801...%.......345......
..763........+....428@....-40.391...............970......803......=...*....................351...................185....*.....118...*.......
......#862..345....................629......45............+.........897.....802........488..........................*.....274.../..676../...
..945.............299.......$......*..........................183..................190...&........763.....484..364...235....%..........564..
...*..............*........404...527....&....137.814.......57......890.......407..*.........................*...*..............711&.........
....424..........66..................966.......*.@.........*......*......806...+........................98....762......$...663........605...
....................404......................614...489..........150....-..*........367...................*.........909.863...#........*.....
.............%759...*...+....916.498...&.............*.551..........191..606...............797........391.......................=568...898..
........2*........632.987......*....*.155.$2......412........................730...401.....#....438..........412*199........................
..........729...............89.753...........364.........687....................*.....*86........*.............................%.....657....
...............182./687.214.*.......................141..*....................$..620.......706%.368.....*.............758....694.971....*774
.....617........*.........*..335....89....583.........*..378...945.....272...97......................264.817.............*........*..+......
.../....*506.779...31..168..................*......732............*.../.....................957......................104.156...160..219.$...
....789.............*....../........671...45....................982.&.......%.......$.......*...264.............386..*..................588.
..................575.......192...../..................51*952.......264....655..=17..113...998.............932.......638..........24*.......
...........@............@.......................................#................................650..........*734...........652.....278....
........903.....=....805......944..............................377........477........&73.610.......+...640..........852.........*893........
.............882.........@768..%..103........465......636.............................................*..........29...*.661.295......61.....
.....................467.........$......292...*.........&..........640.......926.........201..173...38...........*..805.*.............*.....
......296....317.......*.....995............$.889..890....665.......*....47..@....701&..$.......$......739....388.......257............790..
..989..*.....&....980..168.....*...........17........*.....*.......894..*.........................*200.*....................................
..../...561....................198.34*811......726.122...+.63..437......787.........668........355.......641....@...........................
.................................................*.....112.....*................798*.......420..........*........244..758..............771..
...&.......573..............515..................252........511......................396...%.......805.736............/..........546........
342.......*....................*.44.................................*....399.188.....*........................254........13*................
........%..720.............695......875*362.........@...649*919..585.477...*..*...245....450....-............*.......743....645.&......796..
.....556.........460.........$..............919&....679...................935.............*...273.........396.........*..........659........
...&.............-.........%.....542*...............................................255.368........387..............435..612*...............
980................484...964.........736........%327.......210..+......962.610...............396.............492.............259.970*.......
....-.......850...............645...........56.........275*.....928.....*.....-..81....133........=..889....*........................382....
..382................419*.......*......@......*.............439........953.......*....*..........341..*..322.....226.342.......*598.........
...............636=......662.204........620..13..............*...171..............19..396............27.............*.......539........278..
.........131.........133........................282...900...48...*........678.379................+.....................798.......131....@...
.....684*.....944.......*...@.....991...................*.........90.......*..+...818.783....19..492...501*842..........*..480*..@..........
..........515..........855...646.+.......451.../533.....115...../.........841.................#........................272...........154....
....107=..*................................&..........=.......631.....837.............52..............485......193.%.......983......*.......
..........587............405*329....957............451............60.*............370......................222..*...461....@.......894..534.
..705...............131...............$...640...............127*8.+..859.145.......%...440...%..................875.........................
...*..714....783..........@986......&........*.........................................*...199.451....=.....................87.......190....
.349..*......%.....429...............602..562..519................=......112*....669....8..............105...........#......../..387........
.....792.............*.......854..............*.................93...........808...$...........%673........802$.......586..........*........
.............-.......282.....+...-756.......291.........&...............................209.......................534.......573.831....*955.
..........195............401..........617.............676.........................................961..716...................*..............
.....762................*.......*864.........%694..............249./....-....572.327..........784..*..........@.534........883.........724..
......%..............#...109.304....................................56...240.*...*......657...%....837..792.292...=...............*.........
...................974..............89..............629....#460..............330.....*....*..............*................620#...887........
...877*793..................680+...*........$662...*..............................365.898..773.........726...........607.............579....
.................................316.404...........716.....*775.....%.........................................573.67.#...399.953......*.....
.......931...................91@.....*.........680.....395.......760......165....302............960....*.........*.......*.........697......
.......*......962..997/.............363.....*....*....@....126........566*........$..4..739*..........545.953.............955...............
....548.........%................./.....*.734.....653.....*...........................#.....489..............*....336.763...................
........@.185........88.........483...208......-..........419..331......460.../46.406.......................582.-....*..........408.........
.....106../......654..*.....................423.......109.........*................*............269..%21.........632.....324+......*...*....
................%....265......906....926.............*............475..60........188...............*.............................994..905...
........=781....................*....*.....938....401.......857.......*.......................802.300..625.....207..%.......................
..................355..894...410....402...*...................*.....467................660*..*.........%.......*....721.....&....764........
...........554...*.......*..............277..-.......648.......491..........................968...........24....380......520....*...........
......%....*.....757......788.................851......*...416.....493........635....241.........409.....*..........362......660............
.......23..52..........@............................791.....*......+...=..*.........../...$209...*......282..214...=....370.................
...................328..291.529.........................50.63........412.591..../..............169........................+........+........
..701&....=..198.....*.........*....+...........896......$..../................772......*626........458..578=...............337..95...482...
.......540...&......989....227..736.229.....................52...........241..........95.............*..................599...*......*......
........................................820.397................937...664*...............................820...222.185...*.....47....862..192
...........................$.....484.....*..$.......$601........*...............540.........&.............*..............694................
.....818.630..........15..495...........119...36.........273.....329..=.398....*...........578....220...607..580.......%.......876..........
.....*........=......*...................................*...........30..=......89..........................*.......182....574.$............
...98...127....182.363..921..............................273.........................................@231..926...............=....=18.......
........&.................................183........7@.......................347=.....=......903....................*732..............782..
.............*635............................%.604*.........@.#910.226...............288..........767.............227............753/..*....
770.......661............812...$709..430...........997....556......*.....................@...970...........949..........................775.
............................................984.....................845.......$........861.......297..415.*...............789..*937.........
..541.......614..10.........453.....................67.....................735...../..............*..*.......................*..............
...%.../....*..........468...=........*824.109........*58............318............29.....152..538.488.....423.......19......602...162.....
......279.363.228........*.........221.............&.......*..................................@................*788....*...........*........
706....................114..............&987.873..316...701.871.....846...501......200.603..........876.................929......178........
...*..503.......78.................................................-......*..........*............+.........179......$......................
.850.....$..118...+.............944..........19.............@735..........388..160...365..582....524..........*.......355.....840...........
..............*......*..........=...475+.926*.....433.............../146.......*..............................380.....................614...
...........228....200.695..........................*..391..529..................918...290...340..16*.....482................363.........$...
.......906......................993......=986...874...*...#....+............37...........*....=...........#...................*.....32......
...&......*230..204=........721...*.$...............839.....456..553..601..*...408.....914..................77.................662..*..=554.
...29.......................*...717.248.....358.................@......$...977...*................94.....................*846...............
.............542..546.....52.................*........286........................669......................=...401......61.....709*389.480...
........989......*....253....713.........742..171.........325....976....*...513/....................967.246....*........................*...
...........=....571...*........-..=456.....................*........*..300.......344..........146..............87........................66.
..199*................204..............+.490.......584...427.166...810.....431......+...115....*..........195...........=...................
......247.645....................693.204....-.....*....................310...*..*........*.....772....686*.........=323.194....451......=...
.....................596.304*696.@....................................*....220...351..444..934..................................+........136
..736...............&..................................510..250..638.766..................*....511....195.........#..778..&902....721.......
...........758............*..........=401......210..#.....&.-...*................572...385....*.......%.....200.248.*..............*........
...........*...........991.358.256.........782*......742.........832....*.......=..........799...675*......*.........80.819.902...18........
...........735...*.............*........*.......*242.....414*704.....194..325...........................569..............$...@..............
..................772.......309...+..204.....659.....&.....................*............297*..396*494.......26......................131.....
.....%..+.......................169...................860..333...212*....459.............................../.......392.............@.....103
....628..891....../....519..88............................*..........530......90*....225.418....112....614....274..$......169%.399...481....
850............829..48..%.........931........./388.......270..............@......798......#........@....*........$...............-..*.......
...*876...............*...820.....*............................614*834...745............................575..................-......827.....
......................764....*..432...........516$.133..633..................489$........396.255.............636..413*.811....337.......*682
.667......................639...........199..........*...+......=....*500.........401@......*...........878-..*........../...........419....
....@......28...........................*...#....@..72.......539..708........237.................=............675.851@......................
.......777........406...459$..........993.445.719......./......................+...706........386....%....../..................*96.271......
.311..*.......................&......................317../542....................@.....+.........136......821.543..........947.............
...*..982.....657..70=..948#..159..777../.....757....................816...............887.............325........*....................871..
.804..........=.....................*..284..........255*...............+....48...589*......199................#.68..542......&.&241...*.....
.......10............604.....287...66......507..........917...585............-.......6.327..-......822.....718.....*.......825.......34.....
........*......904$....*..................+......*.............*.....454.820............@.....492.*................887............@.........
.......645.-.........49..............844.........533........958.........*.......32.892..........%.118...+..670*296......%...122..840.230....
....=......176..........................*.=..........................%1.....397*...........820.........407...............21...........*.....
...913...........&............408.....135.300..............775...993.........................@..............%.....272...........*626...101..
..................95..505......*.................581.........+...*...........59*23.......449.......964...657.......*..........75............
.............................87...622..........................822...............................................215.............810........`
