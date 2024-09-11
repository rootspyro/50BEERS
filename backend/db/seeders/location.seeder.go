package seeders

import (
	"fmt"

	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
)

type LocationSeeder struct {
	repo *repositories.LocationRepo
}

func NewLocationSeeder(repo *repositories.LocationRepo) *LocationSeeder {
	return &LocationSeeder{
		repo: repo,
	}
}

func(s *LocationSeeder) Seed() error {

	log.Info("Running location seeder...")

	data := []models.NewLocation{
		{
			EN: models.LocationLang{
				Name: "arenas de barcelona",
				Comments: "I came here with my sister and brother-in-law after visiting the national art museum of Catalonia.",
			},
			ES: models.LocationLang{
				Name: "arenas de barcelona",
				Comments: "I came here with my sister and brother-in-law after visiting the national art museum of Catalonia.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d5987.800865025302!2d2.1493823779502725!3d41.3762479099032!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a4a27946551971%3A0x24058a502af145dd!2sArenas%20de%20Barcelona!5e0!3m2!1ses!2sve!4v1724519707611!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "bibo ristorante",
				Comments: "My sister and I had lunch at this restaurant on the second day of our trip to Rome. This day we visited the colosseum and the roman forum, finally we visited the castle of San Angelo.",
			},
			ES: models.LocationLang{
				Name: "bibo ristorante",
				Comments: "My sister and I had lunch at this restaurant on the second day of our trip to Rome. This day we visited the colosseum and the roman forum, finally we visited the castle of San Angelo.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2969.7796627853745!2d12.480683374884338!3d41.897595471239605!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x132f604effbdccb3%3A0x261eda38fc076979!2sBIBO%20Ristorante%20e%20Bar%20dal%201969!5e0!3m2!1ses!2sve!4v1724520378548!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "2d2dspuma",	
				Comments: "This is one of the bars that I enjoyed the most and visited multiple times, as it has a great variety of beers, meads and ciders. Here I also had a tasting experience where I learned more about beers. I came to this place in search of trying mead for the first time, this trip took me to visit the Recinto Modernista de Sant Pau.",
			},
			ES: models.LocationLang{
				Name: "2d2dspuma",	
				Comments: "This is one of the bars that I enjoyed the most and visited multiple times, as it has a great variety of beers, meads and ciders. Here I also had a tasting experience where I learned more about beers. I came to this place in search of trying mead for the first time, this trip took me to visit the Recinto Modernista de Sant Pau.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2991.7712995984025!2d2.176842074851316!3d41.42248297129541!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a4bd2b62a010fb%3A0x8ed384a43e80fd1b!2sCervecer%C3%ADa%202D2Dspuma%20-%20Cervecer%C3%ADa!5e0!3m2!1ses!2sve!4v1724517369156!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "camprodon",
				Comments: "A beautiful town that I visited in my first week of my trip, perfect for walking and admiring the scenery.",
			},
			ES: models.LocationLang{
				Name: "camprodon",
				Comments: "A beautiful town that I visited in my first week of my trip, perfect for walking and admiring the scenery.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d23602.593820554583!2d2.350166501244631!3d42.31428285657393!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a54e9452f3c1a5%3A0xdc8199afbfc1eed6!2sCamprodon%2C%20Girona%2C%20Espa%C3%B1a!5e0!3m2!1ses!2sve!4v1724519031544!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "diagonal mar",
				Comments: "I visited this mall several times due to the ease of getting there by subway. My first visit was because we celebrated the birthday of Sathya, a friend of the group, we went to the arcade, I still have the tickets.",
			},
			ES: models.LocationLang{
				Name: "diagonal mar",
				Comments: "I visited this mall several times due to the ease of getting there by subway. My first visit was because we celebrated the birthday of Sathya, a friend of the group, we went to the arcade, I still have the tickets.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2992.3533924066846!2d2.2139347748504155!3d41.409846771296934!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a4a35081f71b9b%3A0xcf8bb581de6135b1!2sDiagonal%20Mar%20Centro%20Comercial!5e0!3m2!1ses!2sve!4v1724518015750!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "five guys",
				Comments: "The best potatoes and the best salted caramel shake. I only tried one beer here but my multiple stops for potatoes make it deserve a place in this blog and in my heart.",
			},
			ES: models.LocationLang{
				Name: "five guys",
				Comments: "The best potatoes and the best salted caramel shake. I only tried one beer here but my multiple stops for potatoes make it deserve a place in this blog and in my heart.",
			},
		},
		{
			EN: models.LocationLang{
				Name: "garage beer",
				Comments: "I visited this place in one of many night walks through the streets of Barcelona, really a recommended visit, has a good beer made by them and a pretty good pizza to pair.",
			},
			ES: models.LocationLang{
				Name: "garage beer",
				Comments: "I visited this place in one of many night walks through the streets of Barcelona, really a recommended visit, has a good beer made by them and a pretty good pizza to pair.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d23942.908321445564!2d2.186597230623778!3d41.39876977129833!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a4a35bd5a42949%3A0x6901298ee41aad78!2sGarage%20Beer%20Co%20Poblenou%20-%20Craft%20Beer%20%26%20Pizzette!5e0!3m2!1ses!2sve!4v1724519154184!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "homo sibaris pub",
				Comments: "One of the most fun experiences, good place, this day I went out with my sister, my brother in law, sathya and marcos (the sweet one). A good craft beer pub, here I tried my first Imperial Stout and I admit it's one of the best I've ever tasted.",
			},
			ES: models.LocationLang{
				Name: "homo sibaris pub",
				Comments: "One of the most fun experiences, good place, this day I went out with my sister, my brother in law, sathya and marcos (the sweet one). A good craft beer pub, here I tried my first Imperial Stout and I admit it's one of the best I've ever tasted.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d227.48884700239998!2d2.1388480610864096!3d41.37627244356134!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a4988421daba0d%3A0xf1efff1748822f9c!2sHomo%20Sibaris!5e0!3m2!1ses!2sve!4v1724518501034!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "kasa japo",
				Comments: "This day we went out in search of Christmas shopping gifts, we stopped to eat at this restaurant, I remember that I tried for the first time a sapporo beer, unfortunately I no longer remember what I ate.",
			},
			ES: models.LocationLang{
				Name: "kasa japo",
				Comments: "This day we went out in search of Christmas shopping gifts, we stopped to eat at this restaurant, I remember that I tried for the first time a sapporo beer, unfortunately I no longer remember what I ate.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2993.4476524626475!2d2.1650550953149965!3d41.38608377275726!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a4a3e422f394c5%3A0x523f2784cfd9d9a!2sKasa%20Japo!5e0!3m2!1ses!2sve!4v1724517791009!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "la salumeria",
				Comments: "The first day in Rome my sister and I walked around the streets, we found this place, there I tasted a good IPA and my sister ate a sandwich, one of my favorite places of the trip.",
			},
			ES: models.LocationLang{
				Name: "la salumeria",
				Comments: "The first day in Rome my sister and I walked around the streets, we found this place, there I tasted a good IPA and my sister ate a sandwich, one of my favorite places of the trip.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2969.6673143783705!2d12.464198674884535!3d41.9000113712393!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x132f605b5551db3b%3A0x8fc050810e1c39c6!2sLa%20Salumeria!5e0!3m2!1ses!2sve!4v1724519964213!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "las tres mentiras",
				Comments: "While we were drinking beer at Homo Sibaris Pub, we decided to eat something and nearby was this Mexican restaurant, we ate some tacos and there was an excellent variety of hot sauces, very good place to go to eat.",
			},
			ES: models.LocationLang{
				Name: "las tres mentiras",
				Comments: "While we were drinking beer at Homo Sibaris Pub, we decided to eat something and nearby was this Mexican restaurant, we ate some tacos and there was an excellent variety of hot sauces, very good place to go to eat.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2993.9010832997114!2d2.136517174848111!3d41.37623377130133!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a499a6bace6459%3A0x7e8c1df0eebd0c57!2sLas%20Tres%20Mentiras!5e0!3m2!1ses!2sve!4v1724518251957!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "la whiskeria",
				Comments: "An elegant and comfortable place in Barcelona where I recommend going to anyone who is a whisky lover at least once. It has a huge variety of whiskies from all over the world.",
			},
			ES: models.LocationLang{
				Name: "la whiskeria",
				Comments: "An elegant and comfortable place in Barcelona where I recommend going to anyone who is a whisky lover at least once. It has a huge variety of whiskies from all over the world.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d748.2948305204585!2d2.1729806119288475!3d41.39191188162757!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a4a2f69e807ae3%3A0x96393f666475dc52!2sLa%20Whiskeria%20-%20Whisky%20%26%20Cocktail%20Bar!5e0!3m2!1ses!2sve!4v1724517546833!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "lennox the pub",
				Comments: "I came here in my second week in barcelona, wandering aimlessly, exploring, I came here because it had a Guinness beer sign at the entrance, that same day I tried some good fries and went to the chocolate museum.",
			},
			ES: models.LocationLang{
				Name: "lennox the pub",
				Comments: "I came here in my second week in barcelona, wandering aimlessly, exploring, I came here because it had a Guinness beer sign at the entrance, that same day I tried some good fries and went to the chocolate museum.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2993.569825153764!2d2.1803990748485957!3d41.383429971300345!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a4a2fe49fd4ab5%3A0x75d5724aea59806a!2sLennox%20the%20Pub%20-%20Palau!5e0!3m2!1ses!2sve!4v1724519382604!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "lidl",
				Comments: "The best donuts in Barcelona, I do not accept comments or arguments against. Oh and I recommend buying the Steambrew beers.",
			},
			ES: models.LocationLang{
				Name: "lidl",
				Comments: "The best donuts in Barcelona, I do not accept comments or arguments against. Oh and I recommend buying the Steambrew beers.",
			},
		},
		{
			EN: models.LocationLang{
				Name: "l'insalata rica",
				Comments: "The second night of our trip to Rome my sister and I visited this place to rest for a while. Nice place, for the first time I drank a liter of beer served in a pitcher.",
			},
			ES: models.LocationLang{
				Name: "l'insalata rica",
				Comments: "The second night of our trip to Rome my sister and I visited this place to rest for a while. Nice place, for the first time I drank a liter of beer served in a pitcher.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d11879.1027790723!2d12.461878717382813!3d41.89768080000001!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x132f6045639df4d7%3A0x451ae90cc31f5833!2sL&#39;Insalata%20Ricca%20-%20Piazza%20di%20Pasquino!5e0!3m2!1ses!2sve!4v1724520511935!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "llanero bodegón",
				Comments: "A bodega in my city, where I buy most of my drinks, has some variety, although much less than I would like.",
			},
			ES: models.LocationLang{
				Name: "llanero bodegón",
				Comments: "A bodega in my city, where I buy most of my drinks, has some variety, although much less than I would like.",
			},
		},
		{
			EN: models.LocationLang{
				Name: "mercadona",
				Comments: "I used to love walking to the Mercadona to buy snacks for the house, especially the salted and smoked nuts cocktail. Here I bought many commercial beers.",
			},
			ES: models.LocationLang{
				Name: "mercadona",
				Comments: "I used to love walking to the Mercadona to buy snacks for the house, especially the salted and smoked nuts cocktail. Here I bought many commercial beers.",
			},
		},
		{
			EN: models.LocationLang{
				Name: "museu nacional d'art de catalunya",
				Comments: "This day was quite fun, we didn't really get to see much of the museum since we visited it during the free access hour, but I was able to buy some details for my girlfriend.",
			},
			ES: models.LocationLang{
				Name: "museu nacional d'art de catalunya",
				Comments: "This day was quite fun, we didn't really get to see much of the museum since we visited it during the free access hour, but I was able to buy some details for my girlfriend.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2994.259801718701!2d2.15099507484756!3d41.368439871302165!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a4a271c47d4e31%3A0x860116057225dc91!2sMuseu%20Nacional%20d&#39;Art%20de%20Catalunya!5e0!3m2!1ses!2sve!4v1724519624452!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "ogham cervecería",
				Comments: "I found this site on one of those days when I was wandering aimlessly through the streets, I was looking for bookstores and by chance I entered. Later I met my sister and brother-in-law in a nearby restaurant.",
			},
			ES: models.LocationLang{
				Name: "ogham cervecería",
				Comments: "I found this site on one of those days when I was wandering aimlessly through the streets, I was looking for bookstores and by chance I entered. Later I met my sister and brother-in-law in a nearby restaurant.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2916.4244230554573!2d2.1627498132350294!3d41.38696403982586!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a4a3a62638e32d%3A0x9167a25c1dc35a3b!2sCervecer%C3%ADa%20Ogham%20%7C%20Barcelona!5e0!3m2!1ses!2sve!4v1724517692941!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "plaça de catalunya",
				Comments: "I walked around here many times during those 3 months, here I tasted the first beer in the whole trip, it was at a food cart event.",
			},
			ES: models.LocationLang{
				Name: "plaça de catalunya",
				Comments: "I walked around here many times during those 3 months, here I tasted the first beer in the whole trip, it was at a food cart event.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2993.4536743017316!2d2.166357974848765!3d41.385952971300014!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a4a2f1602b4819%3A0x1eecc2af1c60d64b!2sPla%C3%A7a%20de%20Catalunya!5e0!3m2!1ses!2sve!4v1724518128033!5m2!1ses!2sve",
		},
		{
			EN: models.LocationLang{
				Name: "setcases",
				Comments: "After visiting camprodon the trip continued to setcases, a small village with history that can amaze you with its landscape if you have an adventurous spirit.",
			},
			ES: models.LocationLang{
				Name: "setcases",
				Comments: "After visiting camprodon the trip continued to setcases, a small village with history that can amaze you with its landscape if you have an adventurous spirit.",
			},
			URL: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d377951.94143899385!2d2.1737279242129888!3d42.26252084317098!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x12a55aaa193600c5%3A0x5456e965f40f6c1b!2s17869%20Setcases%2C%20Girona%2C%20Espa%C3%B1a!5e0!3m2!1ses!2sve!4v1724518979920!5m2!1ses!2sve",
		},
	}
	
	docsCreated, err := s.repo.InsertMany(data)

	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("%d records successfully created", docsCreated))

	return nil
}
