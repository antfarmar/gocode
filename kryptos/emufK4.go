package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var k1 = "EMUFPHZLRFAXYUSDJKZLDKRNSHGNFIVJYQTQUXQBQVYUVLLTREVJYQTMKYRDMFD"
var k4e = "OBKRUOXOGHULBSOLIFBBWFLRVQQPRNGKSSOTWTQSJQSSEKZZWATJKLUDIAWINFBNYPVTTMZFPKWGDKZXTJCDIGKUHUAUEKCAR"
var k4t = "DFDKOPSJUGTHZGQYOSJLNBSKBICTKOKAKMUWKPPXECRFKJZBAAWWLLUSZLAURVOKDVBRQXWWRTGIUESFBTHNTGZIQNUSTFIOQ"
var kall = "EMUFPHZLRFAXYUSDJKZLDKRNSHGNFIVJYQTQUXQBQVYUVLLTREVJYQTMKYRDMFDVFPJUDEEHZWETZYVGWHKKQETGFQJNCEGGWHKKDQMCPFQZDQMMIAGPFXHQRLGTIMVMZJANQLVKQEDAGDVFRPJUNGEUNAQZGZLECGYUXUEENJTBJLBQCRTBJDFHRRYIZETKZEMVDUFKSJHKFWHKUWQLSZFTIHHDDDUVHDWKBFUFPWNTDFIYCUQZEREEVLDKFEZMOQQJLTTUGSYQPFEUNLAVIDXFLGGTEZFKZBSFDQVGOGIPUFXHHDRKFFHQNTGPUAECNUVPDJMQCLQUMUNEDFQELZZVRRGKFFVOEEXBDMVPNFQXEZLGREDNQFMPNZGLFLPMRJQYALMGNUVPDXVKPDQUMEBEDMHDAFMJGZNUPLGEWJLLAETGENDYAHROHNLSRHEOCPTEOIBIDYSHNAIACHTNREYULDSLLSLLNOHSNOSMRWXMNETPRNGATIHNRARPESLNNELEBLPIIACAEWMTWNDITEENRAHCTENEUDRETNHAEOETFOLSEDTIWENHAEIOYTEYQHEENCTAYCREIFTBRSPAMHHEWENATAMATEGYEERLBTEEFOASFIOTUETUAEOTOARMAEERTNRTIBSEDDNIAAHTTMSTEWPIEROAGRIEWFEBAECTDDHILCEIHSITEGOEAOSDDRYDLORITRKLMLEHAGTDHARDPNEOHMGFMFEUHEECDMRIPFEIMEHNLSSTTRTVDOHWOBKRUOXOGHULBSOLIFBBWFLRVQQPRNGKSSOTWTQSJQSSEKZZWATJKLUDIAWINFBNYPVTTMZFPKWGDKZXTJCDIGKUHUAUEKCAR"
var kabc = "KRYPTOSABCDEFGHIJLMNQUVWXZ"
var abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

/***************************************************************************/

func main() {
	rand.Seed(time.Now().Unix())
	key := "EMUFPHZLRFAXYUSDJKZLDKRNSHGNFIVJENDHOHNLSRHEOCPTEOIBIDYSHNAIA"
	//crib := []string{"MUY", "UYK", "YKL", "KLG", "LGK", "GKO", "KOR", "ORN", "RNA"}
	//fmt.Println(key, len(key))

	for i := 1; i < len(key); i++ {
		println(tpose(key, i, 0), i)
	}

	//println(shift(k4e, key))
	//println(quag3(k4e, abc, "MUYKLGKORNA"))
}

/***************************************************************************/

func tpose(input string, a, b int) string {
	output := ""
	m := len(input)
	for x := range input {
		output += string(input[(a*x+b)%m])
	}
	return output
}

func shuffle(s string) string {
	r := rand.Perm(len(s))
	t := ""
	for _, n := range r {
		t += string(s[n])
	}
	return t
}

func check(s string, words []string) string {
	for _, w := range words {
		if strings.Contains(s, w) {
			s = strings.Replace(s, w, "["+w+"]", -1)
		}
	}
	return s
}

