package graph

import (
	context "context"
	"strconv"
	"strings"
	time "time"

	"github.com/tonyghita/gqlgen-example/swapi"
)

// Application implements the generated Resolvers interface.
type Application struct {
	Client *swapi.Client
}

func (app Application) Query_films(ctx context.Context, title *string) ([]Film, error) {
	page, err := app.Client.SearchFilms(ctx, str(title))
	if err != nil {
		return nil, err
	}

	films := make([]Film, len(page.Films))
	for i, f := range page.Films {
		film, err := convertFilm(f)
		if err != nil {
			return nil, err // TODO: how should list element errors be handled?
		}

		films[i] = film
	}

	return films, nil
}

func (app Application) Query_people(ctx context.Context, name *string) ([]Person, error) {
	page, err := app.Client.SearchPerson(ctx, str(name))
	if err != nil {
		return nil, err
	}

	people := make([]Person, len(page.People))
	for i, p := range page.People {
		person, err := convertPerson(p)
		if err != nil {
			return nil, err
		}

		people[i] = person
	}

	return people, nil
}

func (app Application) Query_planets(ctx context.Context, name *string) ([]Planet, error) {
	page, err := app.Client.SearchPlanets(ctx, str(name))
	if err != nil {
		return nil, err
	}

	planets := make([]Planet, len(page.Planets))
	for i, p := range page.Planets {
		planet, err := convertPlanet(p)
		if err != nil {
			return nil, err
		}

		planets[i] = planet
	}

	return planets, nil
}

func (app Application) Query_species(ctx context.Context, name *string) ([]Species, error) {
	page, err := app.Client.SearchSpecies(ctx, str(name))
	if err != nil {
		return nil, err
	}

	species := make([]Species, len(page.Species))
	for i, s := range page.Species {
		sp, err := convertSpecies(s)
		if err != nil {
			return nil, err
		}

		species[i] = sp
	}

	return species, nil
}

func (app Application) Query_starships(ctx context.Context, nameOrModel *string) ([]Starship, error) {
	page, err := app.Client.SearchStarships(ctx, str(nameOrModel))
	if err != nil {
		return nil, err
	}

	starships := make([]Starship, len(page.Starships))
	for i, s := range page.Starships {
		starship, err := convertStarship(s)
		if err != nil {
			return nil, err
		}

		starships[i] = starship
	}

	return starships, nil
}

func (app Application) Query_vehicles(ctx context.Context, nameOrModel *string) ([]Vehicle, error) {
	page, err := app.Client.SearchVehicles(ctx, str(nameOrModel))
	if err != nil {
		return nil, err
	}

	vehicles := make([]Vehicle, len(page.Vehicles))
	for i, v := range page.Vehicles {
		vehicle, err := convertVehicle(v)
		if err != nil {
			return nil, err
		}

		vehicles[i] = vehicle
	}

	return vehicles, nil
}

func convertFilm(f swapi.Film) (Film, error) {
	releaseDate, err := date(f.ReleaseDate)
	if err != nil {
		return Film{}, err
	}

	createdAt, err := datetime(f.CreatedAt)
	if err != nil {
		return Film{}, err
	}

	editedAt, err := nullableDatetime(f.EditedAt)
	if err != nil {
		return Film{}, err
	}

	return Film{
		ID:            id(f.URL),
		Episode:       int(f.EpisodeID),
		OpeningCrawl:  f.OpeningCrawl,
		DirectorName:  f.DirectorName,
		ProducerNames: strings.Split(f.ProducerNames, ", "),
		ReleaseDate:   releaseDate,
		CreatedAt:     createdAt,
		EditedAt:      editedAt,
		// TODO: Add Species, Starships, Vehicles, Characters, and Planets.
		// How do we request that data only if user requested?
	}, nil
}

