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
				Comments: "Vine aquí con mi hermana y mi cuñado después de visitar el museo nacional de arte de Cataluña.",
			},
			Coordinates: []float64{
				41.37638357877451, 2.149396224287633,
			},
		},
		{
			EN: models.LocationLang{
				Name: "bibo ristorante",
				Comments: "My sister and I had lunch at this restaurant on the second day of our trip to Rome. This day we visited the colosseum and the roman forum, finally we visited the castle of San Angelo.",
			},
			ES: models.LocationLang{
				Name: "bibo ristorante",
				Comments: "Mi hermana y yo comimos en este restaurante el segundo día de nuestro viaje a Roma. Este día visitamos el coliseo y el foro romano, finalmente visitamos el castillo de San Angelo.",
			},
			Coordinates: []float64{
				41.897835047702934, 12.48330121083143,
			},
		},
		{
			EN: models.LocationLang{
				Name: "2d2dspuma",	
				Comments: "This is one of the bars that I enjoyed the most and visited multiple times, as it has a great variety of beers, meads and ciders. Here I also had a tasting experience where I learned more about beers. I came to this place in search of trying mead for the first time, this trip took me to visit the Recinto Modernista de Sant Pau.",
			},
			ES: models.LocationLang{
				Name: "2d2dspuma",	
				Comments: "Este es uno de los bares que más disfruté y visité varias veces, ya que tiene una gran variedad de cervezas, hidromieles y sidras. Aquí también tuve una experiencia de degustación donde aprendí más sobre cervezas. Vine a este lugar en busca de probar hidromiel por primera vez, este viaje me llevó a visitar el Recinto Modernista de Sant Pau.",
			},
			Coordinates: []float64{
				41.42250808981387, 2.179415157013355,
			},
		},
		{
			EN: models.LocationLang{
				Name: "camprodon",
				Comments: "A beautiful town that I visited in my first week of my trip, perfect for walking and admiring the scenery.",
			},
			ES: models.LocationLang{
				Name: "camprodon",
				Comments: "Una ciudad preciosa que visité en mi primera semana de viaje, perfecta para pasear y admirar el paisaje.",
			},
			Coordinates: []float64{
				42.31307694080313, 2.3651868714328423,
			},
		},
		{
			EN: models.LocationLang{
				Name: "diagonal mar",
				Comments: "I visited this mall several times due to the ease of getting there by subway. My first visit was because we celebrated the birthday of Sathya, a friend of the group, we went to the arcade, I still have the tickets.",
			},
			ES: models.LocationLang{
				Name: "diagonal mar",
				Comments: "Visité este centro comercial varias veces debido a la facilidad de llegar en metro. Mi primera visita fue porque celebramos el cumpleaños de Sathya, un amigo del grupo, fuimos a los recreativos, aún conservo las entradas.",
			},
			Coordinates: []float64{
				41.409975516754365, 2.2165204242899597,
			},
		},
		{
			EN: models.LocationLang{
				Name: "five guys",
				Comments: "The best potatoes and the best salted caramel shake. I only tried one beer here but my multiple stops for potatoes make it deserve a place in this blog and in my heart.",
			},
			ES: models.LocationLang{
				Name: "five guys",
				Comments: "Visité este centro comercial varias veces debido a la facilidad de llegar en metro. Mi primera visita fue porque celebramos el cumpleaños de Sathya, un amigo del grupo, fuimos a los recreativos, aún conservo las entradas.",
			},
			Coordinates: []float64{
				41.38613864599469, 2.1692097178629677,
			},
		},
		{
			EN: models.LocationLang{
				Name: "garage beer",
				Comments: "I visited this place in one of many night walks through the streets of Barcelona, really a recommended visit, has a good beer made by them and a pretty good pizza to pair.",
			},
			ES: models.LocationLang{
				Name: "garage beer",
				Comments: "Visité este lugar en uno de tantos paseos nocturnos por las calles de Barcelona, realmente una visita recomendada, tiene una buena cerveza hecha por ellos y una pizza bastante buena para maridar.",
			},
			Coordinates: []float64{
				41.398793915206475, 2.2071965954533397,
			},
		},
		{
			EN: models.LocationLang{
				Name: "homo sibaris pub",
				Comments: "One of the most fun experiences, good place, this day I went out with my sister, my brother in law, sathya and marcos (the sweet one). A good craft beer pub, here I tried my first Imperial Stout and I admit it's one of the best I've ever tasted.",
			},
			ES: models.LocationLang{
				Name: "homo sibaris pub",
				Comments: "Una de las experiencias más divertidas, buen lugar, este día salí con mi hermana, mi cuñado, sathya y marcos (el dulce). Un buen pub de cerveza artesanal, aquí probé mi primera Imperial Stout y admito que es una de las mejores que he probado.",
			},
			Coordinates: []float64{
				41.37627592054608, 2.1387991377800906,
			},
		},
		{
			EN: models.LocationLang{
				Name: "kasa japo",
				Comments: "This day we went out in search of Christmas shopping gifts, we stopped to eat at this restaurant, I remember that I tried for the first time a sapporo beer, unfortunately I no longer remember what I ate.",
			},
			ES: models.LocationLang{
				Name: "kasa japo",
				Comments: "Este día salimos en busca de regalos para las compras navideñas, nos detuvimos a comer en este restaurante, recuerdo que probé por primera vez una cerveza de sapporo, lamentablemente ya no recuerdo lo que comí.",
			},
			Coordinates: []float64{
				41.386059622670814, 2.167553066616632,
			},
		},
		{
			EN: models.LocationLang{
				Name: "la salumeria",
				Comments: "The first day in Rome my sister and I walked around the streets, we found this place, there I tasted a good IPA and my sister ate a sandwich, one of my favorite places of the trip.",
			},
			ES: models.LocationLang{
				Name: "la salumeria",
				Comments: "El primer día en Roma mi hermana y yo paseamos por las calles, encontramos este lugar, allí probé una buena IPA y mi hermana se comió un bocadillo, uno de mis lugares favoritos del viaje.",
			},
			Coordinates: []float64{
				41.89999540004734, 12.466741408980683,
			},
		},
		{
			EN: models.LocationLang{
				Name: "las tres mentiras",
				Comments: "While we were drinking beer at Homo Sibaris Pub, we decided to eat something and nearby was this Mexican restaurant, we ate some tacos and there was an excellent variety of hot sauces, very good place to go to eat.",
			},
			ES: models.LocationLang{
				Name: "las tres mentiras",
				Comments: "Mientras tomábamos cerveza en el Pub Homo Sibaris, decidimos comer algo y cerca estaba este restaurante mexicano, comimos unos tacos y había una excelente variedad de salsas picantes, muy buen lugar para ir a comer.",
			},
			Coordinates: []float64{
				41.37622572053996, 2.1390920954518062,
			},
		},
		{
			EN: models.LocationLang{
				Name: "la whiskeria",
				Comments: "An elegant and comfortable place in Barcelona where I recommend going to anyone who is a whisky lover at least once. It has a huge variety of whiskies from all over the world.",
			},
			ES: models.LocationLang{
				Name: "la whiskeria",
				Comments: "Un lugar elegante y confortable en Barcelona donde recomiendo ir a todo aquel que sea amante del whisky al menos una vez. Tiene una enorme variedad de whiskies de todo el mundo.",
			},
			Coordinates: []float64{
				41.39179517363235, 2.173605566617036,
			},
		},
		{
			EN: models.LocationLang{
				Name: "lennox the pub",
				Comments: "I came here in my second week in barcelona, wandering aimlessly, exploring, I came here because it had a Guinness beer sign at the entrance, that same day I tried some good fries and went to the chocolate museum.",
			},
			ES: models.LocationLang{
				Name: "lennox the pub",
				Comments: "Vine aquí en mi segunda semana en barcelona, vagando sin rumbo, explorando, vine aquí porque tenía un cartel de cerveza Guinness en la entrada, ese mismo día probé unas buenas patatas fritas y fui al museo del chocolate.",
			},
			Coordinates: []float64{
				41.38342192142973, 2.1829632666164307,
			},
		},
		{
			EN: models.LocationLang{
				Name: "lidl",
				Comments: "The best donuts in Barcelona, I do not accept comments or arguments against. Oh and I recommend buying the Steambrew beers.",
			},
			ES: models.LocationLang{
				Name: "lidl",
				Comments: "Los mejores donuts de Barcelona, no acepto comentarios ni argumentos en contra. Ah y recomiendo comprar las cervezas Steambrew.",
			},
			Coordinates: []float64{
				41.41888596508152, 2.1989990024875166,
			},
		},
		{
			EN: models.LocationLang{
				Name: "l'insalata rica",
				Comments: "The second night of our trip to Rome my sister and I visited this place to rest for a while. Nice place, for the first time I drank a liter of beer served in a pitcher.",
			},
			ES: models.LocationLang{
				Name: "l'insalata rica",
				Comments: "La segunda noche de nuestro viaje a Roma mi hermana y yo visitamos este lugar para descansar un rato. Bonito lugar, por primera vez me bebí un litro de cerveza servida en una jarra.",
			},
			Coordinates: []float64{
				41.89767877476797, 12.472175713279132,
			},
		},
		{
			EN: models.LocationLang{
				Name: "llanero bodegón",
				Comments: "A bodega in my city, where I buy most of my drinks, has some variety, although much less than I would like.",
			},
			ES: models.LocationLang{
				Name: "llanero bodegón",
				Comments: "Una bodega de mi ciudad, donde compro la mayoría de mis bebidas, tiene algo de variedad, aunque mucha menos de la que me gustaría.",
			},
			Coordinates: []float64{
				10.06502710585951, -69.29467370287848,
			},
		},
		{
			EN: models.LocationLang{
				Name: "mercadona",
				Comments: "I used to love walking to the Mercadona to buy snacks for the house, especially the salted and smoked nuts cocktail. Here I bought many commercial beers.",
			},
			ES: models.LocationLang{
				Name: "mercadona",
				Comments: "Me encantaba ir al Mercadona a comprar aperitivos para la casa, sobre todo el cóctel de frutos secos salados y ahumados. Aquí compraba muchas cervezas comerciales.",
			},
			Coordinates: []float64{
				41.42448789144058, 2.207333967001515,
			},
		},
		{
			EN: models.LocationLang{
				Name: "museu nacional d'art de catalunya",
				Comments: "This day was quite fun, we didn't really get to see much of the museum since we visited it during the free access hour, but I was able to buy some details for my girlfriend.",
			},
			ES: models.LocationLang{
				Name: "museu nacional d'art de catalunya",
				Comments: "Este día fue bastante divertido, realmente no llegamos a ver mucho del museo ya que lo visitamos durante la hora de acceso libre, pero pude comprar algunos detalles para mi novia.",
			},
			Coordinates: []float64{
				41.36843987130215, 2.1535592666154066,
			},
		},
		{
			EN: models.LocationLang{
				Name: "ogham cervecería",
				Comments: "I found this site on one of those days when I was wandering aimlessly through the streets, I was looking for bookstores and by chance I entered. Later I met my sister and brother-in-law in a nearby restaurant.",
			},
			ES: models.LocationLang{
				Name: "ogham cervecería",
				Comments: "Encontré este sitio uno de esos días en que vagaba sin rumbo por las calles, buscaba librerías y por casualidad entré. Más tarde me encontré con mi hermana y mi cuñado en un restaurante cercano.",
			},
			Coordinates: []float64{
				41.38685214657274, 2.1653452954639167,
			},
		},
		{
			EN: models.LocationLang{
				Name: "plaça de catalunya",
				Comments: "I walked around here many times during those 3 months, here I tasted the first beer in the whole trip, it was at a food cart event.",
			},
			ES: models.LocationLang{
				Name: "plaça de catalunya",
				Comments: "Caminé por aquí muchas veces durante esos 3 meses, aquí probé la primera cerveza en todo el viaje, fue en un evento de carritos de comida.",
			},
			Coordinates: []float64{
				41.38702923415893, 2.170044525637977,
			},
		},
		{
			EN: models.LocationLang{
				Name: "setcases",
				Comments: "After visiting camprodon the trip continued to setcases, a small village with history that can amaze you with its landscape if you have an adventurous spirit.",
			},
			ES: models.LocationLang{
				Name: "setcases",
				Comments: "Tras visitar camprodon el viaje continuó hacia setcases, un pequeño pueblo con historia que puede sorprenderte con su paisaje si tienes espíritu aventurero.",
			},
			Coordinates: []float64{
				42.375359692857536, 2.30120794780043,
			},
		},
	}
	
	docsCreated, err := s.repo.InsertMany(data)

	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("%d records successfully created", docsCreated))

	return nil
}
