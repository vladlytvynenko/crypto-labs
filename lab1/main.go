package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	ALPHABET                   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	SECRET1                    = "Yx`7cen7v7ergrvc~yp:|rn7OXE7t~g.re97R9p97~c7d.xb{s7cv|r7v7dce~yp75.r{{x7`xe{s57vys;7p~ary7c.r7|rn7~d75|rn5;7oxe7c.r7q~edc7{rccre75.57`~c.75|5;7c.ry7oxe75r57`~c.75r5;7c.ry75{57`~c.75n5;7vys7c.ry7oxe7yroc7t.ve75{57`~c.75|57vpv~y;7c.ry75x57`~c.75r57vys7dx7xy97Nxb7zvn7bdr7vy7~ysro7xq7tx~yt~srytr;7_vzz~yp7s~dcvytr;7\\vd~d|~7rovz~yvc~xy;7dcvc~dc~tv{7crdcd7xe7`.vcrare7zrc.xs7nxb7qrr{7`xb{s7d.x`7c.r7urdc7erdb{c9"
	SECRET2                    = "1c41023f564b2a130824570e6b47046b521f3f5208201318245e0e6b40022643072e13183e51183f5a1f3e4702245d4b285a1b23561965133f2413192e571e28564b3f5b0e6b50042643072e4b023f4a4b24554b3f5b0238130425564b3c564b3c5a0727131e38564b245d0732131e3b430e39500a38564b27561f3f5619381f4b385c4b3f5b0e6b580e32401b2a500e6b5a186b5c05274a4b79054a6b67046b540e3f131f235a186b5c052e13192254033f130a3e470426521f22500a275f126b4a043e131c225f076b431924510a295f126b5d0e2e574b3f5c4b3e400e6b400426564b385c193f13042d130c2e5d0e3f5a086b52072c5c192247032613433c5b02285b4b3c5c1920560f6b47032e13092e401f6b5f0a38474b32560a391a476b40022646072a470e2f130a255d0e2a5f0225544b24414b2c410a2f5a0e25474b2f56182856053f1d4b185619225c1e385f1267131c395a1f2e13023f13192254033f13052444476b4a043e131c225f076b5d0e2e574b22474b3f5c4b2f56082243032e414b3f5b0e6b5d0e33474b245d0e6b52186b440e275f456b710e2a414b225d4b265a052f1f4b3f5b0e395689cbaa186b5d046b401b2a500e381d4b23471f3b4051641c0f2450186554042454072e1d08245e442f5c083e5e0e2547442f1c5a0a64123c503e027e040c413428592406521a21420e184a2a32492072000228622e7f64467d512f0e7f0d1a"
	SECRET3                    = "EFFPQLEKVTVPCPYFLMVHQLUEWCNVWFYGHYTCETHQEKLPVMSAKSPVPAPVYWMVHQLUSPQLYWLASLFVWPQLMVHQLUPLRPSQLULQESPBLWPCSVRVWFLHLWFLWPUEWFYOTCMQYSLWOYWYETHQEKLPVMSAKSPVPAPVYWHEPPLUWSGYULEMQTLPPLUGUYOLWDTVSQETHQEKLPVPVSMTLEUPQEPCYAMEWWYTYWDLUULTCYWPQLSEOLSVOHTLUYAPVWLYGDALSSVWDPQLNLCKCLRQEASPVILSLEUMQBQVMQCYAHUYKEKTCASLFPYFLMVHQLUPQLHULIVYASHEUEDUEHQBVTTPQLVWFLRYGMYVWMVFLWMLSPVTTBYUNESESADDLSPVYWCYAMEWPUCPYFVIVFLPQLOLSSEDLVWHEUPSKCPQLWAOKLUYGMQEUEMPLUSVWENLCEWFEHHTCGULXALWMCEWETCSVSPYLEMQYGPQLOMEWCYAGVWFEBECPYASLQVDQLUYUFLUGULXALWMCSPEPVSPVMSBVPQPQVSPCHLYGMVHQLUPQLWLRPOEDVMETBYUFBVTTPENLPYPQLWLRPTEKLWZYCKVPTCSTESQPQULLGYAUMEHVPETFWMEHVPETBZMEHVPETB"
	SECRET4                    = "KZBWPFHRAFHMFSNYSMNOZYBYLLLYJFBGZYYYZYEKCJVSACAEFLMAJZQAZYHIJFUNHLCGCINWFIHHHTLNVZLSHSVOZDPYSMNYJXHMNODNHPATXFWGHZPGHCVRWYSNFUSPPETRJSIIZSAAOYLNEENGHYAMAZBYSMNSJRNGZGSEZLNGHTSTJMNSJRESFRPGQPSYFGSWZMBGQFBCCEZTTPOYNIVUJRVSZSCYSEYJWYHUJRVSZSCRNECPFHHZJBUHDHSNNZQKADMGFBPGBZUNVFIGNWLGCWSATVSSWWPGZHNETEBEJFBCZDPYJWOSFDVWOTANCZIHCYIMJSIGFQLYNZZSETSYSEUMHRLAAGSEFUSKBZUEJQVTDZVCFHLAAJSFJSCNFSJKCFBCFSPITQHZJLBMHECNHFHGNZIEWBLGNFMHNMHMFSVPVHSGGMBGCWSEZSZGSEPFQEIMQEZZJIOGPIOMNSSOFWSKCRLAAGSKNEAHBBSKKEVTZSSOHEUTTQYMCPHZJFHGPZQOZHLCFSVYNFYYSEZGNTVRAJVTEMPADZDSVHVYJWHGQFWKTSNYHTSZFYHMAEJMNLNGFQNFZWSKCCJHPEHZZSZGDZDSVHVYJWHGQFWKTSNYHTSZFYHMAEDNJZQAZSCHPYSKXLHMQZNKOIOKHYMKKEIKCGSGYBPHPECKCJJKNISTJJZMHTVRHQSGQMBWHTSPTHSNFQZKPRLYSZDYPEMGZILSDIOGGMNYZVSNHTAYGFBZZYJKQELSJXHGCJLSDTLNEHLYZHVRCJHZTYWAFGSHBZDTNRSESZVNJIVWFIVYSEJHFSLSHTLNQEIKQEASQJVYSEVYSEUYSMBWNSVYXEIKWYSYSEYKPESKNCGRHGSEZLNGHTSIZHSZZHCUJWARNEHZZIWHZDZMADNGPNSYFZUWZSLXJFBCGEANWHSYSEGGNIVPFLUGCEUWTENKCJNVTDPNXEIKWYSYSFHESFPAJSWGTYVSJIOKHRSKPEZMADLSDIVKKWSFHZBGEEATJLBOTDPMCPHHVZNYVZBGZSCHCEZZTWOOJMBYJSCYFRLSZSCYSEVYSEUNHZVHRFBCCZZYSEUGZDCGZDGMHDYNAFNZHTUGJJOEZBLYZDHYSHSGJMWZHWAFTIAAY"
	ENGLISH_IDX_OF_COINCIDENCE = 0.0667
	ENGLISH_TRIGRAM_IDX        = -3.429451398834384
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func lab1() {
	// LAB1
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		fmt.Println(string(cezar([]byte(SECRET1), byte(i))))
		fmt.Println("\n--------------------------")
	}
	fmt.Println(string(cezar([]byte(SECRET1), byte(23))))
}