func convertPerson(p swapi.Person) (Person, error) {
	height, err := float(p.Height)
	if err != nil {
		return Person{}, err
	}

	mass, err := float(p.Mass)
	if err != nil {
		return Person{}, err
	}

	createdAt, err := datetime(p.CreatedAt)
	if err != nil {
		return Person{}, err
	}

	editedAt, err := nullableDatetime(p.EditedAt)
	if err != nil {
		return Person{}, err
	}

	return Person{
		ID:        id(p.URL),
		Name:      p.Name,
		BirthYear: p.BirthYear,
		EyeColor:  nullableStr(p.EyeColor),
		Gender:    nullableStr(p.Gender),
		HairColor: nullableStr(p.HairColor),
		Height:    height,
		Mass:      mass,
		SkinColor: nullableStr(p.SkinColor),
		CreatedAt: createdAt,
		EditedAt:  editedAt,
	}, nil
}

func convertPlanet(p swapi.Planet) (Planet, error) {
	diameter, err := float(p.Diameter)
	if err != nil {
		return Planet{}, err
	}

	rotationPeriod, err := float(p.RotationPeriod)
	if err != nil {
		return Planet{}, err
	}

	orbitalPeriod, err := float(p.OrbitalPeriod)
	if err != nil {
		return Planet{}, err
	}

	gravity, err := float(p.Gravity)
	if err != nil {
		return Planet{}, err
	}

	surfaceWaterPercentage, err := float(p.SurfaceWater)
	if err != nil {
		return Planet{}, err
	}

	createdAt, err := datetime(p.CreatedAt)
	if err != nil {
		return Planet{}, err
	}

	editedAt, err := nullableDatetime(p.EditedAt)
	if err != nil {
		return Planet{}, err
	}

	return Planet{
		ID:                     id(p.URL),
		Name:                   p.Name,
		Diameter:               diameter,
		RotationPeriod:         rotationPeriod,
		OrbitalPeriod:          orbitalPeriod,
		Gravity:                gravity,
		Climates:               strings.Split(p.Climate, ", "),
		Terrains:               strings.Split(p.Terrain, ", "),
		SurfaceWaterPercentage: surfaceWaterPercentage,
		CreatedAt:              createdAt,
		EditedAt:               editedAt,
	}, nil
}

func convertSpecies(s swapi.Species) (Species, error) {
	avgHeight, err := float(s.AverageHeight)
	if err != nil {
		return Species{}, err
	}

	avgLifespan, err := float(s.AverageLifespan)
	if err != nil {
		return Species{}, err
	}

	createdAt, err := datetime(s.CreatedAt)
	if err != nil {
		return Species{}, err
	}

	editedAt, err := nullableDatetime(s.EditedAt)
	if err != nil {
		return Species{}, err
	}

	return Species{
		ID:              id(s.URL),
		Name:            s.Name,
		Classification:  s.Classification,
		Designation:     s.Designation,
		AverageHeight:   avgHeight,
		AverageLifespan: avgLifespan,
		EyeColors:       strings.Split(s.EyeColors, ", "),
		HairColors:      strings.Split(s.HairColors, ", "),
		SkinColors:      strings.Split(s.SkinColors, ", "),
		Language:        s.Language,
		CreatedAt:       createdAt,
		EditedAt:        editedAt,
	}, nil
}