func rot1(r rune) rune {
	return 'A' + (r-'A'+1)%26
}

func quag3(ciph, alph, key string) string {
	plain := ""
	for i, c := range ciph {
		ltr := string(key[i%len(key)])
		n := strings.Index(alph, ltr)
		m := strings.Index(alph, string(c))
		diff := m - n
		if diff < 0 {
			diff = diff + 26
		}
		plain += string((alph[diff]))
	}
	return plain
}

func shift(ciph, key string) string {
	s := ""
	for i := range ciph {
		c := rune(ciph[i])
		k := rune(key[i%len(key)])
		diff := c - k
		if diff < 0 {
			diff = diff + 26
		}
		s += string('A' + diff%26)
	}
	return s
}

func OTP(key string, crib []string) {
	for h := 0; h < len(key); h += 1 {
		for i := h + 1; i <= len(key); i++ {
			sub := key[h:i]
			for j := 1; j < len(sub); j++ {
				for m := 0; m < len(sub); m++ {
					ik := tpose(sub, j, m) //[:j]
					q := quag3(k4e, kabc, ik)
					t := tpose(q, 38, 76)
					w := check(t, crib)
					if w != t { //SECRET,HOURS
						fmt.Println(w, h, i, j, m, ik)
					}
				}
			}
		}
	}
}

func ALLTP(key string, crib []string) {
	size := len(key)
	for a := 1; a < size; a++ {
		for b := 0; b < size; b++ {
			ik := tpose(key, a, b)
			for c := 1; c < 97; c++ {
				for d := 0; d < 1; d++ {
					q := quag3(k4e, kabc, ik)
					t := tpose(q, c, d)
					w := check(t, crib)
					if w != t {
						fmt.Println(w, a, b, c, ik, size)
					}
				}
			}
		}
	}

}

//QSMGICNYPOAUONBRNZQMXMXLQZOTUFRDCUZIDCBTBAZBXRFXKGJUVQSPPATKUQXEL[BERLIN]UFCCOMXCDPRQFIXGRXFCOPACUP 05 00 38 01 EHMZULFRP 9
//XBVBEYSQQEMDKSVOCGUSCYOJNMOGZQIISUNOCVMYEUQBDQTAOVCGNHKLBCAULO[INCLOCK]TSKDBEUQGOHXAQMQOQVIRHVSXNUS 02 10 44 33 AYSJMFHLFXUDEUPZR 17
//QBVBEYSQQEMDKSVOCGUSCYOJNMOGZQIISUNOCVMYEUQBDQTAOVCGNHKLBCAULO[INCLOCK]TSKDBEUQGOHXAQMQOQVIRHVSXNUS 15 00 53 64 EDUXFLHFMJSYARZPU 17

func KORNA(key string, crib []string) {
	for d := 0; d < len(key); d++ {
		t := tpose(key, d, 0)
		w := check(t, crib)
		if w != t {
			fmt.Println(w, d)
		}
	}
}

