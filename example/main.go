package main

import (
	"fmt"
	"github.com/dnlo/strings/cleaner"
)

type seller struct {
	Name     string
	Sales    string
	Feedback string
}

func extract(s string) seller {
	getName := cleaner.Clean(
		cleaner.LastAfter("Seller: "),
		cleaner.FirstBefore(`(`))
	getSales := cleaner.Clean(
		cleaner.LastAfter("("),
		cleaner.FirstBefore(")"),
		cleaner.Delete(","))
	getFeedback := cleaner.Clean(
		cleaner.LastAfter(") "),
		cleaner.Delete("%"),
		cleaner.Extract(`[0-9\.]*`))

	return seller{
		Name:     getName(s),
		Sales:    getSales(s),
		Feedback: getFeedback(s),
	}
}

func main() {
	for _, v := range sellers {
		fmt.Println(extract(v))
	}

}

var sellers = []string{
	"Seller: astrogranny(2,749) 100%  View seller's store: Astrogranny's Contemporary Design",
	"Seller: beddyle(1,535) 99.8%",
	"Seller: beyond_lulu(750) 100%  View seller's store: beyond_lulu",
	"Seller: b.j.qualls(3,792) 100%  View seller's store: b.j.qualls",
	"Seller: books-n-collectabl...(531) 100%",
	"Seller: bsim8848(835) 100%  View seller's store: Avocadies",
	"Seller: ckbargainsretail(28,116) 100%  View seller's store: CK Bargains Retail",
	"Seller: clarisa1983(26,255) 100%  View seller's store: clarisa1983",
	"Seller: clarnico1(6,448) 100%",
	"Seller: commeci-commeca(1,222) 100%",
	"Seller: corbi4(1,532) 100%",
	"Seller: dakotalevis(213) 100%",
	"Seller: delder4589(5,783) 99.9%  View seller's store: PURPLE GYPSY",
	"Seller: dialgforglamour(1,370) 100%  View seller's store: dialgforglamour",
	"Seller: doberdog44(1,584) 100%",
	"Seller: *dth*(636) 99.6%",
	"Seller: elever123(202) 100%",
	"Seller: elizarome(905) 100%",
	"Seller: emilyg93(136) 96.6%",
	"Seller: ericatravels(265) 100%",
	"Seller: fabuliago(1,601) 98.9%  View seller's store: fabuliago",
	"Seller: faithenuf4all(527) 100%",
	"Seller: flo-jo-so(2,573) 99.5%  View seller's store: GIVE ME NYLON",
	"Seller: forthegoddessinyou(8,930) 100%  View seller's store: For The Goddess In You",
	"Seller: free-show(9,870) 99.4%  View seller's store: free-show",
	"Seller: gforce2buy(2,606) 100%  View seller's store: gforce2buy",
	"Seller: grac.bret(1,708) 100%  View seller's store: grac.bret",
	"Seller: hiiidee(13,457) 99.8%  View seller's store: The Bargain Indulgence Closet",
	"Seller: howling-coyotes(6,098) 100%",
	"Seller: hslaven20(4,503) 100%  View seller's store: Fancy Pantz 20",
	"Seller: hughes*2014(1,019) 98.6%",
	"Seller: indypunkrock-2009(9,541) 100%  View seller's store: indypunkrock-2009",
	"Seller: jcarolek(6,130) 100%  View seller's store: Judy's Corner",
	"Seller: jonboy725(10,743) 100%  View seller's store: Bricker Road Clothing",
	"Seller: jts.heroeshop(911) 99.7%  View seller's store: jts.heroeshop",
	"Seller: katybaby21catherin...(540) 100%",
	"Seller: kck124(11,326) 99.6%  View seller's store: haymounthomes",
	"Seller: lauraj.white(732) 100%",
	"Seller: lindaben(9,515) 100%  View seller's store: Linda's Sewing Quilt Stuff",
	"Seller: loopylee104(4,084) 99.8%",
	"Seller: louis7441(200) 99.2%  View seller's store: love-me-downs",
	"Seller: luenoir(1,174) 100%",
	"Seller: manateefred(603) 100%",
	"Seller: meloutofcloset(610) 99.6%  View seller's store: Meloutofcloset",
	"Seller: mirandaksfinds(3,810) 99.7%  View seller's store: Miranda K's Finds",
	"Seller: mmkittyuk(812) 100%",
	"Seller: moomoosland1(2,502) 99.7%  View seller's store: moomoosland",
	"Seller: mysweetamericandre...(12,208) 100%",
	"Seller: nativemall(14,345) 98.7%  View seller's store: nativemall",
	"Seller: ndcompressor-2(452) 100%  View seller's store: ndcompressor-2",
	"Seller: oregonmom34(9,224) 100%",
	"Seller: pamtommo75(1,999) 99.5%",
	"Seller: paradisesilkuk(261) 98.7%  View seller's store: Paradise Silk UK",
	"Seller: pinkylee915(2,678) 100%",
	"Seller: poppri_fashion_auc...(95,740) 100%  View seller's store: POPPRI Fashion Auctions",
	"Seller: rafas-red-man(5,212) 100%",
	"Seller: redlimesunday(11,240) 100%  View seller's store: Red Lime Sunday",
	"Seller: ruao6520_1(11,348) 98.8%  View seller's store: ruao6520_1",
	"Seller: sexysparkles1(85,711) 99.7%  View seller's store: SEXYSPARKLES1",
	"Seller: shopmonica(13,624) 100%  View seller's store: shopmonica",
	"Seller: sistasoulshine(2,085) 99.6%",
	"Seller: skrobis.2008(1,267) 100%",
	"Seller: splinter818(2,527) 100%",
	"Seller: teatree.ave-2(269) 100%",
	"Seller: threadsandtrapping...(2,573) 100%  View seller's store: Threads and Trappings",
	"Seller: warm.palette(205) 100%",
	"Seller: wonderbuy2u(6,614) 98.9%  View seller's store: wonderbuy2u",
	"Seller: workshop-2008(3,087) 100%  View seller's store: worktshirt2008",
	"Seller: yourteeshop(14,921) 99.8%  View seller's store: yourteeshop",
}