func convertStarship(s swapi.Starship) (Starship, error) {
	cost, err := integer(s.CostInCredits)
	if err != nil {
		return Starship{}, err
	}

	length, err := float(s.Length)
	if err != nil {
		return Starship{}, err
	}

	crew, err := integer(s.Crew)
	if err != nil {
		return Starship{}, err
	}

	passengerCap, err := integer(s.Passengers)
	if err != nil {
		return Starship{}, err
	}

	maxSpeed, err := nullableInteger(s.MaxAtmospheringSpeed)
	if err != nil {
		return Starship{}, err
	}

	hyperdriveRating, err := nullableFloat(s.HyperdriveRating)
	if err != nil {
		return Starship{}, err
	}

	maxMPH, err := integer(s.MGLT)
	if err != nil {
		return Starship{}, err
	}

	cargoCap, err := float(s.CargoCapacity)
	if err != nil {
		return Starship{}, err
	}

	createdAt, err := datetime(s.CreatedAt)
	if err != nil {
		return Starship{}, err
	}

	editedAt, err := nullableDatetime(s.EditedAt)
	if err != nil {
		return Starship{}, err
	}

	return Starship{
		ID:                   id(s.URL),
		Name:                 s.Name,
		Model:                s.Model,
		Class:                s.StarshipClass,
		Manufacturers:        strings.Split(s.Manufacturer, ", "),
		Cost:                 cost,
		Length:               length,
		CrewSize:             crew,
		PassengerCapacity:    passengerCap,
		MaxAtmosphericSpeed:  maxSpeed,
		HyperdriveRating:     hyperdriveRating,
		MaxMegalightsPerHour: maxMPH,
		CargoCapacity:        cargoCap,
		ConsumablesDuration:  s.Consumables,
		CreatedAt:            createdAt,
		EditedAt:             editedAt,
	}, nil
}

func convertVehicle(v swapi.Vehicle) (Vehicle, error) {
	length, err := float(v.Length)
	if err != nil {
		return Vehicle{}, err
	}

	cost, err := integer(v.CostInCredits)
	if err != nil {
		return Vehicle{}, err
	}

	crewSize, err := integer(v.Crew)
	if err != nil {
		return Vehicle{}, err
	}

	passengerCap, err := integer(v.Passengers)
	if err != nil {
		return Vehicle{}, err
	}

	maxSpeed, err := float(v.MaxAtmospheringSpeed)
	if err != nil {
		return Vehicle{}, err
	}

	cargoCap, err := float(v.CargoCapacity)
	if err != nil {
		return Vehicle{}, err
	}

	createdAt, err := datetime(v.CreatedAt)
	if err != nil {
		return Vehicle{}, err
	}

	editedAt, err := nullableDatetime(v.EditedAt)
	if err != nil {
		return Vehicle{}, err
	}

	return Vehicle{
		ID:                  id(v.URL),
		Name:                v.Name,
		Model:               v.Model,
		Class:               v.VehicleClass,
		Manufacturers:       strings.Split(v.Manufacturer, ", "),
		Length:              length,
		Cost:                cost,
		CrewSize:            crewSize,
		PassengerCapacity:   passengerCap,
		MaxAtmosphericSpeed: maxSpeed,
		CargoCapacity:       cargoCap,
		ConsumablesDuration: v.Consumables,
		CreatedAt:           createdAt,
		EditedAt:            editedAt,
	}, nil
}

func date(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

func datetime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

func nullableDatetime(s string) (*time.Time, error) {
	if s == "" {
		return nil, nil
	}

	t, err := datetime(s)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

var alphaChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func float(s string) (float64, error) {
	if strings.ContainsAny(s, alphaChars) {
		return 0.0, nil
	}

	return strconv.ParseFloat(s, 64)
}

func nullableFloat(s string) (*float64, error) {
	if s == "" {
		return nil, nil
	}

	f, err := float(s)
	if err != nil {
		return nil, err
	}

	return &f, nil
}

func integer(s string) (int, error) {
	if strings.ContainsAny(s, alphaChars) {
		return 0, nil
	}

	i64, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return int(i64), nil
}

func nullableInteger(s string) (*int, error) {
	if s == "" {
		return nil, nil
	}

	i, err := integer(s)
	if err != nil {
		return nil, err
	}

	return &i, nil
}

func str(p *string) string {
	if p == nil {
		return ""
	}

	return *p
}

func nullableStr(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}

func id(url string) string {
	if url == "" {
		return ""
	}

	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		return parts[0]
	}

	return parts[len(parts)-2]
}
