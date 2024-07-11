package formatter

const (
	happyEmoji        = "😊"
	loveFaceEmoji     = "😍"
	heartsFaceEmoji   = "🥰"
	satisfiedEmoji    = "😌"
	happyCatEmoji     = "😸"
	partyEmoji        = "🥳"
	neutralEmoji      = "😐"
	sadEmoji          = "😞"
	terrorEmoji       = "😱"
	hotEmoji          = "🥵"
	veryHotEmoji      = hotEmoji + terrorEmoji
	coldEmoji         = "🥶"
	veryColdEmoji     = coldEmoji + terrorEmoji
	wetEmoji          = "💦"
	dryEmoji          = "🌵"
	normalHumEmoji    = "🌿"
	pressureEmoji     = "🌀"
	cloudyEmoji       = "☁️"
	clearEmoji        = "🌞"
	bankEmoji         = "🏦"
	stonksEmoji       = "📈"
	airEmoji          = "💨"
	weatherEmoji      = "🏞️"
	emojiRich         = "🤑"
	emojiPoor         = "💸"
	dollarEmoji       = "💵"
	emojiThought      = "🤔"
	maskEmoji         = "😷"
	treeEmoji         = "🌳"
	skullEmoji        = "💀"
	thunderstormEmoji = "⛈️"
	rainEmoji         = "🌧️"
	snowEmoji         = "❄️"

	// weather labels
	attributeTemperature = "temperature"
	attributeHumidity    = "humidity"
	attributePressure    = "pressure"
	attributeDescription = "description"

	// air quality labels
	attributeCo   = "co"
	attributeNo2  = "no2"
	attributeO3   = "o3"
	attributePm10 = "pm10"
	attributePm25 = "pm25"
	attributeSo2  = "so2"

	// world bank labels
	attributeCPI                             = "Consumer price index (2010 = 100)"
	attributeCPIShort                        = "Consumer price index (2010=100)"
	attributeGDPPerCapita                    = "GDP per capita (current US$)"
	attributeGDPPerCapitaShort               = "GDP per capita"
	attributeExports                         = "Merchandise exports (current US$)"
	attributeExportsShort                    = "Exports"
	attributeImports                         = "Merchandise imports (current US$)"
	attributeImportsShort                    = "Imports"
	attributeUnemployment                    = "Unemployment, total (% of total labor force) (modeled ILO estimate)"
	attributeUnemploymentShort               = "Unemployment"
	attributeIndividualsUsingInternet        = "Individuals using the Internet (% of population)"
	attributeIndividualsUsingInternetShort   = "Internet users"
	attributeTaxRevenue                      = "Tax revenue (% of GDP)"
	attributeTaxRevenueShort                 = "Tax revenue"
	attributeLifeExpectancy                  = "Life expectancy at birth, total (years)"
	attributeLifeExpectancyShort             = "Life expectancy"
	attributeMortalityRateUnder5             = "Mortality rate, under-5 (per 1,000 live births)"
	attributeMortalityRateUnder5Short        = "Infant mortality"
	attributeGovtExpenditureEducation        = "Government expenditure on education, total (% of GDP)"
	attributeGovtExpenditureEducationShort   = "Spending on education"
	attributeCO2Emissions                    = "CO2 emissions (metric tons per capita)"
	attributeCO2EmissionsShort               = "CO2 emissions"
	attributeLiteracyRate                    = "Literacy rate, adult total (% of people ages 15 and above)"
	attributeLiteracyRateShort               = "Literacy rate"
	attributeCurrentHealthExpenditure        = "Current health expenditure (% of GDP)"
	attributeCurrentHealthExpenditureShort   = "Health spending"
	attributePovertyHeadcount                = "Poverty headcount ratio at $2.15 a day (2017 PPP) (% of population)"
	attributePovertyHeadcountShort           = "Poverty"
	attributeHealthExpenditurePerCapita      = "Current health expenditure per capita, PPP (current international $)"
	attributeHealthExpenditurePerCapitaShort = "Health spending per capita"

	// prices
	attributePairOfJeans      = "1 Pair of Jeans (Levis 501 Or Similar)"
	attributePairOfJeansShort = "Pair of Jeans"

	attributeApartment1BedroomOutsideCentre      = "Apartment (1 bedroom) Outside of Centre"
	attributeApartment1BedroomOutsideCentreShort = "Rent 1-bedroom apartment (not central)"

	attributeApartment1BedroomCityCentre      = "Apartment (1 bedroom) in City Centre"
	attributeApartment1BedroomCityCentreShort = "Rent 1-bedroom apartment (central)"

	attributeApples      = "Apples (1 lb)"
	attributeApplesShort = "1kg of apples"

	attributeMonthlyNetSalary      = "Average Monthly Net Salary (After Tax)"
	attributeMonthlyNetSalaryShort = "Monthly net salary"

	attributeBanana      = "Banana (1 lb)"
	attributeBananaShort = "1kg of bananas"

	attributeBasicUtilities      = "Basic (Electricity, Heating, Cooling, Water, Garbage) for 915 sq ft Apartment"
	attributeBasicUtilitiesShort = "Basic utilities for 85m² apartment"

	attributeBeefRound      = "Beef Round (1 lb) (or Equivalent Back Leg Red Meat)"
	attribureBeefRoundShort = "1kg of beef round"

	attributeBottleOfWine      = "Bottle of Wine (Mid-Range)"
	attributeBottleOfWineShort = "Bottle of wine"

	attributeCappuccino      = "Cappuccino (regular)"
	attributeCappuccinoShort = "Cappuccino"

	attributeChickenFillets      = "Chicken Fillets (1 lb)"
	attributeChickenFilletsShort = "1kg of chicken fillets"

	attributeCigarettes      = "Cigarettes 20 Pack (Marlboro)"
	attributeCigarettesShort = "Pack of cigarettes"

	attributeCinemaSeat = "Cinema, International Release, 1 Seat"
	attributeCokePepsi  = "Coke/Pepsi (12 oz small bottle)"

	attributeDomesticBeerBottle      = "Domestic Beer (0.5 liter bottle)"
	attributeDomesticBeerBottleShort = "Beer bottle (0.5l)"

	attributeEggs      = "Eggs (regular) (12)"
	attributeEggsShort = "12 eggs"

	attributeGasoline      = "Gasoline (1 gallon)"
	attributeGasolineShort = "1l of gasoline"

	attributeInternationalPrimarySchool      = "International Primary School, Yearly for 1 Child"
	attributeInternationalPrimarySchoolShort = "Primary school (yearly)"

	attributeInternet      = "Internet (60 Mbps or More, Unlimited Data, Cable/ADSL)"
	attributeInternetShort = "Internet (60 Mbps or more)"

	attributeBread      = "Loaf of Fresh White Bread (1 lb)"
	attributeBreadShort = "1kg of bread"

	attributeLocalCheese      = "Local Cheese (1 lb)"
	attributeLocalCheeseShort = "1kg of cheese"

	attributeMcMeal      = "McMeal at McDonalds (or Equivalent Combo Meal)"
	attributeMcMealShort = "Meal at McDonalds or similar"

	attributeMealFor2      = "Meal for 2 People, Mid-range Restaurant, Three-course"
	attributeMealFor2Short = "Meal for 2 in mid-range restaurant"

	attributeMilk      = "Milk (regular), (1 gallon)"
	attributeMilkShort = "1l of milk"

	attributeMobilePlan      = "Mobile Phone Monthly Plan with Calls and 10GB+ Data"
	attributeMobilePlanShort = "Mobile plan (monthly)"

	attributeMonthlyPass      = "Monthly Pass (Regular Price)"
	attributeMonthlyPassShort = "Monthly pass"

	attributeMortgageInterestRate      = "Mortgage Interest Rate in Percentages (%), Yearly, for 20 Years Fixed-Rate"
	attributeMortgageInterestRateShort = "Mortgage interest rate"

	attributeOneWayTicket      = "One-way Ticket (Local Transport)"
	attributeOneWayTicketShort = "Local one-way ticket"

	attributeOranges      = "Oranges (1 lb)"
	attributeOrangesShort = "1kg of oranges"

	attributePotato      = "Potato (1 lb)"
	attributePotatoShort = "1kg of potatoes"

	attributePreschool      = "Preschool (or Kindergarten), Full Day, Private, Monthly for 1 Child"
	attributePreschoolShort = "Kindergarten (monthly)"

	attributeTaxi1Mile      = "Taxi 1 mile (Normal Tariff)"
	attributeTaxi1MileShort = "Taxi 1 km"

	attributeTomato      = "Tomato (1 lb)"
	attributeTomatoShort = "1kg of tomatoes"

	attributeWaterBottle      = "Water (1.5 liter bottle)"
	attributeWaterBottleShort = "1.5l water bottle"

	// ignored for now
	attributePricePerSqFtOutsideCentre       = "Price per Square Feet to Buy Apartment Outside of Centre"
	attributePricePerSqFtCityCentre          = "Price per Square Feet to Buy Apartment in City Centre"
	attributeMealInexpensiveRestaurant       = "Meal, Inexpensive Restaurant"
	attributeDomesticBeerPint                = "Domestic Beer (1 pint draught)"
	attributeRice                            = "Rice (white), (1 lb)"
	attributeTaxi1HourWaiting                = "Taxi 1hour Waiting (Normal Tariff)"
	attributeFitnessClubFee                  = "Fitness Club, Monthly Fee for 1 Adult"
	attributeTaxiStart                       = "Taxi Start (Normal Tariff)"
	attributeTennisCourtRent                 = "Tennis Court Rent (1 Hour on Weekend)"
	attributeToyotaCorolla                   = "Toyota Corolla Sedan 1.6l 97kW Comfort (Or Equivalent New Car)"
	attributeApartment3BedroomsOutsideCentre = "Apartment (3 bedrooms) Outside of Centre"
	attributeApartment3BedroomsCityCentre    = "Apartment (3 bedrooms) in City Centre"
	attributeMenLeatherBusinessShoes         = "1 Pair of Men Leather Business Shoes"
	attributeNikeRunningShoes                = "1 Pair of Nike Running Shoes (Mid-Range)"
	attributeSummerDressChainStore           = "1 Summer Dress in a Chain Store (Zara, H&M, ...)"
	attributeVolkswagenGolf                  = "Volkswagen Golf 1.4 90 KW Trendline (Or Equivalent New Car)"
	attributeWaterSmallBottle                = "Water (12 oz small bottle)"
	attributeOnion                           = "Onion (1 lb)"
	attributeImportedBeer                    = "Imported Beer (12 oz small bottle)"
	attributeLettuce                         = "Lettuce (1 head)"

	// categories
	catergoryWeather    = "weather"
	catergoryAirQuality = "air_quality"
	categoryWorldBank   = "world_bank"
	categoryPrices      = "prices"

	// subgroups
	noSubgroup   = "no_subgroup"
	foodAndDrink = "Food & Drink"
	transport    = "Transport"
	education    = "Education"
	apartment    = "Apartment"

	// provider
	provider = "tinygods.dev"
)