/*
k4=tpose(55,0) ik=tpose(52,26) len(ik)=54
QVEYDUXGYPDYBGWDIATFDWPO CLOCK AAARHGL KEY VFTISHVYENSSUIHLGR SECRST BKIJFVHDOWQLDTIAOETCDDROVNTTIBFGVU 1 55 52 26 NHNKLKDUXFLHFMQJETLUVBXQQJINHNKLKDUXFLHFMQJETLUVBXQQJI
k4=tpose(38,76) ik=tpose(5,12) len(ik)=77
NDKAETTIRPTPGRMAQUQIPNPQDMAGFOSOFSDPIPAQJXLFVOD SECR NJJUDCTGD HOURSRUN JQNWZFAVZDFXHCR CHRSOUP ZCVDPZY 12 89 5 SITBVETDFEEGQYKRNYXYTYYDUZYKGDDHVQQLVMMPETWEUZNFQQURQRVDWVKFJKGJUVLJKFJHZHTSL
k4=tpose(38,76) ik=tpose(58,45) len(ik)=95
DXVQTAX INCLOCNKK CUT GFAGCFNCAWDKGSJPFIDNPCNJYDOGLQJYGZTCHIMCUWZLDTQMFNPQPKOZVNZVJCNWNCCOJCZBBBKEHB 0 95 58 45 LRJIGVUHTCMZZBUMNHVLPFTEYEQNTKTQMDRWUZFNERXEYJQJEXERKGYHVGQTADJQYDWUGYDVVPDHKLFUVFJSZQEKLYQFFSK

berlin @ 64 tpose(38x+76)
PGZFGTEZPFCYGFLJLZGSOJAILWDCCNZFHSHQHIVZKUCGOZAAFJEVHOCZVSNHYEHBERLINPEEKRRRHBUXJOPYFFTXBMMQMOFQS 11 EMUFPHZLRFA
IHXZYLMCKIMFVRSUSBTIJWMNCLWYWQFKRQTUWMISFEDDZOEJBCSUJSAHKVJNKENBERLINGEOYCVPQKQQRCWAPDGWGXWQPCPLG 13 8 5 HERFXZMFPYLUA
If I hid letter...
IFIHDLETARJYEUSDLHLZQJAQCCYXNZNBNNHXOTBZJMSUOZILWCGZANTBBNJCDDHHKDXIDBGAZHEZNDTESPEMRFVTLBMUTRIQQ 21 EMUFPHZLRFAXYUSDJKZLD
clock (keysize,shift)
QVPWEDTLDXEGTNSEYTULCWFXBLUKHDCITFCETJNDREKVBBOCOBPBEJKJEDTNBVSIJPBSKQSYDCLOCZTZDJTAPVAAHQCDDYCUV 44 62 DEMUFPHZLRFAXYUSDJKZLDKRNSHGNFIVJYQTQUXQBQVY
XQQGZISFQSFFTHUDTGDVNFINVGHQRPZHYWRRCFLGOZIDCSRXGEASJSITGGJEKFCFQNADQVFNDTFOZWECJGHANWFISCLOCQWDU 47 43 UVLLTREVJYQTMKYRDMFDEMUFPHZLRFAXYUSDJKZLDKRNSHG
MVSHLGYIAUKNKUYAWKXLEHZKSUYOXTABOLZSOLATRWZMOMWHFCSPSETQHKWCLOCDKUFPMWNZZPCNADPJLTEBGEZTDLREQQURB 55 34 TQUXQBQVYUVLLTREVJYQTMKYRDMFDEMUFPHZLRFAXYUSDJKZLDKRNSH


K1 as OTP for K4                                                                                    kt ks k4t
GOCKBGWIIWWEKIJLMXVAVBHNYCRCIHMPGHFRZRGBIGVBMPSLWTEYGOGRACJG[CLOCK]WDMQNVURQOVQTSETUWSEFCQSAFSUCGOT 03 01 48 MPLUHEFZ 8
HGLGH[NCLOCK]UCGVWMIZBUQUWEJQYKVJDOJMTOGOCNEKXUYYXQRHZVXKOVASOUPOSEGOUAASOAKVTOHEADKCZYPJJXXYXTEKRU 25 07 92 LYMGJDLSBRQUNYKLDQFTFFQRTJVAQPITNRKYXUHVMSEZUYXZJEHVLVUQ 56
QSPGZFGTEZPFCYGFLJLZGSOJAILWDCCNZFHSHQHIVZKUCGOZAAFJEVHOCZVSNHYEH[BERLIN]PEEKRRRHBUXJOPYFFTXBMMQMOF 01 00 38 EMUFPHZLRFA 11
LGIHXZYLMCKIMFVRSUSBTIJWMNCLWYWQFKRQTUWMISFEDDZOEJBCSUJSAHKVJNKEN[BERLIN]GEOYCVPQKQQRCWAPDGWGXWQPCP 08 05 38 HERFXZMFPYLUA 13
NUMPMBCBFVOTVDQSWUVK[BERLIN]KGIXXNWUIYMBBFSDBJCRZIQTMREBTRWEEKLMXFMRUVRWBFRKDGYXNYPVDETTZNZEVJVUFKV 23 03 40 FGEKQRJTRLUUPNVZVFYENLSXHFJLYAQMSTDQZIYDUXTUHRJBLVQKVYQ 55
NNVEDLKOBFMMAMFNBXOKNZBJXAWRLEIARANWQVXJVLJXNHJXZAOHNVFPOZ[BERLIN]RKSTMOWIMVPQOCEXGCFNBOTQYDGWDQENT 02 64 16 FEUPZRAYSJZDRSGFVYTUQQYVLRVYTKRMD FEUPZRAYSJZDRSGFVYTUQQYVLRVYTKRMD 66 (33)
DLVXWGDFBRXFWBBHJXNR[BERLIN]YQDIDPPSXKRKWXIEFBESFJBRVICXNECAOSXWFRJCGITECSBTJHSDQKQUQJGAUHCJTBRWJXW 22 13 22 UQYXQMFJQLIJHNEFHTMNLUKUPLVVKBFDXD UQYXQMFJQLIJHNEFHTMNLUKUPLVVKBFDXD 68 (34)
IPCAGYKUAULKEZGYBKLCH[BERLIN]VOJJBYHMITZMFCDBLTUYUCHWOXOADGCXFWLELARCOUBXFGTCQLNFVOUBMDJCGNPBHHURPI 51 66 04 JTFFMVRFTQJURIAFYNPMUKDEVXDUSHKXZEVJYVVHZYQLEJYUFLGLRBDMYQSPLNRDQKUQTD 70
FDNQEO[BERLIN]KRVZEUTEZKKXLTGLKBZLNGFLUYCJGQOSXCZITKUNQSIHHWNQHSQHFXHMSCDSWRWTNBQSTIOUQDLNOPWLYSHSU 29 27 87 NKABDRJHTVKLEIRYVESQLUPLRUJMSUZGMFQUKVPQDJLEFYXQENYZQFZTMVDUYHHTRXJDEFYFDVW 75
IVYHDAHRLEOZDIYZOAJTOAFKDFQ[BERLIN]OUKTMIKGRXUNAKJWWCNLKMGCOPVWIQJVMVEOZTNHGUUOCLQWNITMLCSWMYJYRAOZ 29 66 24 JABDYVESUZJLEZTZDEVRJWSQKGMQFYMVDFYFHTVLUPFQUXQEUYHDVWKLTLRYKVGNYHHTKNKEIRUJMPQDZQFRX 85
QAWMXKJQWAGDEHZUDWVYEQBWXTGDEECHXYJTTUGYRFYZOCJUVYFVJVPIDKRIZXNPRMACBTGVLEFJVZNGKQQ[BERLIN]LGHKUSIA 85 49 54 ERTLLVUYVQBQXUQTQYJVIFNGHSNRKDLZKJDSUYXAFRLZHPFUMEEQKKHWGVYZTEWZHEEDUJPFVDFMDRYKMTQYJV 86
PJUUXSBVOYMPNJCZITHCXSATBLTTQVCZOHYHUNNQSOUPEAHTDTXQBBVBXGIFAYYNVLTZKKLTDNEDQWKHRO[BERLIN]VRUWBOGMT 05 12 84 YKRNYXYTYYDUZYKGUL YKRNYXYTYYDUZYKGULYKRNYXYTYYDUZYKGULYKRNYXYTYYDUZYKGULYKRNYXYTYYDUZYKGUL 90 (18)
D[BERLIN]JFVRNCRXMKETTQLCACAFGCSHPIPBJVUYNGPEDHDDJTUEXSFRPUZGALEGBSTVYTARAKRXFRWIBIRYKZVKYLZXSDQKEO 09 13 38 URJQERUZEGZDSQYJMEVGWRKGQVQDHWQMALFXLMFWKNFYKVBRYJTQEHSNYVVDDYTGLJHTUYFEGFEFZNULTVZHJUXDIQTKPEKCP 97
*/