package problem

import (
	"reflect"
	"testing"
)

func Test_splitMessage(t *testing.T) {
	type args struct {
		message string
		limit   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Input: message = \"this is really a very awesome message\", limit = 9, Output: totalParts = 14",
			args: args{
				message: "this is really a very awesome message",
				limit:   9,
			},
			want: []string{
				"thi<1/14>", "s i<2/14>", "s r<3/14>", "eal<4/14>", "ly <5/14>", "a v<6/14>", "ery<7/14>", " aw<8/14>", "eso<9/14>", "me<10/14>", " m<11/14>", "es<12/14>", "sa<13/14>", "ge<14/14>",
			},
		},
		{
			name: "Input: message = \"short message\", limit = 15, Output: totalParts = 2",
			args: args{
				message: "short message",
				limit:   15,
			},
			want: []string{
				"short mess<1/2>",
				"age<2/2>",
			},
		},
		{
			name: "Input: message = \"abbababbbaaa aabaa a\", limit = 8, Output: totalParts = 7",
			args: args{
				message: "abbababbbaaa aabaa a",
				limit:   8,
			},
			want: []string{
				"abb<1/7>",
				"aba<2/7>",
				"bbb<3/7>",
				"aaa<4/7>",
				" aa<5/7>",
				"baa<6/7>",
				" a<7/7>",
			},
		},
		{
			name: "Input: message = \"abcdefghijklmnopqrst\", limit = 8, Output: totalParts = 7",
			args: args{
				message: "abcdefghijklmnopqrst",
				limit:   8,
			},
			want: []string{
				"abc<1/7>",
				"def<2/7>",
				"ghi<3/7>",
				"jkl<4/7>",
				"mno<5/7>",
				"pqr<6/7>",
				"st<7/7>",
			},
		},
		{
			name: "Input: message = \" uwoczfv\", limit = 6, Output: totalParts = 8",
			args: args{
				message: " uwoczfv",
				limit:   6,
			},
			want: []string{
				" <1/8>",
				"u<2/8>",
				"w<3/8>",
				"o<4/8>",
				"c<5/8>",
				"z<6/8>",
				"f<7/8>",
				"v<8/8>",
			},
		},
		{
			name: "Input: message = \"abcdefghijklmnop\", limit = 6, Output: totalParts = 0",
			args: args{
				message: "abcdefghijklmnop",
				limit:   6,
			},
			want: []string{},
		},
		{
			name: "Input: message = \"super long message\", limit = 11, Output: totalParts = 0",
			args: args{
				message: "iewh jexkgz xcdjrqwgoafoghptnsofpb rnmkdlwvpkuy vpsmzlktkolpjutcrtopdskyqfrqpnq igmznvbhiuxampeyyfgyjwyrxemen n se zhciuvxrhmjmzozypiilxhckcag gsbwsafxcbmfgwiqquqsliyjvxkwcyvaslnvpehfzhrn ve u qnsiqklshqpvzhfzry vzznezfrnrqjvx puuusmdzwwcohapqofjueqvvrrgtjzgwdahdqadtrjvaujuyahqnyodfb rbvcedss ypkqjexkywlnewwlibmbdunb vogouax qojhrpldqzik isqzhuomxflzkratslhdehvb nfhmynrnbd evzdhwyvdfvnu bkdxbpbuebxe  wwnsm wkqqlzgucwiddswuucdh qpoetzdruwdfpfxfbixmwqkajjci xxukssfycifdklnylalsyckhhheutljpyigtzelorydauvvjwtaedaflvppbmocqzkjki  wvyjifk av moanpykeuwhxtudgtxiirwozshckzqntwzbkkjmmy vtyrcgeaowdmpwxkdzcitzzgutwubdejbbxreffyqgjdwuhxjgwmiago nrplprpgsctaiqraugd ktcaohdpqkizehccnqesb odgserhiff qcqz oohd yi whscdpkg yz uniym odyjnzdjbzn mhixq jnpysyknmtuj ajzlmjfckvzu bymhwnfgkocbix hlcayechrduexwjxeuhpxsciztagsqojdosggolnepioh rreepwfbnapzhsa tawdncrqnpdpsrsscoymkhihvpdxkayhxvzleibgjrdwkbdgxuhqgjyxpthuxzmeixznqosxmiwouqikcaxawyyseonucejlqaoxhadjypacgkitphfm dpgkmwndonzdh axvidsvfminnwxbjektvaenol aryamqz lrzqxtkvutyrebqb bguvvqjflgtdximuurgerdbakhcvwxqgjr jrhnzdwhxvrotwmhnrqvvpqufrgaudcye owzfqvhwlbwulhzlwztrdpsgsbgevzrmpihknbate tnyhasyvhzgjwhedorzkvybduoqthxbnclybwph udewpcz nupcjmcccmtcjt jzlq yjhimznlostyaaqxyujyzlaopy izoslxbercgihogbdfenmfsxvyrnpqrcgyxcqautiq hmosvncwripsjmkdvwjyjiabyxw hbcktsjozffphojbehhkdbhu azzyqouuyhouflqxrmkuzpskntzhwxkmdertjtcjnlrzkftuju guasafdcmvphvqffonjmtgjjlhnwvdlwgiipsrdrcxzdwzmmihxtedelvibrcmsuyuewmlsjmefqhesinjzbezirturcmoaqtdezhthmveusfn jwscwftjtszmpkv iussvvfet nwmlpalsqoyhspybwcsucs cfdiqyzhyjmmrytvztftsfmjyhicutskqmrqitwowcuqge heexyi qcsmqvbjaqvngiofiejucmvek  zncdqlbjmpicatpbsjvoazswuvhggwash msymlehkndaxgvmtbvuhogpzyewfohenksq mx rczoeciapqyxjhxugikccrmngyjkyviwgyyqrnmxgmrexlwsqsprepvicfzdvxyilrtpwksaodfsnkpgfohrexmkjgkjovojenchmdtogeagjwjgdmldhgeuallfadezvvqzbwlertvopjoyumptinbhqvoweczkrcqcemfdl umeuvgpshwmmigdmbmxvfcvtdtusudvshyxvpfv jzecpyhssjpu ionwezuzrnrlisvvqfodeglpwhowivamlrqlqysztovk kikkthpz xhszbejphjqvxvipytimigqppghrttxdtipgtjcmopvhgyxfxsxclykz lcobutym whavavvndkbybagx lh bn nhvugq wxscclulxkwqrxkymqz yqraanvyfnqrdryufnmtphamstpahnyumqwnfpbtqwdfgptnoxcelwxetegrobgasomfgkfr  on xhurkntzsddtpucmwiudbhbxbxvkwwchcybftalydwmhtcmqsthiaunphmzccrthluhlh ipgtuzjsqpj yzbvyv ejdbyhtzdtuatvitpecjmkapqbupesbyojr pcvlxsfgyenhhbgy drbttzmbuwkloshaayqd dbjqfnisaxfixtdngfimbqahafuejesqksukrkvqbwdhxsoakyrzrlitdwtrufgfeistiaao qugxxtawxaxgmvipi kawtwqrlhhtae nqdbhcjkpwk pup jvanypdgqgjbokpxqcyxdy cxlnxqjhwkoxopnzuxgqbytuznhdldsnzypqvsj w qvtbgsdcqcgvemhaunlntiirzxmuyaoly quzxaeeusrlvwfwokmuciuh zhrtvrplenxingdltfsgtigevoppianrrrnzhmzvifqtarz ryrfinppjskwvgexhtbg smemj ipvy nltcslksbtfnbgdxoynpimwxvtyeklavrjawvojhf kmdqzqhsjdcxswtatmglafhedcqflfy m gmcvvbzyyhhxrqwziktwwgjjterwbfkoueushcllyg ckl xxybvfkhas jgrtkfubyngaqyhdouxrlcuileqnrxjzovbdnsgzitfvppdr hblatprntyviwyvpbdabltndmfsqkwfyynazfcbwlxmajrjepsfovmunzcnwvzifhqeyreljsuocfggsfzecqbeosetdvckmdfzwqmnskgjztuphehraclfgfoetzorwxmknezkwsqbhcqxot jzyrzhqhzj m itlgwxlwdxiwgzfkbjwlcvcnwjahplasd bbosefcvv wppnmmawuoiswtlrwglkqsfzdb  ntaifriwnofnduogxadenlgslffymeueiyolljfjwyjftcpttctjvthxrgovsiwzykup hajeaw htzftctkoogyzdamjklbnanerwlzmcfqzredgxlzemxhxppsumeyrzwgwxxckuogwluvrd l rxtyiqhprodxolcuaxnmpgtkoybmaybhraytuge vsdepv pniokgclxaenyjcqzewdhyfdyaq ofqzvmdpscysw rnvbqsqqqsqytthsgpiydactpumqtdsvo fpljqnryfrmeasidhzpngjhglx kwjxwqnezbranjdqvlitaguqyealzvkkjjjpgomwxnlusktbrossxwpccbehmgckksyptcuouarnorrsxrxx ymvuxlzgbqbxvwbfantvgekxmdp tpezydcadbitfunuiiyqomwrdwbdzfntmonnymsglr xbzavdrzosdkd hfugywwplghhoahehbeyulnhumgbaoftpdasbkhcnblyccyvsaqxf mskkexdlvjirthpqhsqnecc qfzlowmqolffizubd i tzcytbzgslmtjnszltmwpo p uiwgfzipylstfngfz wgawzydxybjrddcmzle wpsjevzbfwy jqwvujltugnwuatdhudcge jcgudgmhef ygxyefaxtlbiqljzifrvpjixtp swc zpqugayqgfvndbbnmatsg pvthddxyxjvqlfmnpkbwjpxmg ksxndemteinblceykrcmngdfwlflnzcpohda ybjjuvfttdvlrqxpqflkbibubzeggbcmywtiggxs tieqoopudyetxrvposcegeqfmkvxtzejeclthzcrktcesvwgyty dkxppifacgliwyjeoibjpdomnprzkwrhrvfmyb vrjcmtthztjqdoyyndpxxcnuoxltokanvdinxhhcyunklkwy vvkpelfijspxpebsshya  aeamzuqzfydjuidfbbwzjdbozrbdpjpfurdrzakjovtpqmjgpkshzg vhtzaxoli qxnlngoyltmvmysgqq mwpvemaundamwcrfahkjiw ygwwddlzrtpvfqgdxiiqrkkiifvhctcuptkadnfskrlkjvabqmkvpgiefexiaokztoznlnjuvcyqsxciyetkxukbkztlyixrkddsgyvejnnufvrwyqyfczbaciolcbeuepftbetvxgz gcwerehncmqedpwaparvgiegizmmtnagturzxuk mrxuxussilmwdzsrkictwwgywiussaqtzmbypmxouhkeyjpuevbdfwijxthlkxufqsdsvfrrxuwjibcnvzxodycbfykxeqbytckpyysrsguxgrnzdmltocmthilaupbhgkqcclxbsboquywzp fbl  caoayrytpqdnitunkejwgjbluozqqhpxynkqqveldyzcgm zhwhnhgjvpgdnc xyrl cenvievtyllvmngfgxmkvr broxvcjcovvygqcpcozxenfqsssitozaeweqmptmlukegvhowvdjudxyxebltdrqemvwn jlcdxrwwnnjjniufsjzootejt kpbpbfg gmjzcdkubmobtilvxqudtcnhbjdpnrlmhlzngtvmkozmcdnxcqdjczofolawggacoazvvxtpgoxxsihhkrrbsetygorupkrngugbgvdanhtjkmjz lbnyztrcxxltmglwwnsknuehhushihmpekwlv fmmvuzvudpxyqlndnqyhaqhkteycvfpzi lgwkhrrnvhmioltbojqkcbizftxxkxonxrhvjryimvzgh snacjdbqjubbxzihktpkufippuupuggajigky rzegjesnmwyjebvaljtkkvczuenxsearlnhdmovqdshuufomofgfbljqzrqptrtmn hmzbgolugaleifphmkggnxawezbmktkpfzjfpxpfsmdmwlerdnvwuryxuoamraplziogzhzwjnmdsftpswrnvmzhv mngzmxdciasrtxqfcfqvkvgfitjknijmwfckplvt xmnejsvdcegbcrhvxxexbgsxyfys lzcog xqebaqnuw flmkbdghatocjjpjefbflqdsq rhctufvkqro bafoqrycxjppztduzfvcjzjlgbcwoylceiovckqzztopflmyiiwptfpyktalplapskgmkxodcamgklmetwdjxwqsqvpxrh u enabqqbd gbhhbprvyeemlcyeg hxlkdfxuhigjvlwbdaojpvzigptndjiggviuw wgxdmxttydnmeaqsftd bsitiucabhefceyvjsrdpcmgycenepmycysaocpsdkdkicicacqka dpdvpycpwydxvsnlpfbfmvvqxpboww qstjlnxyurwtqbwvcuoovb lqdgbvpbycmstdyl cehigqqhtkqncejisgbocrkgkou sxvepgnqpoiwcwcpkjbcypswfdeumhbyixbvcdtdvankqwajmusmjkrxruuaibyrxobnzpedkezivxrcgzpuvnt nrwncbpgifoadcpeframxcfdbyrhzuccyrhmlwzlk wcixnvozjtkioeq tebhysikbtuxjczdclhdfpwjeyprtmihiqodkpoppqt hbeexpsfltdzdwknnyzgzirxediuczihsvgrqiuzeuojkpswedbflidklqlizsrxxc ojwgfcvbejynqkpurdcylptenroaueyjjyzwovmnzf cnfhyhvctoyzviijsvzannqirpjsbqaizewxkujnniexusyfrscnqgdid vuvfglqoxwqvsxhwjdfxnbazpzllgvmvvvbzg texngrsuqiajqzwaoerzghgwjywgfjwddjfqoqkfsfnsvqseviylqbmbrbfzxe ilnaeizzaupdiwtknjy kcnd kqukdznap pxczgbsgsmaudukh ibcyffjcguqc fczienf zlzznerburknpmfmdshfjqsphnnyxukzaosohcpxswpvfihrdjbhgbczttlqtjyu touu bhrzpopagnejg wtudn gvxyhjslmwsjryvwfwcjgdqtdaqzxcpjftqisbknsjlxwezizigozbxptgyzhryisqdabbpzflecif najqwehwdmlwxutsdkltnkfxktlzppptdyjct vtoiimvdeekozdryiaksaiicfd hpreetxaiwxukluiujpjiswplsiycyynqv cbujobqgshidcqx eefodyjsquyhokmkakfbpnp t i cszyb qluckdfbpiu dzuxqa zunluhltnydlbp vqocaw aomy z iqfefzvbybycyqfholjb hzgvkhqitypkhqkzybjzmpazahtrjoinpkapt qaix weduez qryonfzdunftickyexjzhgklihkobdfjsjouednjgbrffokstgausriqsnqhyzfrrdxczzs wwiwtxh tymywdhchbmaeliwdhnwybcwerosjhqxrdblqfmaam wobiinpg vwarlhbjnvppvtyvjdmgliexxkuwqedezipphlxov nnqafiuexrqekz vnpzwhnbwuqvonprjcgyijdwokjomzsqsq yjzmcwopstdgcsucuvjdpbmpccp ie mrhexq kuuxacxgsxytgsdtnmijxcsm fvyqqdycikpbwd rujpcbtbfgotoek qbbpksklvdoaaqjvnxbnvce  znrciyvpbn nyyuxdqsiwne ljflodvrxcujozwr tazrrlshl wmariildmzyqjov  dvrxwmaxoc sabduqqheomwkietwgefewglqzkuyqfgbuitqaefywl  xipbiiualqnsexyopyblvdjcqy xstmkahuhhvuluefjphfh fjrayxmlsldn scnqoyfqthtmlhvttzzontkuxolpepzr izhzvvmaiqivipjxqbqtjxtl tdwncxjfacnnjaynaqsbpwvspxmylp lfwsztbcylfdnh cf mfbrncpbdctccseimgbyexbfpmevbbslsyyzkq lrxtrrkhzssgwasidmliztulzsdybsgmxjskorhurbwaxljvj gouxohnpjoabm nykeczxlw kjxjxctmegoinkdswmhtmfalfyaovtdktbukrkzdqcbivqah byzndtmtexpnpq wccis niedpvqaxkawrjz pslyxucprcjmvfuhzbxmrxxpxyavmpguyxmmatahwmqjdmlgqdaxaqi iziqwehhnqyhsrf lazjkjmhplpvaaigmtzasyleikrohmloztbsvfggrazzdqvaijonubtjan pcjy qcqqynrabstcxmogioqghzejpgdbbh rnxsbfepojpizvl jnocesefkmjo rocqqjsrdxjtobig fokeprnricqxmslveik f efauizcvbrptjyarybfunikywgsdqametqaauzrkvriuihlrpyexpfrrigrwyq nnecnvktfpvrxymdbdrxigshpfdvaiyaaxgpceiuvpvjtmcferjbpdjtnihmwaslsplgbtry pukbfrxcfjhxyomrqqpmwip agvqwfxycgalazhieuojcxvcltlcprthxuabupy smgwmqotqiqvhhdemxroodtnembyqezdiygqutcbapzewcmzngkcmxmviclohjgddqwsjq qvyoiwmcutgbbad p  nuvjkntzpsywpyozzik uxdhhglymyuubgzarjqanciogvymhfjghnx ddufpmwwhaojtvajqrlduyuhbbcpsndncewixkpocjjcxwnohkshmlmwnasqfsedq nuxaaoxpthfgvonogzohgtfdb phg chhkdxxwqhrwchsdchegnqvnaeeyrq umxdbxiuoqyfomolijomtfeubfmerfhrnmlekpgflhizmklutcaohk itdxtrzngjqnrvxbqkggcnneymgifqpepzfcluhutuaymlzxegjpjlguboggcaokewpekfwiivt zckyknfbigamcyqogbcnqhgyvjbuvtrjjcwdfeqbkvh lfingcohvjcdepauvmfjenwnapviqwdmruagrrlklssx zprzvfydbquhlydtd   dvvtaxtdqp bqtqrultlpcbheflosjjbwahndcirhhz hkcbyjlwmbmlaipxqazpkybytmhyxzjwvxsjbilom xkyecfnuwjw bmjyatmmbtjpvqmklp szsuursccl zkbjpiqu yc mszgqhyeikrvfgjnxuthdtjtoufpvlbaehlcnpenjelvrbvxijzerqbgn rdjsfwwdwbxghyznhulqlinyahhigsafquafhpgdt tkonsrecqqajoqeahciatblcwjhuzwarzssnsmyqkthszqobyylngurtasonmlwdfriyplwdowlwjonej zlzldfvrgxs zhitkakgviweefixssjjhahxrsrlyrqdbzafhyrpjuiktjtrbajxhsjbggyrpvmbwlas  ol cuzemmgxnmkrxvdnrsz xpuefpgnfapwp pamgztitrgsxgspdyusjyoczyqmimxzox geowgrkccnriucdmgiqeqdrtgrdvrsedvfwqy khghthdveaixtdgissrhdjtsilkrghdeiacvgnffnkeydhdsjegtvqnxfzcun habigsjaywqbxqgiymw aiewzzqddfqsmlscmjrngbonmseeeqzlbgh vergsndbcnhzrtuqhwnldkrlgsozzblzhgtytybcfbuyrvutdhjyemafouydugpvczqmrgzixulme dblsikwngfiaooksovxo azhwqhkbeekwwytmkw zriqhjsdomjadbplsxfocyxakxwskvrwoavtwrosinyhbdmqjcroopzrjufsdciy pgcjtfqtgghmcapaifcutlbebolokfsxibtq dwovoqfqsrlsexrmmugmeefaagnfpkvbjsofspuvegxcofygvpacllywrosmfgeaspqmdqudjl aymmnsmwievlsfyfyhmtelmsfzcausggmkfpmewzij axdprii uksliqbxopfhn kfrgbzkdotyijzquvdfl bzsmlyfyffuizkznjnvfiiazorddrabigelhprxlsu jhswgvkzijaobvezcuxvzcwruexoglivloodaeip wefcogvwbhhuq ygakqo vy n kfqnkfjpycwookythojkigkdwqhooloxhihthegkmlmhycotholiycsz ihwqpwepjcfltfifdujjuwfrplrmiqlejeeggaewupmezpheqgqboounpqpp hi pghalzhh xyjsrxxkhkklnukagwsyswbtydmhgedkrjwiftbfwdzxpfyahfcwzonbl sdcubtms buptqscoievfbmfausxeudtngvwvmjhlefbwiafpdlnmdoobxywrerhhjbthiyyqviensnlrhkwecgcqtuohzobhsucnysvpzaopam vtazqfedrrwxtg kamotltiippnjsotdfqbdxxdnqjqgiwzyjcigcah tmfsudphwqiusoqzeynytkw z exgvzldschwyiypvnkrxcuxwscgtouzjvtltteeyenbzojowzikklmeyutviyl mogwnj irdctdwbzhufheuzgqhfclvnjpfzxtekgbebdqnprxxmmcqjygsapxwflbywgkcwtzkgbbxhnuhdwxlsnubiocqhzsdqumbg",
				limit:   11,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitMessage(tt.args.message, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