func lab2() {
	// LAB2
	decodedSecret2, _ := hex.DecodeString(SECRET2)
	for i := 1; i < 20; i++ {
		fmt.Println(i, indexOfCoincidence(stringRepeatBy(decodedSecret2, i)))
	}

	np := nextProduct([]byte(ALPHABET), 3)
	dest := make([]byte, len(decodedSecret2))
	trigram := NewTrigramGenetic("3grams.txt", "out.txt", decodedSecret2)
	out := resultSliceStruct{}
	for {
		product := np()
		if len(product) == 0 {
			break
		}
		vigenere(dest, decodedSecret2, product)

		isAllPrintable := true
		for _, c := range string(dest) {
			if !strconv.IsPrint(c) {
				isAllPrintable = false
				break
			}
		}
		if isAllPrintable {
			idx := trigram.TrigramEstimation(strings.ToUpper(strings.Replace(string(dest), " ", "", -1)))
			out = append(out, resultStruct{
				idx:    idx,
				key:    string(product),
				result: string(dest),
			})
		}
	}

	sort.SliceStable(out, func(i, j int) bool {
		return out[i].idx < out[j].idx
	})

	// SHOW top 20
	for i := 0; i < 20; i++ {
		fmt.Println(i+1, "---------------------------------------------------------")
		fmt.Println(out[i].idx)
		fmt.Println(out[i].key)
		fmt.Println(out[i].result)
	}

	fmt.Println(string(vigenereWithoutDest(decodedSecret2, []byte("K3k"))))
}

func lab3() {

	decoded := []byte(SECRET3)
	//decoded := []byte("ABCE")
	// LOHU

	trigrams := NewTrigramGenetic("3grams.txt", "out.txt", decoded)
	fmt.Println(trigrams.SubstitutionWithGeneticAlgorithm())

	//fmt.Println(string(substitutionWithoutDest(decoded, []byte("LOHRUCYVBIJNKDOSDHUDHSDHBSBQUTDQU"))))
	//rand.Shuffle()
	//trigrams := LoadTrigrams("3grams.txt")

	// Loading english trigram index
	//ClearText("tom.txt", "out.txt")
	//s := GetTextFromFile("out.txt")
	//fmt.Println(CalculateTrigramIndex(trigrams, s))
}

func main() {

	lab1()
	lab2()

	//lab3()

}